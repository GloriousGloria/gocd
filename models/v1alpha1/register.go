package v1alpha1

import (
	"fmt"

	"github.com/gloriousgloria/gocd/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupVersion string = "v1alpha1"

func ApiVersion() string {
	return fmt.Sprintf("%s/%s", models.GroupName, GroupVersion)
}

var SchemeGroupVersion = schema.GroupVersion{Group: models.GroupName, Version: GroupVersion}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&ServiceDeclarationList{},
		&ServiceDeclaration{},
		&InstallResourceList{},
		&InstallResource{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
