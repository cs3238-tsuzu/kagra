package analyzer

import (
	"strings"
)

func parseAPIVersino(s string) (group, version string) {
	split := strings.SplitN(s, "/", 2)

	if len(split) > 0 {
		group = split[0]
	}
	if len(split) > 1 {
		version = split[1]
	}

	return
}
