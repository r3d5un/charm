package progressbar

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ProgressBarCmd = &cobra.Command{
	Use:     "progressbar",
	Aliases: []string{"prog", "pb"},
	Short:   "Displays a progressbar, then exits",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented")
	},
}
