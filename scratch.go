package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"github.com/dpastoor/goutils"
	"github.com/metrumresearchgroup/rcmdparser/rcmdparser"
	"github.com/spf13/afero"
	"strings"
)

func main() {
	pathPtr := flag.String("path", "./rcmdparser/testdata/", "full path to test folders, including trailing forward slash")
	filterPtr := flag.String("filter", "", "only run on folders in <path> that match filter string (default is empty)")
	verbosePtr := flag.Bool("verbose", true, "displays prettyPrint output from each folder")
	flag.Parse()

	path, err := ioutil.ReadDir(*pathPtr)
	dirs := goutils.ListDirNames(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range dirs {
		fullPath := filepath.Join(*pathPtr, f)
		if( len(*filterPtr) > 0 ){
			if( !strings.Contains(fullPath, *filterPtr)){
				continue;
			}
		}
		fmt.Println("parsing " + fullPath )
		fs := afero.NewOsFs()
		output, err := rcmdparser.NewCheck(fs, fullPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if *verbosePtr == true {
			prettyPrint(output)
		}
	}
}

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
