package db

import (
	"BeegoProject/BeegoProject0603/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init(){
	fmt.Println("链接数据库")
	con := beego.AppConfig
	drivername := con.String("drivername")
	dd := con.String("dd")
	db ,err := sql.Open(drivername,dd)
	if err != nil{
		panic("数据库链接错误")
	}
	Db = db
}
func InsertUser(user models.User)(int64,error){
	//1、将用户密码进行hash脱敏，使用md5计算密码hash值，并存储hash值
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	fmt.Println("用户名:",user.Nick,"密码:",user.Password)
	result,err := Db.Exec("insert into user(nick, password) values(?,?)",user.Nick,user.Password)
	if err != nil {
		return -1,err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}