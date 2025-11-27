//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build !windows && cgo
// +build !windows,cgo

package darwin

// #include <stdint.h>
//
// extern void* doWVEventCallbackProc(uintptr_t f, void* args, long argcount);
// static void* doWVEventCallbackAddr() {
//    return &doWVEventCallbackProc;
// }
//
// extern void* doWVRemoveEventCallbackProc(uintptr_t ptr);
// static void* doWVRemoveEventCallbackAddr() {
//    return &doWVRemoveEventCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doWVEventCallbackProc
func doWVEventCallbackProc(f C.uintptr_t, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return nil
}

//export doWVRemoveEventCallbackProc
func doWVRemoveEventCallbackProc(ptr C.uintptr_t) unsafe.Pointer {
	removeEventCallbackProc(uintptr(ptr))
	return nil
}

var (
	eventCallback       = uintptr(C.doWVEventCallbackAddr())
	removeEventCallback = uintptr(C.doWVRemoveEventCallbackAddr())
)
