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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ContainerInstance container instance
// swagger:model ContainerInstance
type ContainerInstance struct {

	// e c2 instance ID
	EC2InstanceID string `json:"EC2InstanceID,omitempty"`

	// agent connected
	// Required: true
	AgentConnected *bool `json:"agentConnected"`

	// agent update status
	AgentUpdateStatus string `json:"agentUpdateStatus,omitempty"`

	// attributes
	Attributes []*ContainerInstanceAttribute `json:"attributes"`

	// cluster a r n
	// Required: true
	ClusterARN *string `json:"clusterARN"`

	// container instance a r n
	// Required: true
	ContainerInstanceARN *string `json:"containerInstanceARN"`

	// pending tasks count
	// Required: true
	PendingTasksCount *int64 `json:"pendingTasksCount"`

	// registered resources
	// Required: true
	RegisteredResources []*ContainerInstanceResource `json:"registeredResources"`

	// remaining resources
	// Required: true
	RemainingResources []*ContainerInstanceResource `json:"remainingResources"`

	// running tasks count
	// Required: true
	RunningTasksCount *int64 `json:"runningTasksCount"`

	// status
	// Required: true
	Status *string `json:"status"`

	// version info
	// Required: true
	VersionInfo *ContainerInstanceVersionInfo `json:"versionInfo"`
}

// Validate validates this container instance
func (m *ContainerInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentConnected(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClusterARN(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerInstanceARN(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePendingTasksCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegisteredResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemainingResources(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRunningTasksCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersionInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerInstance) validateAgentConnected(formats strfmt.Registry) error {

	if err := validate.Required("agentConnected", "body", m.AgentConnected); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstance) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {

		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {

			if err := m.Attributes[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ContainerInstance) validateClusterARN(formats strfmt.Registry) error {

	if err := validate.Required("clusterARN", "body", m.ClusterARN); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstance) validateContainerInstanceARN(formats strfmt.Registry) error {

	if err := validate.Required("containerInstanceARN", "body", m.ContainerInstanceARN); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstance) validatePendingTasksCount(formats strfmt.Registry) error {

	if err := validate.Required("pendingTasksCount", "body", m.PendingTasksCount); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstance) validateRegisteredResources(formats strfmt.Registry) error {

	if err := validate.Required("registeredResources", "body", m.RegisteredResources); err != nil {
		return err
	}

	for i := 0; i < len(m.RegisteredResources); i++ {

		if swag.IsZero(m.RegisteredResources[i]) { // not required
			continue
		}

		if m.RegisteredResources[i] != nil {

			if err := m.RegisteredResources[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ContainerInstance) validateRemainingResources(formats strfmt.Registry) error {

	if err := validate.Required("remainingResources", "body", m.RemainingResources); err != nil {
		return err
	}

	for i := 0; i < len(m.RemainingResources); i++ {

		if swag.IsZero(m.RemainingResources[i]) { // not required
			continue
		}

		if m.RemainingResources[i] != nil {

			if err := m.RemainingResources[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ContainerInstance) validateRunningTasksCount(formats strfmt.Registry) error {

	if err := validate.Required("runningTasksCount", "body", m.RunningTasksCount); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstance) validateVersionInfo(formats strfmt.Registry) error {

	if err := validate.Required("versionInfo", "body", m.VersionInfo); err != nil {
		return err
	}

	if m.VersionInfo != nil {

		if err := m.VersionInfo.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}