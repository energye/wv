//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import wvTypes "github.com/energye/wv/types/darwin"

type TWKButtonEvent struct {
	EventType  int32   // Integer
	WindowX    float64 // Double
	WindowY    float64 // Double
	LocalX     float64 // Double
	LocalY     float64 // Double
	Timestamp  float64 // Double
	ClickCount int32   // Integer
	DeltaX     float64 // Double
	DeltaY     float64 // Double
	Event      NSEvent // NSEvent
}

type NSURLRequest uintptr
type NSURLResponse uintptr
type NSObject uintptr
type NSEvent uintptr
type NSHTTPURLResponse uintptr
type NSProgress uintptr
type NSURL uintptr
type NSMutableURLRequest uintptr
type NSURLAuthenticationChallenge uintptr
type NSURLCredential uintptr
type NSURLCredentialPersistence uintptr
type NSURLRequestCachePolicy = wvTypes.NSURLRequestCachePolicy
type NSURLRequestNetworkServiceType = wvTypes.NSURLRequestNetworkServiceType
type NSTimeInterval = float64
