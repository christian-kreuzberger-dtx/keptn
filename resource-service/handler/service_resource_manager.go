package handler

import (
	"github.com/keptn/keptn/resource-service/models"
)

type ServiceResourceManager struct {
}

func NewServiceResourceManager() *ServiceResourceManager {
	serviceResourceManager := &ServiceResourceManager{}
	return serviceResourceManager
}

func (s ServiceResourceManager) CreateResources(params models.CreateResourcesParams) error {
	panic("implement me")
}

func (s ServiceResourceManager) GetResources(params models.GetResourcesParams) (*models.GetResourcesResponse, error) {
	panic("implement me")
}

func (s ServiceResourceManager) UpdateResources(params models.UpdateResourcesParams) error {
	panic("implement me")
}

func (s ServiceResourceManager) GetResource(params models.GetResourceParams) (*models.GetResourceResponse, error) {
	panic("implement me")
}

func (s ServiceResourceManager) UpdateResource(params models.UpdateResourceParams) error {
	panic("implement me")
}

func (s ServiceResourceManager) DeleteResource(params models.DeleteResourceParams) error {
	panic("implement me")
}
