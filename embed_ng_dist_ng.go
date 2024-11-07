package bistack

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-thomaspeugeot-bistack/dist/ng-github.com-thomaspeugeot-bistack
var NgDistNg embed.FS
