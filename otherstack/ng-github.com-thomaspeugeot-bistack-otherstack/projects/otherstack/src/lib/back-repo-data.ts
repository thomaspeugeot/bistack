// generated code - do not edit

//insertion point for imports
import { BarAPI } from './bar-api'


export class BackRepoData {
	// insertion point for declarations
	BarAPIs = new Array<BarAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.BarAPIs = data?.BarAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}