// generated code - do not edit

import { BarAPI } from './bar-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Bar {

	static GONGSTRUCT_NAME = "Bar"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyBarToBarAPI(bar: Bar, barAPI: BarAPI) {

	barAPI.CreatedAt = bar.CreatedAt
	barAPI.DeletedAt = bar.DeletedAt
	barAPI.ID = bar.ID

	// insertion point for basic fields copy operations
	barAPI.Name = bar.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyBarAPIToBar update basic, pointers and slice of pointers fields of bar
// from respectively the basic fields and encoded fields of pointers and slices of pointers of barAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBarAPIToBar(barAPI: BarAPI, bar: Bar, frontRepo: FrontRepo) {

	bar.CreatedAt = barAPI.CreatedAt
	bar.DeletedAt = barAPI.DeletedAt
	bar.ID = barAPI.ID

	// insertion point for basic fields copy operations
	bar.Name = barAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
