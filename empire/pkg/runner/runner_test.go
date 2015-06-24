package runner

import (
	"os"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/remind101/empire/empire/pkg/dockerutil"
)

func TestRunner(t *testing.T) {
	r := newTestRunner(t)

	if err := r.Run(context.Background(), RunOpts{
		Image:   "ubuntu:14.04",
		Command: "/bin/bash",
		Input:   strings.NewReader(`ls`),
		Output:  os.Stdout,
	}); err != nil {
		t.Fatal(err)
	}
}

func newTestRunner(t testing.TB) *Runner {
	c, err := dockerutil.NewClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	return NewRunner(c)
}
