package main

import (
	"github.com/astaxie/beego"
	_ "sdd1/db"
	_ "sdd1/routers"
)

func main() {
	beego.Run()
}

