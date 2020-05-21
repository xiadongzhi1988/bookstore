package main

import (
	"bookstore0612/controller"
	"net/http"
)


func main() {
	// 31
	//设置处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	http.HandleFunc("/main", controller.IndexHandler)
	//去登陆
	http.HandleFunc("/login", controller.Login)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)
	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//添加图书
	//http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//更新或添加图书
	http.HandleFunc("/updateOraddBook", controller.UpdateOrAddBook)
	http.ListenAndServe(":8080", nil)
}