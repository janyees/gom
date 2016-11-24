package gom

import (
	"reflect"
)

type SqlFactory interface {
	Insert(TableModel) (string,[]interface{})
	Delete(TableModel,Condition) (string,[]interface{})
	Update(TableModel,Condition) (string,[]interface{})
	Query(TableModel,Condition) (string,[]interface{})
}
type TableModel struct {
	ModelType reflect.Type
	ModelValue reflect.Value
	TableName string
	Columns []Column
	Primary Column
}
type Column struct {
	ColumnType reflect.Type
	ColumnName string
	FieldName string
	Auto bool
}
type Condition interface {
	State() string
	Value() []interface{}
}
type Conditions struct {
	States string
	Values []interface{}
}

func (c Conditions) State() string {
	return c.States
}
func (c Conditions) Value() [] interface{} {
	return c.Values
}
func (mo TableModel) insertValues() []interface{} {
	var interfaces []interface{}
	results := reflect.Indirect(reflect.ValueOf(&interfaces))
	for _,column:=range mo.Columns{
		vars:=reflect.ValueOf(mo.ModelValue.FieldByName(column.FieldName).Interface())
		if(results.Kind()==reflect.Ptr){
			results.Set(reflect.Append(results,vars.Addr()))
		}else{
			results.Set(reflect.Append(results,vars))
		}
	}
	return interfaces
}
func (m TableModel) getPrimary() interface{}  {
	return m.ModelValue.FieldByName(m.Primary.FieldName).Interface()
}
func (m TableModel) getPrimaryCondition() Condition {
	return Conditions{"where "+m.Primary.ColumnName+" = ?",m.getPrimary()}
}