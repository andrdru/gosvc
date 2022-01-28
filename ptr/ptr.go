// Package ptr constant pointer helper
package ptr

import "time"

// Bool return bool pointer
func Bool(from bool) *bool {
	return &from
}

// Int return int pointer
func Int(from int) *int {
	return &from
}

// Int64 return int64 pointer
func Int64(from int64) *int64 {
	return &from
}

// String return string pointer
func String(from string) *string {
	return &from
}

// Time return time.Time pointer
func Time(from time.Time) *time.Time {
	return &from
}
