package what

import (
	"bytes"
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	got := &bytes.Buffer{}
	log.SetOutput(got)
	log.SetFlags(0)
	n := 23

	// no package name set - all packages are enabled
	enabled = map[string]bool{}
	Happens("what.Happens - all packages enabled")
	If(true, "If true")
	If(false, "If false")
	Is(n)
	Func()
	Package()

	enabled = map[string]bool{
		"what": true,
	}
	Happens("what.Happens - package 'what' enabled")

	enabled = map[string]bool{
		"appliedgo.net/what": true,
	}
	Happens("what.Happens - package 'appliedgo.net/what' enabled")

	enabled = map[string]bool{
		"someotherpackage": true,
	}
	Happens("what.Happens - package 'what' NOT enabled") // this should not print

	want := `appliedgo.net/what.TestAll: what.Happens - all packages enabled
appliedgo.net/what.If: If true
(int) 23
Func appliedgo.net/what.TestAll in line 21 of file /Users/christoph/dev/go/what/what_test.go
Package appliedgo.net/what
appliedgo.net/what.TestAll: what.Happens - package 'what' enabled
appliedgo.net/what.TestAll: what.Happens - package 'appliedgo.net/what' enabled
`

	if got.String() != want {
		t.Errorf("Got: %s\n\nWant: %s", got, want)
	}
}
