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
			name, data := GoStr(getVal(2)), GoStr(getVal(3))
			fn.(TWkProcessMessageEvent)(AsObject(getPtr(0)), WKUserContentController(getVal(1)), name, data)
		case TWKDecidePolicyForNavigationActionPreferences:
			actionPolicy := (*WKNavigationActionPolicy)(getPtr(2))
			preferences := (*WKWebpagePreferences)(getPtr(3))
			fn.(TWKDecidePolicyForNavigationActionPreferences)(AsObject(getPtr(0)), WKNavigationAction(getVal(1)),
				actionPolicy, preferences)
		case TWKDecidePolicyForNavigationResponse:
			fn.(TWKDecidePolicyForNavigationResponse)(AsObject(getPtr(0)), WKNavigationResponse(getVal(1)))
		case TWkStartProvisionalNavigation:
			fn.(TWkStartProvisionalNavigation)(AsObject(getPtr(0)))
		case TWkReceiveServerRedirectForProvisionalNavigation:
			fn.(TWkReceiveServerRedirectForProvisionalNavigation)(AsObject(getPtr(0)))
		case TWkFailProvisionalNavigationWithError:
			fn.(TWkFailProvisionalNavigationWithError)(AsObject(getPtr(0)))
		case TWkCommitNavigation:
			fn.(TWkCommitNavigation)(AsObject(getPtr(0)))
		case TWkFinishNavigation:
			fn.(TWkFinishNavigation)(AsObject(getPtr(0)))
		case TWkFailNavigationWithError:
			fn.(TWkFailNavigationWithError)(AsObject(getPtr(0)))
		case TWkWebContentProcessDidTerminate:
			fn.(TWkWebContentProcessDidTerminate)(AsObject(getPtr(0)))
		case TWkNavigationActionDidBecomeDownload:
			fn.(TWkNavigationActionDidBecomeDownload)(AsObject(getPtr(0)))
		case TWkNavigationResponseDidBecomeDownload:
			fn.(TWkNavigationResponseDidBecomeDownload)(AsObject(getPtr(0)))
		case TWKStartURLSchemeTask:
			fn.(TWKStartURLSchemeTask)(AsObject(getPtr(0)))
		case TWKStopURLSchemeTask:
			fn.(TWKStopURLSchemeTask)(AsObject(getPtr(0)))
		case TWKCreateWebView:
			fn.(TWKCreateWebView)(AsObject(getPtr(0)))
		case TWKRunJavaScriptAlert:
			fn.(TWKRunJavaScriptAlert)(AsObject(getPtr(0)))
		case TWKRunJavaScriptConfirmCompletion:
			fn.(TWKRunJavaScriptConfirmCompletion)(AsObject(getPtr(0)))
		case TWKRunJavaScriptTextInputCompletion:
			fn.(TWKRunJavaScriptTextInputCompletion)(AsObject(getPtr(0)))
		case TWKWebViewDidClose:
			fn.(TWKWebViewDidClose)(AsObject(getPtr(0)))
		case TWKDownloadCancelCompletionHandler:
			fn.(TWKDownloadCancelCompletionHandler)(AsObject(getPtr(0)))
		case TWKDownloadDecideDestinationUsingResponseSuggestedFilename:
			fn.(TWKDownloadDecideDestinationUsingResponseSuggestedFilename)(AsObject(getPtr(0)))
		case TWKDownloadWillPerformHTTPRedirectionNewRequest:
			fn.(TWKDownloadWillPerformHTTPRedirectionNewRequest)(AsObject(getPtr(0)))
		case TWKDownloadFinish:
			fn.(TWKDownloadFinish)(AsObject(getPtr(0)))
		case TWKDownloadFailWithError:
			fn.(TWKDownloadFailWithError)(AsObject(getPtr(0)))
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
