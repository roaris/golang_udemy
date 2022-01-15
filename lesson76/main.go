package main

import (
	"database/sql"
	"fmt"
	"log"

	// https://daido.hatenablog.jp/entry/2020/09/17/154228
	// https://qiita.com/ryysud/items/1cf66ee4363aec22394a
	_ "github.com/mattn/go-sqlite3"
)

// database + sqlite3
// テーブル作成
var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")

	defer Db.Close()

	/*
		cmd := `CREATE TABLE IF NOT EXISTS persons(name STRING, age INT)`
		_, err := Db.Exec(cmd)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		cmd := `INSERT INTO persons(name, age) VALUES(?, ?)`
		_, err := Db.Exec(cmd, "taro", 20)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		cmd := "UPDATE persons SET age = ? WHERE name = ?"
		_, err := Db.Exec(cmd, "taro", 25)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	/*
		cmd := `INSERT INTO persons(name, age) VALUES(?, ?)`
		_, err := Db.Exec(cmd, "hanako", 10)
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// 特定のデータを取得
	cmd := "SELECT * FROM persons where age = ?"
	// QueryRow 1レコード取得
	row := Db.QueryRow(cmd, 25)
	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// 複数データの取得
	cmd = "SELECT * FROM persons"
	// Queryは条件に合うものを全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// データの削除
	cmd = "DELETE FROM persons WHERE name = ?"
	_, err = Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}

}
