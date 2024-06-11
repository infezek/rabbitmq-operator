//go:build !ignore_autogenerated

/*
Copyright 2024 Ezequiel Lopes.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPASpec) DeepCopyInto(out *HPASpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPASpec.
func (in *HPASpec) DeepCopy() *HPASpec {
	if in == nil {
		return nil
	}
	out := new(HPASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitHPA) DeepCopyInto(out *RabbitHPA) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitHPA.
func (in *RabbitHPA) DeepCopy() *RabbitHPA {
	if in == nil {
		return nil
	}
	out := new(RabbitHPA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitHPA) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitHPAList) DeepCopyInto(out *RabbitHPAList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitHPA, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitHPAList.
func (in *RabbitHPAList) DeepCopy() *RabbitHPAList {
	if in == nil {
		return nil
	}
	out := new(RabbitHPAList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitHPAList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitHPASpec) DeepCopyInto(out *RabbitHPASpec) {
	*out = *in
	out.Rabbit = in.Rabbit
	out.HPA = in.HPA
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitHPASpec.
func (in *RabbitHPASpec) DeepCopy() *RabbitHPASpec {
	if in == nil {
		return nil
	}
	out := new(RabbitHPASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitHPAStatus) DeepCopyInto(out *RabbitHPAStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitHPAStatus.
func (in *RabbitHPAStatus) DeepCopy() *RabbitHPAStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitHPAStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitSpec) DeepCopyInto(out *RabbitSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitSpec.
func (in *RabbitSpec) DeepCopy() *RabbitSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitSpec)
	in.DeepCopyInto(out)
	return out
}
