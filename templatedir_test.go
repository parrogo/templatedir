package templatedir

import (
	"embed"
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/parro-it/vs/memfs"
	"github.com/parro-it/vs/syncfs"
	"github.com/parro-it/vs/writefs"
	"github.com/stretchr/testify/assert"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

func TestWalkDir(t *testing.T) {
	res, errs := walkDir(fixtureFS)
	var actual []string
	for s := range res {
		actual = append(actual, s)
	}
	err := <-errs
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"dir1/dir2/file3.txt.template",
		"dir1/dir3/file4.template",
		"dir1/vars/test.template",
		"{{.Count}}.template",
	}, actual)
}

var args = map[string]int{
	"Count": 42,
}

func TestRenderFile(t *testing.T) {
	r := renderer{
		srcfs:  fixtureFS,
		destfs: syncfs.New(memfs.NewFS()).(writefs.WriteFS),
		args:   args,
	}
	err := r.renderFile("dir1/dir2/file3.txt.template")
	assert.NoError(t, err)

	actual, err := fs.ReadFile(r.destfs, "dir1/dir2/file3.txt")
	assert.NoError(t, err)

	assert.Equal(t, "you pass 42.", string(actual))
}

func TestRenderFileName(t *testing.T) {
	r := renderer{
		srcfs:  fixtureFS,
		destfs: syncfs.New(memfs.NewFS()).(writefs.WriteFS),
		args:   args,
	}
	err := r.renderFile("{{.Count}}.template")
	if !assert.NoError(t, err) {
		return
	}

	actual, err := fs.ReadFile(r.destfs, "42")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "answer is 42", string(actual))

	_, err = fs.Stat(r.destfs, "{{.Count}}.template")
	assert.True(t, errors.Is(err, fs.ErrNotExist))

}

func TestRenderTo(t *testing.T) {

	outfs := memfs.NewFS()
	err := RenderTo(fixtureFS, outfs, args)
	assert.NoError(t, err)

	actual, err := fs.ReadFile(outfs, "dir1/dir2/file3.txt")
	assert.NoError(t, err)
	assert.Equal(t, "you pass 42.", string(actual))

	actual, err = fs.ReadFile(outfs, "dir1/dir3/file4")
	assert.NoError(t, err)
	assert.Equal(t, "another 42.", string(actual))
}

func TestTemplateFilesRemovedFromDest(t *testing.T) {
	outfs := memfs.NewFS()
	err := writefs.MkDir(outfs, "dir1", fs.FileMode(0755))
	assert.NoError(t, err)
	err = writefs.MkDir(outfs, "dir1/dir2", fs.FileMode(0755))
	assert.NoError(t, err)
	_, err = writefs.WriteFile(outfs, "dir1/dir2/file3.txt.template", []byte{0x42})
	assert.NoError(t, err)

	err = RenderTo(fixtureFS, outfs, args)
	assert.NoError(t, err)

	actual, err := fs.ReadFile(outfs, "dir1/dir2/file3.txt")
	assert.NoError(t, err)
	assert.Equal(t, "you pass 42.", string(actual))

	_, err = fs.ReadFile(outfs, "dir1/dir2/file3.txt.template")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, fs.ErrNotExist))

}

func TestVars(t *testing.T) {
	outfs := memfs.NewFS()
	err := writefs.MkDir(outfs, "dir1", fs.FileMode(0755))
	assert.NoError(t, err)
	err = writefs.MkDir(outfs, "dir1/vars", fs.FileMode(0755))
	assert.NoError(t, err)

	err = os.Setenv("GITHUB_REPOSITORY", "parrogo/templatedir")
	assert.NoError(t, err)

	err = os.Setenv("GITHUB_WORKSPACE", "/root")
	assert.NoError(t, err)

	a, err := DefaultArgs()
	if !assert.NoError(t, err) {
		return
	}
	err = RenderTo(fixtureFS, outfs, a)
	assert.NoError(t, err)

	actual, err := fs.ReadFile(outfs, "dir1/vars/test")
	assert.NoError(t, err)

	assert.Equal(t, `Author is parrogo
This repository is named templatedir
Local root of repository is /root
`, string(actual))
}
