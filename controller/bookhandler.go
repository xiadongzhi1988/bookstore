package controller

import (
	"bookstore0612/dao"
	"bookstore0612/model"
	"html/template"
	"net/http"
	"strconv"
)

//GetBooks 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request)  {
	//调用bookdao中获取所有图书的函数
	books, _ := dao.GetBooks()
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, books)
}
// AddBook添加图书
func AddBook(w http.ResponseWriter, r *http.Request)  {
	//获取图书信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//将价格,销量,库存转换
	fPrice, _ := strconv.ParseFloat(price,64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	//创建Book
	book := &model.Book{
		Title: title,
		Author: author,
		Price: fPrice,
		Sales: int(iSales),
		Stock: int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	//调用bookdao中添加图书的函数
	dao.AddBook(book)
	//调用GetBooks函数再查询一次
	GetBooks(w, r)

}