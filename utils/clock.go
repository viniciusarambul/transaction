//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=clock.go -destination=utilsmock/clock.go .
package utils

import "time"

type (
	Clock interface {
		Now() time.Time
	}

	clock struct{}
)

func New() Clock {
	return &clock{}
}

func (clock) Now() time.Time {
	return time.Now().Local().UTC()
}
