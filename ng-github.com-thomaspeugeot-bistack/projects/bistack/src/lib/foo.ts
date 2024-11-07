// generated code - do not edit

import { FooAPI } from './foo-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Foo {

	static GONGSTRUCT_NAME = "Foo"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyFooToFooAPI(foo: Foo, fooAPI: FooAPI) {

	fooAPI.CreatedAt = foo.CreatedAt
	fooAPI.DeletedAt = foo.DeletedAt
	fooAPI.ID = foo.ID

	// insertion point for basic fields copy operations
	fooAPI.Name = foo.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyFooAPIToFoo update basic, pointers and slice of pointers fields of foo
// from respectively the basic fields and encoded fields of pointers and slices of pointers of fooAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyFooAPIToFoo(fooAPI: FooAPI, foo: Foo, frontRepo: FrontRepo) {

	foo.CreatedAt = fooAPI.CreatedAt
	foo.DeletedAt = fooAPI.DeletedAt
	foo.ID = fooAPI.ID

	// insertion point for basic fields copy operations
	foo.Name = fooAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
