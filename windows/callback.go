//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build windows

package windows

import (
	"syscall"

	"github.com/energye/lcl/api"
	"github.com/energye/lcl/tool/ptr"
)

// 移除事件，释放相关的引用
func removeEventCallbackProc(f uintptr) uintptr {
	api.RemoveEventElement(f)
	return 0
}

func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	fn := api.PtrToElementValue(f)
	if fn != nil {
		if cb, ok := fn.(*callback); ok {
			defer func() {
				if err := recover(); err != nil {
					// TODO 增加回调到用户
					println("[ERROR] CallbackEvent:", cb.name, err.(error).Error())
				}
			}()
			getVal := func(i int) uintptr {
				return ptr.GetParamOf(i, args)
			}
			cb.cb(getVal)
		}
	}
	return 0
}

var (
	eventCallback       = syscall.NewCallback(eventCallbackProc)
	removeEventCallback = syscall.NewCallback(removeEventCallbackProc)
)
