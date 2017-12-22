package cache

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

func TestIsExpired(t *testing.T) {
	for _, tt := range isExpiredTest {
		if s := IsExpired(tt.in); s != tt.out {
			t.Errorf("IsExpired(%v) => %v, want %v", tt.in, s, tt.out)
		}
	}
}
