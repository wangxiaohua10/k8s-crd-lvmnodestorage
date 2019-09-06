package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
//+genclient:nonNamespaced
//+genclient:onlyVerbs=get,list,create,update,patch,delete,deleteCollection,watch
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeLVMStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeLVMStorageSpec `json:"spec"`
}

type NodeLVMStorageSpec struct {
	// Message   string `json:"message"`
	// SomeValue *int32 `json:"someValue"`
	// LogicalVolumes  []LogicalVolume  `json:"logic_volumes,omitempty"`
	// PhysicalVolumes []PhysicalVolume `json:"physical_volumes,omitempty"`
	// VolumeGroups    []VolumeGroup    `json:"volume_groups,omitempty"`

	NodeName         string  `json:"node_name,omitempty"`
	LVName           string  `json:"lv_name,omitempty"`
	VGName           string  `json:"vg_name,omitempty"`
	Attrs            string  `json:"attrs,omitempty"`
	VolumeType       int     `json:"volume_type,omitempty"`
	Writable         bool    `json:"writable,omitempty"`
	AllocationPolicy int     `json:"allocation_type,omitempty"`
	Locked           bool    `json:"locked,omitempty"`
	FixedMinor       bool    `json:"fixed_minor,omitempty"`
	State            int     `json:"state,omitempty"`
	LVSize           float64 `json:"lv_size,omitempty"`
	VGSize           float64 `json:"vgsize,omitempty"`
	VGFree           float64 `json:"vgfree,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeLVMStorageList struct {
	// metav1.TypeMeta `json:",inline"`
	// metav1.ListMeta `json:"metadata"`
	// Items           []NodeLVMStorage `json:"items"`

	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeLVMStorage `json:"items"`
}

// type LogicalVolume struct {
// 	NodeName         string  `json:"node_name,omitempty"`
// 	LVName           string  `json:"lv_name,omitempty"`
// 	VGName           string  `json:"vg_name,omitempty"`
// 	Attrs            string  `json:"attrs,omitempty"`
// 	VolumeType       int     `json:"volume_type,omitempty"`
// 	Writable         bool    `json:"writable,omitempty"`
// 	AllocationPolicy int     `json:"allocation_type,omitempty"`
// 	Locked           bool    `json:"locked,omitempty"`
// 	FixedMinor       bool    `json:"fixed_minor,omitempty"`
// 	State            int     `json:"state,omitempty"`
// 	LVSize           float64 `json:"lv_size,omitempty"`
// }

// +genclient
// +genclient:noStatus
//+genclient:nonNamespaced
//+genclient:onlyVerbs=get,list,create,update,patch,delete,deleteCollection,watch
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeLvmPhysicalVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeLvmPhysicalVolumeSpec `json:"spec"`
}

type NodeLvmPhysicalVolumeSpec struct {
	NodeName string  `json:"node_name,omitempty"`
	PVName   string  `json:"pv_name,omitempty"`
	VGName   string  `json:"vg_name,omitempty"`
	Format   string  `json:"format,omitempty"`
	Attr     string  `json:"attr,omitempty"`
	PVSize   float64 `json:"pv_size,omitempty"`
	FreePE   float64 `json:"free_pe,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeLvmPhysicalVolumeList struct {
	// metav1.TypeMeta `json:",inline"`
	// metav1.ListMeta `json:"metadata"`
	// Items           []NodeLVMStorage `json:"items"`

	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeLvmPhysicalVolume `json:"items"`
}

// +genclient
// +genclient:noStatus
//+genclient:nonNamespaced
//+genclient:onlyVerbs=get,list,create,update,patch,delete,deleteCollection,watch
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeLvmVolumeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodeLvmVolumeGroupSpec `json:"spec"`
}

type NodeLvmVolumeGroupSpec struct {
	NodeName         string  `json:"node_name,omitempty"`
	VGName           string  `json:"vg_name,omitempty"`
	PhysicalVolumes  int     `json:"physical_volumes,omitempty"`
	LogicalVolumes   int     `json:"logical_volumes,omitempty"`
	Attrs            string  `json:"attrs,omitempty"`
	Writable         bool    `json:"writable,omitempty"`
	Resizable        bool    `json:"resizable,omitempty"`
	Exported         bool    `json:"exported,omitempty"`
	Partial          bool    `json:"partial,omitempty"`
	AllocationPolicy string  `json:"allocation_policy,omitempty"`
	Clustered        bool    `json:"clustered,omitempty"`
	VSize            float64 `json:"vsize,omitempty"`
	VFree            float64 `json:"vfree,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeLvmVolumeGroupList struct {
	// metav1.TypeMeta `json:",inline"`
	// metav1.ListMeta `json:"metadata"`
	// Items           []NodeLVMStorage `json:"items"`

	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeLvmVolumeGroup `json:"items"`
}
