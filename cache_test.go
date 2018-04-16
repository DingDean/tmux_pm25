package main

import (
	"testing"
	"time"
)

var isExpiredTest = []struct {
	in  int64
	out bool
}{
	{time.Now().Unix() - 7200, true},
	{time.Now().Unix(), false},
}

func TestCacheIsExpired(t *testing.T) {
	for _, tt := range isExpiredTest {
		if s := cacheIsExpired(tt.in); s != tt.out {
			t.Errorf("IsExpired(%v) => %v, want %v", tt.in, s, tt.out)
		}
	}
}
