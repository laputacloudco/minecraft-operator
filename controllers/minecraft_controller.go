/*
Copyright 2020 The Laputa Cloud Co.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	v1alpha2 "github.com/laputacloudco/minecraft-operator/api/v1alpha2"
	"github.com/laputacloudco/minecraft-operator/internal/component"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// OwnerKey index for owner references
	OwnerKey = ".metadata.controller"
)

// MinecraftReconciler reconciles a Minecraft object
type MinecraftReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=game.laputacloud.co,resources=minecrafts,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=game.laputacloud.co,resources=minecrafts/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=game.laputacloud.co,resources=minecrafts/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;delete;get;list;update;watch
//+kubebuilder:rbac:groups="",resources=configmaps;persistentvolumeclaims;services,verbs=create;delete;get;list;update;watch
//+kubebuilder:rbac:groups="",resources=events,verbs=create;update;patch

// Reconcile a Minecraft object
func (r *MinecraftReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("minecraft", req.NamespacedName)

	log.Info("reconciling", "name", req.Name)

	// Get this Minecraft
	mc := &v1alpha2.Minecraft{}
	if err := r.Get(ctx, req.NamespacedName, mc); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	v1alpha2.SetDefaults(mc)
	destroying := mc.Status.Status == v1alpha2.Destroying

	// generate configmap from spec
	cm := component.GenerateConfigMap(*mc)
	if err := ctrl.SetControllerReference(mc, &cm, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// list existing configmaps
	var configMaps corev1.ConfigMapList
	if err := r.List(ctx, &configMaps, client.InNamespace(req.Namespace), client.MatchingFields{OwnerKey: req.Name}); err != nil {
		log.Error(err, "unable to list ConfigMaps", "namespace", req.Namespace, "owner", req.Name)
		return ctrl.Result{}, err
	}

	// create ConfigMap, if it does not exist
	if !destroying {
		if len(configMaps.Items) < 1 {
			log.Info("creating configmap")
			if err := r.setStatus(ctx, mc, v1alpha2.Creating); err != nil {
				return ctrl.Result{}, err
			}
			if err := r.Create(ctx, &cm); err != nil {
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorCreating", "Failed to create ConfigMap %s, err = ", cm.Name, err)
				return ctrl.Result{}, err
			}
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Created", "ConfigMap %s created", cm.Name)
		} else {
			// update configmap if it does not match spec
			if component.NeedsUpdateConfigMap(cm, configMaps.Items[0]) {
				log.Info("updating configmap")
				if err := r.setStatus(ctx, mc, v1alpha2.Updating); err != nil {
					return ctrl.Result{}, err
				}
				if err := r.Update(ctx, &cm); err != nil {
					r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorUpdating", "Failed to update ConfigMap %s, err = ", cm.Name, err)
					return ctrl.Result{}, err
				}
				r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Updated", "ConfigMap %s updated", cm.Name)
			}
		}
	} else {
		for _, cm := range configMaps.Items {
			if err := r.Delete(ctx, &cm); err != nil {
				log.Error(err, "failed to delete cm")
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorDeleting", "Failed to delete ConfigMap %s, err = ", cm.Name, err)
			} else {
				r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Deleted", "ConfigMap %s deleted", cm.Name)
			}
		}
	}

	// list existing pvcs
	var persistentVolumeClaims corev1.PersistentVolumeClaimList
	if err := r.List(ctx, &persistentVolumeClaims, client.InNamespace(req.Namespace), client.MatchingFields{OwnerKey: req.Name}); err != nil {
		log.Error(err, "unable to list PersistentVolumeClaims", "namespace", req.Namespace, "owner", req.Name)
		return ctrl.Result{}, err
	}

	// create PVC, if it does not exist
	// TODO: if it needs to be updated, and the StorageClass supports updates
	// then we could submit a PVC Update Request
	pvc, err := component.GeneratePVC(*mc)
	if err != nil {
		log.Error(err, "unable to parse PVC settings")
		return ctrl.Result{}, err
	}
	if err := ctrl.SetControllerReference(mc, &pvc, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}
	if !destroying {
		if len(persistentVolumeClaims.Items) < 1 {
			log.Info("creating pvc")
			if err := r.setStatus(ctx, mc, v1alpha2.Creating); err != nil {
				return ctrl.Result{}, err
			}
			if err := r.Create(ctx, &pvc); err != nil {
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorCreating", "Failed to create PersistentVolumeClaim %s, err = ", pvc.Name, err)
				return ctrl.Result{}, err
			}
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Created", "PersistentVolumeClaim %s created", pvc.Name)
		}
	} else {
		for _, pvc := range persistentVolumeClaims.Items {
			if err := r.Delete(ctx, &pvc); err != nil {
				log.Error(err, "failed to delete pvc")
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorDeleting", "Failed to delete PersistentVolumeClaim %s, err = ", pvc.Name, err)
			} else {
				r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Deleted", "PersistentVolumeClaim %s deleted", pvc.Name)
			}
		}
	}

	// generate deployment from spec
	deploy, err := component.GenerateDeployment(*mc)
	if err != nil {
		return ctrl.Result{}, err
	}
	if err := ctrl.SetControllerReference(mc, &deploy, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	if !destroying {
		// get the latest configmap for the deployment annotations
		if err := r.Get(ctx, types.NamespacedName{Namespace: cm.Namespace, Name: cm.Name}, &cm); err != nil {
			log.Error(err, "failed to get latest configmap", "namespace", cm.Namespace, "owner", cm.Name)
			return ctrl.Result{}, err
		}
		deploy.Annotations["configmap-revision"] = cm.ResourceVersion
	}

	// list existing deployments
	var deployments appsv1.DeploymentList
	if err := r.List(ctx, &deployments, client.InNamespace(req.Namespace), client.MatchingFields{OwnerKey: req.Name}); err != nil {
		log.Error(err, "unable to list Deployments", "namespace", req.Namespace, "owner", req.Name)
		return ctrl.Result{}, err
	}

	if !destroying {
		// create Deployment, if it does not exist
		if len(deployments.Items) < 1 {
			log.Info("creating deployment")
			if err := r.setStatus(ctx, mc, v1alpha2.Creating); err != nil {
				return ctrl.Result{}, err
			}
			if err := r.Create(ctx, &deploy); err != nil {
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorCreating", "Failed to create Deployment %s, err = ", deploy.Name, err)
				return ctrl.Result{}, err
			}
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Created", "Deployment %s created", deploy.Name)
		} else {
			// update deployment if it does not match spec
			if component.NeedsUpdateDeployment(deploy, deployments.Items[0]) {
				log.Info("updating deployment")
				if err := r.setStatus(ctx, mc, v1alpha2.Updating); err != nil {
					return ctrl.Result{}, err
				}
				if *deploy.Spec.Replicas == 0 && *deployments.Items[0].Spec.Replicas != 0 {
					if err := r.setStatus(ctx, mc, v1alpha2.Stopping); err != nil {
						return ctrl.Result{}, err
					}
				}
				if *deploy.Spec.Replicas != 0 && *deployments.Items[0].Spec.Replicas == 0 {
					if err := r.setStatus(ctx, mc, v1alpha2.Starting); err != nil {
						return ctrl.Result{}, err
					}
				}
				if err := r.Update(ctx, &deploy); err != nil {
					r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorUpdating", "Failed to update Deployment %s, err = ", deploy.Name, err)
					return ctrl.Result{}, err
				}
				r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Updated", "Deployment %s updated", deploy.Name)
			}
		}
	} else {
		for _, deploy := range deployments.Items {
			if err := r.Delete(ctx, &deploy); err != nil {
				log.Error(err, "failed to delete deploy")
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorDeleting", "Failed to delete Deployment %s, err = ", deploy.Name, err)
			}
		}
		r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Deleted", "Deployment %s deleted", deploy.Name)
	}

	// generate service from spec
	service := component.GenerateService(*mc)
	if err := ctrl.SetControllerReference(mc, &service, r.Scheme); err != nil {
		return ctrl.Result{}, err
	}

	// list existing services
	var services corev1.ServiceList
	if err := r.List(ctx, &services, client.InNamespace(req.Namespace), client.MatchingFields{OwnerKey: req.Name}); err != nil {
		log.Error(err, "unable to list Services", "namespace", req.Namespace, "owner", req.Name)
		return ctrl.Result{}, err
	}

	if !destroying {
		// create Service, if it does not exist and Minecraft.Serve is true
		if len(services.Items) == 0 && mc.Spec.Serve {
			log.Info("creating service", "existing", len(services.Items))

			// set status and create
			if err := r.setStatus(ctx, mc, v1alpha2.Starting); err != nil {
				return ctrl.Result{}, err
			}
			if err := r.Create(ctx, &service); err != nil {
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorCreating", "Failed to create Service %s, err = ", service.Name, err)
				return ctrl.Result{}, err
			}
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Created", "Service %s created", service.Name)
		}
		// destroy service, if it does exist and Minecraft.Serve is false
		if len(services.Items) > 0 && !mc.Spec.Serve {
			log.Info("destroying service")
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Deleting", "Deleting service %s", service.Name)
			if err := r.setStatus(ctx, mc, v1alpha2.Stopping); err != nil {
				return ctrl.Result{}, err
			}
			if err := r.Delete(ctx, &service); err != nil {
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorDeleting", "Failed to delete Service %s, err = ", service.Name, err)
				return ctrl.Result{}, err
			}
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Deleted", "Service %s deleted", service.Name)
		}
		if len(services.Items) == 0 && !mc.Spec.Serve {
			if err := r.setStatus(ctx, mc, v1alpha2.Stopped); err != nil {
				return ctrl.Result{}, err
			}
			if err := r.setStatusAddress(ctx, mc, nil); err != nil {
				return ctrl.Result{}, err
			}
		}
		if len(services.Items) >= 1 && mc.Spec.Serve {
			// set the Address status once we have a LoadBalancer IP
			if len(services.Items[0].Status.LoadBalancer.Ingress) > 0 && services.Items[0].Status.LoadBalancer.Ingress[0].IP != "" {
				if err := r.setStatusAddress(ctx, mc, &services.Items[0]); err != nil {
					return ctrl.Result{}, err
				}
				if err := r.setStatus(ctx, mc, v1alpha2.Running); err != nil {
					return ctrl.Result{}, err
				}
			} else {
				if err := r.setStatus(ctx, mc, v1alpha2.Starting); err != nil {
					return ctrl.Result{}, err
				}
				if err := r.setStatusAddress(ctx, mc, nil); err != nil {
					return ctrl.Result{}, err
				}
			}
		}
	} else {
		for _, svc := range services.Items {
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Deleting", "Deleting service %s", service.Name)
			if err := r.Delete(ctx, &svc); err != nil {
				log.Error(err, "failed to delete svc")
				r.Recorder.Eventf(mc, corev1.EventTypeWarning, "ErrorDeleting", "Failed to delete Service %s, err = ", service.Name, err)
			}
			r.Recorder.Eventf(mc, corev1.EventTypeNormal, "Deleted", "Deleting service %s", service.Name)
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MinecraftReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &appsv1.Deployment{}, OwnerKey, component.IndexDeployment); err != nil {
		return err
	}
	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &corev1.ConfigMap{}, OwnerKey, component.IndexConfigMap); err != nil {
		return err
	}
	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &corev1.Service{}, OwnerKey, component.IndexService); err != nil {
		return err
	}
	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &corev1.PersistentVolumeClaim{}, OwnerKey, component.IndexPVC); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha2.Minecraft{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.PersistentVolumeClaim{}).
		Complete(r)
}

func (r *MinecraftReconciler) setStatus(ctx context.Context, mc *v1alpha2.Minecraft, status v1alpha2.ServerStatus) error {
	mc.Status.Status = status
	return r.Status().Update(ctx, mc)
}

func (r *MinecraftReconciler) setStatusAddress(ctx context.Context, mc *v1alpha2.Minecraft, svc *corev1.Service) error {
	mc.Status.Address = ""
	if svc == nil {
		return r.Status().Update(ctx, mc)
	}

	addr := svc.Status.LoadBalancer.Ingress[0].IP
	for _, p := range svc.Spec.Ports {
		if p.Name == "minecraft" {
			addr = fmt.Sprintf("%s:%d", addr, p.Port)
		}
	}
	mc.Status.Address = addr
	return r.Status().Update(ctx, mc)
}
