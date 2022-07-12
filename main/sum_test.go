package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	compareValues := func(tb testing.TB, got, want int, numbers []int) {
		tb.Helper()
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sum(numbers)
		want := 28

		compareValues(t, got, want, numbers)

	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		compareValues(t, got, want, numbers)

	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	/*if len(got) == len(want) {
		for i := 0; i < len(want); i++ {
			if got[i] != want[i] {
				t.Errorf("got %v want %v", got, want)
			}
		}

	} else {
		t.Errorf("got %v want %v", got, want)
	}*/

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
