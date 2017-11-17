// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

<<<<<<< HEAD
=======
// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PartialObjectMetadata).DeepCopyInto(out.(*PartialObjectMetadata))
			return nil
		}, InType: reflect.TypeOf(&PartialObjectMetadata{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PartialObjectMetadataList).DeepCopyInto(out.(*PartialObjectMetadataList))
			return nil
		}, InType: reflect.TypeOf(&PartialObjectMetadataList{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Table).DeepCopyInto(out.(*Table))
			return nil
		}, InType: reflect.TypeOf(&Table{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TableColumnDefinition).DeepCopyInto(out.(*TableColumnDefinition))
			return nil
		}, InType: reflect.TypeOf(&TableColumnDefinition{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TableOptions).DeepCopyInto(out.(*TableOptions))
			return nil
		}, InType: reflect.TypeOf(&TableOptions{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TableRow).DeepCopyInto(out.(*TableRow))
			return nil
		}, InType: reflect.TypeOf(&TableRow{})},
		{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TableRowCondition).DeepCopyInto(out.(*TableRowCondition))
			return nil
		}, InType: reflect.TypeOf(&TableRowCondition{})},
	}
}

>>>>>>> Initial dep workover
// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialObjectMetadata) DeepCopyInto(out *PartialObjectMetadata) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialObjectMetadata.
func (in *PartialObjectMetadata) DeepCopy() *PartialObjectMetadata {
	if in == nil {
		return nil
	}
	out := new(PartialObjectMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartialObjectMetadata) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartialObjectMetadataList) DeepCopyInto(out *PartialObjectMetadataList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*PartialObjectMetadata, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(PartialObjectMetadata)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartialObjectMetadataList.
func (in *PartialObjectMetadataList) DeepCopy() *PartialObjectMetadataList {
	if in == nil {
		return nil
	}
	out := new(PartialObjectMetadataList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartialObjectMetadataList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.ColumnDefinitions != nil {
		in, out := &in.ColumnDefinitions, &out.ColumnDefinitions
		*out = make([]TableColumnDefinition, len(*in))
		copy(*out, *in)
	}
	if in.Rows != nil {
		in, out := &in.Rows, &out.Rows
		*out = make([]TableRow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Table) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableColumnDefinition) DeepCopyInto(out *TableColumnDefinition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableColumnDefinition.
func (in *TableColumnDefinition) DeepCopy() *TableColumnDefinition {
	if in == nil {
		return nil
	}
	out := new(TableColumnDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableOptions) DeepCopyInto(out *TableOptions) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableOptions.
func (in *TableOptions) DeepCopy() *TableOptions {
	if in == nil {
		return nil
	}
	out := new(TableOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TableOptions) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableRow) DeepCopyInto(out *TableRow) {
	clone := in.DeepCopy()
	*out = *clone
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableRowCondition) DeepCopyInto(out *TableRowCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableRowCondition.
func (in *TableRowCondition) DeepCopy() *TableRowCondition {
	if in == nil {
		return nil
	}
	out := new(TableRowCondition)
	in.DeepCopyInto(out)
	return out
}
