package reverse_polish

import (
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	t.Log(strings.Trim("  123  ", " "))
}

func TestRePolish(t *testing.T) {
	t.Log(RePolish("1+3*2+4"))
}
