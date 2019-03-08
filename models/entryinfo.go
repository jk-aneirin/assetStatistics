package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Entryinfo struct {
	Id		int64
	Name    string `orm:"size(20)"`
	Apt     string `orm:"size(15)"`
	Sertag  string `orm:"size(10)"`
	Macaddr string `orm:"size(20)"`
	Mtorsn  string `orm:"size(25)"`
}

func init() {
	orm.RegisterDataBase("default","mysql","dbuser:password@tcp(localhost:3306)/asset?charset=utf8",30)
	orm.RegisterModel(new(Entryinfo))
	//create table
	orm.RunSyncdb("default",false,true)
}

func getOb(db string) orm.Ormer {
	ob := orm.NewOrm()
	ob.Using(db)
	return ob
}

func Addentry(entry *Entryinfo) int64 {
	id,err := getOb("default").Insert(entry)

	if err != nil {
		fmt.Println(err)
	}

	return id
}

func GetAllEntrys() []*Entryinfo {
	var entrys []*Entryinfo
	
	q := getOb("default").QueryTable("entryinfo")
	q.All(&entrys)
	return entrys
}

func GetEntryByName(name string) Entryinfo {
	q := getOb("default").QueryTable("entryinfo")
	entry := Entryinfo{Name: name}

	err := q.Filter("name__istartswith", name).One(&entry)
	if err == orm.ErrNoRows {
		return Entryinfo{}
	}
								
	return entry
}

func GetEntryByMac(mac string) Entryinfo {
	entry := Entryinfo{Macaddr: mac}
	err := getOb("default").Read(&entry,"Macaddr")
	if err == orm.ErrNoRows {
		return Entryinfo{}
	}

	return entry
}

func DelEntry(id int64) error {
	_,err := getOb("default").Delete(&Entryinfo{Id: id})
	
	return err
}
