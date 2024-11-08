//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/inits"
	"unsafe"
)

// getParam 从指定索引和地址获取事件中的参数
func getParamOf(index int, ptr uintptr) uintptr {
	return *(*uintptr)(unsafePointer(ptr + uintptr(index)*unsafe.Sizeof(ptr)))
}

// 移除事件，释放相关的引用
func removeEventCallbackProc(f uintptr) uintptr {
	RemoveEventElement(f)
	return 0
}

// 回调过程
func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	fn := PtrToElementValue(f)
	if fn != nil {
		// 获取值
		getVal := func(i int) uintptr {
			return getParamOf(i, args)
		}

		// 指针
		getPtr := func(i int) unsafePointer {
			return unsafePointer(getVal(i))
		}

		switch fn.(type) {
		case TWkProcessMessageEvent:
			fn.(TWkProcessMessageEvent)(AsWKUserContentController(getPtr(0)), "", "")

		default:
		}
	}
	return 0
}

// Init
//
//	Webkit2初始化
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	inits.Init(libs, resources)
	SetWKEventCallback(eventCallback)
	SetWKRemoveEventCallback(removeEventCallback)
}
