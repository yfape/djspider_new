/*
 * @Author: your name
 * @Date: 2021-07-12 15:31:52
 * @LastEditTime: 2021-07-29 10:48:07
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\site.go
 */
package db

import (
	_ "github.com/go-sql-driver/mysql"
)

type Spi struct {
	SpiId      int    `db:"spi_id"`
	Col0Id     int    `db:"col0_id"`
	Col1Id     int    `db:"col1_id"`
	Col2Id     int    `db:"col2_id"`
	Col3Id     int    `db:"col3_id"`
	Url        string `db:"url"`
	Strategy   string `db:"strategy"`
	Mode       int    `db:"mode"`
	CreateTime int    `db:"create_time"`
	Col0Name   string `db:"col0_name"`
}

type SpiColname struct {
	ColId int    `db:"col_id"`
	Name  string `db:"name"`
}

func GetSpis() []Spi {
	var spis []Spi
	Db := MysqlConnect()
	err := Db.Select(&spis, "SELECT a.*,b.name as col0_name FROM dj_spi a join dj_spi_colname b on a.col0_id=b.col_id")
	if err != nil {
		panic(err)
	}
	defer MysqlClose(Db)
	return spis
}

func GetSpiByUrl(url string) []Spi {
	var spi []Spi
	Db := MysqlConnect()
	err := Db.Select(&spi, "select a.*,b.name as col0_name FROM dj_spi a join dj_spi_colname b on a.col0_id=b.col_id where a.url=?", url)
	if err != nil {
		panic(err)
	}
	defer MysqlClose(Db)
	return spi
}
