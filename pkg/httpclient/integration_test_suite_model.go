package httpclient

type ScenarioExecutionResponse struct {
	Title string `json:"title"`
	Path string `json:"path"`
	Result secenarioResult `json:"result"`
}

type secenarioResult struct {

	Status string `json:"status"`
	StatusMessage string `json:"statusMessage"`
	Steps []step `json:"steps"`
}

type step struct {
	Action action `json:"action"`
	Request string `json:"request"`
	Response string `json:"response"`
	ExpectedResponse string `json:"expectedResponse"`
	Status string `json:"status"`
	StatusMessage string `json:"statusMessage"`
	Errors []stepError `json:"errors"`
}

type action struct {
	Type string `json:"type"`
}

type stepError struct {
	Type string `json:"type"`
	Path string `json:"path"`
	NodeName string `json:"nodeName"`
}