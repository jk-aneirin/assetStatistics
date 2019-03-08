package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	"models"
)

type Assetcontroller struct {
	beego.Controller
}

type Result struct {
	Status string
}

func (a *Assetcontroller) Post() {
	var entry models.Entryinfo
	if err := json.Unmarshal(a.Ctx.Input.RequestBody,&entry);err != nil {
		fmt.Println(err)
	}

	id := models.Addentry(&entry)
	
	a.Data["json"] = id
	a.ServeJSON()
}

func (a *Assetcontroller) DisAll() {
	entrys := models.GetAllEntrys()
	
	a.Data["json"] = entrys
	a.ServeJSON()
}

func (a *Assetcontroller) GetByAny() {
	var rep interface{}
	var entry models.Entryinfo
	
	key := a.GetString(":key")

	if IsMac(key) {
		entry = models.GetEntryByMac(key)
		if len(entry.Sertag) == 0 {
			rep = Result{
				Status:"no record",
			}
		} else {
			rep = entry
		}
	} else if IsIp(key) {
		mac := GetMac(key)
		entry = models.GetEntryByMac(mac)
		if len(entry.Sertag) == 0 {
			rep = Result{
				Status: mac,
			}
		} else {
			rep = entry
		}
	} else {
		entry = models.GetEntryByName(key)
		if len(entry.Sertag) == 0 {
			rep = Result{
				Status:"no record",
			}
		} else {
			rep = entry
		}
	}
	
	a.Data["json"] = rep
	a.ServeJSON()
}

func (a *Assetcontroller) DeletebyId() {
	var rep Result

	id,err := a.GetInt(":id")
	if err != nil {
		fmt.Println(err)
	}

	err = models.DelEntry(int64(id))
	if err == nil {
		rep = Result{
			Status: "success",
		}
	} else {
			rep = Result{
				Status: "fail",
			}
	}
	
	a.Data["json"] = rep
	a.ServeJSON()
}
