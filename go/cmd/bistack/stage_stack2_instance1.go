package main

import (
	"time"

	"github.com/thomaspeugeot/bistack/otherstack/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Bar__000000_Bar_1 := (&models.Bar{}).Stage(stage)

	// Setup of values

	__Bar__000000_Bar_1.Name = `Bar 1`

	// Setup of pointers
}
