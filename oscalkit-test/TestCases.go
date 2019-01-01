package main

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/opencontrol/oscalkit/types/oscal/catalog"
)

// SecurityControlsSubcontrolCheck is a test to verify that all controls from the catalog are being mapped correctly
func SecurityControlsSubcontrolCheck(check []catalog.Catalog, ProfileFile string) error {

	CodeGeneratedControls := ProtocolsMapping(check)

	f, err := os.Open(ProfileFile)
	if err != nil {
		log.Fatal(err)
	}

	parsedProfile, err := GetProfile(f)
	if err != nil {
		log.Fatal(err)
	}

	ProfileControlsDetails := ProfileProcessing(parsedProfile)

	if len(CodeGeneratedControls) == len(ProfileControlsDetails) {
		println("Perfect Count Match")
		println("Go file control, sub-control count: ", len(CodeGeneratedControls))
		println("Profile control, sub-control count: ", len(ProfileControlsDetails))
		CodeGeneratedMapping := ProtocolsMapping(check)
		mapcompareflag := AreMapsSame(ProfileControlsDetails, CodeGeneratedMapping)
		if mapcompareflag {
			color.Green("ID, Class & Title Mapping Correct")
		} else {
			color.Red("ID, Class & Title Mapping Incorrect")
		}
	} else if len(CodeGeneratedControls) > len(ProfileControlsDetails) {
		println("Controls in go file are greater in number then present in profile")
		println("Go file control, sub-control count: ", len(CodeGeneratedControls))
		println("Profile control, sub-control count: ", len(ProfileControlsDetails))
		color.Red("ID, Class & Title Mapping Incorrect")
	} else if len(CodeGeneratedControls) < len(ProfileControlsDetails) {
		println("Controls in profile are greater in number then present in go file")
		println("Go file control, sub-control count: ", len(CodeGeneratedControls))
		println("Profile control, sub-control count: ", len(ProfileControlsDetails))
		color.Red("ID, Class & Title Mapping Incorrect")
	}
	os.RemoveAll("./tmp/")
	return nil
}
