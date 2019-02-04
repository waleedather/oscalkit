package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/docker/oscalkit/types/oscal"
	"github.com/docker/oscalkit/types/oscal/catalog"
	"github.com/docker/oscalkit/types/oscal/profile"
)

// ProtocolsMapping Method To Parse The generated .go file and save the
// mapping of ID, Class & Titles
func ProtocolsMapping(check []catalog.Catalog) map[string][]string {

	SecurityControls := make(map[string][]string)
	for CatalogCount := 0; CatalogCount < len(check); CatalogCount++ {
		for GroupsCount := 0; GroupsCount < len(check[CatalogCount].Groups); GroupsCount++ {
			for ControlsCount := 0; ControlsCount < len(check[CatalogCount].Groups[GroupsCount].Controls); ControlsCount++ {
				if _, ok := SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id]; ok {
				} else {
					SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id], check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Class)
					SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id], string(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Title))
				}

				for ControlPartCount := 0; ControlPartCount < len(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts); ControlPartCount++ {
					if _, ok := SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id]; ok {
					} else {
						if check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id != "" {
							SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id], check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Class)
							SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id], string(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Title))
						} else if check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id == "" && check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Class == "assessment" {
							SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id], check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Class)
							SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Id], string(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Parts[ControlPartCount].Title))
						}
					}
				}

				for SubControlsCount := 0; SubControlsCount < len(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols); SubControlsCount++ {
					if _, ok := SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id]; ok {
					} else {
						SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id], check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Class)
						SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id], string(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Title))
					}
					for SubControlsPartCount := 0; SubControlsPartCount < len(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts); SubControlsPartCount++ {
						if _, ok := SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id]; ok {
						} else {
							if check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id != "" {
								SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id], check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Class)
								SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id], string(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Title))
							} else if check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id == "" && check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Class == "assessment" {
								SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id], check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Class)
								SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id] = append(SecurityControls[check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Id+"?"+check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Id], string(check[CatalogCount].Groups[GroupsCount].Controls[ControlsCount].Subcontrols[SubControlsCount].Parts[SubControlsPartCount].Title))
							}
						}
					}
				}
			}
		}
	}
	return SecurityControls
}

// GetCatalog gets a catalog
func GetCatalog(r io.Reader) (*catalog.Catalog, error) {
	o, err := oscal.New(r)
	if err != nil {
		return nil, err
	}
	if o.Catalog == nil {
		return nil, fmt.Errorf("cannot map profile")
	}
	return o.Catalog, nil
}

// GetProfile gets a profile
func GetProfile(r io.Reader) (*profile.Profile, error) {
	o, err := oscal.New(r)
	if err != nil {
		return nil, err
	}
	if o.Profile == nil {
		return nil, fmt.Errorf("cannot map profile")
	}
	return o.Profile, nil
}

// controlInProfile accepts a Control or SubcontrolID and an array of all
// the controls & subcontrols present in the profile.
func controlInProfile(controlID string, profile []string) bool {
	for _, value := range profile {
		if value == controlID {
			return true
		}
	}
	return false
}

// ParentControlCheck checks if the subcontrol's parent controls exists
// in the provided array on parent controls
func ParentControlCheck(subcontrol string, parentcontrols []string) bool {

	subcontroltrim := strings.Split(subcontrol, ".")

	for _, value := range parentcontrols {
		if value == subcontroltrim[0] {
			return true
		}
	}
	return false
}

// DownloadCatalog writes the JSON of the provided URL into a catalog.json file
func DownloadCatalog(url string) (string, error) {
	URLSplit := strings.Split(url, "/")
	tmpDir, err := ioutil.TempDir(".", "oscaltesttmp")
	if err != nil {
		log.Fatal(err)
	}
	filename := tmpDir + "/" + URLSplit[len(URLSplit)-1]
	println("Catalog will be downloaded to: " + filename)
	catalog, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer catalog.Close()
	println("Downloading catalog from URL: " + url)
	data, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer data.Body.Close()
	_, err = io.Copy(catalog, data.Body)
	if err != nil {
		return "", err
	}
	return tmpDir, nil
}

// ProfileParsing method to parse the profile and return the controls and subcontrols ID's
func ProfileParsing(parsedProfile *profile.Profile) []string {

	SecurityControls := make([]string, 0)

	for ImportCount := 0; ImportCount < len(parsedProfile.Imports); ImportCount++ {
		for IDSelectorCount := 0; IDSelectorCount < len(parsedProfile.Imports[ImportCount].Include.IdSelectors); IDSelectorCount++ {
			if parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].ControlId != "" {
				SecurityControls = append(SecurityControls, parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].ControlId)
			}
			if parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].SubcontrolId != "" {
				SecurityControls = append(SecurityControls, parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].SubcontrolId)
			}
		}
	}
	return SecurityControls
}

