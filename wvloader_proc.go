//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

// GlobalWebView2Loader
// Global instance of TWVLoader used to simplify the WebView2 initialization and destruction.
func GlobalWebView2Loader(AOwner IComponent) IWVLoader {
	r1 := predefImportAPI().SysCallN(7, GetObjectUintptr(AOwner))
	if r1 != 0 {
		return AsWVLoader(r1)
	}
	return nil
}
