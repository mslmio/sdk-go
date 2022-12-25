package mslm

import (
	"time"
)

// String returns a pointer to the string value passed in.
func String(str string) *string {
	return &str
}

// StringSlice converts a slice of string values into a slice of string
// pointers.
func StringSlice(strList []string) []*string {
	newStrList := make([]*string, len(strList))
	for i := 0; i < len(strList); i++ {
		newStrList[i] = &(strList[i])
	}
	return newStrList
}

// StringMap converts a string map of string values into a string map of string
// pointers.
func StringMap(strMap map[string]string) map[string]*string {
	newStrMap := make(map[string]*string)
	for k, val := range strMap {
		v := val
		newStrMap[k] = &v
	}
	return newStrMap
}

// Bool returns a pointer to the bool value passed in.
func Bool(b bool) *bool {
	return &b
}

// BoolSlice converts a slice of bool values into a slice of bool pointers.
func BoolSlice(boolList []bool) []*bool {
	newBoolList := make([]*bool, len(boolList))
	for i := 0; i < len(boolList); i++ {
		newBoolList[i] = &(boolList[i])
	}
	return newBoolList
}

// BoolMap converts a string map of bool values into a string map of bool
// pointers.
func BoolMap(boolMap map[string]bool) map[string]*bool {
	newBoolMap := make(map[string]*bool)
	for k, val := range boolMap {
		v := val
		newBoolMap[k] = &v
	}
	return newBoolMap
}

// Int returns a pointer to the int value passed in.
func Int(v int) *int {
	return &v
}

// IntSlice converts a slice of int values into a slice of int pointers.
func IntSlice(intList []int) []*int {
	newIntList := make([]*int, len(intList))
	for i := 0; i < len(intList); i++ {
		newIntList[i] = &(intList[i])
	}
	return newIntList
}

// IntMap converts a string map of int values into a string map of int
// pointers.
func IntMap(intMap map[string]int) map[string]*int {
	newIntMap := make(map[string]*int)
	for k, val := range intMap {
		v := val
		newIntMap[k] = &v
	}
	return newIntMap
}

// Uint returns a pointer to the uint value passed in.
func Uint(v uint) *uint {
	return &v
}

// UintSlice converts a slice of uint values uinto a slice of uint pointers.
func UintSlice(uintList []uint) []*uint {
	newUintList := make([]*uint, len(uintList))
	for i := 0; i < len(uintList); i++ {
		newUintList[i] = &(uintList[i])
	}
	return newUintList
}

// UintMap converts a string map of uint values uinto a string map of uint
// pointers.
func UintMap(uintMap map[string]uint) map[string]*uint {
	newUintMap := make(map[string]*uint)
	for k, val := range uintMap {
		v := val
		newUintMap[k] = &v
	}
	return newUintMap
}

// Int8 returns a pointer to the int8 value passed in.
func Int8(v int8) *int8 {
	return &v
}

// Int8Slice converts a slice of int8 values into a slice of int8 pointers.
func Int8Slice(int8List []int8) []*int8 {
	newInt8List := make([]*int8, len(int8List))
	for i := 0; i < len(int8List); i++ {
		newInt8List[i] = &(int8List[i])
	}
	return newInt8List
}

// Int8Map converts a string map of int8 values into a string map of int8
// pointers.
func Int8Map(int8Map map[string]int8) map[string]*int8 {
	newInt8Map := make(map[string]*int8)
	for k, val := range int8Map {
		v := val
		newInt8Map[k] = &v
	}
	return newInt8Map
}

// Int16 returns a pointer to the int16 value passed in.
func Int16(v int16) *int16 {
	return &v
}

// Int16Slice converts a slice of int16 values into a slice of int16 pointers.
func Int16Slice(int16List []int16) []*int16 {
	newInt16List := make([]*int16, len(int16List))
	for i := 0; i < len(int16List); i++ {
		newInt16List[i] = &(int16List[i])
	}
	return newInt16List
}

// Int16Map converts a string map of int16 values into a string map of int16
// pointers.
func Int16Map(int16Map map[string]int16) map[string]*int16 {
	newInt16Map := make(map[string]*int16)
	for k, val := range int16Map {
		v := val
		newInt16Map[k] = &v
	}
	return newInt16Map
}

// Int32 returns a pointer to the int32 value passed in.
func Int32(v int32) *int32 {
	return &v
}

// Int32Slice converts a slice of int32 values into a slice of int32 pointers.
func Int32Slice(int32List []int32) []*int32 {
	newInt32List := make([]*int32, len(int32List))
	for i := 0; i < len(int32List); i++ {
		newInt32List[i] = &(int32List[i])
	}
	return newInt32List
}

// Int32Map converts a string map of int32 values into a string map of int32
// pointers.
func Int32Map(int32Map map[string]int32) map[string]*int32 {
	newInt32Map := make(map[string]*int32)
	for k, val := range int32Map {
		v := val
		newInt32Map[k] = &v
	}
	return newInt32Map
}

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Slice converts a slice of int64 values into a slice of int64 pointers.
func Int64Slice(int64List []int64) []*int64 {
	newInt64List := make([]*int64, len(int64List))
	for i := 0; i < len(int64List); i++ {
		newInt64List[i] = &(int64List[i])
	}
	return newInt64List
}

