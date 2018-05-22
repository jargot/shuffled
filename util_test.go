package main

import (
  "testing"
)

func TestIsUniqInSlice(t *testing.T) {
  testData := []struct {
    single int
    slice []int
    expectedRes bool
  } {
    3, []int {3, 4, 6}, true,
    1, []int {3, 4, 6}, false,
  }
  for i := 0; i < len(testData); i++ {
    fmt.Println(testData.single)
  }

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}
