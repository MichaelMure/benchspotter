package commands

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"

	"github.com/MichaelMure/benchspotter/commands/execenv"
)

func newVersionCommand(env *execenv.Env) *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print version information",
		Example: RootCmdName + " version",
		Long: `
Print version information.

Format:
  ` + RootCmdName + ` <version> [commit[/dirty]] <compiler version> <platform> <arch>

Format Description:
  <version> may be one of:
  	- A semantic version string, prefixed with a "v", e.g. v1.2.3
  	- "undefined" (if not provided, or built with an invalid version string)

  [commit], if present, is the commit hash that was checked out during the
  build. This may be suffixed with '/dirty' if there were local file
  modifications. This is indicative of your build being patched, or modified in
  some way from the commit.

  <compiler version> is the version of the go compiler used for the build.

  <platform> is the target platform (GOOS).

  <arch> is the target architecture (GOARCH).
`,
		Run: func(cmd *cobra.Command, args []string) {
			version, err := getVersion()
			if err != nil {
				env.Err.Printf("failed to get version: %v\n", err)
			} else {
				env.Out.Printf("%s %s\n", RootCmdName, version)
			}
		},
	}
}

var version = "undefined"

// getVersion returns a string representing the version information defined when
// the binary was built, or a sane default indicating a local build. a string is
// always returned. an error may be returned along with the string in the event
// that we detect a local build but are unable to get build metadata.
func getVersion() (string, error) {
	var arch string
	var commit string
	var modified bool
	var platform string

	var v strings.Builder

	// automatically add the v prefix if it's missing
	if version != "undefined" && !strings.HasPrefix(version, "v") {
		version = fmt.Sprintf("v%s", version)
	}

	// reset the version string to undefined if it is invalid
	if ok := semver.IsValid(version); !ok {
		version = "undefined"
	}

	v.WriteString(version)

	info, ok := debug.ReadBuildInfo()
	if !ok {
		v.WriteString(fmt.Sprintf(" (no build info)\n"))
		return v.String(), errors.New("unable to read build metadata")
	}

	for _, kv := range info.Settings {
		switch kv.Key {
		case "GOOS":
			platform = kv.Value
		case "GOARCH":
			arch = kv.Value
		case "vcs.modified":
			if kv.Value == "true" {
				modified = true
			}
		case "vcs.revision":
			commit = kv.Value
		}
	}

	if commit != "" {
		v.WriteString(fmt.Sprintf(" %.12s", commit))
	}

	if modified {
		v.WriteString("/dirty")
	}

	v.WriteString(fmt.Sprintf(" %s", info.GoVersion))

	if platform != "" {
		v.WriteString(fmt.Sprintf(" %s", platform))
	}

	if arch != "" {
		v.WriteString(fmt.Sprintf(" %s", arch))
	}

	return v.String(), nil
}
