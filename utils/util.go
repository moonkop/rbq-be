package utils

import "reflect"

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
