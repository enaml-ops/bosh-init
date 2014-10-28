package fakes

import (
	bmagentclient "github.com/cloudfoundry/bosh-micro-cli/microdeployer/agentclient"
	bmstemcell "github.com/cloudfoundry/bosh-micro-cli/stemcell"
)

type FakeApplySpecCreator struct {
	CreateInput     CreateInput
	CreateApplySpec bmagentclient.ApplySpec
	CreateErr       error
}

type CreateInput struct {
	ApplySpec               bmstemcell.ApplySpec
	DeploymentName          string
	JobName                 string
	NetworksSpec            map[string]interface{}
	ArchivedTemplatesBlobID string
	ArchivedTemplatesPath   string
	TemplatesDir            string
}

func NewFakeApplySpecCreator() *FakeApplySpecCreator {
	return &FakeApplySpecCreator{}
}

func (c *FakeApplySpecCreator) Create(
	applySpec bmstemcell.ApplySpec,
	deploymentName string,
	jobName string,
	networksSpec map[string]interface{},
	archivedTemplatesBlobID string,
	archivedTemplatesPath string,
	templatesDir string,
) (bmagentclient.ApplySpec, error) {
	c.CreateInput = CreateInput{
		ApplySpec:               applySpec,
		DeploymentName:          deploymentName,
		JobName:                 jobName,
		NetworksSpec:            networksSpec,
		ArchivedTemplatesBlobID: archivedTemplatesBlobID,
		ArchivedTemplatesPath:   archivedTemplatesPath,
		TemplatesDir:            templatesDir,
	}

	return c.CreateApplySpec, c.CreateErr
}