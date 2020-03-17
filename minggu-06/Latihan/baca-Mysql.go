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
    defer db.Close()
    row, err := db.Query("select * from mahasiswa")
	if err != nil {
		fmt.Println(err.Error())
        return
    }
    defer row.Close()
    for row.Next() {
		var mahasiswa mhs

		err = row.Scan(&mahasiswa.id, &mahasiswa.nama, &mahasiswa.nim, &mahasiswa.jurusan)
		if err != nil {
			fmt.Println(err.Error())
            return
		}
		fmt.Println( mahasiswa.id, mahasiswa.nama, mahasiswa.nim, mahasiswa.jurusan)
	}
}


