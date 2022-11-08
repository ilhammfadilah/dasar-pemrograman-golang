package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "qlhem:qlhem@tcp(127.0.0.1)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 22
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result []student
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}

}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	defer db.Close()

	var result = student{}
	var id = "B001"
	err = db.QueryRow("select name, grade from tb_student where id = ?", id).
		Scan(&result.name, &result.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name : %s\ngrade : %d\n", result.name, result.grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	var result2 = student{}
	var result3 = student{}
	stmt.QueryRow("B001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name : %s \ngrade : %d\n", result1.name, result1.grade)
	stmt.QueryRow("B002").Scan(&result2.name, &result2.grade)
	fmt.Printf("name : %s \ngrade : %d\n", result2.name, result2.grade)
	stmt.QueryRow("B003").Scan(&result3.name, &result3.grade)
	fmt.Printf("name : %s \ngrade : %d\n", result3.name, result3.grade)
}

func main() {
	// sqlQuery()
	// sqlQueryRow()
	sqlPrepare()
}