// ParentControls to get the list of all parent controls in the profile
func ParentControls(parsedProfile *profile.Profile) []string {
	ParentControlsList := make([]string, 0)

	for ImportCount := 0; ImportCount < len(parsedProfile.Imports); ImportCount++ {
		temp := ParseImport(parsedProfile, parsedProfile.Imports[ImportCount].Href.Path, "Parent")
		ParentControlsList = appendslice(ParentControlsList, temp)
	}

	ParentControlsList = unique(ParentControlsList)

	return ParentControlsList
}

// ProfileProcessing is used to generate the mapping of ID Class & Title of
// all the controls subcontrols and parts
func ProfileProcessing(parsedProfile *profile.Profile, ListParentControls []string) map[string][]string {
	SecurityControlsDetails := make(map[string][]string)

	for ImportCounts := 0; ImportCounts < len(parsedProfile.Imports); ImportCounts++ {
		println("Import:", parsedProfile.Imports[ImportCounts].Href.String())
		dirName := "test_util/artifacts/"
		var err error
		if strings.Contains(parsedProfile.Imports[ImportCounts].Href.String(), "http") {
			dirName, err = DownloadCatalog(parsedProfile.Imports[ImportCounts].Href.String())
			if err != nil {
				log.Fatal(err)
			}
		}
		URLSplit := strings.Split(parsedProfile.Imports[ImportCounts].Href.Path, "/")
		filename := dirName + "/" + URLSplit[len(URLSplit)-1]
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		check, _ := ProfileCatalogCheck(f)
		if check == "Catalog" {

			ProfileControls := ParseImport(parsedProfile, parsedProfile.Imports[ImportCounts].Href.Path, "all")

			catalogPath := dirName + "/" + URLSplit[len(URLSplit)-1]
			f, err := os.Open(catalogPath)
			if err != nil {
				log.Fatal(err)
			}

			parsedCatalog, err := GetCatalog(f)
			if err != nil {
				log.Fatal(err)
			}

			CatalogControlsDetails := ParseCatalog(parsedCatalog, ProfileControls, ListParentControls)

			PartsProfileControls := ProfileParsing(parsedProfile)

			Parts := ParseParts(parsedProfile, PartsProfileControls)

			CatalogControlsDetails = appendAlterations(CatalogControlsDetails, Parts)

			println("Size of Catalog: ", len(CatalogControlsDetails))
			if len(SecurityControlsDetails) == 0 {
				SecurityControlsDetails = appendMaps(SecurityControlsDetails, CatalogControlsDetails)
			} else if len(SecurityControlsDetails) > 0 {
				SecurityControlsDetails = appendMaps(SecurityControlsDetails, CatalogControlsDetails)
				SecurityControlsDetails = uniqueMaps(SecurityControlsDetails, CatalogControlsDetails)
			}
			println("Size of SecurityControls: ", len(SecurityControlsDetails))

		} else if check == "Profile" {

			fmt.Println("profile path: " + URLSplit[len(URLSplit)-1])
			f, err := os.Open(dirName + "/" + URLSplit[len(URLSplit)-1])
			if err != nil {
				log.Fatal(err)
			}

			ProfileHref, err := GetProfile(f)
			if err != nil {
				log.Fatal(err)
			}

			ParsedProfile := ProfileProcessing(ProfileHref, ListParentControls)
			ParsedProfileControls := ParseImport(parsedProfile, parsedProfile.Imports[ImportCounts].Href.Path, "all")

			PartsProfileControls := ProfileParsing(parsedProfile)

			Parts := ParseParts(parsedProfile, PartsProfileControls)

			println("Recursive count = ", len(ParsedProfile))
			println("Count of profile = ", len(ParsedProfileControls))

			println("Common = ", len(CommonMap(ParsedProfileControls, ParsedProfile)))
			SecurityControlsDetails = appendMaps(SecurityControlsDetails, CommonMap(ParsedProfileControls, ParsedProfile))

			SecurityControlsDetails = appendAlterations(SecurityControlsDetails, Parts)

			println("Final Count = ", len(SecurityControlsDetails))
		}
	}

	return SecurityControlsDetails
}

