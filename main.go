package main

import (
	"bookstore0612/controller"
	"html/template"
	"net/http"
)

// IndexHandler  去首页
func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, "")
}
func main() {
	//设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	http.HandleFunc("/main", IndexHandler)
	//去登陆
	http.HandleFunc("/login", controller.Login)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	http.HandleFunc("/getBooks", controller.GetBooks)
	http.ListenAndServe(":8080", nil)
}