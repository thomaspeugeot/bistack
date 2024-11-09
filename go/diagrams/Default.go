package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/thomaspeugeot/bistack/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__GongStructShape__000000_Default_Foo := (&models.GongStructShape{}).Stage(stage)

	__Position__000000_Pos_Default_Foo := (&models.Position{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__GongStructShape__000000_Default_Foo.Name = `Default-Foo`

	//gong:ident [ref_models.Foo] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_Foo.Identifier = `ref_models.Foo`
	__GongStructShape__000000_Default_Foo.ShowNbInstances = false
	__GongStructShape__000000_Default_Foo.NbInstances = 0
	__GongStructShape__000000_Default_Foo.Width = 240.000000
	__GongStructShape__000000_Default_Foo.Height = 63.000000
	__GongStructShape__000000_Default_Foo.IsSelected = false

	__Position__000000_Pos_Default_Foo.X = 45.000000
	__Position__000000_Pos_Default_Foo.Y = 70.000000
	__Position__000000_Pos_Default_Foo.Name = `Pos-Default-Foo`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Foo)
	__GongStructShape__000000_Default_Foo.Position = __Position__000000_Pos_Default_Foo
}
