//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import "github.com/energye/lcl/lcl"

// AsStrings Convert a pointer object to an existing class object
func AsStrings(obj uintptr) IStrings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	stream := new(TStrings)
	SetObjectInstance(stream, instance)
	return stream
}

// AsObject Convert a pointer object to an existing class object
func AsObject(obj interface{}) IObject {
	return lcl.AsObject(obj)
}
