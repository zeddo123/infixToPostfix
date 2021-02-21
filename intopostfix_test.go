package intopostfix

import (
	"testing"
)

//TestConvert tests the Convert function
func TestConvert(t *testing.T) {
	result := Convert("(a+b)*c")
	if result.String() != "ab+c*" {
		t.Errorf("convertion failed")
	}
}
