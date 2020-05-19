package dao

import (
	"bookstore0612/model"
	"bookstore0612/utils"
)

// GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	// 写sql
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books, nil
}

// AddBook 向数据库中添加一本图书
func AddBook(b *model.Book) error {
	//写sql
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}