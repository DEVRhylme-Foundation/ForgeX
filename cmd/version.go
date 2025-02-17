/*
Go ForgeX version command
*/
package cmd

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/cobra"
)

// GoForgeXVersion is set by GoReleaser in CI with the release version.
var GoForgeXVersion string

// getGoForgeXVersion retrieves the version information.
func getGoForgeXVersion() string {
	const noVersionMsg = "No version info available for this build, run 'forgex help version' for additional info"

	// Check if GoForgeXVersion is set (injected at build time).
	if GoForgeXVersion != "" {
		return GoForgeXVersion
	}

	// Read Go build info.
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return noVersionMsg
	}

	// Return version from build info if available.
	if bi.Main.Version != "(devel)" {
		return bi.Main.Version
	}

	// Extract VCS revision and timestamp from build info.
	var vcsRevision string
	var vcsTime time.Time

	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		}
	}

	// Return VCS revision and timestamp if available.
	if vcsRevision != "" {
		return fmt.Sprintf("%s (%s)", vcsRevision, vcsTime)
	}

	return noVersionMsg
}

// versionCmd defines the "version" command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information",
	Long: `Displays the current version of Go ForgeX.

The version information is embedded at compile time. 
If unavailable, Go ForgeX attempts to retrieve it from the build environment.
If built within a version control repository, the revision hash will be used.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Go ForgeX CLI version: %v\n", getGoForgeXVersion())
	},
}
