package ToolsDBHelper

import (
	"database/sql"
	"errors"
	"reflect"
	"strconv"
)

func DoQuery(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, _ := rows.Columns()
	columnstype, _ := rows.ColumnTypes()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	for index, _ := range cache {              //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list = make([]map[string]interface{}, 0) //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			switch columnstype[i].DatabaseTypeName() {
			case "CHAR", "VARCHAR":
				{
					item[columns[i]], _ = (*data.(*interface{})).(string)
				}
			default:
				{
					item[columns[i]] = *data.(*interface{}) //取实际类型
				}
			}
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list, nil
}

func DoProcessRow(row *sql.Rows, Afunc func(data map[string]interface{}, Atype []*sql.ColumnType) bool) error {
	columns, _ := row.Columns()
	columnstype, _ := row.ColumnTypes()
	columnLength := len(columns)
	rowdatas := make([]interface{}, columnLength)
	for index, _ := range rowdatas { //为每一列初始化一个指针
		var a interface{}
		rowdatas[index] = &a
	}
	item := make(map[string]interface{})
	err := row.Scan(rowdatas...)
	if err != nil {
		return err
	}
	for i, data := range rowdatas {
		item[columns[i]] = *data.(*interface{}) //取实际类型
	}
	Afunc(item, columnstype)
	return nil
}

func DoProcessRows(rows *sql.Rows, Afunc func(data map[string]interface{}, Atype []*sql.ColumnType) bool) error {
	columns, _ := rows.Columns()
	columnstype, _ := rows.ColumnTypes()
	columnLength := len(columns)
	rowdatas := make([]interface{}, columnLength)
	for index, _ := range rowdatas { //为每一列初始化一个指针
		var a interface{}
		rowdatas[index] = &a
	}
	for rows.Next() {
		err := rows.Scan(rowdatas...)
		if err != nil {
			return err
		}
		item := make(map[string]interface{})
		for i, data := range rowdatas {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		if !Afunc(item, columnstype) {
			break
		}
	}
	return nil
}

func mapping(AfieldName string, Avalue string, v reflect.Value) error {
	t := v.Type()
	val := v.Elem()
	typ := t.Elem()

	if !val.IsValid() {
		return errors.New("数据类型不正确")
	}

	for i := 0; i < val.NumField(); i++ {
		value := val.Field(i)
		kind := value.Kind()
		tag := typ.Field(i).Tag.Get("col")

		if (len(tag) > 0) && (tag == AfieldName) {
			meta := Avalue
			if !value.CanSet() {
				return errors.New("结构体字段没有读写权限")
			}

			if len(meta) == 0 {
				continue
			}

			if kind == reflect.String {
				value.SetString(meta)
			} else if kind == reflect.Float32 {
				f, err := strconv.ParseFloat(meta, 32)
				if err != nil {
					return err
				}
				value.SetFloat(f)
			} else if kind == reflect.Float64 {
				f, err := strconv.ParseFloat(meta, 64)
				if err != nil {
					return err
				}
				value.SetFloat(f)
			} else if kind == reflect.Int64 {
				integer64, err := strconv.ParseInt(meta, 10, 64)
				if err != nil {
					return err
				}
				value.SetInt(integer64)
			} else if kind == reflect.Int {
				integer, err := strconv.Atoi(meta)
				if err != nil {
					return err
				}
				value.SetInt(int64(integer))
			} else if kind == reflect.Bool {
				b, err := strconv.ParseBool(meta)
				if err != nil {
					return err
				}
				value.SetBool(b)
			} else {
				return errors.New("数据库映射存在不识别的数据类型")
			}
			break
		}
	}
	return nil
}

func DoRowToStruct(rows *sql.Rows, objaddr interface{}) error {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	values := make([][]byte, columnLength)     //values是每个列的值，这里获取到byte里
	for index, _ := range cache {              //为每一列初始化一个指针
		cache[index] = &values[index]
	}

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return err
		}
		return sql.ErrNoRows
	}
	err := rows.Scan(cache...)
	if err != nil {
		return err
	}
	k := reflect.ValueOf(objaddr).Elem()
	newObj := reflect.New(k.Type())

	for i, data := range values {
		err = mapping(columns[i], string(data), newObj)
		if err != nil {
			return err
		}
	}
	k.Set(newObj.Elem())
	return nil
}

func DoRowsToStruct(rows *sql.Rows, objaddr interface{}) error {
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	values := make([][]byte, columnLength)     //values是每个列的值，这里获取到byte里
	for index, _ := range cache {              //为每一列初始化一个指针
		cache[index] = &values[index]
	}

	v := reflect.ValueOf(objaddr).Elem()
	newv := reflect.MakeSlice(v.Type(), 0, 0)

	for rows.Next() {
		err := rows.Scan(cache...)
		if err != nil {
			return err
		}
		k := v.Type().Elem()
		newObj := reflect.New(k)

		for i, data := range values {
			err = mapping(columns[i], string(data), newObj)
			if err != nil {
				return err
			}
		}
		newv = reflect.Append(newv, newObj.Elem())
	}
	v.Set(newv)
	return nil
	//_ = rows.Close()   //有可能有多个数据集，由外面关闭
}
