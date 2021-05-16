/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
)

const (
	defaultImage = "docker.io/itzg/minecraft-server"
)

// SetDefaults sets defaults if server options were not specified.
func SetDefaults(mc *Minecraft) {
	if mc.Spec.DeploymentOptions.Image == "" {
		mc.Spec.DeploymentOptions.Image = defaultImage
	}
	if mc.Spec.StorageOptions.StorageSize == "" {
		mc.Spec.StorageOptions.StorageSize = "8Gi"
	}
	if mc.Spec.ServiceOptions.ServiceType == "" {
		mc.Spec.ServiceOptions.ServiceType = string(v1.ServiceTypeLoadBalancer)
	}
	if mc.Spec.ServiceOptions.ServicePort == 0 {
		mc.Spec.ServiceOptions.ServicePort = 25565
	}
	if mc.Spec.DeploymentOptions.ProbePeriod == 0 {
		mc.Spec.DeploymentOptions.ProbePeriod = 30
	}
	if mc.Spec.DeploymentOptions.ProbeDelay == 0 {
		mc.Spec.DeploymentOptions.ProbeDelay = 300
	}
}
