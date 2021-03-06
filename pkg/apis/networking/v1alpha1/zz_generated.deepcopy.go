// +build !ignore_autogenerated

/*
Copyright 2019 The Knative Authors

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	duck_v1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngress) DeepCopyInto(out *ClusterIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngress.
func (in *ClusterIngress) DeepCopy() *ClusterIngress {
	if in == nil {
		return nil
	}
	out := new(ClusterIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngressBackend) DeepCopyInto(out *ClusterIngressBackend) {
	*out = *in
	out.ServicePort = in.ServicePort
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngressBackend.
func (in *ClusterIngressBackend) DeepCopy() *ClusterIngressBackend {
	if in == nil {
		return nil
	}
	out := new(ClusterIngressBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngressBackendSplit) DeepCopyInto(out *ClusterIngressBackendSplit) {
	*out = *in
	out.ClusterIngressBackend = in.ClusterIngressBackend
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngressBackendSplit.
func (in *ClusterIngressBackendSplit) DeepCopy() *ClusterIngressBackendSplit {
	if in == nil {
		return nil
	}
	out := new(ClusterIngressBackendSplit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngressList) DeepCopyInto(out *ClusterIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngressList.
func (in *ClusterIngressList) DeepCopy() *ClusterIngressList {
	if in == nil {
		return nil
	}
	out := new(ClusterIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngressRule) DeepCopyInto(out *ClusterIngressRule) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTPClusterIngressRuleValue)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngressRule.
func (in *ClusterIngressRule) DeepCopy() *ClusterIngressRule {
	if in == nil {
		return nil
	}
	out := new(ClusterIngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngressTLS) DeepCopyInto(out *ClusterIngressTLS) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngressTLS.
func (in *ClusterIngressTLS) DeepCopy() *ClusterIngressTLS {
	if in == nil {
		return nil
	}
	out := new(ClusterIngressTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPClusterIngressPath) DeepCopyInto(out *HTTPClusterIngressPath) {
	*out = *in
	if in.Splits != nil {
		in, out := &in.Splits, &out.Splits
		*out = make([]ClusterIngressBackendSplit, len(*in))
		copy(*out, *in)
	}
	if in.AppendHeaders != nil {
		in, out := &in.AppendHeaders, &out.AppendHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		if *in == nil {
			*out = nil
		} else {
			*out = new(HTTPRetry)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPClusterIngressPath.
func (in *HTTPClusterIngressPath) DeepCopy() *HTTPClusterIngressPath {
	if in == nil {
		return nil
	}
	out := new(HTTPClusterIngressPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPClusterIngressRuleValue) DeepCopyInto(out *HTTPClusterIngressRuleValue) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]HTTPClusterIngressPath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPClusterIngressRuleValue.
func (in *HTTPClusterIngressRuleValue) DeepCopy() *HTTPClusterIngressRuleValue {
	if in == nil {
		return nil
	}
	out := new(HTTPClusterIngressRuleValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRetry) DeepCopyInto(out *HTTPRetry) {
	*out = *in
	if in.PerTryTimeout != nil {
		in, out := &in.PerTryTimeout, &out.PerTryTimeout
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Duration)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRetry.
func (in *HTTPRetry) DeepCopy() *HTTPRetry {
	if in == nil {
		return nil
	}
	out := new(HTTPRetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]ClusterIngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ClusterIngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSpec.
func (in *IngressSpec) DeepCopy() *IngressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressStatus) DeepCopyInto(out *IngressStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(duck_v1alpha1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		if *in == nil {
			*out = nil
		} else {
			*out = new(LoadBalancerStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressStatus.
func (in *IngressStatus) DeepCopy() *IngressStatus {
	if in == nil {
		return nil
	}
	out := new(IngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerIngressStatus) DeepCopyInto(out *LoadBalancerIngressStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerIngressStatus.
func (in *LoadBalancerIngressStatus) DeepCopy() *LoadBalancerIngressStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerIngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]LoadBalancerIngressStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}
