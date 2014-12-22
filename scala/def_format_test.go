package scala

import (
	"testing"

	"sourcegraph.com/sourcegraph/srclib/graph"
)

func TestFormatter(t *testing.T) {
	def := &graph.Def{
		DefKey: graph.DefKey{Repo: "x.com/r", UnitType: "scala", Unit: "u", Path: "p"},
		Name:   "n",
	}

	if got, want := def.Fmt().Name(graph.Unqualified), "n"; got != want {
		t.Errorf("got Name(Unqualified) == %q, want %q", got, want)
	}
	if got, want := def.Fmt().Name(graph.RepositoryWideQualified), "p"; got != want {
		t.Errorf("got Name(RepositoryWideQualified) == %q, want %q", got, want)
	}
}
