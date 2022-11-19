package mathmod

import (
    "testing"
)

// Tests add function
// add(5, 8) should return 13
func TestAdd(t *testing.T) {
    z := Add(5, 8)
    if z != 13 {
        t.Fatalf("add(5, 8) function did not return 13.")
    }
}
