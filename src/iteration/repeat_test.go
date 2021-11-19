package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := repeat("a", 5) //lower case r.  : P
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 5)
	}
}

func ExampleRepeat() {
	returnedValue := repeat("a", 7)
	fmt.Println(returnedValue)
	// Output: aaaaaaa
}
