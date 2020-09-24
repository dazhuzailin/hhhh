package controllers

import (
	"encoding/json"
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

func (c *MainController) Post(){
	var mmm models.Ppsot
	databytes ,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("发生错误")
		return
	}
	err = json.Unmarshal(databytes,&mmm)
	if err != nil {
		c.Ctx.WriteString("发生错误")
		return
	}
	c.Ctx.WriteString("成功")
}
