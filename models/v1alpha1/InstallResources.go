package v1alpha1

import (
	"context"
	"fmt"

	"github.com/jnnkrdb/k8s/operator"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type SDMeta struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type RelatedObjectStruct struct {
	ObjectType string `json:"objecttype"`
	ApiVersion string `json:"apiversion"`
	Namespace  string `json:"namespace"`
	Name       string `json:"name"`
}

type GitlabPointerStruct struct {
	GitlabAPIV4 string `json:"gitlabapiv4"`
	ProjectID   int    `json:"projectid"`
	AccessToken string `json:"accesstoken"`
	File        string `json:"file"`
}

type InstallResourceSpec struct {
	SDMeta        SDMeta                `json:"sdmeta"`
	RelatedObject []RelatedObjectStruct `json:"relatedobject"`
	GitlabPointer []GitlabPointerStruct `json:"gitlabpointer"`
}

// deepcopy
func (in *InstallResource) DeepCopyInto(out *InstallResource) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = InstallResourceSpec{
		SDMeta:        in.Spec.SDMeta,
		RelatedObject: in.Spec.RelatedObject,
		GitlabPointer: in.Spec.GitlabPointer,
	}
}

// ----------------------------------------------------
// kubernetes dependencies
type InstallResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstallResourceSpec `json:"spec"`
}

type InstallResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstallResource `json:"items"`
}

func (in *InstallResource) DeepCopyObject() runtime.Object {
	out := InstallResource{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *InstallResourceList) DeepCopyObject() runtime.Object {
	out := InstallResourceList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]InstallResource, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}

// ----------------------------------------------------
// helper functions

const _INSTALLRESOURCE_STRING string = "installresources"

// list all installresources
func ListInstallResources() (irl InstallResourceList, err error) {
	if err = operator.CRD().Get().Resource(_INSTALLRESOURCE_STRING).Do(context.TODO()).Into(&irl); err != nil {
		err = fmt.Errorf("[%s]:ListInstallResources() error: %#v", _INSTALLRESOURCE_STRING, err)
	}
	return
}

// get specific installresource
func GetInstallResource(_namespace, _name string) (ir InstallResource, err error) {
	if err = operator.CRD().Get().Resource(_INSTALLRESOURCE_STRING).Namespace(_namespace).Name(_name).Do(context.TODO()).Into(&ir); err != nil {
		err = fmt.Errorf("[%s]:GetInstallResource(%s, %s) http error [get]: %#v", _INSTALLRESOURCE_STRING, _namespace, _name, err)
	}
	return
}

// create or update an installresource
func PushInstallResource(new_IR InstallResource) (ir InstallResource, err error) {
	if new_IR.APIVersion == ApiVersion() {
		if ir, err = GetInstallResource(new_IR.Namespace, new_IR.Name); err != nil {
			if err = operator.CRD().Post().Resource(_INSTALLRESOURCE_STRING).Namespace(new_IR.Namespace).Name(new_IR.Name).Body(&new_IR).Do(context.TODO()).Into(&ir); err != nil {
				err = fmt.Errorf("[%s]:PushInstallResource(%v) http error [post]: %#v", _INSTALLRESOURCE_STRING, new_IR, err)
			}
		} else {
			ir.Spec = new_IR.Spec
			if err = operator.CRD().Put().Resource(_INSTALLRESOURCE_STRING).Namespace(new_IR.Namespace).Name(new_IR.Name).Body(&ir).Do(context.TODO()).Into(&ir); err != nil {
				err = fmt.Errorf("[%s]:PushInstallResource(%v) http error [put]: %#v", _INSTALLRESOURCE_STRING, new_IR, err)
			}
		}
	}
	return
}

// delete an installresource
func DeleteInstallResource(_namespace, _name string) (ir InstallResource, err error) {
	if ir_k8s, e := GetInstallResource(_namespace, _name); e == nil {
		if ir_k8s.APIVersion == ApiVersion() {
			if err = operator.CRD().Delete().Resource(_INSTALLRESOURCE_STRING).Namespace(_namespace).Name(_name).Do(context.TODO()).Error(); err != nil {
				err = fmt.Errorf("[%s/%s]:DeleteInstallResource() http error [delete]: %#v", _namespace, _name, err)
			}
		}
	}
	return
}
