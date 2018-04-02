package tools

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"io/ioutil"
	"os"
	"testing"
)

func touchFile(r *require.Assertions, path string){
	f, err := os.Create(path)
	r.Nil(err)
	f.Close()
}

func TestCheckBinaryExistsInPath(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	dir, err := ioutil.TempDir("", "goinit-test")
	r.Nil(err)
	defer os.RemoveAll(dir)

	os.Setenv("PATH", dir)

	a.False(CheckBinaryExistsInPath("sh"))

	touchFile(r, dir+"/sh")

	a.True(CheckBinaryExistsInPath("sh"))
}

func TestCheckGimmeDependencies(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	dir, err := ioutil.TempDir("", "goinit-test")
	r.Nil(err)
	defer os.RemoveAll(dir)

	os.Setenv("PATH", dir)

	b, e := CheckGimmeDependencies()
	a.False(b)
	a.EqualValues(NoDownloaderError, e)

	touchFile(r, dir+"/wget")
	b, e = CheckGimmeDependencies()
	a.False(b)
	a.EqualValues(NoSHAError, e)

	touchFile(r, dir+"/sha256sum")
	b, e = CheckGimmeDependencies()
	a.False(b)
	a.EqualValues(NoGitError, e)

	touchFile(r, dir+"/git")
	b, e = CheckGimmeDependencies()
	a.False(b)
	a.EqualValues(NoTarError, e)

	touchFile(r, dir+"/tar")
	b, e = CheckGimmeDependencies()
	a.False(b)
	a.EqualValues(NoMakeError, e)

	touchFile(r, dir+"/make")
	b, e = CheckGimmeDependencies()
	a.True(b)
	a.Nil(e)
}
