package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"aleaplay.com/alea-cli/pkg/httpclient"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(integrationTestStatusCmd)
}

var integrationTestStatusCmd = &cobra.Command{
	Use:   "tests-status",
	Short: "Get the status of the integration tests",
	Long: `Get the status of the integration tests.
	By default it will run all the tests against DEV environment unless you say otherwise.`,
	Run: func(cmd *cobra.Command, args []string) {

		failedScenarios, err := httpclient.GetIntegrationTestStatus()

		if err != nil {
			log.Fatalf("Error fetching response from integrator-test-suite. %s", err)
		}

		if len(failedScenarios) > 0 {
			json, _ := json.MarshalIndent(failedScenarios, "", "  ")
			fmt.Printf("❌ Some tests failed. \n %s", json)
			return
		}

		fmt.Printf("✅ All tests passed \n")
	},
}
