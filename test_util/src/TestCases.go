package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/docker/oscalkit/types/oscal/catalog"
	"github.com/fatih/color"
)

// SecurityControlsSubcontrolCheck is a test to verify that all controls from the catalog are being mapped correctly
func SecurityControlsSubcontrolCheck(check []catalog.Catalog, ProfileFile string) error {

	codeGeneratedControls := ProtocolsMapping(check)

	f, err := os.Open(ProfileFile)
	if err != nil {
		log.Fatal(err)
	}

	parsedProfile, err := GetProfile(f)
	if err != nil {
		log.Fatal(err)
	}

	ListParentControls := ParentControls(parsedProfile)

	profileControlsDetails := ProfileProcessing(parsedProfile, ListParentControls)

	if Count(codeGeneratedControls, "controls") == Count(profileControlsDetails, "controls") {
		color.Green("Controls & SubControls Count Matched")
		println("Go file control & sub-control count: ", Count(codeGeneratedControls, "controls"))
		println("Profile control & sub-control count: ", Count(profileControlsDetails, "controls"))
	} else if Count(codeGeneratedControls, "controls") > Count(profileControlsDetails, "controls") {
		color.Red("Controls & Subcontrols in go file are greater in number then present in profile")
		println("Go file control & sub-control count: ", Count(codeGeneratedControls, "controls"))
		println("Profile control & sub-control count: ", Count(profileControlsDetails, "controls"))
	} else if Count(codeGeneratedControls, "controls") < Count(profileControlsDetails, "controls") {
		color.Red("Controls & Subcontrols in profile are greater in number then present in go file")
		println("Go file control & sub-control count: ", Count(codeGeneratedControls, "controls"))
		println("Profile control & sub-control count: ", Count(profileControlsDetails, "controls"))
	}

	controlmapcompareflag := AreMapsSame(profileControlsDetails, codeGeneratedControls, "controls")
	if controlmapcompareflag {
		color.Green("ID, Class & Title Mapping Of All Controls & SubControls Correct")
	} else {
		color.Red("ID, Class & Title Mapping Of All Controls & SubControls Incorrect")
	}

	if Count(codeGeneratedControls, "parts") == Count(profileControlsDetails, "parts") {
		color.Green("Parts Count Matched")
		println("Go file parts count: ", Count(codeGeneratedControls, "parts"))
		println("Profile parts count: ", Count(profileControlsDetails, "parts"))
	} else if Count(codeGeneratedControls, "parts") > Count(profileControlsDetails, "parts") {
		color.Red("Parts in go file are greater in number then present in profile")
		println("Go file parts count: ", Count(codeGeneratedControls, "parts"))
		println("Profile parts count: ", Count(profileControlsDetails, "parts"))
	} else if Count(codeGeneratedControls, "parts") < Count(profileControlsDetails, "parts") {
		color.Red("Parts in profile are greater in number then present in go file")
		println("Go file parts count: ", Count(codeGeneratedControls, "parts"))
		println("Profile parts count: ", Count(profileControlsDetails, "parts"))
	}

	partsmapcompareflag := AreMapsSame(profileControlsDetails, codeGeneratedControls, "parts")
	if partsmapcompareflag {
		color.Green("ID, Class & Title Mapping Of Parts Correct")
	} else {
		color.Red("ID, Class & Title Mapping Of All Parts Incorrect")
	}

	file, _ := filepath.Glob("./oscaltesttmp*")
	if file != nil {
		for _, f := range file {
			os.RemoveAll(f)
		}
	}
	return nil
}
