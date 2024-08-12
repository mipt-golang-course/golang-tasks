package testequal

// Important information!
// If you change this file,
// then change 33,34,35 strings
// There should be numbers of
// strings with AssertEqual,
// AssertNotEqual and RequireEqual

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelper(t *testing.T) {
	if os.Getenv("FAIL_ASSERTIONS") == "1" {
		AssertEqual(t, 1, 2, "%d must be equal to %d", 1, 2)
		AssertNotEqual(t, 1, 1, "1 != 1")
		RequireEqual(t, 1, 2)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.v", "-test.run=TestHelper")
	cmd.Env = append(os.Environ(), "FAIL_ASSERTIONS=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf

	err := cmd.Run()
	var exitErr *exec.ExitError
	if ok := errors.As(err, &exitErr); ok && !exitErr.Success() {
		require.Contains(t, buf.String(), "helper_test.go:22")
		require.Contains(t, buf.String(), "helper_test.go:23")
		require.Contains(t, buf.String(), "helper_test.go:24")
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