// ParseCatalog accepts a catalog struct and return the mapping of Control,
// Subcontrols & Parts. ID, Class & Titles
func ParseCatalog(parsedCatalog *catalog.Catalog, ProfileControls []string, ListParentControls []string) map[string][]string {
	CatalogControlsDetails := make(map[string][]string)

	for GroupCount := 0; GroupCount < len(parsedCatalog.Groups); GroupCount++ {
		for ControlCount := 0; ControlCount < len(parsedCatalog.Groups[GroupCount].Controls); ControlCount++ {
			if controlInProfile(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id, ProfileControls) {
				CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id], parsedCatalog.Groups[GroupCount].Controls[ControlCount].Class)
				CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id], string(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Title))
				for ControlPartCount := 0; ControlPartCount < len(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts); ControlPartCount++ {
					if parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id != "" {
						CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id], parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Class)
						CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id], string(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Title))
					} else if parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id == "" && parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Class == "assessment" {
						CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id], parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Class)
						CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Id], string(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Parts[ControlPartCount].Title))
					}
				}
			}

			for SubControlCount := 0; SubControlCount < len(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols); SubControlCount++ {
				if controlInProfile(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id, ProfileControls) && ParentControlCheck(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id, ListParentControls) {
					CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id], parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Class)
					CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id], string(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Title))
					for SubControlPartCount := 0; SubControlPartCount < len(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts); SubControlPartCount++ {
						if parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id != "" {
							CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id], parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Class)
							CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id], string(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Title))
						} else if parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id == "" && parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Class == "assessment" {

							CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id], parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Class)
							CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id] = append(CatalogControlsDetails[parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Id+"?"+parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Id], string(parsedCatalog.Groups[GroupCount].Controls[ControlCount].Subcontrols[SubControlCount].Parts[SubControlPartCount].Title))
						}
					}
				}
			}
		}
	}
	return CatalogControlsDetails
}

// ProfileCatalogCheck checks if the path provided is for a profile or a catolog
func ProfileCatalogCheck(r io.Reader) (string, error) {
	o, err := oscal.New(r)
	if err != nil {
		return "Invalid File", err
	}
	if o.Profile == nil {
		return "Catalog", nil
	}
	if o.Catalog == nil {
		return "Profile", nil
	}
	return "Invalid File", nil
}

// CommonMap returns the elements in Map that are also present in profile
func CommonMap(profile []string, CatalogControlsDetails map[string][]string) map[string][]string {

	CommonMapping := make(map[string][]string)

	for key, mapvalue := range CatalogControlsDetails {
		for _, slicevalue := range profile {
			subcontroltrim := strings.Split(key, "?")

			if slicevalue == key {
				CommonMapping[key] = append(CommonMapping[key], mapvalue[0])
				CommonMapping[key] = append(CommonMapping[key], mapvalue[1])
			} else if slicevalue == subcontroltrim[0] {
				CommonMapping[key] = append(CommonMapping[key], mapvalue[0])
				CommonMapping[key] = append(CommonMapping[key], mapvalue[1])
			}
		}
	}
	return CommonMapping
}

// ParseImport method to parse the profile and return the controls and subcontrols or only controls
func ParseImport(parsedProfile *profile.Profile, link string, token string) []string {

	SecurityControls := make([]string, 0)
	for ImportCount := 0; ImportCount < len(parsedProfile.Imports); ImportCount++ {
		if parsedProfile.Imports[ImportCount].Href.Path == link {
			for IDSelectorCount := 0; IDSelectorCount < len(parsedProfile.Imports[ImportCount].Include.IdSelectors); IDSelectorCount++ {
				if parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].ControlId != "" {
					SecurityControls = append(SecurityControls, parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].ControlId)
				}
				if parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].SubcontrolId != "" && token != "Parent" {
					SecurityControls = append(SecurityControls, parsedProfile.Imports[ImportCount].Include.IdSelectors[IDSelectorCount].SubcontrolId)
				}
			}
		}
	}

	return SecurityControls
}

