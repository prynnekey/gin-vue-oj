package service

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/define"
	"github.com/prynnekey/gin-vue-oj/models"
	"gorm.io/gorm"
)

// GetCategoryList
// @Summary 获取分类列表
// @Param authorization header string false "token"
// @Param page query int false "请输入当前页,默认第一页"
// @Param pageSize query int false "每页多少条数据,默认20条"
// @Param keyWord query string false "关键字"
// @Description 获取分类列表
// @Tags 管理员私有方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /admin/category [get]
func GetCategoryList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", define.PROBLEM_GET_PAGE))
		pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", define.PROBLEM_GET_PAGE_SIZE))
		keyWord := ctx.Query("keyWord")

		if err != nil {
			log.Println("GetProblemList Param strconv Error:", err)
			response.Failed(ctx, "参数类型错误")
			return
		}

		// 查询数据库
		cb, count, err := models.GetCategory(page, pageSize, keyWord)
		if err != nil {
			response.Failed(ctx, "查询失败")
			return
		}

		// 返回信息
		response.Success(ctx, gin.H{
			"count": count,
			"list":  cb,
		}, "查询成功")
	}
}

// AddCategory
// @Summary 新增分类
// @Param authorization header string false "token"
// @Param name formData string false "分类名称 例如:数组"
// @Param parent_id formData int false "父级分类id 默认:0(顶级id)"
// @Description 新增分类
// @Tags 管理员私有方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /admin/category [post]
func AddCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		parentId := ctx.DefaultPostForm("parent_id", "0")

		// 参数校验
		if name == "" {
			response.Failed(ctx, "分类名称不能为空")
			return
		}

		// 添加到数据库
		err := models.AddCategory(name, parentId)
		if err != nil {
			if s := strings.Split(err.Error(), " "); s[3] == "Duplicate" {
				response.Failed(ctx, "添加失败:'"+name+"'已存在")
				return
			}
			response.Failed(ctx, "添加失败:"+err.Error())
			return
		}

		// 返回信息
		response.Success(ctx, nil, "添加成功")
	}
}

// DeleteCategoryById
// @Summary 根据id删除分类
// @Param authorization header string false "token"
// @Param id query int false "要删除的分类id"
// @Description 根据id删除分类
// @Tags 管理员私有方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /admin/category [delete]
func DeleteCategoryById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		id := ctx.Query("id")

		// 校验参数
		if id == "" {
			response.Failed(ctx, "id不能为空")
			return
		}

		// 根据id删除
		i, isDelete, err := models.DeleteCategoryById(id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response.Failed(ctx, "数据不存在")
				return
			}
			response.Failed(ctx, "发生错误:"+err.Error())
			return
		}

		if isDelete {
			response.Failed(ctx, "删除失败:该分类下有子分类")
			return
		}

		if i == 0 {
			response.Failed(ctx, "数据不存在")
			return
		}

		response.Success(ctx, nil, "删除成功")
	}
}

// UpdateCategoryById
// @Summary 根据id修改分类
// @Param authorization header string false "token"
// @Param id query int false "要修改的分类id"
// @Param name query string false "修改后分类名称"
// @Param parent_id query int false "修改后分类父id"
// @Description 根据id修改分类
// @Tags 管理员私有方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /admin/category [put]
func UpdateCategoryById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		id := ctx.Query("id")
		name := ctx.Query("name")
		parentId := ctx.Query("parent_id")

		// 校验参数
		if id == "" {
			response.Failed(ctx, "id不能为空")
			return
		}

		if name == "" {
			response.Failed(ctx, "分类名称不能为空")
			return
		}

		if parentId == "" {
			response.Failed(ctx, "父id不能为空")
			return
		}

		// 根据id修改
		i, err := models.UpdateCategoryById(id, name, parentId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response.Failed(ctx, "数据不存在")
				return
			}
			response.Failed(ctx, "发生错误:"+err.Error())
			return
		}

		if i == 0 {
			response.Failed(ctx, "数据不存在")
			return
		}

		// 返回结果
		response.Success(ctx, nil, "修改成功")
	}
}
