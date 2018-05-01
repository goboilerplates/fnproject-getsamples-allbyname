package main

import (
	"github.com/goboilerplates/core"
)

// GetSamplesAllByNameAPI is the GetSamplesAllByNameAPI interface.
type GetSamplesAllByNameAPI interface {
	AllByName(keyword string) ([]*core.SampleEntity, error)
}

// GetSamplesAllByNameAPIImpl is the implementation of GetSamplesAllByNameAPI interface.
type GetSamplesAllByNameAPIImpl struct {
	Interactor core.GetSamplesInteractor
}

// AllByName gets all samples by name keyword.
func (api GetSamplesAllByNameAPIImpl) AllByName(keyword string) ([]*core.SampleEntity, error) {
	return api.Interactor.AllByName(keyword)
}
