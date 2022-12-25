package mslm

import (
	"time"
)

// String returns the value of the string pointer passed in or "" if the
// pointer is nil.
func ToString(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}

// StringSlice converts a slice of string pointers into a slice of string
// values.
func ToStringSlice(strList []*string) []string {
	newStrList := make([]string, len(strList))
	for i := 0; i < len(strList); i++ {
		if strList[i] != nil {
			newStrList[i] = *(strList[i])
		}
	}
	return newStrList
}

// StringMap converts a string map of string pointers into a string map of
// string values.
func ToStringMap(strMap map[string]*string) map[string]string {
	newStrMap := make(map[string]string)
	for k, val := range strMap {
		if val != nil {
			newStrMap[k] = *val
		}
	}
	return newStrMap
}

// Bool returns the value of the bool pointer passed in or false if the pointer
// is nil.
func ToBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

// BoolSlice converts a slice of bool pointers into a slice of bool values.
func ToBoolSlice(boolList []*bool) []bool {
	newBoolList := make([]bool, len(boolList))
	for i := 0; i < len(boolList); i++ {
		if boolList[i] != nil {
			newBoolList[i] = *(boolList[i])
		}
	}
	return newBoolList
}

// BoolMap converts a string map of bool pointers into a string map of bool
// values.
func ToBoolMap(boolMap map[string]*bool) map[string]bool {
	newBoolMap := make(map[string]bool)
	for k, val := range boolMap {
		if val != nil {
			newBoolMap[k] = *val
		}
	}
	return newBoolMap
}

