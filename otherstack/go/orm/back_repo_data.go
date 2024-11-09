// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	BarAPIs []*BarAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {

	// wait till backRepo is written by commit
	backRepo.rwMutex.RLock()
	defer backRepo.rwMutex.RUnlock()

	// insertion point for slices copies
	for _, barDB := range backRepo.BackRepoBar.Map_BarDBID_BarDB {

		var barAPI BarAPI
		barAPI.ID = barDB.ID
		barAPI.BarPointersEncoding = barDB.BarPointersEncoding
		barDB.CopyBasicFieldsToBar_WOP(&barAPI.Bar_WOP)

		backRepoData.BarAPIs = append(backRepoData.BarAPIs, &barAPI)
	}

}
