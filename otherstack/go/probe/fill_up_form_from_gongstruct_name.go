// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/bistack/otherstack/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Bar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bar Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarFormCallback(
			nil,
			probe,
			formGroup,
		)
		bar := new(models.Bar)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bar, formGroup, probe)
	}
	formStage.Commit()
}
