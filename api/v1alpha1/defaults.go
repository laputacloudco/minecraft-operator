/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

const (
	defaultImage = "quay.io/laputacloudco/minecraft-vanilla:latest"
)

// SetDefaults sets defaults if server options were not specified.
func SetDefaults(mc *Minecraft) {
	if mc.Spec.Image == "" {
		mc.Spec.Image = defaultImage
	}
	if mc.Spec.StorageSize == "" {
		mc.Spec.StorageSize = "8Gi"
	}
	if mc.Spec.ServiceType == "" {
		mc.Spec.ServiceType = string(v1.ServiceTypeLoadBalancer)
	}
	if mc.Spec.ServicePort == 0 {
		mc.Spec.ServicePort = 25565
	}
	if mc.Spec.ProbePeriod == 0 {
		mc.Spec.ProbePeriod = 30
	}
	if mc.Spec.ProbeDelay == 0 {
		mc.Spec.ProbeDelay = 300
	}
}
