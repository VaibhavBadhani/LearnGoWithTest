package integers

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// Notice the special format of the comment, // Output: 6. While the example will always be compiled, adding this comment means the example will also be executed. Go ahead and temporarily remove the comment // Output: 6, then run go test, and you will see ExampleAdd is no longer executed.
