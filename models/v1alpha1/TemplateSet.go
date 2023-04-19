package v1alpha1

import (
	"context"
	"fmt"

	"github.com/jnnkrdb/k8s/operator"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type TemplateSetSpec struct {
	Content string `json:"content"`
}

// deepcopy
func (in *TemplateSet) DeepCopyInto(out *TemplateSet) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = TemplateSetSpec{
		Content: in.Spec.Content,
	}
}

// ----------------------------------------------------
// kubernetes dependencies
type TemplateSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TemplateSetSpec `json:"spec"`
}

type TemplateSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TemplateSet `json:"items"`
}

func (in *TemplateSet) DeepCopyObject() runtime.Object {
	out := TemplateSet{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *TemplateSetList) DeepCopyObject() runtime.Object {
	out := TemplateSetList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		out.Items = make([]TemplateSet, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}

// ----------------------------------------------------
// helper functions

const _TEMPLATESET_STRING string = "templatesets"

// list all existing service declarations
func ListTemplateSets() (sdl TemplateSetList, err error) {
	if err = operator.CRD().Get().Resource(_TEMPLATESET_STRING).Do(context.TODO()).Into(&sdl); err != nil {
		err = fmt.Errorf("[%s]:ListTemplateSets() error listing servicerequests: %#v", _TEMPLATESET_STRING, err)
	}
	return
}

// get a specific servicedeclaration
func GetTemplateSet(namespace, name string) (sd TemplateSet, err error) {
	if err = operator.CRD().Get().Resource(_TEMPLATESET_STRING).Namespace(namespace).Name(name).Do(context.TODO()).Into(&sd); err != nil {
		err = fmt.Errorf("[%s]:GetTemplateSet(%s, %s) error listing servicerequests: %#v", _TEMPLATESET_STRING, namespace, name, err)
	}
	return
}

// get templatesets by labels
func GetTemplateByLabels(labels map[string]string) (sd TemplateSet, err error) {
	var _list TemplateSetList
	if _list, err = ListTemplateSets(); err == nil {
		// check a specific template
		var totalEqual bool = false
		for _, item := range _list.Items {
			var equals bool = true

			// check the labels
			for lblKey := range labels {

				// check if the requested label exists in the item from k8s
				if _, ok := item.Labels[lblKey]; !ok {
					equals = false
					break
				}

				if labels[lblKey] != item.Labels[lblKey] {
					equals = false
					break
				}
			}

			// if the items equal, the for loop will be broken and the current sd will be returned
			if equals {
				sd = item
				totalEqual = true
				break
			}
		}
		if !totalEqual {
			err = fmt.Errorf("there is no templateset, which matches the required labels")
		}
	}

	return
}
