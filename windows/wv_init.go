//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"os"
)

// Init Webview2
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	lcl.Init(libs, resources)
	defer func() {
		if err := recover(); err != nil {
			println(err)
			os.Exit(1)
		}
	}()
	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()

	// 注册 Webview2 对象事件回调
	api.SetEventCallback(eventCallback, api.EctWV)
	// 注册 Webview2 对象事件回调移除
	api.SetEventCallback(removeEventCallback, api.EctWVRemove)

}
