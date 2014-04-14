package main

import (
	"github.com/edmore/esp/network"

	"testing"
)

func TestEvaluate(t *testing.T) {
	t.Skip("skipping test")
}

func TestInitialize(t *testing.T) {
	t.Skip("skipping test")
}

func BenchmarkInitialize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		feedForward := network.NewFeedForward(6, 5, 1, true)
		// Initialization
		_ = initialize(5, 100, feedForward.GeneSize)
	}
}
