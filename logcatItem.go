package main

import (
	"time"
)

// LogcatItem represents a line of logcat log.
type LogcatItem map[string]string

// time return parsed time.
func (item *LogcatItem) time() time.Time {
	t, _ := time.Parse("12-28 18:54:08.043", (*item)["time"])
	return t
}

// Keys returns valid keys for LogcatItem
func (item *LogcatItem) Keys() []string {
	return []string{
		"time", "pid", "tid", "priority", "tag", "message",
	}
}
