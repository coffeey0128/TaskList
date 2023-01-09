package query_condition

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCondition struct {
	UserNameEq                   string `sql_column_name:"name"`
	AgeEq                        int
	AddressLike                  string
	UserNameOrParentUserNameLike string
	AgeOrParentAgeEq             int
}

func TestStructToSQL(t *testing.T) {
	var q = QueryCondition{
		Condition: TestCondition{
			UserNameEq:                   "test",
			AgeEq:                        18,
			AddressLike:                  " QQQ ",
			UserNameOrParentUserNameLike: "test2",
			AgeOrParentAgeEq:             20,
		},
	}
	expected := []string{"1=1", "age = 18", "name = 'test'", "address like '% QQQ %'", "( user_name like '%test2%' OR parent_user_name like '%test2%' )", "( age = 20 OR parent_age = 20 )"}
	actual := strings.Split(q.ToSQL(), "AND")
	// 移除前後空格，避免排序錯誤
	for i, v := range actual {
		actual[i] = strings.TrimSpace(v)
	}
	sort.Strings(expected)
	sort.Strings(actual)
	assert.Equal(t, expected, actual)
}
func TestSimpleMapToSQL(t *testing.T) {
	var q = QueryCondition{
		Condition: map[string]interface{}{
			"AgeEq":                        18,
			"UserNameEq":                   "test",
			"AddressLike":                  "QQQ",
			"UserNameOrParentUserNameLike": "test2",
			"AgeOrParentAgeEq":             20,
		},
	}
	expected := []string{"1=1", "age = 18", "user_name = 'test'", "address like '%QQQ%'", "( user_name like '%test2%' OR parent_user_name like '%test2%' )", "( age = 20 OR parent_age = 20 )"}
	actual := strings.Split(q.ToSQL(), "AND")
	// 移除前後空格，避免排序錯誤
	for i, v := range actual {
		actual[i] = strings.TrimSpace(v)
	}
	sort.Strings(expected)
	sort.Strings(actual)
	assert.Equal(t, expected, actual)
}
