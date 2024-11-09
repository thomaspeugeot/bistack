package otherstack

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-bistack-otherstack/dist/ng-github.com-thomaspeugeot-bistack-otherstack
var NgDistNg embed.FS
