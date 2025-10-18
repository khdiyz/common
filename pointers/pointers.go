package pointers

import "time"

// String returns a pointer to the string value.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer.
// If the pointer is nil, it returns an empty string.
func StringValue(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// Int returns a pointer to the int value.
func Int(v int) *int {
	return &v
}

// IntValue returns the value of the int pointer.
// If the pointer is nil, it returns 0.
func IntValue(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

// Int8 returns a pointer to the int8 value.
func Int8(v int8) *int8 {
	return &v
}

// Int8Value returns the value of the int8 pointer.
// If the pointer is nil, it returns 0.
func Int8Value(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

// Int16 returns a pointer to the int16 value.
func Int16(v int16) *int16 {
	return &v
}

// Int16Value returns the value of the int16 pointer.
// If the pointer is nil, it returns 0.
func Int16Value(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

// Int32 returns a pointer to the int32 value.
func Int32(v int32) *int32 {
	return &v
}

// Int32Value returns the value of the int32 pointer.
// If the pointer is nil, it returns 0.
func Int32Value(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

// Int64 returns a pointer to the int64 value.
func Int64(v int64) *int64 {
	return &v
}

// Int64Value returns the value of the int64 pointer.
// If the pointer is nil, it returns 0.
func Int64Value(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

// Uint returns a pointer to the uint value.
func Uint(v uint) *uint {
	return &v
}

// UintValue returns the value of the uint pointer.
// If the pointer is nil, it returns 0.
func UintValue(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

// Uint8 returns a pointer to the uint8 value.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint8Value returns the value of the uint8 pointer.
// If the pointer is nil, it returns 0.
func Uint8Value(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

// Uint16 returns a pointer to the uint16 value.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint16Value returns the value of the uint16 pointer.
// If the pointer is nil, it returns 0.
func Uint16Value(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

// Uint32 returns a pointer to the uint32 value.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint32Value returns the value of the uint32 pointer.
// If the pointer is nil, it returns 0.
func Uint32Value(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

// Uint64 returns a pointer to the uint64 value.
func Uint64(v uint64) *uint64 {
	return &v
}

// Uint64Value returns the value of the uint64 pointer.
// If the pointer is nil, it returns 0.
func Uint64Value(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

// Float32 returns a pointer to the float32 value.
func Float32(v float32) *float32 {
	return &v
}

// Float32Value returns the value of the float32 pointer.
// If the pointer is nil, it returns 0.0.
func Float32Value(v *float32) float32 {
	if v == nil {
		return 0.0
	}
	return *v
}

// Float64 returns a pointer to the float64 value.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer.
// If the pointer is nil, it returns 0.0.
func Float64Value(v *float64) float64 {
	if v == nil {
		return 0.0
	}
	return *v
}

// Bool returns a pointer to the bool value.
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer.
// If the pointer is nil, it returns false.
func BoolValue(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

// Byte returns a pointer to the byte value.
func Byte(v byte) *byte {
	return &v
}

// ByteValue returns the value of the byte pointer.
// If the pointer is nil, it returns 0.
func ByteValue(v *byte) byte {
	if v == nil {
		return 0
	}
	return *v
}

// Rune returns a pointer to the rune value.
func Rune(v rune) *rune {
	return &v
}

// RuneValue returns the value of the rune pointer.
// If the pointer is nil, it returns 0.
func RuneValue(v *rune) rune {
	if v == nil {
		return 0
	}
	return *v
}

// Time returns a pointer to the time.Time value.
func Time(v time.Time) *time.Time {
	return &v
}

// TimeValue returns the value of the time.Time pointer.
// If the pointer is nil, it returns the zero time.
func TimeValue(v *time.Time) time.Time {
	if v == nil {
		return time.Time{}
	}
	return *v
}

// Duration returns a pointer to the time.Duration value.
func Duration(v time.Duration) *time.Duration {
	return &v
}

// DurationValue returns the value of the time.Duration pointer.
// If the pointer is nil, it returns 0.
func DurationValue(v *time.Duration) time.Duration {
	if v == nil {
		return 0
	}
	return *v
}
