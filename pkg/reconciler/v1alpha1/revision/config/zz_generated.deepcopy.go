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

package config

import (
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Controller) DeepCopyInto(out *Controller) {
	*out = *in
	if in.RegistriesSkippingTagResolving != nil {
		in, out := &in.RegistriesSkippingTagResolving, &out.RegistriesSkippingTagResolving
		*out = make(sets.String, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Controller.
func (in *Controller) DeepCopy() *Controller {
	if in == nil {
		return nil
	}
	out := new(Controller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observability) DeepCopyInto(out *Observability) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observability.
func (in *Observability) DeepCopy() *Observability {
	if in == nil {
		return nil
	}
	out := new(Observability)
	in.DeepCopyInto(out)
	return out
}
