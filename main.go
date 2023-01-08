package main

import (
	"github.com/prynnekey/gin-vue-oj/router"
)

func main() {
	r := router.Init()

	r.Run()
}
