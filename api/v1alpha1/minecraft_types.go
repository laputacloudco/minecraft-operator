/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MinecraftSpec defines the desired state of Minecraft
type MinecraftSpec struct {
	// Config to be passed to the server container as environment variables.
	Config map[string]string `json:"config,omitempty"`

	// ServicePort the port the server service will be reachable on.
	ServicePort int32 `json:"servicePort,omitempty"`

	// ServiceType the type of Kubernetes Service to create.
	ServiceType string `json:"serviceType,omitempty"`

	// StorageClassName the storage class for creating the game data PVC.
	StorageClassName string `json:"storageClassName,omitempty"`

	// StorageSize the capacity of the game data PVC.
	StorageSize string `json:"storageSize,omitempty"`

	// Image the server container to pull and run.
	Image string `json:"image,omitempty"`

	// Serve tells the controller to run or stop this Server.
	Serve bool `json:"serve,omitempty"`

	// CPU is the CPU resource limit of the created instance.
	// +kubebuilder:validation:Required
	LimitCPU string `json:"limitCpu,omitempty"`

	// Memory is the Memory resource limit of the created instance.
	// +kubebuilder:validation:Required
	LimitMemory string `json:"limitMemory,omitempty"`

	// ProbeDelay gives the Pod time to start before healthchecking begins.
	ProbeDelay int32 `json:"probeDelay,omitempty"`

	// ProbePeriod indicates how often to do liveness probing.
	// Changing this period will effect how aggresively the instance is
	// restarted if it starts to become resource-constrained.
	ProbePeriod int32 `json:"probePeriod,omitempty"`

	// CPU is the CPU resource limit of the created instance.
	// +kubebuilder:validation:Required
	RequestCPU string `json:"requestCpu,omitempty"`

	// Memory is the Memory resource limit of the created instance.
	// +kubebuilder:validation:Required
	RequestMemory string `json:"requestMemory,omitempty"`
}

// ServerStatus indicates the Server Status
// +kubebuilder:validation:Enum=Creating;Destroying;Running;Starting;Stopped;Stopping;Unknown;Updating
type ServerStatus string

const (
	// Creating status
	Creating ServerStatus = "Creating"
	// Destroying status
	Destroying ServerStatus = "Destroying"
	// Running status
	Running ServerStatus = "Running"
	// Starting status
	Starting ServerStatus = "Starting"
	// Stopped status
	Stopped ServerStatus = "Stopped"
	// Stopping status
	Stopping ServerStatus = "Stopping"
	// Unknown status
	Unknown ServerStatus = "Unknown"
	// Updating status
	Updating ServerStatus = "Updating"
)

// MinecraftStatus defines the observed state of Minecraft
type MinecraftStatus struct {
	// Status indicates the Server Status
	Status ServerStatus `json:"status,omitempty"`
	// Address the public server address
	Address string `json:"address,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`
// +kubebuilder:printcolumn:name="Address",type=string,JSONPath=`.status.address`

// Minecraft is the Schema for the minecrafts API
type Minecraft struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MinecraftSpec   `json:"spec,omitempty"`
	Status MinecraftStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MinecraftList contains a list of Minecraft
type MinecraftList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Minecraft `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Minecraft{}, &MinecraftList{})
}
