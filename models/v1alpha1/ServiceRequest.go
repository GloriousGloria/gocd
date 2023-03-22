package v1alpha1

import (
	"context"
	"fmt"

	"github.com/jnnkrdb/k8s/operator"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ServiceRequestSpec struct {
	Maintainer         string   `json:"maintainer"`
	Environment        string   `json:"environment"`
	ImageName          string   `json:"imagename"`
	ImageTag           string   `json:"imagetag"`
	Replicas           int      `json:"replicas"`
	Subdomains         []string `json:"subdomains"`
	Subpath            string   `json:"subpath"`
	AppSettingsContent string   `json:"appsettingscontent"`
	ServiceType        string   `json:"servicetype"`
	Envs               []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"envs"`
}

// deepcopy
func (in *ServiceRequest) DeepCopyInto(out *ServiceRequest) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = ServiceRequestSpec{
		Maintainer:         in.Spec.Maintainer,
		Environment:        in.Spec.Environment,
		ImageName:          in.Spec.ImageName,
		ImageTag:           in.Spec.ImageTag,
		Replicas:           in.Spec.Replicas,
		Envs:               in.Spec.Envs,
		Subdomains:         in.Spec.Subdomains,
		Subpath:            in.Spec.Subpath,
		AppSettingsContent: in.Spec.AppSettingsContent,
		ServiceType:        in.Spec.ServiceType,
	}
}

// ----------------------------------------------------
// kubernetes dependencies
type ServiceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceRequestSpec `json:"spec"`
}

type ServiceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceRequest `json:"items"`
}

func (in *ServiceRequest) DeepCopyObject() runtime.Object {
	out := ServiceRequest{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *ServiceRequestList) DeepCopyObject() runtime.Object {
	out := ServiceRequestList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]ServiceRequest, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}

// ----------------------------------------------------
// helper functions

const _SERVICEREQUEST_STRING string = "servicerequests"

// list all existing service declarations
func ListServiceRequests() (sdl ServiceRequestList, err error) {
	if err = operator.CRD().Get().Resource(_SERVICEREQUEST_STRING).Do(context.TODO()).Into(&sdl); err != nil {
		err = fmt.Errorf("[%s]:ListServiceRequests() error listing servicerequests: %#v", _SERVICEREQUEST_STRING, err)
	}
	return
}

// get a specific servicedeclaration
func GetServiceRequest(namespace, name string) (sd ServiceRequest, err error) {
	if err = operator.CRD().Get().Resource(_SERVICEREQUEST_STRING).Namespace(namespace).Name(name).Do(context.TODO()).Into(&sd); err != nil {
		err = fmt.Errorf("[%s]:GetServiceRequest(%s, %s) error listing servicerequests: %#v", _SERVICEREQUEST_STRING, namespace, name, err)
	}
	return
}
