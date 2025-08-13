package version

// Version information that can be set at build time with ldflags
var (
	// Version is the current version of gTunnel
	Version = "dev"

	// GitCommit is the git commit hash
	GitCommit = "unknown"

	// BuildDate is when the binary was built
	BuildDate = "unknown"

	// GoVersion is the Go version used to build
	GoVersion = "unknown"
)

// Info represents version information
type Info struct {
	Version   string `json:"version"`
	GitCommit string `json:"git_commit"`
	BuildDate string `json:"build_date"`
	GoVersion string `json:"go_version"`
}

// Get returns the version information
func Get() Info {
	return Info{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		GoVersion: GoVersion,
	}
}

// GetVersion returns just the version string
func GetVersion() string {
	return Version
}
