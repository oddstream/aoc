package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		if got := part1(); got != 138 {
			t.Errorf("part1() = %v, want %v", got, 138)
		}
	})
}

func Test_part2(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		if got := part2(); got != 1771 {
			t.Errorf("part1() = %v, want %v", got, 1771)
		}
	})
}
