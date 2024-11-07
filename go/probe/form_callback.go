// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/bistack/go/models"
	"github.com/thomaspeugeot/bistack/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__FooFormCallback(
	foo *models.Foo,
	probe *Probe,
	formGroup *table.FormGroup,
) (fooFormCallback *FooFormCallback) {
	fooFormCallback = new(FooFormCallback)
	fooFormCallback.probe = probe
	fooFormCallback.foo = foo
	fooFormCallback.formGroup = formGroup

	fooFormCallback.CreationMode = (foo == nil)

	return
}

type FooFormCallback struct {
	foo *models.Foo

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (fooFormCallback *FooFormCallback) OnSave() {

	log.Println("FooFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	fooFormCallback.probe.formStage.Checkout()

	if fooFormCallback.foo == nil {
		fooFormCallback.foo = new(models.Foo).Stage(fooFormCallback.probe.stageOfInterest)
	}
	foo_ := fooFormCallback.foo
	_ = foo_

	for _, formDiv := range fooFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(foo_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if fooFormCallback.formGroup.HasSuppressButtonBeenPressed {
		foo_.Unstage(fooFormCallback.probe.stageOfInterest)
	}

	fooFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Foo](
		fooFormCallback.probe,
	)
	fooFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if fooFormCallback.CreationMode || fooFormCallback.formGroup.HasSuppressButtonBeenPressed {
		fooFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(fooFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__FooFormCallback(
			nil,
			fooFormCallback.probe,
			newFormGroup,
		)
		foo := new(models.Foo)
		FillUpForm(foo, newFormGroup, fooFormCallback.probe)
		fooFormCallback.probe.formStage.Commit()
	}

	fillUpTree(fooFormCallback.probe)
}
