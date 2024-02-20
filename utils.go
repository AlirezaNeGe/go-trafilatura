// This file is part of go-trafilatura, Go package for extracting readable
// content, comments and metadata from a web page. Source available in
// <https://github.com/AlirezaNeGe/go-trafilatura>.
// Copyright (C) 2021 Markus Mobius
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
// for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program. If not, see <https://www.gnu.org/licenses/>.

package trafilatura

import (
	"mime"
	"path/filepath"
	"strings"
)

// trim removes unnecessary spaces within a text string.
func trim(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	return strings.TrimSpace(s)
}

func strWordCount(s string) int {
	return len(strings.Fields(s))
}

func strOr(args ...string) string {
	for i := 0; i < len(args); i++ {
		if args[i] != "" {
			return args[i]
		}
	}
	return ""
}

func strIn(s string, args ...string) bool {
	for i := 0; i < len(args); i++ {
		if args[i] == s {
			return true
		}
	}
	return false
}

func getRune(s string, idx int) rune {
	for i, r := range s {
		if i == idx {
			return r
		}
	}

	return -1
}

func isImageFile(imageSrc string) bool {
	if imageSrc == "" {
		return false
	}

	ext := filepath.Ext(imageSrc)
	mimeType := mime.TypeByExtension(ext)
	return strings.HasPrefix(mimeType, "image")
}

func uniquifyLists(currents ...string) []string {
	var finalTags []string
	tracker := map[string]struct{}{}

	for _, current := range currents {
		separator := ","
		if strings.Count(current, ";") > strings.Count(current, ",") {
			separator = ";"
		}

		for _, entry := range strings.Split(current, separator) {
			entry = trim(entry)
			entry = strings.ReplaceAll(entry, `"`, "")
			entry = strings.ReplaceAll(entry, `'`, "")

			if _, tracked := tracker[entry]; entry != "" && !tracked {
				finalTags = append(finalTags, entry)
				tracker[entry] = struct{}{}
			}
		}
	}

	return finalTags
}
