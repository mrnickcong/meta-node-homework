package main

import (
	"fmt"
	"log"
	// 添加正确的 sqlx 包导入
	"github.com/jmoiron/sqlx"
	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

/**

假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。

要求 ：

- 定义一个 Book 结构体，包含与 books 表对应的字段。
- 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func queryBooksByPrice(db *sqlx.DB, price float64) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM t_books WHERE price > ?", price)
	return books, err

}

func main() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/ethan_blog")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	books, err := queryBooksByPrice(db, 10.0)
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

}
