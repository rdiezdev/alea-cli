package cmd

import (
	"fmt"

	"aleaplay.com/alea/pkg/httpclient"
	"github.com/spf13/cobra"
)

func init() {
	integrationTestCmd.Flags().StringP("integrator", "i", "", "Integrator's name in kebab-case. E.g. 'my-integrator'")
	
	integrationTestCmd.Flags().StringP("scenario", "s", "", "Scenario's name in kebab-case. E.g. 'happy-path'")

	rootCmd.AddCommand(integrationTestCmd)
}

var integrationTestCmd = &cobra.Command{
	Use:   "tests",
	Short: "Perform an integration test with integratior-test-suite",
	Long: `Perform an integration test with integratior-test-suite.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		integrator, _ := cmd.Flags().GetString("integrator")			
		scenario, _ := cmd.Flags().GetString("scenario")

		if integrator == "" && scenario == "" {
			if len(args) == 0 {
				fmt.Println("❌ You must specify an a path scenario, like 'amatic/scenarios/happy-path' or an integrator and a scenario using the flags --i and --s")
				return
			}
			httpclient.ExecuteScenarioPath(args[0])
		}
		

		response, err := httpclient.ExecuteScenario(integrator, scenario)

		if err != nil {
			fmt.Printf("Error executing scenario \n %+v \n", err)
		}

		if response.Result.Status != "OK" {
			fmt.Printf("❌ Scenario failed \n %+v \n", response)
		}

		fmt.Println("✅ All steps passed")
	},
}
