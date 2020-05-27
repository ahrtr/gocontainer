// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package utils_test

import (
	"testing"
	"time"

	"github.com/ahrtr/gocontainer/utils"
)

func TestCompare(t *testing.T) {
	// bool
	if ret, _ := utils.Compare(false, true); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(false, false); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(true, true); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(true, false); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// int
	if ret, _ := utils.Compare(1, 2); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(2, 2); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(2, 1); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// int8
	if ret, _ := utils.Compare(int8(1), int8(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int8(2), int8(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int8(2), int8(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// int16
	if ret, _ := utils.Compare(int16(1), int16(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int16(2), int16(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int16(2), int16(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// int32
	if ret, _ := utils.Compare(int32(1), int32(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int32(2), int32(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int32(2), int32(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// rune
	if ret, _ := utils.Compare(rune(1), rune(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(rune(2), rune(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(rune(2), rune(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// int64
	if ret, _ := utils.Compare(int64(1), int64(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int64(2), int64(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(int64(2), int64(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// uint
	if ret, _ := utils.Compare(uint(1), uint(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint(2), uint(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint(2), uint(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// uint8
	if ret, _ := utils.Compare(uint8(1), uint8(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint8(2), uint8(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint8(2), uint8(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// byte
	if ret, _ := utils.Compare(byte(1), byte(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(byte(2), byte(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(byte(2), byte(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// uint16
	if ret, _ := utils.Compare(uint16(1), uint16(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint16(2), uint16(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint16(2), uint16(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// uint32
	if ret, _ := utils.Compare(uint32(1), uint32(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint32(2), uint32(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint32(2), uint32(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// uint64
	if ret, _ := utils.Compare(uint64(1), uint64(2)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint64(2), uint64(2)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(uint64(2), uint64(1)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// float32
	if ret, _ := utils.Compare(float32(1.0), float32(2.0)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(float32(2.0), float32(2.0)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(float32(2.0), float32(1.0)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// float64
	if ret, _ := utils.Compare(float64(1.0), float64(2.0)); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(float64(2.0), float64(2.0)); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(float64(2.0), float64(1.0)); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// string
	if ret, _ := utils.Compare("abc", "ade"); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare("ade", "ade"); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare("ade", "abc"); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}

	// time.Time
	t1, t2 := time.Now(), time.Now().Add(10*time.Hour)
	if ret, _ := utils.Compare(t1, t2); ret != -1 {
		t.Errorf("Compare returns an unexpected value, expected: -1, actual: %d", ret)
	}
	if ret, _ := utils.Compare(t1, t1); ret != 0 {
		t.Errorf("Compare returns an unexpected value, expected: 0, actual: %d", ret)
	}
	if ret, _ := utils.Compare(t2, t1); ret != 1 {
		t.Errorf("Compare returns an unexpected value, expected: 1, actual: %d", ret)
	}
}
