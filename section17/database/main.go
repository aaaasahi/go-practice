package main

import (
	"database/sql"
	"log"

	// コード内で使わない場合 _ を使う
	_ "github.com/mattn/go-sqlite3"
)

//コネクションプール グローバルに宣言
var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

// database + sqlite3
// sqlite3 example.sql でDBの中に入れる
func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// テーブル追加
	cmd := `CREATE TABLE IF NOT EXISTS persons(
                name STRING,
				age  INT)`
	_, err := Db.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// データの追加
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err := Db.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}

	// データの更新
	cmd := "UPDATE persons SET age = ? WHERE name = ?"
	_, err := Db.Exec(cmd, 25, "Mike")
	if err != nil {
		log.Fatalln(err)
	}

	// 特定のデータを取得
	cmd := "SELECT * FROM persons where age = ?"
	// QueryRow 1レコード取得
	row := Db.QueryRow(cmd, 25)
	var p Person
	// 取得データをPersonに入れる
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// 複数データの取得
	cmd := "SELECT * FROM persons"
	// Query 条件に合うものを全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	// スライスで宣言
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		// スライスに追加
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// データ削除
	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := Db.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}
