//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build linux && !cgo

package linux

import (
	"github.com/energye/lcl/api/imports"
)

var (
	eventCallback       = imports.NewCallback(eventCallbackProc)
	removeEventCallback = imports.NewCallback(removeEventCallbackProc)
)
