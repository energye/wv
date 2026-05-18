//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package darwin

import (
	"sync"

	"github.com/energye/lcl/api"
)

var initOnce sync.Once

// Init Webview2
func Init() {
	initOnce.Do(func() {
		// 注册 Webkit 对象事件回调
		api.SetEventCallback(eventCallback, api.EctWV)
		// 注册 Webkit 对象事件回调移除
		api.SetEventCallback(removeEventCallback, api.EctWVRemove)
	})
}
