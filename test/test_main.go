package main

import (
	"fmt"
	_ "github.com/janyees/gom/factory/mysql"
	"github.com/janyees/gom"
	"time"
)
type Log struct {
	Id string `json:"id" gom:"primary,id"`
	Level int `gom:"ignore"`
	Info string
	Date time.Time `gom:"column,date"`
}
func (Log) TableName() string {
	return "system_log"
}

func main() {
	var logs Log
	dsn:=`root:Nuaxxxxxzy@tcp(1xx.25.2xx.xxx:3306)/moren`
	db,err:=gom.Open("mysql",dsn)
	if err!=nil{
		fmt.Println(err)
	}
	db.Query(&logs,gom.Cnds("id=?","0d9c1726873f4bc3b6fb955877e5a082"))
	idelte,ed:=db.Delete(logs)
	fmt.Println(idelte,ed)
	logs.Date=time.Now()
	ii,ie:=db.Insert(logs)
	fmt.Println(ii,ie)

}