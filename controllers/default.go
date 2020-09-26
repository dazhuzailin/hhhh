package controllers

import (
	"BeegoProject/BeegoProject0603/db_mysql"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"sdd1/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	var mmm models.Ppsot
	databytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("发生错误")
		return
	}
	err = json.Unmarshal(databytes, &mmm)
	if err != nil {
		c.Ctx.WriteString("发生错误")
		return
	}
	c.Ctx.WriteString("成功")

}

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post(){
	fmt.Println(r == nil)
	fmt.Println(r.Ctx == nil)
	fmt.Println(r.Ctx.Request == nil)
	bodyBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误,请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(bodyBytes,&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		return
	}

	id, err := db_mysql.InsertUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败.")
		return
	}

	fmt.Println(id)

	result := models.ResponseResult{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	r.Data["json"] = &result
	r.ServeJSON()
}
