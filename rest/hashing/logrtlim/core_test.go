package logrtlim

import "testing"

type Test struct {
	ts   int
	msg  string
	want bool
}

type Seq struct {
	tests []Test
}

func TestLoggerShouldPrintMessage(t *testing.T) {
	var testSeqs = []Seq{{[]Test{{1, "foo", true},
		{2, "bar", true},
		{3, "foo", false},
		{8, "bar", false},
		{10, "foo", false},
		{11, "foo", true}}}}

	for _, seq := range testSeqs {
		logger := Constructor()
		for _, test := range seq.tests {
			if got := logger.ShouldPrintMessage(test.ts, test.msg); got != test.want {
				t.Errorf("ShouldPrintMessage(%d, %s) = %t, want %t", test.ts, test.msg, got, test.want)
			}
		}
	}
}
