package controller

import (
	"bookstore0612/dao"
	"bookstore0612/model"
	"html/template"
	"net/http"
	"strconv"
)

// IndexHandler  去首页
func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中获取带分页的图书函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}

//GetBooks 获取所有图书
//func GetBooks(w http.ResponseWriter, r *http.Request)  {
//	//调用bookdao中获取所有图书的函数
//	books, _ := dao.GetBooks()
//	//解析模板文件
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	//执行
//	t.Execute(w, books)
//}

// GetPageBooks 获取带分页的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request)  {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中获取带分页的图书函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

// GetPageBooksByPrice 获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request)  {
	//获取页码
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao中获取带分页的图书函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//调用bookdao中获取带分页的图书函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

// AddBook添加图书
//func AddBook(w http.ResponseWriter, r *http.Request)  {
//	//获取图书信息
//	title := r.PostFormValue("title")
//	author := r.PostFormValue("author")
//	price := r.PostFormValue("price")
//	sales := r.PostFormValue("sales")
//	stock := r.PostFormValue("stock")
//	//将价格,销量,库存转换
//	fPrice, _ := strconv.ParseFloat(price,64)
//	iSales, _ := strconv.ParseInt(sales, 10, 0)
//	iStock, _ := strconv.ParseInt(stock, 10, 0)
//	//创建Book
//	book := &model.Book{
//		Title: title,
//		Author: author,
//		Price: fPrice,
//		Sales: int(iSales),
//		Stock: int(iStock),
//		ImgPath: "/static/img/default.jpg",
//	}
//	//调用bookdao中添加图书的函数
//	dao.AddBook(book)
//	//调用GetBooks函数再查询一次
//	GetBooks(w, r)
//
//}

// DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request)  {
	//获取要删除的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中删除图书函数
	dao.DeleteBook(bookID)
	//调用GetBooks函数再查询一次
	GetPageBooks(w, r)

}

// ToUpdateBookPage 去更新或添加图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request)  {
	//获取要更新的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中获取图书函数
	book, _ := dao.GetBookByID(bookID)
	if book.Id > 0 {
		//在更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)
	} else {
		//添加图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, "")
	}

}

// UpdateOrAddBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request)  {
	//获取图书信息
	bookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//将价格,销量,库存转换
	fPrice, _ := strconv.ParseFloat(price,64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	ibookId, _ := strconv.ParseInt(bookId, 10, 0)
	//创建Book
	book := &model.Book{
		Id: int(ibookId),
		Title: title,
		Author: author,
		Price: fPrice,
		Sales: int(iSales),
		Stock: int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.Id > 0 {
		//在更新图书
		//调用bookdao中更新图书的函数
		dao.UpdateBook(book)
	} else {
		//在添加图书
		//调用bookdao中添加图书函数
		dao.AddBook(book)
	}
	//调用GetBooks函数再查询一次
	GetPageBooks(w, r)
}

