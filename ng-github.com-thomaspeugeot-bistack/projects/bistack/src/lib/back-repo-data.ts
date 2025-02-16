// generated code - do not edit

//insertion point for imports
import { FooAPI } from './foo-api'


export class BackRepoData {
	// insertion point for declarations
	FooAPIs = new Array<FooAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.FooAPIs = data?.FooAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}