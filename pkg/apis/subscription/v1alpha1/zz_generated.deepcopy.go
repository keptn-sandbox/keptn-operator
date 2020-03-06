// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnSubscription) DeepCopyInto(out *KeptnSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnSubscription.
func (in *KeptnSubscription) DeepCopy() *KeptnSubscription {
	if in == nil {
		return nil
	}
	out := new(KeptnSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnSubscriptionList) DeepCopyInto(out *KeptnSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnSubscriptionList.
func (in *KeptnSubscriptionList) DeepCopy() *KeptnSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(KeptnSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnSubscriptionSpec) DeepCopyInto(out *KeptnSubscriptionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnSubscriptionSpec.
func (in *KeptnSubscriptionSpec) DeepCopy() *KeptnSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnSubscriptionStatus) DeepCopyInto(out *KeptnSubscriptionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnSubscriptionStatus.
func (in *KeptnSubscriptionStatus) DeepCopy() *KeptnSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}