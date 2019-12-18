// +build !ignore_autogenerated

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

// autogenerated by controller-gen object, do not modify manually

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAWS) DeepCopyInto(out *MinimalAWS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAWS.
func (in *MinimalAWS) DeepCopy() *MinimalAWS {
	if in == nil {
		return nil
	}
	out := new(MinimalAWS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinimalAWS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAWSList) DeepCopyInto(out *MinimalAWSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MinimalAWS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAWSList.
func (in *MinimalAWSList) DeepCopy() *MinimalAWSList {
	if in == nil {
		return nil
	}
	out := new(MinimalAWSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MinimalAWSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAWSSpec) DeepCopyInto(out *MinimalAWSSpec) {
	*out = *in
	out.CredentialsSecretRef = in.CredentialsSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAWSSpec.
func (in *MinimalAWSSpec) DeepCopy() *MinimalAWSSpec {
	if in == nil {
		return nil
	}
	out := new(MinimalAWSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinimalAWSStatus) DeepCopyInto(out *MinimalAWSStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinimalAWSStatus.
func (in *MinimalAWSStatus) DeepCopy() *MinimalAWSStatus {
	if in == nil {
		return nil
	}
	out := new(MinimalAWSStatus)
	in.DeepCopyInto(out)
	return out
}
