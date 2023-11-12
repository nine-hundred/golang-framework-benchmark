package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"strconv"
)

type ItemController struct {
	web.Controller
}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (i *ItemController) Get() {
	var item Item
	err := i.Ctx.Input.Bind(&item.Id, "id")
	if err != nil {
		i.Abort(strconv.Itoa(http.StatusBadRequest))
		return
	}
	item.Name = "hello"
	i.Data["json"] = item
	i.ServeJSON()
}

func (i *ItemController) Post() {
	item := new(Item)
	err := json.Unmarshal(i.Ctx.Input.RequestBody, item)
	if err != nil {
		i.Abort(strconv.Itoa(http.StatusBadRequest))
		return
	}

	i.Data["json"] = item
	i.ServeJSON()
}
