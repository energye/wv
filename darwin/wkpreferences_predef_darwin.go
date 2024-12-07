//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package wv

/*
#cgo CFLAGS: -mmacosx-version-min=10.12 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.12 -framework Cocoa

#include "WebKit/WebKit.h"


void EnableDeveloperExtras(void* wkPreferences) {
    WKPreferences* preferences = (WKPreferences*)wkPreferences;
	[preferences setValue:@YES forKey:@"developerExtrasEnabled"];
}

*/
import "C"
import "unsafe"

func enableDevtools(wkPreferences WKPreferences) {
	C.EnableDeveloperExtras(unsafe.Pointer(wkPreferences))
}

func (m *TWKPreferences) EnableDevtools() {
	enableDevtools(m.Data())
}
