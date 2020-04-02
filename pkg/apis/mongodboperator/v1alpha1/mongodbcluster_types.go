package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Type of mongodb statefulset
type Type string

//Phase of mongodb statefulset
type Phase string

//Restore information struct
type Restore struct {
	BackupName string `json:"backupName"`
	Secret     string `json:"secret"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MongodbClusterSpec defines the desired state of MongodbCluster
// +k8s:openapi-gen=true
type MongodbClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Type defines which type of MongoDB deployment the resource should create
	Type Type `json:"type"`
	// Version defines which version of MongoDB will be used
	Version         string        `json:"version"`
	Shrads          int32         `json:"shrads,omitempty"`
	MongodbImage    string        `json:"mongodbImage,omitempty"`
	SidecarImage    string        `json:"sidecarImage,omitempty"`
	MetricsImage    string        `json:"metricsImage,omitempty"`
	ImagePullPolicy v1.PullPolicy `json:"imagePullPolicy,omitempty"`
	// +listType
	ImagePullSecrets    []v1.LocalObjectReference     `json:"imagePullSecrets,omitempty"`
	DataVolumeClaimSpec *v1.PersistentVolumeClaimSpec `json:"dataVolumeClaimSpec,omitempty"`
	DeletePVCs          bool                          `json:"deletePVCs,omitempty"`
	Resources           *v1.ResourceRequirements      `json:"resources,omitempty"`
	SidecarResources    *v1.ResourceRequirements      `json:"sidecarResources,omitempty"`
	// +listType
	SidecarEnv []v1.EnvVar `json:"sidecarEnv,omitempty"`
	// +listType
	MongodbEnv         []v1.EnvVar `json:"mongodbEnv,omitempty"`
	ServiceAccountName string      `json:"serviceAccountName,omitempty"`
	Restore            *Restore    `json:"restore,omitempty"`
}

// MongodbClusterStatus defines the observed state of MongodbCluster
// +k8s:openapi-gen=true
type MongodbClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	MongoURI string `json:"mongoUri"`
	Phase    Phase  `json:"phase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongodbCluster is the Schema for the mongodbclusters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mongodbclusters,scope=Namespaced
type MongodbCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Cluster string `json:"cluster"`

	Spec   MongodbClusterSpec   `json:"spec,omitempty"`
	Status MongodbClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MongodbClusterList contains a list of MongodbCluster
type MongodbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MongodbCluster{}, &MongodbClusterList{})
}