// Int64Map converts a string map of int64 values into a string map of int64
// pointers.
func Int64Map(int64Map map[string]int64) map[string]*int64 {
	newInt64Map := make(map[string]*int64)
	for k, val := range int64Map {
		v := val
		newInt64Map[k] = &v
	}
	return newInt64Map
}

// Uint8 returns a pointer to the uint8 value passed in.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint8Slice converts a slice of uint8 values into a slice of uint8 pointers.
func Uint8Slice(uint8List []uint8) []*uint8 {
	newUint8List := make([]*uint8, len(uint8List))
	for i := 0; i < len(uint8List); i++ {
		newUint8List[i] = &(uint8List[i])
	}
	return newUint8List
}

// Uint8Map converts a string map of uint8 values into a string map of uint8
// pointers.
func Uint8Map(uint8Map map[string]uint8) map[string]*uint8 {
	newUint8Map := make(map[string]*uint8)
	for k, val := range uint8Map {
		v := val
		newUint8Map[k] = &v
	}
	return newUint8Map
}

// Uint16 returns a pointer to the uint16 value passed in.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint16Slice converts a slice of uint16 values into a slice of uint16
// pointers.
func Uint16Slice(uint16List []uint16) []*uint16 {
	newUint16List := make([]*uint16, len(uint16List))
	for i := 0; i < len(uint16List); i++ {
		newUint16List[i] = &(uint16List[i])
	}
	return newUint16List
}

// Uint16Map converts a string map of uint16 values into a string map of uint16
// pointers.
func Uint16Map(uint16Map map[string]uint16) map[string]*uint16 {
	newUint16Map := make(map[string]*uint16)
	for k, val := range uint16Map {
		v := val
		newUint16Map[k] = &v
	}
	return newUint16Map
}

// Uint32 returns a pointer to the uint32 value passed in.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint32Slice converts a slice of uint32 values into a slice of uint32
// pointers.
func Uint32Slice(uint32List []uint32) []*uint32 {
	newUint32List := make([]*uint32, len(uint32List))
	for i := 0; i < len(uint32List); i++ {
		newUint32List[i] = &(uint32List[i])
	}
	return newUint32List
}

// Uint32Map converts a string map of uint32 values into a string map of uint32
// pointers.
func Uint32Map(uint32Map map[string]uint32) map[string]*uint32 {
	newUint32Map := make(map[string]*uint32)
	for k, val := range uint32Map {
		v := val
		newUint32Map[k] = &v
	}
	return newUint32Map
}

// Uint64 returns a pointer to the uint64 value passed in.
func Uint64(v uint64) *uint64 {
	return &v
}

// Uint64Slice converts a slice of uint64 values into a slice of uint64
// pointers.
func Uint64Slice(uint64List []uint64) []*uint64 {
	newUint64List := make([]*uint64, len(uint64List))
	for i := 0; i < len(uint64List); i++ {
		newUint64List[i] = &(uint64List[i])
	}
	return newUint64List
}

// Uint64Map converts a string map of uint64 values into a string map of uint64
// pointers.
func Uint64Map(uint64Map map[string]uint64) map[string]*uint64 {
	newUint64Map := make(map[string]*uint64)
	for k, val := range uint64Map {
		v := val
		newUint64Map[k] = &v
	}
	return newUint64Map
}

// Float32 returns a pointer to the float32 value passed in.
func Float32(v float32) *float32 {
	return &v
}

// Float32Slice converts a slice of float32 values into a slice of float32
// pointers.
func Float32Slice(float32List []float32) []*float32 {
	newFloat32List := make([]*float32, len(float32List))
	for i := 0; i < len(float32List); i++ {
		newFloat32List[i] = &(float32List[i])
	}
	return newFloat32List
}

// Float32Map converts a string map of float32 values into a string map of
// float32 pointers.
func Float32Map(float32Map map[string]float32) map[string]*float32 {
	newFloat32Map := make(map[string]*float32)
	for k, val := range float32Map {
		v := val
		newFloat32Map[k] = &v
	}
	return newFloat32Map
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Slice converts a slice of float64 values into a slice of float64
// pointers.
func Float64Slice(float64List []float64) []*float64 {
	newFloat64List := make([]*float64, len(float64List))
	for i := 0; i < len(float64List); i++ {
		newFloat64List[i] = &(float64List[i])
	}
	return newFloat64List
}

// Float64Map converts a string map of float64 values into a string map of
// float64 pointers.
func Float64Map(float64Map map[string]float64) map[string]*float64 {
	newFloat64Map := make(map[string]*float64)
	for k, val := range float64Map {
		v := val
		newFloat64Map[k] = &v
	}
	return newFloat64Map
}

// Time returns a pointer to the time.Time value passed in.
func Time(v time.Time) *time.Time {
	return &v
}

// TimeSlice converts a slice of time.Time values into a slice of time.Time
// pointers.
func TimeSlice(timeList []time.Time) []*time.Time {
	newTimeList := make([]*time.Time, len(timeList))
	for i := 0; i < len(timeList); i++ {
		newTimeList[i] = &(timeList[i])
	}
	return newTimeList
}

// TimeMap converts a string map of time.Time values into a string map of
// time.Time pointers.
func TimeMap(timeMap map[string]time.Time) map[string]*time.Time {
	newTimeMap := make(map[string]*time.Time)
	for k, val := range timeMap {
		v := val
		newTimeMap[k] = &v
	}
	return newTimeMap
}
