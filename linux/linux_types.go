//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import wvTypes "github.com/energye/wv/types/linux"

type TWkButtonEvent struct {
	EventType TGdkEventType    // TGdkEventType
	State     TGdkModifierType // TGdkModifierType
	Button    int32            // Integer
	X         int32            // Integer
	Y         int32            // Integer
	XRoot     int32            // Integer
	YRoot     int32            // Integer
	Time      int32            // Integer
}

type TGdkModifierType = wvTypes.TGdkModifierType
type TGdkEventType = wvTypes.TGdkModifierType
