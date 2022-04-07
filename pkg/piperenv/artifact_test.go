package piperenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Artifacts_FindByKind(t *testing.T) {
	artifacts := Artifacts([]Artifact{{
		Path: "a.jar",
		Kind: "jar",
	}, {
		Path: "b",
		Kind: "elf",
	}})

	assert.Len(t, artifacts.FindByKind("garbage"), 0)

	filtered := artifacts.FindByKind("jar")
	require.Len(t, filtered, 1)
	assert.Equal(t, "a.jar", filtered[0].Path)
}
