/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package component

import (
	"errors"

	"github.com/laputacloudco/minecraft-operator/api/v1alpha2"
	"github.com/laputacloudco/minecraft-operator/internal/sort"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GenerateService creates a Service for the Minecraft server.
func GenerateService(mc v1alpha2.Minecraft) v1.Service {
	return v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: make(map[string]string),
			Labels:      standardLabels(mc),
			Name:        mc.Name,
			Namespace:   mc.Namespace,
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:       "minecraft",
					Protocol:   v1.ProtocolTCP,
					Port:       25565,
					TargetPort: intstr.FromString("minecraft"),
				},
			},
			Selector: standardLabels(mc),
			Type:     v1.ServiceType(mc.Spec.ServiceType),
		},
	}
}

// ExtractExistingLoadbalancerIP searches the passed slice of Services for
// existing loadbalancers and returns an IP address of those loadbalancers.
// If the Service slice is empty, returns empty string and no error.
// If there are Services of type LoadBalancer but they do not (yet?) have IP
// addresses assigned, return an error to indicate that this should be aborted
// and tried again later.
func ExtractExistingLoadbalancerIP(extant []v1.Service) (string, error) {
	foundLB := false
	for _, svc := range extant {
		if svc.Spec.Type != v1.ServiceTypeLoadBalancer {
			continue
		}
		foundLB = true
		if svc.Spec.LoadBalancerIP != "" {
			return svc.Spec.LoadBalancerIP, nil
		}
		if len(svc.Status.LoadBalancer.Ingress) > 0 && svc.Status.LoadBalancer.Ingress[0].IP != "" {
			return svc.Status.LoadBalancer.Ingress[0].IP, nil
		}
	}
	if foundLB {
		return "", errors.New("loadbalancers exist but do not yet have IPs")
	}
	return "", nil

}

// AssignServicePort walks the existing service definitions and assigns the
// first available port with the specified inclusive range, or -1 if the range
// is fully populated.
func AssignServicePort(extant []v1.Service, minPort, maxPort int32) int32 {
	used := sort.Int32Slice{}
	for _, svc := range extant {
		for _, p := range svc.Spec.Ports {
			if p.Name == "minecraft" {
				used = append(used, p.Port)
			}
		}
	}
	used.Sort()
	return sort.FirstUnusedInRange(used, minPort, maxPort)
}

// IndexService owner indexer func for controller-runtime.
func IndexService(o client.Object) []string {
	svc := o.(*v1.Service)
	owner := metav1.GetControllerOf(svc)
	if owner == nil {
		return nil
	}
	if owner.APIVersion != v1alpha2.GroupVersion.String() || owner.Kind != "Minecraft" {
		return nil
	}
	return []string{owner.Name}
}
