package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"

	_ "routers"
)

func main() {
	beego.Run()
}
