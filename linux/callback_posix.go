//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows && cgo
// +build !windows,cgo

package wv

//// #cgo darwin CFLAGS: -mmacosx-version-min=10.10 -DMACOSX_DEPLOYMENT_TARGET=10.10
// #cgo darwin CFLAGS: -mmacosx-version-min=10.10
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.10
// #include <stdint.h>
//
// extern void* doWKEventCallbackProc(uintptr_t f, void* args, long argcount);
// static void* doGetWKEventCallbackAddr() {
//    return &doWKEventCallbackProc;
// }
//
// extern void* doRemoveWKEventCallbackProc(uintptr_t ptr);
// static void* doRemoveWKEventCallbackAddr() {
//    return &doRemoveWKEventCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doWKEventCallbackProc
func doWKEventCallbackProc(f C.uintptr_t, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return nil
}

//export doRemoveWKEventCallbackProc
func doRemoveWKEventCallbackProc(ptr C.uintptr_t) unsafe.Pointer {
	removeEventCallbackProc(uintptr(ptr))
	return nil
}

var (
	eventCallback       = uintptr(C.doGetWKEventCallbackAddr())
	removeEventCallback = uintptr(C.doRemoveWKEventCallbackAddr())
)
