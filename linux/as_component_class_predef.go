//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

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

func AsWkCookieManagerDelegateEvent(obj uintptr) IWkCookieManagerDelegateEvent {
	return AsWkWebview(obj)
}

func AsWkSchemeRequestDelegateEvent(obj uintptr) IWkSchemeRequestDelegateEvent {
	return AsWkWebview(obj)
}
