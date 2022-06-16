package versions

import (
	"VersionCraft2/internal"
	"regexp"
)

var ListArrangement = map[string][]string{
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

type Menu map[string]Versions

func SplitVersionsByMenu(versions *Versions) (menus Menu, order []string) {
	menus = make(Menu)

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

	order = make([]string, 0, len(ListArrangement))
	for k := range ListArrangement {
		order = append(order, k)
	}

	return
}
