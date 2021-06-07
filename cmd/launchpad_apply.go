//
// Rover - Launchpad action sub command
// * This carries out actions like plan, apply or destroy via terrafrom
// * TODO: Work in progress
// * Ben C, May 2021
//

package cmd

import (
	"github.com/aztfmod/rover/pkg/landingzone"
	"github.com/spf13/cobra"
)

var lpApplyCmd = &cobra.Command{
	Use:   landingzone.ActionApply.Name(),
	Short: landingzone.ActionApply.Description(),
	Run: func(cmd *cobra.Command, args []string) {
		// Build config from command flags
		opt := landingzone.NewOptionsFromCmd(cmd)
		// And execute the relevant action
		opt.Execute(landingzone.ActionApply)
	},
}

func init() {
	// Place this command under the main `rover launchpad` command
	launchpadCmd.AddCommand(lpApplyCmd)
}
