package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetIntegrationTestStatus() (failedScenarios []string, err error) {

	baseUrl := getBaseUrl()

	res, err := http.Get(baseUrl + "/app/health")

	if err != nil {
		fmt.Printf("Error fetching response from integrator-test-suite. %s", err)
	}

	bytes, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(bytes, &failedScenarios)

	if err != nil {
		fmt.Printf("Error reading response \n %s \n from integrator-test-suite. %s", string(bytes), err)
	}
	return failedScenarios, err
}

func ExecuteScenario(integrator string, scenario string) (scenarioResult ScenarioExecutionResponse, err error) {
	
	return execute(integrator + "/scenarios/" + scenario)
}

func ExecuteScenarioPath(scenarioPath string) (scenarioResult ScenarioExecutionResponse, err error) {
	
	return execute(scenarioPath)
}

func execute(integratorAndScenario string) (scenarioResult ScenarioExecutionResponse, err error) {
	
	baseUrl := getBaseUrl()

	res, err := http.Get(baseUrl + "/api/v1/" + integratorAndScenario)

	if err != nil {
		fmt.Printf("Error fetching response from integrator-test-suite. %s", err)
	}

	bytes, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(bytes, &scenarioResult)

	return scenarioResult, err
}

func getBaseUrl() string {
	env := viper.GetString("env")
	return viper.GetStringMapString("urls-" + env)["integrator-test-suite"]
}
