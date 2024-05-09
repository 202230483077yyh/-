package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{})
	//路由分发
	// routers.AdminRoutersInit(r)
	// routers.APIRoutersInit(r)
	// routers.DefaultRoutersInit(r)

	r.Run()

	//数据库交互测试用例
	// import "db_inter"
	// str1 := make([]string, 0)
	// str1 = append(str1, " name ")
	// m := make(map[string]string, 0)
	// //注意名字前不能多空格！
	// m["tag"] = `"清真菜 "`
	// str := db_inter.Get_caipu_conlum(str1, m)
	// for id, val := range str {
	// 	fmt.Println(id, val)
	// }

}
