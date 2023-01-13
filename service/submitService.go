package service

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/define"
	"github.com/prynnekey/gin-vue-oj/models"
	"github.com/prynnekey/gin-vue-oj/utils"
	"gorm.io/gorm"
)

// GetSubmitList
// @Summary 获取提交记录列表
// @Param page query int false "请输入当前页,默认第一页"
// @Param pageSize query int false "每页多少条数据,默认20条"
// @Param problem_identity query string false "问题的唯一标识"
// @Param user_identity query string false "用户的唯一标识"
// @Param status query int false "提交的状态【-1-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存】"
// @Description 获取问题列表
// @Tags 公共方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /submit-list [get]
func GetSubmitList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", define.PROBLEM_GET_PAGE))
		pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", define.PROBLEM_GET_PAGE_SIZE))
		problemIdentity := ctx.Query("problem_identity")
		userIdentity := ctx.Query("user_identity")
		status, _ := strconv.Atoi(ctx.Query("status"))

		if err != nil {
			log.Println("GetProblemList Param strconv Error:", err)
			response.Failed(ctx, "参数类型错误")
			return
		}

		// 查询数据库
		submitList, count, err := models.GetSubmitList(page, pageSize, problemIdentity, userIdentity, status)
		if err != nil {
			response.Failed(ctx, "查询数据库出错:"+err.Error())
			return
		}

		response.Success(ctx, gin.H{"count": count, "list": submitList}, "查询成功")
	}
}

