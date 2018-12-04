package pb

import (
	"bytes"
	"fmt"
	"strings"
)

type Model struct {
	Field   string
	Type    string
	Comment string
}

func ToModel(tableName string, list []*Model) string {
	maxNameLen := fieldMaxLen(list)
	maxFieldLen := fieldTypeMaxLen(list)
	maxTagLen := fieldTagMaxLen(list)

	var buf bytes.Buffer
	for _, v := range list {

		jsonName := field2JSON(v.Field)
		name := head2Upper(jsonName)
		fieldName := field2Type(v.Type)
		tag := fieldTag(fieldJSONTag(jsonName), fieldGormTag(v.Field))

		buf.WriteString(space(5))
		//field
		buf.WriteString(name)
		buf.WriteString(space(maxNameLen - len(name) + 1))
		//type
		buf.WriteString(fieldName)
		buf.WriteString(space(maxFieldLen - len(fieldName) + 1))
		//tag
		buf.WriteString(tag)
		buf.WriteString(space(maxTagLen - len(tag) + 1))
		//comment
		buf.WriteString(field2Comment(v.Comment))
		buf.WriteString("\n")
	}

	str := fmt.Sprintf("type %s struct {\n%s}", head2Upper(field2JSON(tableName)), buf.String())
	return str
}

func field2Comment(comment string) string {
	if comment == "" {
		return ""
	}
	return fmt.Sprintf("//%s", comment)
}

func fieldTag(tag ...string) string {
	var str string
	for _, v := range tag {
		str += v + " "
	}
	return fmt.Sprintf("`%s`", str)
}

func fieldGormTag(name string) string {
	return fmt.Sprintf(`gorm:"column:%s"`, name)
}

func fieldJSONTag(jsonName string) string {
	return fmt.Sprintf(`json:"%s"`, jsonName)
}

func field2Type(tmp string) string {

	if strings.Contains(tmp, "int") {
		return "int64"
	}

	if strings.Contains(tmp, "varchar") {
		return "string"
	}

	if strings.Contains(tmp, "decimal") {
		return "float64"
	}

	if strings.Contains(tmp, "timestamp") || strings.Contains(tmp, "date") {
		return "*time.Time"
	}

	return "interface{}"
}

func fieldMaxLen(list []*Model) int {
	var max, size int
	for _, v := range list {
		size = len(v.Field)
		if len(v.Field) > max {
			max = size
		}
	}
	return max
}

func fieldTagMaxLen(list []*Model) int {
	var max, size int
	for _, v := range list {
		size = len(fieldTag(fieldJSONTag(field2JSON(v.Field)), fieldGormTag(v.Field)))
		if size > max {
			max = size
		}
	}
	return max
}

func fieldTypeMaxLen(list []*Model) int {
	var max int
	for _, v := range list {
		if len(field2Type(v.Type)) > max {
			max = len(v.Field)
		}
	}
	return max
}

func space(num int) string {
	tmp := make([]rune, num)
	for i, _ := range tmp {
		tmp[i] = 32
	}
	return string(tmp)
}

func field2JSON(field string) string {

	if !strings.Contains(field, "_") {
		return field
	}

	var str string
	for i, tmp := range strings.Split(field, "_") {
		if i == 0 {
			str += tmp
		} else {
			str += head2Upper(tmp)
		}
	}

	return str
}

func head2Upper(str string) string {
	tmp := []rune(str)
	return strings.ToUpper(string(tmp[0])) + string(tmp[1:])
}
