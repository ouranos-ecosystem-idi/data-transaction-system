package common

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// StringPtr
// Summary: This is function which returns the pointer to the string value.
// input: s(string) string value
// output: (*string) pointer to the string value
func StringPtr(s string) *string {
	return &s
}

// IntPtr
// Summary: This is function which returns the pointer to the integer value.
// input: i(int) integer value
// output: (*int) pointer to the integer value
func IntPtr(i int) *int {
	return &i
}

// Float64Ptr
// Summary: This is function which returns the pointer to the float64 value.
// input: f(float64) float64 value
// output: (*float64) pointer to the float64 value
func Float64Ptr(f float64) *float64 {
	return &f
}

// BoolPtr
// Summary: This is function which returns the pointer to the Boolean value.
// input: f(bool) bool value
// output: (*bool) pointer to the bool value
func BoolPtr(b bool) *bool {
	return &b
}

// UUIDPtr
// Summary: This is function which returns the pointer to the UUID value.
// input: f(uuid.UUID) UUID value
// output: (*uuid.UUID) pointer to the UUID value
func UUIDPtr(u uuid.UUID) *uuid.UUID {
	return &u
}

// UUIDPtrToStringPtr
// Summary: This is function which converts the UUID pointer to the string pointer.
// input: u(*uuid.UUID) UUID pointer
// output: (*string) string pointer
func UUIDPtrToStringPtr(u *uuid.UUID) *string {
	if u == nil {
		return nil
	}
	s := u.String()
	return &s
}

// IsStringsEqual
// Summary: This is function which checks whether the two string slices are equal.
// input: a([]string) string slice
// input: b([]string) string slice
// output: (bool) true if the two string slices are equal, false otherwise
func IsStringsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// IsStrPtrValEqual
// Summary: This is function which checks whether the two string pointers are equal.
// input: a(*string) string pointer
// input: b(*string) string pointer
// output: (bool) true if the two string pointers are equal, false otherwise
func IsStrPtrValEqual(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		return *a == *b
	}
	return false
}

// UUIDsToStrings
// Summary: This is function which converts the UUID slice to the string slice.
// input: uuids([]uuid.UUID) UUID slice
// output: ([]string) string slice
func UUIDsToStrings(uuids []uuid.UUID) []string {
	ss := make([]string, len(uuids))
	for i, v := range uuids {
		ss[i] = v.String()
	}
	return ss
}

// JoinUUIDs
// Summary: This is function which joins the UUID slice with the separator.
// input: uuids([]uuid.UUID) UUID slice
// input: sep(string) separator
// output: (string) joined string
func JoinUUIDs(uuids []uuid.UUID, sep string) string {
	ss := make([]string, len(uuids))
	for i, v := range uuids {
		ss[i] = v.String()
	}
	return strings.Join(ss, sep)
}

// JoinUUIDsAsPtr
// Summary: This is function which joins the UUID slice with the separator and returns the pointer to the joined string.
// input: uuids([]uuid.UUID) UUID slice
// input: sep(string) separator
// output: (*string) pointer to the joined string
func JoinUUIDsAsPtr(uuids []uuid.UUID, sep string) *string {
	if len(uuids) == 0 {
		return nil
	}
	ss := make([]string, len(uuids))
	for i, v := range uuids {
		ss[i] = v.String()
	}
	s := strings.Join(ss, sep)
	return &s
}

// GenerateUUIDString
// Summary: This is function which generates the UUID string.
// input: n(int) number of UUID strings to generate
// output: (string) generated UUID string with comma separator
func GenerateUUIDString(n int) string {
	UUIDs := make([]string, 0, n)
	for i := 0; i < n; i++ {
		newUUID, _ := uuid.NewUUID()
		UUIDs = append(UUIDs, newUUID.String())
	}
	return strings.Join(UUIDs, ",")
}

// GenerateCurrentTime
// Summary: This is function which convert string pointer from current time
// output: (string) string of current time
func GenerateCurrentTime() string {
	now := time.Now()
	utcNow := now.UTC()
	return GenerateTime(utcNow)
}

// GenerateTime
// Summary: This is function which convert string pointer from Time
// output: (string) string of time
func GenerateTime(baseTime time.Time) string {
	// Convert to timestamp compliant with "ISO8601" compliant string and "UTC"
	iso8601Format := "2006-01-02T15:04:05Z" // ISO8601
	return baseTime.Format(iso8601Format)
}
