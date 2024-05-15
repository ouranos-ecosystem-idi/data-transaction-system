package mocks

import (
	"database/sql/driver"
	"time"
)

type Anytime struct{}

func (a Anytime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
