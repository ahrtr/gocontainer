// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Comparator imposes a total ordering on some collection of objects, and it allows precise control over the sort order.
type Comparator interface {
	// Compare compares its two arguments for order.
	// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	Compare(v1 interface{}, v2 interface{}) (int, error)
}

// Compare compares two arguments using the given Comparator. If the Comparator isn't provided, then the two values are compared according to their natural ordering.
// They must be the same type, otherwise returns an error in the second return value.
// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
func Compare(v1 interface{}, v2 interface{}, cmp Comparator) (int, error) {
	if nil == v1 && nil == v2 {
		return 0, nil
	}
	if nil == v1 || nil == v2 {
		return 0, errors.New("a nil value can't be compared to a non-nil value")
	}

	k1, k2 := reflect.TypeOf(v1).Kind(), reflect.TypeOf(v2).Kind()
	if k1 != k2 {
		return 0, fmt.Errorf("two values of different type can't be compared, %s: %s", k1, k2)
	}

	// Compare the two values using the given customized comparator
	if cmp != nil {
		return cmp.Compare(v1, v2)
	}

	cmpRet := 0
	switch k1 {
	case reflect.Int:
		cv1, cv2 := v1.(int), v2.(int)
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
	case reflect.String:
		cv1, cv2 := v1.(string), v2.(string)
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
	case reflect.Bool:
		// false < true
		b1, b2 := v1.(bool), v2.(bool)
		if !b1 && b2 { // b1 == false && b2 == true
			cmpRet = -1
		} else if b1 && !b2 { // b1 == true && b2 == false
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
	case reflect.Int32: // valid for both int32 and rune
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
	case reflect.Uint8: // valid for both uint8 and byte
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
	case reflect.Struct:
		// compare time
		isBothTime, timeCmpRet := CompareTime(v1, v2)
		if isBothTime {
			return timeCmpRet, nil
		}
		return 0, errors.New("please define a customized sort.Comparator for your struct")
	default:
		return 0, fmt.Errorf("type '%s' can't be compared", k1)
	}

	return cmpRet, nil
}

// CompareTime compares its two arguments if both of them are time.Time, and returns true
// and the comparison result; otherwise return false in the first return argument.
func CompareTime(v1 interface{}, v2 interface{}) (bool, int) {
	time1, ok1 := v1.(time.Time)
	time2, ok2 := v2.(time.Time)
	if ok1 && ok2 {
		if time1.Before(time2) {
			return true, -1
		}
		if time1.After(time2) {
			return true, 1
		}
		return true, 0
	}
	return false, 0
}
