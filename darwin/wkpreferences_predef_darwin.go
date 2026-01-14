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

// :predefine:

package darwin

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "WebKit/WebKit.h"

void EnableDeveloperExtras(void* wkPreferences) {
    WKPreferences* preferences = (WKPreferences*)wkPreferences;
	[preferences setValue:@YES forKey:@"developerExtrasEnabled"];
}
*/
import "C"
import (
	"unsafe"

	wvTypes "github.com/energye/wv/types/darwin"
)

func EnableDevtools(wkPreferences wvTypes.WKPreferences) {
	C.EnableDeveloperExtras(unsafe.Pointer(wkPreferences))
}

func (m *TWkPreferences) EnableDevtools() {
	EnableDevtools(m.Data())
}
