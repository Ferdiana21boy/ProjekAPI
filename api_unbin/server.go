package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type pembayaran struct {
	ID      uint   `json:"id"`
	Nama    string `json:"nama"`
	Tanggal string `json:"tanggal"`
	Nominal string `json:"nominal"`
}

func main() {

	// database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbs_api")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	// database connection

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Service API!")
	})

	// Mahasiswa
	e.GET("/tbl_pembayaran", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM tbl_pembayaran")

		defer res.Close()

		if err != nil {
			log.Fatal(err)
		}
		var tbl_pembayaran []pembayaran
		for res.Next() {
			var m pembayaran
			_ = res.Scan(&m.ID, &m.Nama, &m.Tanggal, &m.Nominal)
			tbl_pembayaran = append(tbl_pembayaran, m)
		}

		return c.JSON(http.StatusOK, tbl_pembayaran)
	})

	e.POST("/tbl_pembayaran", func(c echo.Context) error {
		var tbl_pembayaran pembayaran
		c.Bind(&tbl_pembayaran)

		sqlStatement := "INSERT INTO tbl_pembayaran (id,nama, tanggal,nominal)VALUES (?,?, ?, ?)"
		res, err := db.Query(sqlStatement, tbl_pembayaran.ID, tbl_pembayaran.Nama, tbl_pembayaran.Tanggal, tbl_pembayaran.Nominal)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, tbl_pembayaran)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.PUT("/mahasiswa/:npm", func(c echo.Context) error {
		var mahasiswa pembayaran
		c.Bind(&mahasiswa)

		sqlStatement := "UPDATE tbl_pembayaran SET nama=?,nama=?,tanggal=? WHERE id=?"
		res, err := db.Query(sqlStatement, mahasiswa.Nama, mahasiswa.Tanggal, mahasiswa.Nominal, c.Param("id"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})

	e.DELETE("/tbl_pembayaran/:npm", func(c echo.Context) error {
		var mahasiswa pembayaran
		c.Bind(&mahasiswa)

		sqlStatement := "DELETE FROM mahasiswa WHERE npm=?"
		res, err := db.Query(sqlStatement, c.Param("npm"))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, mahasiswa)
		}
		return c.String(http.StatusOK, "ok")
	})
	// Mahasiswa

	e.Logger.Fatal(e.Start(":1323"))
}
