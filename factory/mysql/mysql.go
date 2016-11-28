package gom

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gom"
)

func init()  {
	gom.Register("mysql",&MySqlFactory{})
}

type MySqlFactory struct {

}

func Insert(model gom.TableModel) (string,[]interface{}) {
	var datas []interface{}
	sql:="insert into \\`"+model.TableName+"\\` ("
	values:=""
	for i,v:=range model.Columns{
		value:=model.ModelValue.FieldByName(v.ColumnName).Interface()
		if value !=nil{
			if i>0{
				sql+=","
				values+=","
			}
			append(datas,value)
			values+=" ? "
			sql+=v.ColumnName
		}

	}
	sql+=") VALUES ("+values+")"
	return sql,datas
}
func Delete(model gom.TableModel) (string,[]interface{}) {
	sql:="delete from \\`"+model.TableName+"\\` where "
	if model.Cnd != nil{
		sql+=model.Cnd.State()+";"
		return sql,model.Cnd.Value()
	}else{
		sql+=model.GetPrimaryCondition().State()+";"
		return sql,model.GetPrimaryCondition().Value()
	}

}
func Update(model gom.TableModel) (string,[]interface{}) {


}
func Query(model gom.TableModel) (string,[]interface{}) {

}
