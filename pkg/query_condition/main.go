package query_condition

import (
	"reflect"
	"strings"
)

/*
EXAMPLE:

type RequestCondition struct {
	NameEq string `form:"lucky_draw_number_eq"`
	PhoneEq string `form:"phone_eq" sql_column_name:"mobile"`
}

var reqCondition RequestCondition
c.BindQuery(&reqCondition)  // c is *gin.Context
queryCondition := query_condition.QueryCondition{Condition: reqCondition}

queryCondition.ToSQL() => "name=[FOO] AND mobile = [BAR]"
*/
type QueryCondition struct {
	Condition interface{}
}

func (q *QueryCondition) ToSQL() string {

	var builder strings.Builder
	builder.WriteString(" 1=1 ")

	// ignore empty value
	if q.Condition == nil {
		return builder.String()
	}
	switch q.Condition.(type) {
	case map[string]interface{}:
		for key, value := range q.Condition.(map[string]interface{}) {
			node := q.parseSimpleFieldToNode(key, value)
			node.ToSQL(&builder)
		}
	default:
		rType := reflect.TypeOf(q.Condition)
		rVal := reflect.ValueOf(q.Condition)
		for i := 0; i < rType.NumField(); i++ {
			field := rType.Field(i)

			node := q.parseFieldToNode(field, rVal.Field(i).Interface())
			node.ToSQL(&builder)
		}
	}

	return builder.String()
}

//for map[string]interface{}
func (q *QueryCondition) parseSimpleFieldToNode(fieldName string, value interface{}) Node {
	var node = Node{}
	underscoreFieldName := toSnakeCase(fieldName) // UserNameEq => user_name_eq
	lastUnderScoreIndex := strings.LastIndex(underscoreFieldName, "_")
	node.ColumnName = underscoreFieldName[0:lastUnderScoreIndex]
	node.Operator = underscoreFieldName[lastUnderScoreIndex+1:]
	node.Value = value
	return node

}

func (q *QueryCondition) parseFieldToNode(field reflect.StructField, value interface{}) Node {
	var node = Node{}
	underscoreFieldName := toSnakeCase(field.Name) // UserNameEq => user_name_eq
	lastUnderScoreIndex := strings.LastIndex(underscoreFieldName, "_")

	node.ColumnName = field.Tag.Get("sql_column_name")
	if node.ColumnName == "" && lastUnderScoreIndex > 0 {
		//user_name_eq => user_name
		node.ColumnName = underscoreFieldName[0:lastUnderScoreIndex]
	}
	node.Operator = underscoreFieldName[lastUnderScoreIndex+1:]
	node.Value = value
	return node

}
