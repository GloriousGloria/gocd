package v1alpha1

import (
	"context"
	"fmt"

	"github.com/jnnkrdb/k8s/operator"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ServiceDeclarationSpec struct {
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
func (in *ServiceDeclaration) DeepCopyInto(out *ServiceDeclaration) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = ServiceDeclarationSpec{
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
type ServiceDeclaration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDeclarationSpec `json:"spec"`
}

type ServiceDeclarationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceDeclaration `json:"items"`
}

func (in *ServiceDeclaration) DeepCopyObject() runtime.Object {
	out := ServiceDeclaration{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *ServiceDeclarationList) DeepCopyObject() runtime.Object {
	out := ServiceDeclarationList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]ServiceDeclaration, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}

// ----------------------------------------------------
// helper functions

const _SERVICEDECLARATION_STRING string = "servicedeclarations"

// list all existing service declarations
func ListServiceDeclarations() (sdl ServiceDeclarationList, err error) {
	if err = operator.CRD().Get().Resource(_SERVICEDECLARATION_STRING).Do(context.TODO()).Into(&sdl); err != nil {
		err = fmt.Errorf("[%s]:ListServiceDeclarations() error listing servicedeclarations: %#v", _SERVICEDECLARATION_STRING, err)
	}
	return
}

// get a specific servicedeclaration
func GetServiceDeclaration(namespace, name string) (sd ServiceDeclaration, err error) {
	if err = operator.CRD().Get().Resource(_SERVICEDECLARATION_STRING).Namespace(namespace).Name(name).Do(context.TODO()).Into(&sd); err != nil {
		err = fmt.Errorf("[%s]:GetServiceDeclaration(%s, %s) error listing servicedeclarations: %#v", _SERVICEDECLARATION_STRING, namespace, name, err)
	}
	return
}
