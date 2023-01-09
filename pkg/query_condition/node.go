package query_condition

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/guregu/null.v4"
)

type Node struct {
	ColumnName string      // db 的欄位名稱
	Operator   string      // 過濾的操作，目前支援 eq/like/lte/gte
	Value      interface{} // 要過濾的值
}

func (node *Node) ToSQL(builder *strings.Builder) {
	// ignore empty value
	switch node.Operator {
	case "eq":
		switch node.Value.(type) {
		case string:
			if node.Value.(string) == "" {
				return
			}
		case int:
			if node.Value.(int) == 0 {
				return
			}
		case int64:
			if node.Value.(int64) == 0 {
				return
			}
		case null.Time:
			if node.Value.(null.Time).Valid == false {
				return
			}
		case *time.Time:
			if node.Value.(*time.Time) == nil {
				return
			}
		case float32:
			if node.Value.(float32) == 0 {
				return
			}
		case float64:
			if node.Value.(float64) == 0 {
				return
			}
		case *int:
			if node.Value.(*int) == nil {
				return
			}
		case *int64:
			if node.Value.(*int64) == nil {
				return
			}
		}
	case "like":
		switch node.Value.(type) {
		case string:
			if node.Value.(string) == "" {
				return
			}
		}
	case "lte":
		switch node.Value.(type) {
		case time.Time:
			if node.Value.(time.Time).IsZero() {
				return
			}
		}
	case "gte":
		switch node.Value.(type) {
		case time.Time:
			if node.Value.(time.Time).IsZero() {
				return
			}
		}
	}

	builder.WriteString(" AND ")
	if strings.Contains(node.ColumnName, "_or_") {
		// 處理對應多個 column name 的 case
		columnNames := strings.Split(node.ColumnName, "_or_")
		builder.WriteString("( ")
		for i, columnName := range columnNames {
			node.ColumnName = columnName
			node.toSQL(builder)
			if i != len(columnNames)-1 {
				builder.WriteString(" OR ")
			}
		}
		builder.WriteString(" )")
	} else {
		// 處理對應單個 column name 的 case
		node.toSQL(builder)
	}
}

func (node *Node) toSQL(builder *strings.Builder) {
	builder.WriteString(node.ColumnName)
	switch node.Operator {
	case "eq":
		builder.WriteString(" = ")
	case "like":
		builder.WriteString(" like ")
	case "lte":
		builder.WriteString(" <= ")
	case "gte":
		builder.WriteString(" >= ")
	}
	switch node.Value.(type) {
	case string:
		builder.WriteString("'")
		if node.Operator == "like" {
			builder.WriteString("%")
		}
		builder.WriteString(escapeStringBackslash(node.Value.(string)))
		if node.Operator == "like" {
			builder.WriteString("%")
		}
		builder.WriteString("'")
	case int:
		builder.WriteString(fmt.Sprintf("%d", node.Value.(int)))
	case *int:
		builder.WriteString(fmt.Sprintf("%d", *node.Value.(*int)))
	case int64:
		builder.WriteString(fmt.Sprintf("%d", node.Value.(int64)))
	case *int64:
		builder.WriteString(fmt.Sprintf("%d", *node.Value.(*int64)))
	case time.Time:
		builder.WriteString("'" + node.Value.(time.Time).String() + "'")
	}

}
