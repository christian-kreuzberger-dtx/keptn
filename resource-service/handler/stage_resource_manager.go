package handler

import (
	"github.com/keptn/keptn/resource-service/models"
)

type StageResourceManager struct {
}

func NewStageResourceManager() *StageResourceManager {
	stageResourceManager := &StageResourceManager{}
	return stageResourceManager
}

func (s StageResourceManager) CreateResources(params models.CreateResourcesParams) error {
	panic("implement me")
}

func (s StageResourceManager) GetResources(params models.GetResourcesParams) (*models.GetResourcesResponse, error) {
	panic("implement me")
}

func (s StageResourceManager) UpdateResources(params models.UpdateResourcesParams) error {
	panic("implement me")
}

func (s StageResourceManager) GetResource(params models.GetResourceParams) (*models.GetResourceResponse, error) {
	panic("implement me")
}

func (s StageResourceManager) UpdateResource(params models.UpdateResourceParams) error {
	panic("implement me")
}

func (s StageResourceManager) DeleteResource(params models.DeleteResourceParams) error {
	panic("implement me")
}
