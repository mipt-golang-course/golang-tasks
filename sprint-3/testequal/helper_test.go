package testequal_test

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"testing"

	assertion "github.com/mipt-golang-course/golang-tasks/sprint-3/testequal"
	"github.com/stretchr/testify/require"
)

func TestHelper(t *testing.T) {
	if os.Getenv("FAIL_ASSERTIONS") == "1" {
		assertion.AssertEqual(t, 1, 2, "%d must be equal to %d", 1, 2)
		assertion.AssertNotEqual(t, 1, 1, "1 != 1")
		assertion.RequireEqual(t, 1, 2)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.v", "-test.run=TestHelper")
	cmd.Env = append(os.Environ(), "FAIL_ASSERTIONS=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf

	err := cmd.Run()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) && !exitErr.Success() {
		require.Contains(t, buf.String(), "helper_test.go:16")
		require.Contains(t, buf.String(), "helper_test.go:17")
		require.Contains(t, buf.String(), "helper_test.go:18")
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
