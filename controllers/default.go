package controllers

import (
	"io/ioutil"
	"log"
	"fmt"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {

	fmt.Println("hello")
	allCities, err := ioutil.ReadFile("./cities.csv")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("CIties: ", string(allCities))
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"


}
