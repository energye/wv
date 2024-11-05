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
	"github.com/energye/lcl/lcl"
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
		case TWkLoadChangeEvent:
			fn.(TWkLoadChangeEvent)(lcl.AsObject(getPtr(0)), WebKitLoadEvent(getVal(1)))

		case TWkExecuteScriptFinishedEvent:
			fn.(TWkExecuteScriptFinishedEvent)(lcl.AsObject(getPtr(0)), AsWkJSValue(getVal(1)))

		case TWkLoadFailedEvent:
			failingUri := GoStr(getVal(2))
			error_ := GoStr(getVal(3))
			result := (*bool)(getPtr(4))
			*result = fn.(TWkLoadFailedEvent)(lcl.AsObject(getPtr(0)), WebKitLoadEvent(getVal(1)), failingUri, error_)

		case TWkURISchemeRequestEvent:
			fn.(TWkURISchemeRequestEvent)(lcl.AsObject(getPtr(0)), WebKitURISchemeRequest(getVal(1)))

		case TWkProcessMessageEvent:
			fn.(TWkProcessMessageEvent)(lcl.AsObject(getPtr(0)), AsWkJSValue(getVal(1)), TWkProcessId(getVal(2)))

		case TWkMousePressEvent:
			event := *(*TWkButtonEvent)(getPtr(1))
			result := (*bool)(getPtr(2))
			*result = fn.(TWkMousePressEvent)(lcl.AsObject(getPtr(0)), event)

		case TWkMouseReleaseEvent:
			event := *(*TWkButtonEvent)(getPtr(1))
			result := (*bool)(getPtr(2))
			*result = fn.(TWkMouseReleaseEvent)(lcl.AsObject(getPtr(0)), event)

		case TWkGetAcceptPolicyFinishEvent:
			fn.(TWkGetAcceptPolicyFinishEvent)(lcl.AsObject(getPtr(0)), WebKitCookieAcceptPolicy(getVal(1)), GoStr(getVal(2)))

		case TWkAddCookieFinishEvent:
			fn.(TWkAddCookieFinishEvent)(lcl.AsObject(getPtr(0)), GoBool(getVal(1)), GoStr(getVal(2)))

		case TWkGetCookiesFinishEvent:
			fn.(TWkGetCookiesFinishEvent)(lcl.AsObject(getPtr(0)), PList(getVal(1)), GoStr(getVal(2)))

		case TWkDeleteCookieFinishEvent:
			fn.(TWkDeleteCookieFinishEvent)(lcl.AsObject(getPtr(0)), GoBool(getVal(1)), GoStr(getVal(2)))

		case TWkDecidePolicyEvent:
			result := (*bool)(getPtr(3))
			*result = fn.(TWkDecidePolicyEvent)(lcl.AsObject(getPtr(0)), WebKitPolicyDecision(getVal(1)), WebKitPolicyDecisionType(getVal(2)))

		case TWkWebProcessTerminatedEvent:
			fn.(TWkWebProcessTerminatedEvent)(lcl.AsObject(getPtr(0)), WebKitWebProcessTerminationReason(getVal(1)))

		case TWkContextMenuEvent:
			result := (*bool)(getPtr(3))
			*result = fn.(TWkContextMenuEvent)(lcl.AsObject(getPtr(0)), WebKitContextMenu(getVal(1)), PWkAction(getVal(2)))

		case TWkContextMenuCommandEvent:
			fn.(TWkContextMenuCommandEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)))

		case TWkContextMenuDismissedEvent:
			fn.(TWkContextMenuDismissedEvent)(lcl.AsObject(getPtr(0)))

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
