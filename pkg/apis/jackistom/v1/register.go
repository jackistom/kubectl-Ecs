package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/jackistom/kubectl-Ecs/pkg/apis/jackistom"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   jackistom.GroupName,
	Version: jackistom.Version,
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&Sylixos{},
		&SylixosList{},
	)

	// 在scheme注册相应类型资源
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

