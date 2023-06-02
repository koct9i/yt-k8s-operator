//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChytSpec) DeepCopyInto(out *ChytSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChytSpec.
func (in *ChytSpec) DeepCopy() *ChytSpec {
	if in == nil {
		return nil
	}
	out := new(ChytSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerAgentsSpec) DeepCopyInto(out *ControllerAgentsSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerAgentsSpec.
func (in *ControllerAgentsSpec) DeepCopy() *ControllerAgentsSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerAgentsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNodesSpec) DeepCopyInto(out *DataNodesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNodesSpec.
func (in *DataNodesSpec) DeepCopy() *DataNodesSpec {
	if in == nil {
		return nil
	}
	out := new(DataNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoverySpec) DeepCopyInto(out *DiscoverySpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoverySpec.
func (in *DiscoverySpec) DeepCopy() *DiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(DiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedObjectMetadata) DeepCopyInto(out *EmbeddedObjectMetadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedObjectMetadata.
func (in *EmbeddedObjectMetadata) DeepCopy() *EmbeddedObjectMetadata {
	if in == nil {
		return nil
	}
	out := new(EmbeddedObjectMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedPersistentVolumeClaim) DeepCopyInto(out *EmbeddedPersistentVolumeClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.EmbeddedObjectMetadata.DeepCopyInto(&out.EmbeddedObjectMetadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedPersistentVolumeClaim.
func (in *EmbeddedPersistentVolumeClaim) DeepCopy() *EmbeddedPersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(EmbeddedPersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecNodesSpec) DeepCopyInto(out *ExecNodesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecNodesSpec.
func (in *ExecNodesSpec) DeepCopy() *ExecNodesSpec {
	if in == nil {
		return nil
	}
	out := new(ExecNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxiesSpec) DeepCopyInto(out *HTTPProxiesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxiesSpec.
func (in *HTTPProxiesSpec) DeepCopy() *HTTPProxiesSpec {
	if in == nil {
		return nil
	}
	out := new(HTTPProxiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]LocationSpec, len(*in))
		copy(*out, *in)
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]EmbeddedPersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableAntiAffinity != nil {
		in, out := &in.EnableAntiAffinity, &out.EnableAntiAffinity
		*out = new(bool)
		**out = **in
	}
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make([]LoggerSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationSpec) DeepCopyInto(out *LocationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationSpec.
func (in *LocationSpec) DeepCopy() *LocationSpec {
	if in == nil {
		return nil
	}
	out := new(LocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggerSpec) DeepCopyInto(out *LoggerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggerSpec.
func (in *LoggerSpec) DeepCopy() *LoggerSpec {
	if in == nil {
		return nil
	}
	out := new(LoggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MastersSpec) DeepCopyInto(out *MastersSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MastersSpec.
func (in *MastersSpec) DeepCopy() *MastersSpec {
	if in == nil {
		return nil
	}
	out := new(MastersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryTrackerSpec) DeepCopyInto(out *QueryTrackerSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryTrackerSpec.
func (in *QueryTrackerSpec) DeepCopy() *QueryTrackerSpec {
	if in == nil {
		return nil
	}
	out := new(QueryTrackerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RPCProxiesSpec) DeepCopyInto(out *RPCProxiesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
	if in.ServiceType != nil {
		in, out := &in.ServiceType, &out.ServiceType
		*out = new(corev1.ServiceType)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RPCProxiesSpec.
func (in *RPCProxiesSpec) DeepCopy() *RPCProxiesSpec {
	if in == nil {
		return nil
	}
	out := new(RPCProxiesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulersSpec) DeepCopyInto(out *SchedulersSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulersSpec.
func (in *SchedulersSpec) DeepCopy() *SchedulersSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpytSpec) DeepCopyInto(out *SpytSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpytSpec.
func (in *SpytSpec) DeepCopy() *SpytSpec {
	if in == nil {
		return nil
	}
	out := new(SpytSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TabletCellBundleInfo) DeepCopyInto(out *TabletCellBundleInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TabletCellBundleInfo.
func (in *TabletCellBundleInfo) DeepCopy() *TabletCellBundleInfo {
	if in == nil {
		return nil
	}
	out := new(TabletCellBundleInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TabletNodesSpec) DeepCopyInto(out *TabletNodesSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TabletNodesSpec.
func (in *TabletNodesSpec) DeepCopy() *TabletNodesSpec {
	if in == nil {
		return nil
	}
	out := new(TabletNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UISpec) DeepCopyInto(out *UISpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UISpec.
func (in *UISpec) DeepCopy() *UISpec {
	if in == nil {
		return nil
	}
	out := new(UISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YQLAgentSpec) DeepCopyInto(out *YQLAgentSpec) {
	*out = *in
	in.InstanceSpec.DeepCopyInto(&out.InstanceSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YQLAgentSpec.
func (in *YQLAgentSpec) DeepCopy() *YQLAgentSpec {
	if in == nil {
		return nil
	}
	out := new(YQLAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ytsaurus) DeepCopyInto(out *Ytsaurus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ytsaurus.
func (in *Ytsaurus) DeepCopy() *Ytsaurus {
	if in == nil {
		return nil
	}
	out := new(Ytsaurus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ytsaurus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YtsaurusList) DeepCopyInto(out *YtsaurusList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ytsaurus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YtsaurusList.
func (in *YtsaurusList) DeepCopy() *YtsaurusList {
	if in == nil {
		return nil
	}
	out := new(YtsaurusList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *YtsaurusList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YtsaurusSpec) DeepCopyInto(out *YtsaurusSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.AdminCredentials != nil {
		in, out := &in.AdminCredentials, &out.AdminCredentials
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.ExtraPodAnnotations != nil {
		in, out := &in.ExtraPodAnnotations, &out.ExtraPodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Discovery.DeepCopyInto(&out.Discovery)
	in.PrimaryMasters.DeepCopyInto(&out.PrimaryMasters)
	if in.SecondaryMasters != nil {
		in, out := &in.SecondaryMasters, &out.SecondaryMasters
		*out = make([]MastersSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPProxies != nil {
		in, out := &in.HTTPProxies, &out.HTTPProxies
		*out = make([]HTTPProxiesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RPCProxies != nil {
		in, out := &in.RPCProxies, &out.RPCProxies
		*out = make([]RPCProxiesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataNodes != nil {
		in, out := &in.DataNodes, &out.DataNodes
		*out = make([]DataNodesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExecNodes != nil {
		in, out := &in.ExecNodes, &out.ExecNodes
		*out = make([]ExecNodesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schedulers != nil {
		in, out := &in.Schedulers, &out.Schedulers
		*out = new(SchedulersSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerAgents != nil {
		in, out := &in.ControllerAgents, &out.ControllerAgents
		*out = new(ControllerAgentsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TabletNodes != nil {
		in, out := &in.TabletNodes, &out.TabletNodes
		*out = make([]TabletNodesSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Chyt != nil {
		in, out := &in.Chyt, &out.Chyt
		*out = new(ChytSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryTrackers != nil {
		in, out := &in.QueryTrackers, &out.QueryTrackers
		*out = new(QueryTrackerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Spyt != nil {
		in, out := &in.Spyt, &out.Spyt
		*out = new(SpytSpec)
		**out = **in
	}
	if in.YQLAgents != nil {
		in, out := &in.YQLAgents, &out.YQLAgents
		*out = new(YQLAgentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UI != nil {
		in, out := &in.UI, &out.UI
		*out = new(UISpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YtsaurusSpec.
func (in *YtsaurusSpec) DeepCopy() *YtsaurusSpec {
	if in == nil {
		return nil
	}
	out := new(YtsaurusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *YtsaurusStatus) DeepCopyInto(out *YtsaurusStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpdateConditions != nil {
		in, out := &in.UpdateConditions, &out.UpdateConditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SavedTabletCellBundles != nil {
		in, out := &in.SavedTabletCellBundles, &out.SavedTabletCellBundles
		*out = make([]TabletCellBundleInfo, len(*in))
		copy(*out, *in)
	}
	if in.MasterMonitoringPaths != nil {
		in, out := &in.MasterMonitoringPaths, &out.MasterMonitoringPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new YtsaurusStatus.
func (in *YtsaurusStatus) DeepCopy() *YtsaurusStatus {
	if in == nil {
		return nil
	}
	out := new(YtsaurusStatus)
	in.DeepCopyInto(out)
	return out
}
