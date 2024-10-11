package cli

import (
	"fmt"
	"os"

	"github.com/r3d5un/charm/cmd/charming/cli/progressbar"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:     "charming",
	Short:   "CLI tools for learning the charmbracelet stack",
	Version: "v0.0.1",
}

func init() {
	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
	rootCmd.AddCommand(progressbar.ProgressBarCmd)
}
