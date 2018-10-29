package main

import "testing"

func TestPalyndrome(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		t.Helper()
		if palyndromeCheck("abcba")!=true {
			t.Errorf("not equal")
		}
	})

	t.Run("abcba", func(t *testing.T) {
		t.Helper()
		if palyndromeCheck("abccba")!=true {
			t.Errorf("not equal")
		}
	})
}