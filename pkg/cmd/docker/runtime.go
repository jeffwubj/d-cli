package docker

import (
	"fmt"
	"github.com/jeffwubj/d-cli/pkg/printconsole"
	"k8s.io/apimachinery/pkg/api/errors"
	"sort"
	"time"

	"github.com/jeffwubj/d-cli/pkg/driver"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

var d* driver.Driver

func init() {
	d = driver.InitDriver("", DefaultManagementNamespace)
}

func EnsureRuntime() error {
	var needwait = false
	_, err := d.DeploymentsClient.Get(DefaultRuntimeDeploymentName, metav1.GetOptions{})
	if (errors.IsNotFound(err)) {
		needwait = true
		err = bootStrapRuntime()
		if err != nil {
			return err
		}
	}
	if needwait {
		err = wait()
		if err != nil {
			return err
		}
	}

	return nil
}

func GetRunningRuntimePod() (string, error) {
	labelSelector := &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app": DefaultAppName,
		},
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return "", err
	}
	listOpts := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	podList, err := d.PodClient.List(listOpts)
	if err != nil {
		return "", err
	}
	var runningPods []*apiv1.Pod
	for i := range podList.Items {
		pod := &podList.Items[i]
		if pod.Status.Phase == apiv1.PodRunning && !IsPodTerminating(pod) {
			runningPods = append(runningPods, pod)
		}
	}
	sort.Slice(runningPods, func(i, j int) bool {
		return runningPods[i].Name < runningPods[j].Name
	})

	if len(runningPods) == 0 {
		return "", fmt.Errorf("there is no running runtime")
	} else if len(runningPods) > 1 {
		//printconsole.PrintWarning(fmt.Sprintf("there are %d runtimes, will use %q.", len(runningPods), runningPods[0].Name))
		return "", fmt.Errorf("there are %d runtimes", len(runningPods))
	}

	return runningPods[0].Name, nil
}

func bootStrapRuntime() error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: DefaultRuntimeDeploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": DefaultAppName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": DefaultAppName,
					},
				},
				Spec: apiv1.PodSpec{
					Volumes: [] apiv1.Volume{
						{
							Name: "docker-graph-storage",
							VolumeSource: apiv1.VolumeSource{
								HostPath: &apiv1.HostPathVolumeSource{
									Path: "/var/lib/dind",
								},
							},
						},
					},
					Containers: []apiv1.Container{
						{
							SecurityContext: &apiv1.SecurityContext {
								Privileged: pointer.BoolPtr(true),
							},
							Name:  "dind",
							Image: "docker:1.12.6-dind",
							VolumeMounts: [] apiv1.VolumeMount{
								{
									Name: "docker-graph-storage",
									MountPath: "/var/lib/docker",
								},
							},
						},
					},
				},
			},
		},
	}
	printconsole.PrintHint("deploy runtime in cluster...")
	result, err := d.DeploymentsClient.Create(deployment)
	if err != nil {
		return err
	}
	printconsole.PrintHint(fmt.Sprintf("runtime %q deployed.", result.GetObjectMeta().GetName()))
	return nil
}

func wait() error {
	_, err := d.DeploymentsClient.Get(DefaultRuntimeDeploymentName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return err
	}

	printconsole.PrintHint("wait runtime...")
	for i := 0; i < 100; i++ {
		deployment, err := d.DeploymentsClient.Get(DefaultRuntimeDeploymentName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		if deployment.Status.ReadyReplicas == 1 {
			printconsole.PrintHint("runtime is ready")
			return nil
		} else {
			time.Sleep(2 * time.Second)
		}
	}
	return errors.NewTimeoutError("wait runtime timeout", 200)
}