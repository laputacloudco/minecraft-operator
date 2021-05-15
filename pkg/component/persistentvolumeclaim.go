/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package component

import (
	"github.com/laputacloudco/minecraft-operator/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// GeneratePVC creates a PVC for the Minecraft server
func GeneratePVC(mc v1alpha1.Minecraft) (v1.PersistentVolumeClaim, error) {
	pvcSizeRequest, err := resource.ParseQuantity(mc.Spec.StorageSize)
	if err != nil {
		return v1.PersistentVolumeClaim{}, err
	}
	return v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: make(map[string]string),
			Labels:      standardLabels(mc),
			Name:        mc.Name,
			Namespace:   mc.Namespace,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
			StorageClassName: &mc.Spec.StorageClassName,
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: pvcSizeRequest,
				},
			},
		},
	}, nil
}

// IndexPVC indexer func for controller-runtime
func IndexPVC(o runtime.Object) []string {
	pvc := o.(*v1.PersistentVolumeClaim)
	owner := metav1.GetControllerOf(pvc)
	if owner == nil {
		return nil
	}
	if owner.APIVersion != v1alpha1.GroupVersion.String() || owner.Kind != "Minecraft" {
		return nil
	}
	return []string{owner.Name}
}
