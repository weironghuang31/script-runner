package scripts

import (
	"bytes"
	"strings"
	"testing"
)

func TestGetShell(t *testing.T) {
	shell, err := getShell()
	if err != nil {
		t.Fatal(err)
	}

	suffix := "sh"
	if IsWindows {
		suffix += ".exe"
	}

	if !strings.HasSuffix(shell, suffix) {
		t.Fatal("Cannot not file the path of shell")
	}
}

func TestRun(t *testing.T) {
	buffer := &bytes.Buffer{}
	spec, _ := ParseSpec("../test", "good.scripts.yml")
	spec.Stdout = buffer

	err := spec.Run([]string{"foo"})

	if err != nil {
		t.Fatal(err)
	}

	// not sure why buffer.String() causes error on my machine, when the content is from Stdout
	bytes := buffer.Bytes()
	out := strings.TrimSpace(string(bytes))

	if out != "0.0.0" {
		t.Fatal("The run result isn't expected")
	}
}
