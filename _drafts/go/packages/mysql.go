package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type User struct {
	id int64
	name string
	age int
	gender int
	intro sql.NullString
	createdAt sql.NullTime
	updatedAt sql.NullTime
}

func main() {
	tdb, err := sql.Open("mysql", "root:123456@/go_test")
	if err != nil {
		panic(err.Error())
	}
	db = tdb
	defer db.Close()

	createTable()
	insert()
	queryRow()
	update()
	queryRow()
	begin()
	queryRow()
	delete()
	prepare()
	query()
	truncate()
}

func createTable () {
	q := `
	CREATE TABLE IF NOT EXISTS users (
	  id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	  name varchar(255) NOT NULL DEFAULT '', 
	  age int(10) NOT NULL DEFAULT '0',
	  gender int(10) NOT NULL DEFAULT '0',
	  intro TEXT DEFAULT NULL,
	  created_at TIMESTAMP NULL DEFAULT NULL,
	  updated_at TIMESTAMP NULL DEFAULT NULL,
	  PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`

	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table.")
}

func insert() {
	q := "INSERT INTO users (name, age, intro)VALUES(?,?,?)"
	ret, err := db.Exec(q, "张三", 23, "insert")
	if err != nil {
		log.Fatal(err)
	}
	lastInsertId, err := ret.LastInsertId()
	fmt.Printf("insert LastInsertId: %d\n", lastInsertId)
}

func queryRow() {
	var user User
	q := "SELECT id,name,age,gender,intro FROM users WHERE id=?"
	id := 1
	err := db.QueryRow(q, id).Scan(&user.id, &user.name, &user.age, &user.gender,&user.intro)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("queryRow")
	fmt.Println(user)
}


func update() {
	q := "UPDATE users SET age=?, intro=? WHERE id=?"
	age := 24
	intro := "update"
	id := 1
	ret, err := db.Exec(q, age, intro, id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := ret.RowsAffected()
	fmt.Printf("update RowsAffected: %d\n", rowsAffected)
}

func query() {
	age := 20
	q := "SELECT id,name,age,gender,intro,created_at,updated_at FROM users WHERE age>?"
	rows, err := db.Query(q, age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.id, &user.name, &user.age, &user.gender, &user.intro, &user.createdAt, &user.updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	fmt.Println("query")
	fmt.Println(users)
}


func begin()  {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	id := 1
	age := 25
	intro := "begin"
	q := `UPDATE users SET age=?, intro=? WHERE id=?`
	ret, execErr := tx.Exec(q, age, intro, id)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal(rollbackErr)
		}
		log.Fatal(execErr)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := ret.RowsAffected()
	fmt.Printf("begin RowsAffected: %d\n", rowsAffected)
}

func delete() {
	q := "DELETE FROM users WHERE id=?"
	ret, err := db.Exec(q, "1")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := ret.RowsAffected()
	fmt.Printf("delete RowsAffected: %d\n", rowsAffected)
}

func prepare()  {
	users := []struct {
		name string
		age int
	}{
		{"李四", 24},
		{"王五", 25},
	}

	q := "INSERT INTO users(name, age) VALUES(?, ?)"
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for _, user := range users {
		if _, err := stmt.Exec(user.name, user.age); err != nil {
			log.Fatal(err)
		}
	}
}

func truncate()  {
	q := "TRUNCATE TABLE users"
	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("truncate")
}

/*

Package sql: https://golang.google.cn/pkg/database/sql/

Go-MySQL-Driver: https://github.com/go-sql-driver/mysql


Output:

create table.
insert LastInsertId: 1
queryRow
{1 张三 23 0 {insert true} {{0 0 <nil>} false} {{0 0 <nil>} false}}
update RowsAffected: 1
queryRow
{1 张三 24 0 {update true} {{0 0 <nil>} false} {{0 0 <nil>} false}}
begin RowsAffected: 1
queryRow
{1 张三 25 0 {begin true} {{0 0 <nil>} false} {{0 0 <nil>} false}}
delete RowsAffected: 1
query
[{2 李四 24 0 { false} {{0 0 <nil>} false} {{0 0 <nil>} false}} {3 王五 25 0 { false} {{0 0 <nil>} false} {{0 0 <nil>} false}}]
truncate
Process finished with exit code 0
*/
