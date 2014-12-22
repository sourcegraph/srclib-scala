// Package scala provides a Go def formatter for Scala definitions. It
// is used by srclib (and srclib consumers, such as Sourcegraph) to
// display human-readable names for definitions and types with varying
// levels of qualification.
//
// It is written in Go and called on-the-fly (instead of pregenerating
// its output for each def) for speed (a single search result set
// might need to format tens of defs in multiple languages, which
// means an RPC would be too slow) and flexibility (it is easy to add
// new outputs or change existing outputs to make them look nicer).
package scala

import (
	"fmt"

	"sourcegraph.com/sourcegraph/srclib/graph"
)

const scalaSourceUnitType = "scala"

func init() {
	graph.RegisterMakeDefFormatter(scalaSourceUnitType, func(d *graph.Def) graph.DefFormatter { return scalaFormatter{d} })
}

type scalaFormatter struct{ def *graph.Def }

// TODO(sqs): The actual implementation of these methods will depend
// on what data about each def is produced by the grapher and how it's
// stored. Once I see what info the Scala compiler outputs, I can
// update these to produce more useful formatted output.

func (f scalaFormatter) Name(qual graph.Qualification) string {
	switch qual {
	case graph.Unqualified:
		return f.def.Name
	case graph.ScopeQualified:
		return f.def.Name
	case graph.DepQualified:
		return string(f.def.Path)
	case graph.RepositoryWideQualified:
		return string(f.def.Path)
	case graph.LanguageWideQualified:
		return f.def.Repo + "." + string(f.def.Path)
	}
	panic("Name: unrecognized Qualification: " + fmt.Sprint(qual))
}

func (f scalaFormatter) Type(qual graph.Qualification) string {
	switch qual {
	case graph.Unqualified:
		return f.def.Name
	case graph.ScopeQualified:
		return f.def.Name
	case graph.DepQualified:
		return string(f.def.Path)
	case graph.RepositoryWideQualified:
		return string(f.def.Path)
	case graph.LanguageWideQualified:
		return f.def.Repo + "." + string(f.def.Path)
	}
	panic("Type: unrecognized Qualification: " + fmt.Sprint(qual))
}

func (scalaFormatter) Language() string             { return "Scala" }
func (scalaFormatter) DefKeyword() string           { return "def" /* TODO(sqs): or class/val/var/etc. */ }
func (scalaFormatter) NameAndTypeSeparator() string { return ": " }
func (scalaFormatter) Kind() string                 { return "" /* TODO(sqs): implement this */ }
