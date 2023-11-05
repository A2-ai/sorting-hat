package main

import (
	"fmt"
	"os"
	"time"

	"github.com/a2-ai/sorting-hat/cmd"
)

// https://goreleaser.com/cookbooks/using-main.version
var (
	version string = fmt.Sprintf("dev-%v", time.Now().Unix())
)

func main() {
	// normalize the config home on osx to linux and get rid of the path spacing
	cmd.Execute(version, os.Args[1:])
}
