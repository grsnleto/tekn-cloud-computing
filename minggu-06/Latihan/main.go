package main

import (
        "database/sql"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
        )

type mhs struct {
    id    int
    nama  string
    nim   string
    jurusan string
}
func main() {
    db, err := sql.Open("mysql", "admin:root@tcp(127.0.0.1:3306)/tcc_mhs")
    if err != nil {
        fmt.Println(err.Error())
        return
	}
	fmt.Println("Koneksi Ke",db,"Berhasil!!!")
}