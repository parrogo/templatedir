package templatedir

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockFS(t *testing.T) {

	t.Run("StatFS", func(t *testing.T) {
		fsys := &FS{}

		expected, err := fs.Stat(fixtureFS, "dir1")
		require.NoError(t, err)
		require.NotNil(t, expected)

		fsys.On("Stat", "dir1").Return(expected, nil)

		// call the code we are testing
		info, err := fs.Stat(fsys, "dir1")
		require.NoError(t, err)
		require.NotNil(t, info)

		assert.Equal(t, expected, info)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("MkDir", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("MkDir", "dir1", os.FileMode(0)).Return(nil)

		// call the code we are testing
		err := fsys.MkDir("dir1", os.FileMode(0))
		require.NoError(t, err)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("Remove", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("Remove", "dir1").Return(nil)

		// call the code we are testing
		err := fsys.Remove("dir1")
		require.NoError(t, err)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("OpenFile", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("OpenFile", "dir1", 22, fs.FileMode(12)).Return(nil, nil)

		// call the code we are testing
		f, err := fsys.OpenFile("dir1", 22, fs.FileMode(12))
		require.NoError(t, err)
		require.Nil(t, f)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("ReadFile", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("ReadFile", "dir1").Return(nil, nil)

		// call the code we are testing
		bytes, err := fsys.ReadFile("dir1")
		require.NoError(t, err)
		require.Nil(t, bytes)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("Sub", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("Sub", "dir1").Return(nil, nil)

		// call the code we are testing
		bytes, err := fsys.Sub("dir1")
		require.NoError(t, err)
		require.Nil(t, bytes)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("Open", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("Open", "dir1").Return(nil, nil)

		// call the code we are testing
		f, err := fsys.Open("dir1")
		require.NoError(t, err)
		require.Nil(t, f)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("ReadDir", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("ReadDir", "dir1").Return(nil, nil)

		// call the code we are testing
		f, err := fsys.ReadDir("dir1")
		require.NoError(t, err)
		require.Nil(t, f)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

	t.Run("Glob", func(t *testing.T) {
		fsys := &FS{}

		fsys.On("Glob", "dir1").Return(nil, nil)

		// call the code we are testing
		f, err := fsys.Glob("dir1")
		require.NoError(t, err)
		require.Nil(t, f)

		// assert that the expectations were met
		fsys.AssertExpectations(t)
	})

}
