package main

import (
	"github.com/jamesruan/golf"
	"github.com/spf13/cobra"
)

var logger = golf.DefaultEntry

var rootCmd = &cobra.Command{
	Use:   "confgen",
	Short: "Config Generator",
}

func initGen() *cobra.Command {
	genCmd := &cobra.Command{
		Use:     "gen",
		Short:   "generate",
		Example: "confgen gen --from file://a.tmpl --exec TemplateName",
		Run:     runGenCmd,
	}
	genCmd.Flags().String("from", "", "Required: from path (file://path)")
	genCmd.Flags().String("exec", "", "Optional: template to execute")
	genCmd.MarkFlagRequired("from")
	return genCmd
}

func runGenCmd(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	from, err := flags.GetString("from")
	if err != nil {
		logger.Fatalf("%s", err)
	}
	exec, err := flags.GetString("exec")
	if err != nil {
		logger.Fatalf("%s", err)
	}
	if err = generate(from, exec); err != nil {
		logger.Fatalf("%s", err)
	}
}

func main() {
	genCmd := initGen()
	rootCmd.AddCommand(genCmd)
	rootCmd.Execute()
}
