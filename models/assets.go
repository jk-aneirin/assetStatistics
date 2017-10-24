package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Assetinfo struct {
	Name	string `orm:"pk"`
	Apt		string 
	Sertag	string 
	Macaddr string 
	Mtorsn	string
}

func GetAllAssets() []*Assetinfo {
	o := orm.NewOrm()
	o.Using("default")
	var assets []*Assetinfo
	q := o.QueryTable("assetinfo")
	q.All(&assets)
	return assets
}

func GetOneByName(name string) Assetinfo {
	o := orm.NewOrm()
	o.Using("default")
	q := o.QueryTable("assetinfo")

	a := Assetinfo{Name:name}

	err := q.Filter("name__istartswith",name).One(&a)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("缺少主键")
	}
	return a
	
}

func GetOneByMac(mac string) Assetinfo {
	a := Assetinfo{Macaddr:mac}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&a,"Macaddr")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	}
	return a
}

func AddAsset(asset *Assetinfo) string {
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(asset)
	return asset.Name
}

func DelAsset(name string) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Assetinfo{Name:name})
}

func UpdateAsset(asset *Assetinfo) {
	o := orm.NewOrm()
	o.Using("default")
	o.Update(asset)
}

func init() {
	orm.RegisterModel(new(Assetinfo))
}
