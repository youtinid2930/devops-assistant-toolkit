package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Dockerfile",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🚀 Running Dockerfile wizard...")

		
		err := RunDockerWizard()
		if err != nil {
			return fmt.Errorf("failed to create Dockerfile: %w", err)
		}

		return nil
	},
}