package tools

import (
	"fmt"
	"github.com/warrn/goinit/scripts"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func GimmeKnown() ([]string, error) {
	return GimmeList("known")
}

func GimmeInstalled() ([]string, error) {
	return GimmeList("list")
}

func GimmeList(list string) ([]string, error) {
	dir, err := ioutil.TempDir("", "goinit-scripts")
	if err != nil {
		return []string{}, err
	}
	defer os.RemoveAll(dir)

	if err := scripts.RestoreAsset(dir, "scripts/gimme/gimme"); err != nil {
		return []string{}, err
	}

	c := exec.Command(dir+"/scripts/gimme/gimme", list)

	output, err := c.Output()
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}

	vtext := strings.TrimRight(string(output), "\n")
	versions := strings.Split(vtext, "\n")
	fmt.Println("Size: ", len(versions))
	return versions, nil
}
