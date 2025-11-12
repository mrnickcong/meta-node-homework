package main

import (
	"fmt"
	"log"
	// 添加正确的 sqlx 包导入
	"github.com/jmoiron/sqlx"
	// 数据库驱动
	_ "github.com/go-sql-driver/mysql"
)

/*
题目1：使用SQL扩展库进行查询

假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：

- 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
- 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func queryEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM t_employees WHERE department = ? ORDER BY salary DESC LIMIT 1", department)
	return employees, err

}
func queryHighestSalaryEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM t_employees ORDER BY salary DESC LIMIT 1")
	return &employee, err

}

func main() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/ethan_blog")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	employees, err := queryEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(employees)

	highestSalaryEmployee, err := queryHighestSalaryEmployee(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(highestSalaryEmployee)

}
