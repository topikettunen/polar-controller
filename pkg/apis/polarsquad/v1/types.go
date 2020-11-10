package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PolarDeployment struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// PolarDeploymentSpec is the custom resource spec
	Spec   PolarDeploymentSpec   `json:"spec"`
	Status PolarDeploymentStatus `json:"status"`
}

// PolarDeploymentSpec is the spec for a PolarDeployment resource
type PolarDeploymentSpec struct {
	AppName        string `json:"appName"`
	ContainerName  string `json:"containerName"`
	ImageName      string `json:"imageName"`
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// PolarDeploymentStatus is the status for a PolarDeployment resource
type PolarDeploymentStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PolarDeploymentList is a list of PolarDeployment resources
type PolarDeploymentList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []PolarDeployment `json:"items"`
}