// ParseParts method to parse the profile and return the mapping of all the parts
func ParseParts(parsedProfile *profile.Profile, list []string) map[string][]string {

	SecurityControls := make(map[string][]string)

	for ModifyCount := 0; ModifyCount < len(parsedProfile.Modify.Alterations); ModifyCount++ {
		for AlterCount := 0; AlterCount < len(parsedProfile.Modify.Alterations[ModifyCount].Additions); AlterCount++ {
			for PartCount := 0; PartCount < len(parsedProfile.Modify.Alterations[ModifyCount].Additions[AlterCount].Parts); PartCount++ {
				for _, s1element := range list {
					if parsedProfile.Modify.Alterations[ModifyCount].ControlId == s1element {
						if parsedProfile.Modify.Alterations[ModifyCount].ControlId != "" && parsedProfile.Modify.Alterations[ModifyCount].Additions[AlterCount].Parts[PartCount].Class == "guidance" {
							SecurityControls[parsedProfile.Modify.Alterations[ModifyCount].ControlId+"?"+parsedProfile.Modify.Alterations[ModifyCount].ControlId+"_gdn"] = append(SecurityControls[parsedProfile.Modify.Alterations[ModifyCount].ControlId+"?"+parsedProfile.Modify.Alterations[ModifyCount].ControlId+"_gdn"], parsedProfile.Modify.Alterations[ModifyCount].Additions[AlterCount].Parts[PartCount].Class)
						}
					} else if parsedProfile.Modify.Alterations[ModifyCount].SubcontrolId == s1element {
						if parsedProfile.Modify.Alterations[ModifyCount].SubcontrolId != "" && parsedProfile.Modify.Alterations[ModifyCount].Additions[AlterCount].Parts[PartCount].Class == "guidance" {
							SecurityControls[parsedProfile.Modify.Alterations[ModifyCount].SubcontrolId+"?"+parsedProfile.Modify.Alterations[ModifyCount].SubcontrolId+"_gdn"] = append(SecurityControls[parsedProfile.Modify.Alterations[ModifyCount].SubcontrolId+"?"+parsedProfile.Modify.Alterations[ModifyCount].SubcontrolId+"_gdn"], parsedProfile.Modify.Alterations[ModifyCount].Additions[AlterCount].Parts[PartCount].Class)
						}
					}
				}
			}
		}
	}

	return SecurityControls
}

// appendslice appends two slices
func appendslice(slice []string, slice1 []string) []string {

	for sliceCount := 0; sliceCount < len(slice1); sliceCount++ {
		slice = append(slice, slice1[sliceCount])
	}

	return slice
}

// AreMapsSame compares the values of two  same length maps and returns true if both the maps have the same key value pairs
func AreMapsSame(profileControlsDetails map[string][]string, codeGeneratedMapping map[string][]string, token string) bool {
	for key := range profileControlsDetails {
		if !strings.Contains(key, "?") && token == "controls" {
			if profileControlsDetails[key][0] != codeGeneratedMapping[key][0] && profileControlsDetails[key][1] != codeGeneratedMapping[key][1] {
				println("Mapping for " + key + " incorrect.")
				return false
			}
			// if !reflect.DeepEqual(profileControlsDetails[key], codeGeneratedMapping[key]) {
			// 	println("Mapping for " + key + " incorrect.")
			// 	return false
			// }
		} else if strings.Contains(key, "?") && token == "parts" {
			if profileControlsDetails[key][0] != codeGeneratedMapping[key][0] && profileControlsDetails[key][1] != codeGeneratedMapping[key][1] {
				println("Mapping for " + key + " incorrect.")
				return false
			}
			// if !reflect.DeepEqual(profileControlsDetails[key], codeGeneratedMapping[key]) {
			// 	println("Mapping for " + key + " incorrect.")
			// 	return false
			// }
		}
	}
	return true
}

// unique returns unique values in the slice
func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// appendMaps appends two maps
func appendMaps(SecurityControlsDetails map[string][]string, CatalogControlsDetails map[string][]string) map[string][]string {

	for key, value := range CatalogControlsDetails {
		SecurityControlsDetails[key] = value
	}

	return SecurityControlsDetails
}

func appendAlterations(SecurityControlsDetails map[string][]string, PartsDetails map[string][]string) map[string][]string {

	for key, value := range PartsDetails {
		if _, ok := SecurityControlsDetails[key]; ok {
			delete(SecurityControlsDetails, key)
			SecurityControlsDetails[key+"_1"] = value
			SecurityControlsDetails[key+"_2"] = value
		}
	}

	return SecurityControlsDetails
}

func uniqueMaps(SecurityControlsDetails map[string][]string, CatalogControlsDetails map[string][]string) map[string][]string {

	for key, value := range CatalogControlsDetails {
		if _, ok := SecurityControlsDetails[key]; !ok {
			SecurityControlsDetails[key] = value
		}
	}

	return SecurityControlsDetails
}

// Count to take count of either parts of controls & subcontrols
func Count(SecurityControlsDetails map[string][]string, token string) int {

	count := 0

	for key := range SecurityControlsDetails {
		if token == "parts" {
			count++
		} else if token == "controls" {
			if !strings.Contains(key, "?") {
				count++
			}
		}
	}

	return count
}
