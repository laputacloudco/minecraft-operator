/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package component

import (
	"github.com/laputacloudco/minecraft-operator/api/v1alpha2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NeedsUpdateDeployment returns if the passed Services are out of sync
func NeedsUpdateDeployment(want, have appsv1.Deployment) bool {
	return *want.Spec.Replicas != *have.Spec.Replicas ||
		want.Annotations["configmap-revision"] != have.Annotations["configmap-revision"] ||
		!want.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().Equal(*have.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu()) ||
		!want.Spec.Template.Spec.Containers[0].Resources.Limits.Memory().Equal(*have.Spec.Template.Spec.Containers[0].Resources.Limits.Memory()) ||
		want.Spec.Template.Spec.Containers[0].StartupProbe.InitialDelaySeconds != have.Spec.Template.Spec.Containers[0].StartupProbe.InitialDelaySeconds ||
		want.Spec.Template.Spec.Containers[0].StartupProbe.PeriodSeconds != have.Spec.Template.Spec.Containers[0].StartupProbe.PeriodSeconds ||
		want.Spec.Template.Spec.Containers[0].LivenessProbe.InitialDelaySeconds != have.Spec.Template.Spec.Containers[0].LivenessProbe.InitialDelaySeconds ||
		want.Spec.Template.Spec.Containers[0].LivenessProbe.PeriodSeconds != have.Spec.Template.Spec.Containers[0].LivenessProbe.PeriodSeconds
}

// GenerateDeployment creates a Minecraft server deployment
func GenerateDeployment(mc v1alpha2.Minecraft) (appsv1.Deployment, error) {
	r := int32(0)
	if mc.Spec.Serve {
		r = 1
	}
	limitCPU, _ := resource.ParseQuantity(mc.Spec.LimitCPU)
	limitMem, _ := resource.ParseQuantity(mc.Spec.LimitMemory)
	requestCPU, _ := resource.ParseQuantity(mc.Spec.RequestCPU)
	requestMem, _ := resource.ParseQuantity(mc.Spec.RequestMemory)
	return appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: make(map[string]string),
			Labels:      standardLabels(mc),
			Name:        mc.Name,
			Namespace:   mc.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &r,
			Selector: &metav1.LabelSelector{
				MatchLabels: standardLabels(mc),
			},
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.RecreateDeploymentStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: make(map[string]string),
					Labels:      standardLabels(mc),
					Name:        mc.Name,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            mc.Name,
							Image:           mc.Spec.Image,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									Name:          "minecraft",
									ContainerPort: 25565,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "EULA",
									Value: "TRUE",
								},
							},
							EnvFrom: []corev1.EnvFromSource{
								{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: mc.Name,
										},
									},
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    limitCPU,
									corev1.ResourceMemory: limitMem,
								},
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    requestCPU,
									corev1.ResourceMemory: requestMem,
								},
							},
							StartupProbe: &corev1.Probe{
								Handler: corev1.Handler{
									Exec: &corev1.ExecAction{
										Command: []string{
											"mc-monitor",
											"status",
											"--host",
											"localhost",
										},
									},
								},
								FailureThreshold:    3,
								InitialDelaySeconds: mc.Spec.ProbeDelay,
								PeriodSeconds:       mc.Spec.ProbePeriod,
								SuccessThreshold:    1,
							},
							LivenessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									Exec: &corev1.ExecAction{
										Command: []string{
											"mc-monitor",
											"status",
											"--host",
											"localhost",
										},
									},
								},
								FailureThreshold:    3,
								InitialDelaySeconds: mc.Spec.ProbeDelay,
								PeriodSeconds:       mc.Spec.ProbePeriod,
								SuccessThreshold:    1,
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "data",
									MountPath: "/data",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "data",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: mc.Name,
								},
							},
						},
					},
				},
			},
		},
	}, nil
}

// IndexDeployment indexer func for controller-runtime
func IndexDeployment(o client.Object) []string {
	deploy := o.(*appsv1.Deployment)
	owner := metav1.GetControllerOf(deploy)
	if owner == nil {
		return nil
	}
	if owner.APIVersion != v1alpha2.GroupVersion.String() || owner.Kind != "Minecraft" {
		return nil
	}
	return []string{owner.Name}
}
