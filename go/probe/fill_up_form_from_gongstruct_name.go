// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/thomaspeugeot/bistack/go/models"
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
	case "Foo":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Foo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FooFormCallback(
			nil,
			probe,
			formGroup,
		)
		foo := new(models.Foo)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(foo, formGroup, probe)
	}
	formStage.Commit()
}
