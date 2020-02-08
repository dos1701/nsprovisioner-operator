package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NameSpaceProvisionerSpec defines the desired state of NameSpaceProvisioner
type NameSpaceProvisionerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// NameSpaceProvisionerStatus defines the observed state of NameSpaceProvisioner
type NameSpaceProvisionerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NameSpaceProvisioner is the Schema for the namespaceprovisioners API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=namespaceprovisioners,scope=Namespaced
type NameSpaceProvisioner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NameSpaceProvisionerSpec   `json:"spec,omitempty"`
	Status NameSpaceProvisionerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NameSpaceProvisionerList contains a list of NameSpaceProvisioner
type NameSpaceProvisionerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NameSpaceProvisioner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NameSpaceProvisioner{}, &NameSpaceProvisionerList{})
}
