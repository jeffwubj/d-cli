package driver

import (
	"github.com/davecgh/go-spew/spew"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	clientappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	clientcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type Driver struct {
	PodClient clientcorev1.PodInterface
	DeploymentsClient     clientappsv1.DeploymentInterface
	PersistentVolumeClient clientcorev1.PersistentVolumeInterface
	PersistentVolumeClaimClient clientcorev1.PersistentVolumeClaimInterface
	Config *rest.Config
	DesktopVMClient DesktopVMInterface
}

func InitDriver(kubeconfig string, namespace string) *Driver {
	if kubeconfig == "" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
	}
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	podClient := clientset.CoreV1().Pods(namespace)
	deploymentsClient := clientset.AppsV1().Deployments(namespace)

	persistentVolumeClient := clientset.CoreV1().PersistentVolumes()
	persistentVolumeClaimClient := clientset.CoreV1().PersistentVolumeClaims(namespace)

	desktopVMClientV, err := NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	desktopVMClient := desktopVMClientV.DesktopVMs("default")

	return &Driver{
		PodClient: podClient,
		DeploymentsClient: deploymentsClient,
		PersistentVolumeClient: persistentVolumeClient,
		PersistentVolumeClaimClient: persistentVolumeClaimClient,
		Config: config,
		DesktopVMClient: desktopVMClient,
	}
}

func (d *Driver) GetPods() {
	pods, _ := d.PodClient.List(metav1.ListOptions{})
	spew.Dump(pods)
}

func (d *Driver) GetDeployments() {
	deployements, _ := d.DeploymentsClient.List(metav1.ListOptions{})
	spew.Dump(deployements)
}