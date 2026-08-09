package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestFlakyNetworkCall(t *testing.T) {
	if rand.Float64() < 0.3 {
		time.Sleep(100 * time.Millisecond)
		t.Fatalf("connection to auth-service timed out after 2000ms")
	}
}
