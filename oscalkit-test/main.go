package main

import (
	"flag"
	"os"

	"github.com/opencontrol/oscalkit"
)

func main() {

	var check = oscalkit.ApplicableControls
	profile := flag.String("p", "", "Path of the profile")
	flag.Parse()
	os.RemoveAll("./tmp/")
	SecurityControlsSubcontrolCheck(check, *profile)
}
