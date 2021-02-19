package spec

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestParseAll(t *testing.T) {
	_, filePath, _, ok := runtime.Caller(0)
	if ! ok { t.Fatal(ok) }
	dir := filepath.Dir(filePath) + "/"

	paths, err := filepath.Glob(dir + "spec/*/*/*.txt")
	if err != nil { t.Fatal(err) }
	paths2, err := filepath.Glob(dir + "spec/*/*.txt")
	if err != nil { t.Fatal(err) }

	paths = append(paths, paths2...)

	for _, path := range paths {
		fmt.Printf(" > [%v]\n", path[len(dir):])

		b, err := ioutil.ReadFile(path)
		if err != nil { t.Fatal(err) }

		specBody := string(b)

		_, Err := Parse(specBody)
		if Err != nil {
			t.Fatal(Err, path)
		}
	}

	fmt.Printf("[%v] specifications tested\n", len(paths))
}

