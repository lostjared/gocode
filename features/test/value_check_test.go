package value_check

import (
	"testing"
	"fmt"
)

func TestSpaceCheck(t *testing.T) {
	if(hasSpace("test")==false) {
		fmt.Println("Good Passed\n")
	} else {
		t.Errorf("%s\n", "Did not pass test, does not contain space")
	}
	if(hasSpace("test 1 2 3")==true) {
		fmt.Println("Good Passed")
	} else {
		t.Errorf("%s\n", "Did not pass test, contains space")
	}
}