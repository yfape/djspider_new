/*
 * @Author: your name
 * @Date: 2021-07-29 10:41:53
 * @LastEditTime: 2021-07-29 17:12:24
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\http\col.go
 */
package http

import (
	"djspider/db"
	"encoding/json"
	"log"

	"djspider/myredis"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

var Spis []db.Spi
var Spics []db.SpiColname

type Col struct {
	Spi_id   int
	Name     string
	Children []Col
}

func GetCols(c *gin.Context) {
	rd = myredis.Connect()
	defer (*rd).Close()
	res, err := redis.Bytes((*rd).Do("Get", "cols"))
	if err != nil {
		log.Println(err)
	} else {
		c.String(201, string(res))
		return
	}

	Db = db.MysqlConnect()
	defer Db.Close()
	err = Db.Select(&Spis, "select * from dj_spi order by col0_id asc,col1_id asc,col2_id asc,col3_id asc")
	if err != nil {
		log.Println(err)
		c.JSON(520, gin.H{"msg": "内部错误"})
		return
	}
	err = Db.Select(&Spics, "select * from dj_spi_colname")
	if err != nil {
		log.Println(err)
		c.JSON(520, gin.H{"msg": "内部错误"})
		return
	}
	var sites []Col
	for i := 0; i < len(Spis); i++ {
		if len(sites) <= 0 || Spis[i].Col0Id != sites[len(sites)-1].Spi_id {
			col1 := Col{
				Spi_id:   Spis[i].Col0Id,
				Name:     GetName(Spis[i].Col0Id),
				Children: []Col{},
			}
			if Spis[i].Col1Id != 0 {
				for i1 := 0; i1 < len(Spis); i1++ {
					if Spis[i1].Col0Id == Spis[i].Col0Id && (len(col1.Children) == 0 || Spis[i1].Col1Id != col1.Children[len(col1.Children)-1].Spi_id) {
						col2 := Col{
							Spi_id:   Spis[i1].Col1Id,
							Name:     GetName(Spis[i1].Col1Id),
							Children: []Col{},
						}
						if Spis[i1].Col2Id != 0 {
							for i2 := 0; i2 < len(Spis); i2++ {
								if Spis[i2].Col1Id == Spis[i1].Col1Id && (len(col2.Children) == 0 || Spis[i2].Col2Id != col2.Children[len(col2.Children)-1].Spi_id) {
									col3 := Col{
										Spi_id:   Spis[i2].Col2Id,
										Name:     GetName(Spis[i2].Col2Id),
										Children: []Col{},
									}
									if Spis[i2].Col3Id != 0 {
										for i3 := 0; i3 < len(Spis); i3++ {
											if Spis[i3].Col2Id == Spis[i2].Col2Id && (len(col3.Children) == 0 || Spis[i3].Col3Id != col3.Children[len(col3.Children)-1].Spi_id) {
												col4 := Col{
													Spi_id:   Spis[i3].Col3Id,
													Name:     GetName(Spis[i3].Col3Id),
													Children: []Col{},
												}
												col3.Children = append(col3.Children, col4)
											}
										}
									}
									col2.Children = append(col2.Children, col3)
								}
							}
						}
						col1.Children = append(col1.Children, col2)
					}
				}
			}
			sites = append(sites, col1)
		}
	}

	resout := gin.H{
		"msg":  "获取成功",
		"data": sites,
	}
	sitesredis, _ := json.Marshal(resout)
	_, err = (*rd).Do("Set", "cols", sitesredis)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, gin.H{"data": sites})
}

func GetName(ColId int) string {
	for i := 0; i < len(Spics); i++ {
		if ColId == Spics[i].ColId {
			return Spics[i].Name
		}
	}
	return ""
}
