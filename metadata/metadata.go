//Package metadata
// param in this package can be set by -ldflags when building
package metadata

import "fmt"

var (
	Version = "Unknown"
	OS      = "?"
	ARCH    = "?"
)

func GetVersionString() string {
	return fmt.Sprintf("nets-streaming-cli %s (%s/%s)", Version, OS, ARCH)
}