// SubmitCode()
// @Summary 用户提交代码
// @Description 用户提交代码
// @Tags 用户私有方法
// @Param authorization header string false "用户token"
// @Param problem_identity query string false "问题的唯一标识"
// @Param code body string false "用户提交的代码"
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /user/submit [post]
func SubmitCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		problemIdentity := ctx.Query("problem_identity")
		code, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			response.Failed(ctx, "读取代码失败"+err.Error())
			return
		}

		// 参数校验
		if len(problemIdentity) != 36 {
			response.Failed(ctx, "问题唯一标识不正确")
			return
		}

		_userClaim, ok := ctx.Get("user")
		if !ok {
			response.Failed(ctx, "请重新登录")
			return
		}
		userClaim := _userClaim.(*utils.UserClaims)

		// 将code保存到指定目录
		codePath, err := utils.SaveCode(userClaim.Identity, problemIdentity, code)
		if err != nil {
			// 同时删除创建的文件
			err2 := utils.DeleteSaveCode(codePath)
			if err2 != nil {
				response.Failed(ctx, " 删除代码失败:"+err2.Error())
				return
			}
			response.Failed(ctx, "提交失败,保存代码时出错:"+err.Error())
			return
		}

		// 判断该问题是否存在
		problem := models.ProblemBasic{}
		err = models.DB.Preload("TestCases").Where("identity = ?", problemIdentity).First(&problem).Error
		if err != nil {
			// 同时删除创建的文件
			err2 := utils.DeleteSaveCode(codePath)
			if err2 != nil {
				response.Failed(ctx, " 删除代码失败:"+err2.Error())
				return
			}
			if err == gorm.ErrRecordNotFound {
				response.Failed(ctx, "该问题不存在")
				return
			}
			response.Failed(ctx, "问题查询时失败:"+err.Error())
			return
		}

		// ----------------------代码判断的核心-------------------------

		// 答案错误的channel
		answerError := make(chan int)

		// 编译错误channel
		compileError := make(chan int)

		// 运行超内存
		outOfMemory := make(chan int)

		// 答案正确
		answerCorrect := make(chan int)

		// 非法代码
		illegalCode := make(chan int)

		// 答案通过的个数
		passCount := 0
		// 读写锁 防止passCount的并发问题
		lock := sync.RWMutex{}

		// 最终答案【-1-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存，5编译错误】
		status := -1

		// 向用户提示最后的状态
		msg := ""

		// 运行code
		for _, testCase := range problem.TestCases {
			// 将code与标准输入输出对比

			// 开启goroutine多协程 异步执行代码
			go func(testCase *models.TestCase) {
				// 执行用户的代码 go run userCode/main.go
				cmd := exec.Command("go", "run", codePath)
				var stdout, stderr bytes.Buffer
				cmd.Stderr = &stderr
				cmd.Stdout = &stdout

				// 进行标准输入
				stdinPipe, err := cmd.StdinPipe()
				if err != nil {
					log.Println(err)
				}
				defer stdinPipe.Close()

				_, err = io.WriteString(stdinPipe, testCase.Input)
				if err != nil {
					log.Println(err)
				}

				var beginMem runtime.MemStats
				runtime.ReadMemStats(&beginMem)
				// 读取标准输出
				if err := cmd.Run(); err != nil {
					log.Println(err, stderr.String())
					// 编译错误
					if err.Error() == "exit status 2" {
						/* s := strings.Split(stderr.String(), "\\")
						var sb strings.Builder
						for index, value := range s {
							if index%4 == 0 {
								sb.WriteString(value)
							}
						}
						msg = sb.String() */
						msg = stderr.String()
						compileError <- 1
						return
					}
					illegalCode <- 1
					msg = err.Error()
					return
				}
				var endMem runtime.MemStats
				runtime.ReadMemStats(&endMem)

				// 答案错误
				if testCase.Output != stdout.String() {
					msg = "答案错误。输入:[" + testCase.Input + "]" + "期望输出:[" + testCase.Output + "]" + "实际输出:[" + stdout.String() + "]"
					answerError <- 1
					return
				}

				// 运行超内存
				if endMem.Alloc/1024-beginMem.Alloc/1024 > uint64(problem.MaxMem) {
					msg = "运行超内存"
					outOfMemory <- 1
					return
				}

				// 编译没错 答案没错 运行没超 只能答案正确
				lock.Lock()
				passCount++
				if passCount == len(problem.TestCases) {
					// 答案正确
					answerCorrect <- 1
				}
				lock.Unlock()

			}(testCase)

		}

		// 在这里阻塞
		// 得出结论 submit_basic.status = 【-1-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存，5编译错误】
		select {
		case <-answerError:
			// 答案错误
			status = 2
		case <-outOfMemory:
			// 运行超内存
			status = 4
		case <-compileError:
			// 编译错误
			status = 5
		case <-illegalCode:
			// 非法代码
			status = 5
		case <-answerCorrect:
			// 答案正确
			status = 1
			msg = "答案正确！"
		case <-time.After(time.Millisecond * time.Duration(problem.MaxRuntime)):
			// 运行超时
			status = 3
			msg = "运行超时"
		}

		// -------------------------------------------------------------------

		// 开启事务
		err = models.DB.Transaction(func(tx *gorm.DB) error {

			// 将数据保存到submit_basic表
			submitBasic := models.SubmitBasic{
				Identity:        utils.GenerateUUID(),
				ProblemIdentity: problemIdentity,
				UserIdentity:    userClaim.Identity,
				CodePath:        codePath,
				Status:          status,
			}
			err = tx.Create(&submitBasic).Error
			if err != nil {
				// 只要有err自动回滚事务
				return err
			}

			// 该用户的提交问题的个数+1 根据submit_basic.status判断答对的个数是否+1
			// 查询用户数据
			userBasic := models.UserBasic{}
			err = tx.Where("identity = ?", userClaim.Identity).First(&userBasic).Error
			if err != nil {
				// 只要有err自动回滚事务
				return err
			}

			// 要更新的字段
			updateUser := models.UserBasic{
				SubmitProblemNum: userBasic.SubmitProblemNum + 1,
			}

			if status == 1 {
				updateUser.FinishProblemNum = userBasic.FinishProblemNum + 1
			}

			// 更新数据
			err = tx.Model(&userBasic).Updates(updateUser).Error
			if err != nil {
				return err
			}

			// 返回nil 提交事务
			return nil
		})

		// 事务出现错误 回滚
		if err != nil {
			// 同时删除创建的文件
			err2 := utils.DeleteSaveCode(codePath)
			if err2 != nil {
				response.Failed(ctx, " 删除代码失败:"+err2.Error())
				return
			}
			response.Failed(ctx, "提交失败:"+err.Error())
			return
		}

		// 返回结果
		response.Success(ctx, gin.H{
			"status":    status,
			"msg":       msg,
			"passCount": passCount,
		}, "提交成功")
	}
}
