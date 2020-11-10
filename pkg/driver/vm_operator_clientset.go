package driver

import (
	"github.com/jeffwubj/d-vm-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type DesktopVMInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.DesktopVMList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.DesktopVM, error)
	Create(vm *v1alpha1.DesktopVM) (*v1alpha1.DesktopVM, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type desktopVMClient struct {
	restClient rest.Interface
	ns         string
}

type DesktopVMV1Alpha1Interface interface {
	DesktopVMs(namespace string) DesktopVMInterface
}

type DesktopVMV1Alpha1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*DesktopVMV1Alpha1Client, error) {
	v1alpha1.AddToScheme(scheme.Scheme)
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupVersion.Group, Version: v1alpha1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DesktopVMV1Alpha1Client{restClient: client}, nil
}

func (c *DesktopVMV1Alpha1Client) DesktopVMs(namespace string) DesktopVMInterface {
	return &desktopVMClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}

func (c *desktopVMClient) List(opts metav1.ListOptions) (*v1alpha1.DesktopVMList, error) {
	result := v1alpha1.DesktopVMList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("desktopvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *desktopVMClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.DesktopVM, error) {
	result := v1alpha1.DesktopVM{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("desktopvms").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *desktopVMClient) Create(project *v1alpha1.DesktopVM) (*v1alpha1.DesktopVM, error) {
	result := v1alpha1.DesktopVM{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("desktopvms").
		Body(project).
		Do().
		Into(&result)

	return &result, err
}

func (c *desktopVMClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("desktopvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
