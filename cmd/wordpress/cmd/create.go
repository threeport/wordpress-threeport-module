// generated by 'threeport-sdk gen' but will not be regenerated - intended for modification

package cmd

import cobra "github.com/spf13/cobra"

// CreateCmd represents the create command
var CreateCmd = &cobra.Command{
	Long:  "Create a Threeport Wordpress object.\n\n\tThe create command does nothing by itself.  Use one of the available subcommands\n\tto create different objects from the system.",
	Short: "Create a Threeport Wordpress object",
	Use:   "create",
}

func init() {
	rootCmd.AddCommand(CreateCmd)
}
