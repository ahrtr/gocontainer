package sort

import (
	"errors"
	"fmt"
	"reflect"
)

// Comparator imposes a total ordering on some collection of objects.
// Comparators can be passed to the construction function of a container(such as ArrayList, LinkedList or PriorityQueue) to allow precise control over the sort order.
type Comparator interface {
	// Compare compares its two arguments for order.
	// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	Compare(v1 interface{}, v2 interface{}) (int, error)
}

// Compare compares its two arguments if they have the same type and are comparable, otherwise returns an error in the second return value.
// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
func Compare(v1 interface{}, v2 interface{}) (int, error) {
	if nil == v1 && nil == v2 {
		return 0, nil
	}
	if nil == v1 || nil == v2 {
		return 0, errors.New("A nil value can't be compared to a non-nil value")
	}

	k1, k2 := reflect.TypeOf(v1).Kind(), reflect.TypeOf(v2).Kind()
	if k1 != k2 {
		return 0, fmt.Errorf("Two values of different type can't be compared, %s: %s", k1, k2)
	}

	cmpRet := 0
	switch k1 {
	case reflect.Bool:
		// false < true
		b1, b2 := v1.(bool), v2.(bool)
		if !b1 && b2 { // b1 == false && b2 == true
			cmpRet = -1
		} else if b1 && !b2 { // b1 == true && b2 == false
			cmpRet = 1
		}
	case reflect.Int:
		cv1, cv2 := v1.(int), v2.(int)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Int8:
		cv1, cv2 := v1.(int8), v2.(int8)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Int16:
		cv1, cv2 := v1.(int16), v2.(int16)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Int32:
		cv1, cv2 := v1.(int32), v2.(int32)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Int64:
		cv1, cv2 := v1.(int64), v2.(int64)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Uint:
		cv1, cv2 := v1.(uint), v2.(uint)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Uint8:
		cv1, cv2 := v1.(uint8), v2.(uint8)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Uint16:
		cv1, cv2 := v1.(uint16), v2.(uint16)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Uint32:
		cv1, cv2 := v1.(uint32), v2.(uint32)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Uint64:
		cv1, cv2 := v1.(uint64), v2.(uint64)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Float32:
		cv1, cv2 := v1.(float32), v2.(float32)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.Float64:
		cv1, cv2 := v1.(float64), v2.(float64)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	case reflect.String:
		cv1, cv2 := v1.(string), v2.(string)
		if cv1 < cv2 {
			cmpRet = -1
		} else if cv1 > cv2 {
			cmpRet = 1
		}
	default:
		return 0, fmt.Errorf("Type '%s' can't be compared", k1)
	}

	return cmpRet, nil
}
