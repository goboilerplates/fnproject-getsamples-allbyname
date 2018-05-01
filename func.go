package main

import (
	"context"
	"encoding/json"
	"io"

	fdk "github.com/fnproject/fdk-go"
	"github.com/goboilerplates/core"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(HandleRequest))
}

var (
	api GetSamplesAllByNameAPI
	// TestResult for checking tests.
	TestResult bool
)

// CreateAPI creates API for GetSamplesAllByName.
func CreateAPI() GetSamplesAllByNameAPI {
	if api == nil {
		return GetSamplesAllByNameAPIImpl{
			Interactor: core.CreateDefaultGetSamples("Kaka", "Ronaldo"),
		}
	}
	return api
}

// SampleRequest is the request for samples.
type SampleRequest struct {
	Keyword string `json:"keyword"`
}

// HandleRequest handles request.
func HandleRequest(ctx context.Context, in io.Reader, out io.Writer) {
	api = CreateAPI()
	sampleRequest := &SampleRequest{}
	json.NewDecoder(in).Decode(sampleRequest)
	result, err := api.AllByName(sampleRequest.Keyword)
	if err != nil {
		TestResult = false
		return
	}
	json.NewEncoder(out).Encode(&result)
	TestResult = true
}
