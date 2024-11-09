// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	FooAPIs []*FooAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, fooDB := range backRepo.BackRepoFoo.Map_FooDBID_FooDB {

		var fooAPI FooAPI
		fooAPI.ID = fooDB.ID
		fooAPI.FooPointersEncoding = fooDB.FooPointersEncoding
		fooDB.CopyBasicFieldsToFoo_WOP(&fooAPI.Foo_WOP)

		backRepoData.FooAPIs = append(backRepoData.FooAPIs, &fooAPI)
	}

}
