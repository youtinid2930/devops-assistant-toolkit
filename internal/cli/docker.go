package cli

import "github.com/spf13/cobra"

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related commands",
}

func init() {
	dockerCmd.AddCommand(initCmd)
}

