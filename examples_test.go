package templatedir_test

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/parrogo/templatedir"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

// This example show how to use templatedir.Func()
func ExampleDefaultArgs() {
	args, err := templatedir.DefaultArgs()
	if err != nil {
		panic(err)
	}
	fmt.Println(args["RepoName"])
	fmt.Println(args["Author"])
	// Output: templatedir
	// parrogo
}
