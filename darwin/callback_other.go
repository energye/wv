//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

// debug build

//go:build !darwin && cgo
// +build !darwin,cgo

package darwin

import "C"

var (
	eventCallback       = uintptr(0)
	removeEventCallback = uintptr(0)
)
