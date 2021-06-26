package scripts

import (
	"testing"
)

func TestParseGoodSpec(t *testing.T) {
	spec, err := ParseSpec("../test", "good.scripts.yml")

	if err != nil {
		t.Fatal(err)
	}

	if spec.Envs["TEST_VERSION"] != "0.0.0" {
		t.Fatal("TEST_VERSION default isn't loaded")
	}

	if spec.Scripts["foo"] == "" {
		t.Fatal("Script is empty")
	}

	if spec.Scripts["bar"] == "" {
		t.Fatal("Script is empty")
	}

	if spec.Scripts["baz"] != "" {
		t.Fatal("Script isn't empty")
	}
}

func TestParseBadSpec(t *testing.T) {
	_, err := ParseSpec("../test", "bad.scripts.yml")

	if err == nil {
		t.Fatal("ParseSpec should fail")
	}
}
