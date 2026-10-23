package test

import (
	"github.com/openimsdk/protocol/constant"
	"testing"
)

func TestGLBHasIndependentPlatformAndTerminalClass(t *testing.T) {
	if got := constant.PlatformIDToName(13); got != "GLB" {
		t.Fatalf("platform 13 name = %q, want GLB", got)
	}
	if got := constant.PlatformNameToID("GLB"); got != 13 {
		t.Fatalf("GLB platform = %d, want 13", got)
	}
	if constant.PlatformIDToClass(13) != "GLB" || constant.PlatformNameToClass("GLB") != "GLB" {
		t.Fatal("GLB must have its own terminal class")
	}
	for id := 1; id <= 12; id++ {
		if constant.PlatformIDToName(id) == "" || constant.PlatformIDToClass(id) == "GLB" {
			t.Fatalf("existing platform %d lost its identity or shares GLB class", id)
		}
	}
}
