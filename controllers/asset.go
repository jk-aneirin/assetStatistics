package controllers

import (
	"github.com/astaxie/beego"
	"assetStatistics/models"
	"encoding/json"
	"regexp"
)
		    
// AssetController definition.
type AssetController struct {
	beego.Controller
}
					    
// Get method.
func (c *AssetController) GetAll() {
	ss := models.GetAllAssets()
	c.Data["json"] = ss
	c.ServeJSON()
}

func (c *AssetController) GetByNameOrMacOrIp() {
	var a models.Assetinfo
	k := c.GetString(":key")

	if IsMacAddr(k) {
		if k[2] == ':' {
			re ,_ := regexp.Compile(":")
			k = re.ReplaceAllString(k,"-")
		}
		a = models.GetOneByMac(k)
	} else if IsIpAddr(k) {
		macaddr := GetMacFromIp(k)
		a = models.GetOneByMac(macaddr)
	} else {
		a = models.GetOneByName(k)
	}

	c.Data["json"] = a 
	c.ServeJSON()
}

func (c *AssetController) Post() {
	var a models.Assetinfo
	json.Unmarshal(c.Ctx.Input.RequestBody, &a)
    if !IsMacAddr(a.Macaddr) {
		c.Ctx.WriteString("MAC地址写错了，标准格式为aa-bb-cc-dd-ee-dd")	
		return 
	}

	if a.Macaddr[2] == ':' {
		re ,_ := regexp.Compile(":")
		a.Macaddr = re.ReplaceAllString(a.Macaddr,"-")
	}
	
	name := models.AddAsset(&a)
	c.Data["json"] = name
	c.ServeJSON()
}

func (c *AssetController) Delete() {
	name := c.GetString(":name")
	models.DelAsset(name)
	c.Data["json"] = true
	c.ServeJSON()
}

func (c *AssetController) Update() {
	var a models.Assetinfo
	json.Unmarshal(c.Ctx.Input.RequestBody,&a)
    if !IsMacAddr(a.Macaddr) {
		c.Ctx.WriteString("MAC地址写错了，标准格式为aa-bb-cc-dd-ee-dd")	
		return 
	}

	if a.Macaddr[2] == ':' {
		re ,_ := regexp.Compile(":")
		a.Macaddr = re.ReplaceAllString(a.Macaddr,"-")
	}

	models.UpdateAsset(&a)
	c.Data["json"] = a
	c.ServeJSON()
}

