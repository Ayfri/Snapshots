package versions

import (
	"VersionCraft2/internal"
	"regexp"
)

type versionList map[string][]string

var ListArrangement = versionList{
	"Pre-Alpha": {
		"pre-Classic", "Classic", "Indev", "Infdev",
	},
	"Alpha & Beta": {
		"Alpha", "Beta",
	},
	"Releases": {
		"release",
	},
	"Others": {
		"Combat Tests", "April Fools",
	},
}

func SplitVersionsByMenu(versions *Versions) map[string]Versions {
	menus := make(map[string]Versions)

	for name, menu := range ListArrangement {
		var menuVersions []Version
		for _, version := range *versions {
			if internal.Any(menu, func(a string) bool {
				return regexp.MustCompile(a).MatchString(version.Name)
			}) {
				menuVersions = append(menuVersions, version)
			}

			if menu[0] == "release" && regexp.MustCompile("^Java Edition 1\\.\\d{1,2}$").MatchString(version.Name) {
				menuVersions = append(menuVersions, version)
			}
		}

		menus[name] = menuVersions
	}

	return menus
}
