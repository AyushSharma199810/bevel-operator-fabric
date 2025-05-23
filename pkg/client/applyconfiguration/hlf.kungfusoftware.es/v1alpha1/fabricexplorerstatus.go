/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/pkg/apis/hlf.kungfusoftware.es/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
)

// FabricExplorerStatusApplyConfiguration represents a declarative configuration of the FabricExplorerStatus type for use
// with apply.
type FabricExplorerStatusApplyConfiguration struct {
	Conditions *status.Conditions         `json:"conditions,omitempty"`
	Message    *string                    `json:"message,omitempty"`
	Status     *v1alpha1.DeploymentStatus `json:"status,omitempty"`
}

// FabricExplorerStatusApplyConfiguration constructs a declarative configuration of the FabricExplorerStatus type for use with
// apply.
func FabricExplorerStatus() *FabricExplorerStatusApplyConfiguration {
	return &FabricExplorerStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricExplorerStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricExplorerStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FabricExplorerStatusApplyConfiguration) WithMessage(value string) *FabricExplorerStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricExplorerStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricExplorerStatusApplyConfiguration {
	b.Status = &value
	return b
}
