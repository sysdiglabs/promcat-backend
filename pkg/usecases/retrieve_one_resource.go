package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneResource struct {
	ResourceRepository resource.Repository
	Hash               string
}

func (useCase *RetrieveOneResource) Execute() (res *resource.Resource, err error) {
	return useCase.ResourceRepository.FindById(useCase.Hash)
}