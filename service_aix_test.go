//go:build aix && ppc64
// +build aix,ppc64

package syscore_test

import (
	"os"
	"testing"
)

func interruptProcess(t *testing.T) {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Signal(os.Interrupt); err != nil {
		t.Fatal(err)
	}
}
