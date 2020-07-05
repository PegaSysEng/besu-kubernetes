package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BesuNodeSpec defines the desired state of BesuNode
type BesuNodeSpec struct {

	// Name of the node
	// +optional
	Name string `json:"name"`

	// Type of node, takes one of the values : Bootnode, Validator, Member
	// +kubebuilder:validation:Enum:["Member", "Bootnode", "Validator"]
	// +kubebuider:default:"Member"
	// +optional
	Type string `json:"type"`

	// Number of replica pods corresponding to this node
	// +optional
	// +kubebuider:default:1
	Replicas int32 `json:"replicas"`

	// Public key
	// +optional
	PubKey string `json:"pubkey"`

	// Private key
	// +optional
	PrivKey string `json:"privkey"`

	// Besu Image Configuration
	// +kubebuider:default:{repository: hyperledger/besu; tag: 1.4.6; pullPolicy: IfNotPresent}
	// +optional
	Image Image `json:"image"`

	// 	Size of the Volume
	// +kubebuider:default:1Gi
	// +optional
	PVCSizeLimit string `json:"pvcSizeLimit"`

	// 	Storage class of the Volume
	// +kubebuider:default:standard
	// +optional
	PVCStorageClass string `json:"pvcStorageClass"`

	// Requests and limits
	// +optional
	// +kubebuider:default:{memRequest: 1024Mi; cpuRequest: 100m; memLimit: 2048Mi; cpuLimit: 500m}
	Resources Resources `json:"resources"`

	// P2P
	// +optional
	P2P PortConfig `json:"p2p"`

	// RPC
	// +optional
	RPC PortConfig `json:"rpc"`

	// WS
	// +optional
	WS PortConfig `json:"ws"`

	// GraphQl
	// +optional
	GraphQl PortConfig `json:"graphql"`

	// Defaults to ["*"]
	// +optional
	HTTPWhitelist string `json:"httpwhitelist"`

	// +optional
	Metrics PortConfig `json:"metrics"`

	// +optional
	Bootnodes int `json:"bootnodes"`
}

// PortConfig defines port configurations of different types of ports
type PortConfig struct {
	// Port is enabled or not
	// +kubebuider:default:true
	Enabled bool `json:"enabled"`

	// Host
	// +kubebuider:default:0.0.0.0
	Host string `json:"host"`

	// Port
	Port int `json:"port"`

	// +optional
	API string `json:"api"`

	CorsOrigins string `json:"corsOrigins"`

	// +optional
	AuthenticationEnabled bool `json:"authenticationEnabled"`

	// +optional
	Discovery bool `json:"discovery"`
}

// Resources defines requests and limits of CPU and memory
type Resources struct {

	// Memory Request
	// +kubebuider:default:1024Mi
	MemRequest string `json:"memRequest"`

	// CPU Request
	// +kubebuider:default:100m
	CPURequest string `json:"cpuRequest"`

	// Memory Limit
	// +kubebuider:default:2048Mi
	MemLimit string `json:"memLimit"`

	// CPU Limit
	// +kubebuider:default:500m
	CPULimit string `json:"cpuLimit"`
}

// BesuNodeStatus defines the observed state of BesuNode
type BesuNodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BesuNode is the Schema for the besunodes API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=besunodes,scope=Namespaced
type BesuNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BesuNodeSpec   `json:"spec,omitempty"`
	Status BesuNodeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BesuNodeList contains a list of BesuNode
type BesuNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BesuNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BesuNode{}, &BesuNodeList{})
}