/*
 * @Author: your name
 * @Date: 2021-07-15 10:18:33
 * @LastEditTime: 2021-08-02 10:40:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \djspider\app\db\user.go
 */
package db

type User struct {
	UserId     int    `db:"user_id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Pass       string `db:"pass"`
	Headimg    string `db:"headimg"`
	Identify   int    `db:"identify"`
	Spi_ids    string `db:"spi_ids"`
	Col_ids    string `db:"col_ids"`
	Keywords   string `db:"keywords"`
	CreateTime int    `db:"create_time"`
}

func GetUsers() []User {
	var users []User
	Db := MysqlConnect()
	Db.Select(&users, "select * from dj_user order by user_id asc")
	defer MysqlClose(Db)
	return users
}

func (user *User) IdentifyName() string {
	is := []string{"未知", "用户", "管理员"}
	return is[user.Identify]
}
