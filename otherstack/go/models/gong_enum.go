// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for StacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stacksnames StacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch stacksnames {
	// insertion code per enum code
	case Otherstack_Instance1:
		res = "otherstack instance 1"
	case Otherstack_Instance2:
		res = "otherstack instance 2"
	}
	return
}

func (stacksnames *StacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "otherstack instance 1":
		*stacksnames = Otherstack_Instance1
		return
	case "otherstack instance 2":
		*stacksnames = Otherstack_Instance2
		return
	default:
		return errUnkownEnum
	}
}

func (stacksnames *StacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Otherstack_Instance1":
		*stacksnames = Otherstack_Instance1
	case "Otherstack_Instance2":
		*stacksnames = Otherstack_Instance2
	default:
		return errUnkownEnum
	}
	return
}

func (stacksnames *StacksNames) ToCodeString() (res string) {

	switch *stacksnames {
	// insertion code per enum code
	case Otherstack_Instance1:
		res = "Otherstack_Instance1"
	case Otherstack_Instance2:
		res = "Otherstack_Instance2"
	}
	return
}

func (stacksnames StacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Otherstack_Instance1")
	res = append(res, "Otherstack_Instance2")

	return
}

func (stacksnames StacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "otherstack instance 1")
	res = append(res, "otherstack instance 2")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
