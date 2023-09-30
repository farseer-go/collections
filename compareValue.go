package collections

import (
	"fmt"
	"reflect"
)

// CompareLeftGreaterThanRight 比较两个值，左值是否大于右值
func CompareLeftGreaterThanRight(leftValue any, rightValue any) bool {
	kind := reflect.TypeOf(rightValue).Kind()
	switch kind {
	case reflect.Int8:
		return rightValue.(int8) <= leftValue.(int8)
	case reflect.Int16:
		return rightValue.(int16) <= leftValue.(int16)
	case reflect.Int32:
		return rightValue.(int32) <= leftValue.(int32)
	case reflect.Int64:
		return rightValue.(int64) <= leftValue.(int64)
	case reflect.Int:
		return rightValue.(int) <= leftValue.(int)
	case reflect.Uint:
		return rightValue.(uint) <= leftValue.(uint)
	case reflect.Uint8:
		return rightValue.(uint8) <= leftValue.(uint8)
	case reflect.Uint16:
		return rightValue.(uint16) <= leftValue.(uint16)
	case reflect.Uint32:
		return rightValue.(uint32) <= leftValue.(uint32)
	case reflect.Uint64:
		return rightValue.(uint64) <= leftValue.(uint64)
	case reflect.Float32:
		return rightValue.(float32) <= leftValue.(float32)
	case reflect.Float64:
		return rightValue.(float64) <= leftValue.(float64)
	case reflect.String:
		strRight := rightValue.(string)
		strLeft := leftValue.(string)
		// 空的字符串，认为排在前面，所以返回true
		if strLeft == "" {
			return true
		}

		for i := 0; i < len(strLeft); i++ {
			// 右边字符串比较短，则右边排前面
			if len(strRight) == i {
				return true
			}
			if strLeft[i] == strRight[i] {
				continue
			}

			return strLeft[i] > strRight[i]
		}
		// 说明left长度比right短。则短的排前面
		return false
	}
	panic(fmt.Errorf("该类型无法比较：%s", kind))
}
