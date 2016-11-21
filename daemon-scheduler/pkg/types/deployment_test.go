// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package types

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	taskArn          = "arn:aws:ecs:us-east-1:12345678912:task/c024d145-093b-499a-9b14-5baf273f5835"
	instanceArn      = "arn:aws:us-east-1:123456789123:container-instance/4b6d45ea-a4b4-4269-9d04-3af6ddfdc597"
	desiredTaskCount = 5
)

type DeploymentTestSuite struct {
	suite.Suite
	deployment   *Deployment
	currentTasks []*ecs.Task
	failures     []*ecs.Failure
}

func (suite *DeploymentTestSuite) SetupTest() {
	var err error
	suite.deployment, err = NewDeployment(taskDefinition)
	assert.Nil(suite.T(), err, "Cannot initialize DeploymentTestSuite")

	task := ecs.Task{
		TaskArn:           aws.String(taskArn),
		TaskDefinitionArn: aws.String(taskDefinition),
	}
	suite.currentTasks = []*ecs.Task{&task}

	failedInstance := ecs.Failure{
		Arn: aws.String(instanceArn),
	}
	suite.failures = []*ecs.Failure{&failedInstance}
}

func TestDeploymentTestSuite(t *testing.T) {
	suite.Run(t, new(DeploymentTestSuite))
}

func (suite *DeploymentTestSuite) TestNewDeploymentEmptyTaskDefinition() {
	_, err := NewDeployment("")
	assert.Error(suite.T(), err, "Expected an error when task definition is empty")
}

func (suite *DeploymentTestSuite) TestNewDeployment() {
	d, err := NewDeployment(taskDefinition)
	assert.Nil(suite.T(), err, "Unexpected error when creating a deployment")
	assert.NotNil(suite.T(), d, "Deployment should not be nil")
	assert.NotEmpty(suite.T(), d.ID, "Deployment ID should not be empty")
	assert.Exactly(suite.T(), DeploymentPending, d.Status, "Deployment status should be pending")
	assert.Exactly(suite.T(), DeploymentHealthy, d.Health, "Deployment should be healthy")
	assert.NotNil(suite.T(), d.StartTime, "Deployment startTime should not be empty")
	assert.Empty(suite.T(), d.EndTime, "Deployment endtime should be empty")
	assert.Exactly(suite.T(), taskDefinition, d.TaskDefinition, "Deployment taskDefintion does not match expected")
}

func (suite *DeploymentTestSuite) TestUpdateDeploymentInProgressDeploymentCompleted() {
	suite.deployment.Status = DeploymentCompleted

	_, err := suite.deployment.UpdateDeploymentInProgress(desiredTaskCount, suite.currentTasks, suite.failures)
	assert.Error(suite.T(), err, "Expected an error when deployment is complete")
}

func (suite *DeploymentTestSuite) TestUpdateDeploymentInProgressUnhealthy() {
	d, err := suite.deployment.UpdateDeploymentInProgress(desiredTaskCount, suite.currentTasks, suite.failures)
	assert.Nil(suite.T(), err, "Unexpected error when setting deployment in progress")
	assert.NotNil(suite.T(), d, "Deployment should not be nil")
	assert.NotEmpty(suite.T(), d.ID, "Deployment ID should not be empty")
	assert.Exactly(suite.T(), DeploymentInProgress, d.Status, "Deployment status should be inprogress")
	assert.Exactly(suite.T(), DeploymentUnhealthy, d.Health, "Deployment should be unhealthy")
	assert.Exactly(suite.T(), desiredTaskCount, d.DesiredTaskCount, "Deployment desired task count should match expected")
	assert.NotNil(suite.T(), d.StartTime, "Deployment startTime should not be empty")
	assert.Empty(suite.T(), d.EndTime, "Deployment endtime should be empty")
	assert.Exactly(suite.T(), taskDefinition, d.TaskDefinition, "Deployment taskDefintion does not match expected")
	assert.Exactly(suite.T(), suite.currentTasks, d.CurrentTasks, "Deployment current tasks does not match expected")
	assert.Exactly(suite.T(), suite.failures, d.FailedInstances, "Deployment failed instances does not match expected")
}

func (suite *DeploymentTestSuite) TestUpdateDeploymentInProgressHealthy() {
	d, err := suite.deployment.UpdateDeploymentInProgress(desiredTaskCount, suite.currentTasks, []*ecs.Failure{})
	assert.Nil(suite.T(), err, "Unexpected error when setting deployment in progress")
	assert.NotNil(suite.T(), d, "Deployment should not be nil")
	assert.NotEmpty(suite.T(), d.ID, "Deployment ID should not be empty")
	assert.Exactly(suite.T(), DeploymentInProgress, d.Status, "Deployment status should be inprogress")
	assert.Exactly(suite.T(), DeploymentHealthy, d.Health, "Deployment should be healthy")
	assert.Exactly(suite.T(), desiredTaskCount, d.DesiredTaskCount, "Deployment desired task count should match expected")
	assert.NotNil(suite.T(), d.StartTime, "Deployment startTime should not be empty")
	assert.Empty(suite.T(), d.EndTime, "Deployment endtime should be empty")
	assert.Exactly(suite.T(), taskDefinition, d.TaskDefinition, "Deployment taskDefintion does not match expected")
	assert.Exactly(suite.T(), suite.currentTasks, d.CurrentTasks, "Deployment current tasks does not match expected")
	assert.Empty(suite.T(), d.FailedInstances, "Deployment failed instances does not match expected")
}

func (suite *DeploymentTestSuite) TestUpdateDeploymentCompleted() {
	unhealthyDeployment, err := suite.deployment.UpdateDeploymentInProgress(
		desiredTaskCount, []*ecs.Task{}, suite.failures)

	d := unhealthyDeployment.UpdateDeploymentCompleted(suite.currentTasks)
	assert.Nil(suite.T(), err, "Unexpected error when setting deployment to completed")
	assert.NotNil(suite.T(), d, "Deployment should not be nil")
	assert.NotEmpty(suite.T(), d.ID, "Deployment ID should not be empty")
	assert.Exactly(suite.T(), DeploymentCompleted, d.Status, "Deployment status should be completed")
	assert.Exactly(suite.T(), DeploymentHealthy, d.Health, "Deployment should be healthy")
	assert.Exactly(suite.T(), desiredTaskCount, d.DesiredTaskCount, "Deployment desired task count should match expected")
	assert.NotNil(suite.T(), d.StartTime, "Deployment startTime should not be empty")
	assert.NotNil(suite.T(), d.EndTime, "Deployment endtime should not be empty")
	assert.Exactly(suite.T(), taskDefinition, d.TaskDefinition, "Deployment taskDefintion does not match expected")
	assert.Exactly(suite.T(), suite.currentTasks, d.CurrentTasks, "Deployment current tasks does not match expected")
	assert.Nil(suite.T(), d.FailedInstances, "Deployment failed instances does not match expected")
}