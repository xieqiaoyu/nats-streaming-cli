// Package metadata
// param in this package can be set by -ldflags when building

package metadata

import "fmt"

var (
	//Version version string
	Version = "Unknown"
	//OS OS string
	OS = "?"
	//ARCH ARCH string
	ARCH = "?"
)

//GetVersionString return the version
func GetVersionString() string {
	return fmt.Sprintf("nets-streaming-cli %s (%s/%s)", Version, OS, ARCH)
}