// Int returns the value of the int pointer passed in or 0 if the pointer is
// nil.
func ToInt(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// IntSlice converts a slice of int pointers into a slice of int values.
func ToIntSlice(intList []*int) []int {
	newIntList := make([]int, len(intList))
	for i := 0; i < len(intList); i++ {
		if intList[i] != nil {
			newIntList[i] = *(intList[i])
		}
	}
	return newIntList
}

// IntMap converts a string map of int pointers into a string map of int
// values.
func ToIntMap(intMap map[string]*int) map[string]int {
	newIntMap := make(map[string]int)
	for k, val := range intMap {
		if val != nil {
			newIntMap[k] = *val
		}
	}
	return newIntMap
}

// Uint returns the value of the uint pointer passed in or 0 if the pointer is
// nil.
func ToUint(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// UintSlice converts a slice of uint pointers uinto a slice of uint values.
func ToUintSlice(uintList []*uint) []uint {
	newUintList := make([]uint, len(uintList))
	for i := 0; i < len(uintList); i++ {
		if uintList[i] != nil {
			newUintList[i] = *(uintList[i])
		}
	}
	return newUintList
}

// UintMap converts a string map of uint pointers uinto a string map of uint
// values.
func ToUintMap(uintMap map[string]*uint) map[string]uint {
	newUintMap := make(map[string]uint)
	for k, val := range uintMap {
		if val != nil {
			newUintMap[k] = *val
		}
	}
	return newUintMap
}

// Int8 returns the value of the int8 pointer passed in or 0 if the pointer is
// nil.
func ToInt8(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

// Int8Slice converts a slice of int8 pointers into a slice of int8 values.
func ToInt8Slice(int8List []*int8) []int8 {
	newInt8List := make([]int8, len(int8List))
	for i := 0; i < len(int8List); i++ {
		if int8List[i] != nil {
			newInt8List[i] = *(int8List[i])
		}
	}
	return newInt8List
}

// Int8Map converts a string map of int8 pointers into a string map of int8
// values.
func ToInt8Map(int8Map map[string]*int8) map[string]int8 {
	newInt8Map := make(map[string]int8)
	for k, val := range int8Map {
		if val != nil {
			newInt8Map[k] = *val
		}
	}
	return newInt8Map
}

// Int16 returns the value of the int16 pointer passed in or 0 if the pointer
// is nil.
func ToInt16(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

// Int16Slice converts a slice of int16 pointers into a slice of int16 values.
func ToInt16Slice(int16List []*int16) []int16 {
	newInt16List := make([]int16, len(int16List))
	for i := 0; i < len(int16List); i++ {
		if int16List[i] != nil {
			newInt16List[i] = *(int16List[i])
		}
	}
	return newInt16List
}

// Int16Map converts a string map of int16 pointers into a string map of int16
// values.
func ToInt16Map(int16Map map[string]*int16) map[string]int16 {
	newInt16Map := make(map[string]int16)
	for k, val := range int16Map {
		if val != nil {
			newInt16Map[k] = *val
		}
	}
	return newInt16Map
}

// Int32 returns the value of the int32 pointer passed in or 0 if the pointer
// is nil.
func ToInt32(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int32Slice converts a slice of int32 pointers into a slice of int32 values.
func ToInt32Slice(int32List []*int32) []int32 {
	newInt32List := make([]int32, len(int32List))
	for i := 0; i < len(int32List); i++ {
		if int32List[i] != nil {
			newInt32List[i] = *(int32List[i])
		}
	}
	return newInt32List
}

// Int32Map converts a string map of int32 pointers into a string map of int32
// values.
func ToInt32Map(int32Map map[string]*int32) map[string]int32 {
	newInt32Map := make(map[string]int32)
	for k, val := range int32Map {
		if val != nil {
			newInt32Map[k] = *val
		}
	}
	return newInt32Map
}

// Int64 returns the value of the int64 pointer passed in or 0 if the pointer
// is nil.
func ToInt64(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Slice converts a slice of int64 pointers into a slice of int64 values.
func ToInt64Slice(int64List []*int64) []int64 {
	newInt64List := make([]int64, len(int64List))
	for i := 0; i < len(int64List); i++ {
		if int64List[i] != nil {
			newInt64List[i] = *(int64List[i])
		}
	}
	return newInt64List
}

// Int64Map converts a string map of int64 pointers into a string map of int64
// values.
func ToInt64Map(int64Map map[string]*int64) map[string]int64 {
	newInt64Map := make(map[string]int64)
	for k, val := range int64Map {
		if val != nil {
			newInt64Map[k] = *val
		}
	}
	return newInt64Map
}

// Uint8 returns the value of the uint8 pointer passed in or 0 if the pointer
// is nil.
func ToUint8(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint8Slice converts a slice of uint8 pointers into a slice of uint8 values.
func ToUint8Slice(uint8List []*uint8) []uint8 {
	newUint8List := make([]uint8, len(uint8List))
	for i := 0; i < len(uint8List); i++ {
		if uint8List[i] != nil {
			newUint8List[i] = *(uint8List[i])
		}
	}
	return newUint8List
}

// Uint8Map converts a string map of uint8 pointers into a string map of uint8
// values.
func ToUint8Map(uint8Map map[string]*uint8) map[string]uint8 {
	newUint8Map := make(map[string]uint8)
	for k, val := range uint8Map {
		if val != nil {
			newUint8Map[k] = *val
		}
	}
	return newUint8Map
}

// Uint16 returns the value of the uint16 pointer passed in or 0 if the pointer
// is nil.
func ToUint16(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint16Slice converts a slice of uint16 pointers into a slice of uint16
// values.
func ToUint16Slice(uint16List []*uint16) []uint16 {
	newUint16List := make([]uint16, len(uint16List))
	for i := 0; i < len(uint16List); i++ {
		if uint16List[i] != nil {
			newUint16List[i] = *(uint16List[i])
		}
	}
	return newUint16List
}

// Uint16Map converts a string map of uint16 pointers into a string map of
// uint16 values.
func ToUint16Map(uint16Map map[string]*uint16) map[string]uint16 {
	newUint16Map := make(map[string]uint16)
	for k, val := range uint16Map {
		if val != nil {
			newUint16Map[k] = *val
		}
	}
	return newUint16Map
}

// Uint32 returns the value of the uint32 pointer passed in or 0 if the pointer
// is nil.
func ToUint32(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint32Slice converts a slice of uint32 pointers into a slice of uint32
// values.
func ToUint32Slice(uint32List []*uint32) []uint32 {
	newUint32List := make([]uint32, len(uint32List))
	for i := 0; i < len(uint32List); i++ {
		if uint32List[i] != nil {
			newUint32List[i] = *(uint32List[i])
		}
	}
	return newUint32List
}

// Uint32Map converts a string map of uint32 pointers into a string map of
// uint32 values.
func ToUint32Map(uint32Map map[string]*uint32) map[string]uint32 {
	newUint32Map := make(map[string]uint32)
	for k, val := range uint32Map {
		if val != nil {
			newUint32Map[k] = *val
		}
	}
	return newUint32Map
}

// Uint64 returns the value of the uint64 pointer passed in or 0 if the pointer
// is nil.
func ToUint64(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint64Slice converts a slice of uint64 pointers into a slice of uint64
// values.
func ToUint64Slice(uint64List []*uint64) []uint64 {
	newUint64List := make([]uint64, len(uint64List))
	for i := 0; i < len(uint64List); i++ {
		if uint64List[i] != nil {
			newUint64List[i] = *(uint64List[i])
		}
	}
	return newUint64List
}

// Uint64Map converts a string map of uint64 pointers into a string map of
// uint64 values.
func ToUint64Map(uint64Map map[string]*uint64) map[string]uint64 {
	newUint64Map := make(map[string]uint64)
	for k, val := range uint64Map {
		if val != nil {
			newUint64Map[k] = *val
		}
	}
	return newUint64Map
}

// Float32 returns the value of the float32 pointer passed in or 0 if the
// pointer is nil.
func ToFloat32(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32Slice converts a slice of float32 pointers into a slice of float32
// values.
func ToFloat32Slice(float32List []*float32) []float32 {
	newFloat32List := make([]float32, len(float32List))
	for i := 0; i < len(float32List); i++ {
		if float32List[i] != nil {
			newFloat32List[i] = *(float32List[i])
		}
	}
	return newFloat32List
}

// Float32Map converts a string map of float32 pointers into a string map of
// float32 values.
func ToFloat32Map(float32Map map[string]*float32) map[string]float32 {
	newFloat32Map := make(map[string]float32)
	for k, val := range float32Map {
		if val != nil {
			newFloat32Map[k] = *val
		}
	}
	return newFloat32Map
}

// Float64 returns the value of the float64 pointer passed in or 0 if the
// pointer is nil.
func ToFloat64(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Slice converts a slice of float64 pointers into a slice of float64
// values.
func ToFloat64Slice(float64List []*float64) []float64 {
	newFloat64List := make([]float64, len(float64List))
	for i := 0; i < len(float64List); i++ {
		if float64List[i] != nil {
			newFloat64List[i] = *(float64List[i])
		}
	}
	return newFloat64List
}

// Float64Map converts a string map of float64 pointers into a string map of
// float64 values.
func ToFloat64Map(float64Map map[string]*float64) map[string]float64 {
	newFloat64Map := make(map[string]float64)
	for k, val := range float64Map {
		if val != nil {
			newFloat64Map[k] = *val
		}
	}
	return newFloat64Map
}

// Time returns the value of the time.Time pointer passed in or time.Time{} if
// the pointer is nil.
func ToTime(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

// TimeSlice converts a slice of time.Time pointers into a slice of time.Time
// values.
func ToTimeSlice(timeList []*time.Time) []time.Time {
	newTimeList := make([]time.Time, len(timeList))
	for i := 0; i < len(timeList); i++ {
		if timeList[i] != nil {
			newTimeList[i] = *(timeList[i])
		}
	}
	return newTimeList
}

// TimeMap converts a string map of time.Time pointers into a string map of
// time.Time values.
func ToTimeMap(timeMap map[string]*time.Time) map[string]time.Time {
	newTimeMap := make(map[string]time.Time)
	for k, val := range timeMap {
		if val != nil {
			newTimeMap[k] = *val
		}
	}
	return newTimeMap
}
