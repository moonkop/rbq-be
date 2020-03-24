package utils

import (
	"encoding/json"
	"reflect"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

//func Map(arr []interface{},fun func(interface{},int) interface{})[]interface{}  {
//	retArr:=make([]interface{}, len(arr))
//	for index, item := range arr {
//		retArr[index]=fun(item,index)
//	}
//	return retArr
//}

func Map(t interface{}, f func(interface{}) interface{}) []interface{} {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		arr := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			arr[i] = f(s.Index(i).Interface())
		}
		return arr
	}
	return nil
}

func StructToMap(obj interface{}) (newMap map[string]interface{}) {
	data, err := json.Marshal(x) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}
