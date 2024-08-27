// Contains code to filter out file paths that contain any of the prefixes in prefixesToFilter.
package files

import (
	"strings"
)

// Filter out file paths that contain any of the prefixes in prefixesToFilter.
func FilterFilePaths(filePaths []string, prefixesToFilter []string) []string {
	var retainedFilePaths []string
	for _, filePath := range filePaths {
		isPrefixFound := false
		for _, prefixToFilter := range prefixesToFilter {
			if strings.HasPrefix(filePath, prefixToFilter) {
				isPrefixFound = true
				break
			}
		}
		if !isPrefixFound {
			retainedFilePaths = append(retainedFilePaths, filePath)
		}
	}
	return retainedFilePaths
}
