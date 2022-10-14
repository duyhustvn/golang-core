package implement

import (
	"testing"
)

func TestImplementKind(t *testing.T) {
	type DataTest struct {
		Input          rtype
		ExpectedOutput string
	}

	tests := []DataTest{
		{rtype{kind: 2}, "int"},
	}

	for _, test := range tests {
		ImplementKind(test.Input)
	}

}
