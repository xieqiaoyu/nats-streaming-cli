//Package metadata
// param in this package can be set by -ldflags when building
package metadata

var (
	Version = "Unknown"
)

func GetVersionString() string {
	return Version
}
