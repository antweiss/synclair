/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha0

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NamespaceCleanerSpec defines the desired state of NamespaceCleaner
type NamespaceCleanerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	NamespaceAgeDays int32 `json:"namespaceAgeDays,omitempty"`
	// This flag tells the controller if to delete namespaces that have objects other than secrets
    //  Defaults to false.
    // +optional
	DeleteIfNotEmpty bool  `json:"deleteIfNotEmpty,omitempty"`
	//Comma-separated list of regex name patterns to exclude from deletion 
	// +optional
	ExcludePatterns string  `json:"excludePatterns,omitempty"`

}

// NamespaceCleanerStatus defines the observed state of NamespaceCleaner
type NamespaceCleanerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +optional
	LastRun *metav1.Time `json:"lastRun,omitempty"`
	// +optional
	CleanedLast string `json:"cleanedLast,omitempty"`

}

// +kubebuilder:object:root=true

// NamespaceCleaner is the Schema for the namespacecleaners API
type NamespaceCleaner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespaceCleanerSpec   `json:"spec,omitempty"`
	Status NamespaceCleanerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamespaceCleanerList contains a list of NamespaceCleaner
type NamespaceCleanerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespaceCleaner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NamespaceCleaner{}, &NamespaceCleanerList{})
}
