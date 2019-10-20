package mash

import (
	"time"
)

type TimeHandler struct {
}

func NewTimeHandler() *TimeHandler {
	return &TimeHandler{}
}

func (h *TimeHandler) Now() time.Time {
	return time.Now()
}

func (h *TimeHandler) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (h *TimeHandler) FormatDuration(d time.Duration) string {
	return d.String()
}
