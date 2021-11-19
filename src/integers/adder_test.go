package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := add(2, 2) //you can't force me to make functions start upper case : P
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
