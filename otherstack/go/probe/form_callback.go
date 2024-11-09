// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/bistack/otherstack/go/models"
	"github.com/thomaspeugeot/bistack/otherstack/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__BarFormCallback(
	bar *models.Bar,
	probe *Probe,
	formGroup *table.FormGroup,
) (barFormCallback *BarFormCallback) {
	barFormCallback = new(BarFormCallback)
	barFormCallback.probe = probe
	barFormCallback.bar = bar
	barFormCallback.formGroup = formGroup

	barFormCallback.CreationMode = (bar == nil)

	return
}

type BarFormCallback struct {
	bar *models.Bar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (barFormCallback *BarFormCallback) OnSave() {

	log.Println("BarFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	barFormCallback.probe.formStage.Checkout()

	if barFormCallback.bar == nil {
		barFormCallback.bar = new(models.Bar).Stage(barFormCallback.probe.stageOfInterest)
	}
	bar_ := barFormCallback.bar
	_ = bar_

	for _, formDiv := range barFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bar_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bar_.Unstage(barFormCallback.probe.stageOfInterest)
	}

	barFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bar](
		barFormCallback.probe,
	)
	barFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if barFormCallback.CreationMode || barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(barFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BarFormCallback(
			nil,
			barFormCallback.probe,
			newFormGroup,
		)
		bar := new(models.Bar)
		FillUpForm(bar, newFormGroup, barFormCallback.probe)
		barFormCallback.probe.formStage.Commit()
	}

	fillUpTree(barFormCallback.probe)
}
