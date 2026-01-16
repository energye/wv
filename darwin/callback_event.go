//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
	wvTypes "github.com/energye/wv/types/darwin"
)

type callback struct {
	name string
	cb   func(getVal func(i int) uintptr)
}

func getPtr(val uintptr) base.UnsafePointer {
	return base.UnsafePointer(val)
}

func makeTWKCreateWebView(cb TWKCreateWebView) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKCreateWebView",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(sender: TObject; configuration_: WKWebViewConfiguration; navigationAction: WKNavigationAction; windowFeatures: WKWindowFeatures): WKWebView;
			sender := lcl.AsObject(getVal(0))
			configuration := wvTypes.WKWebViewConfiguration(getVal(1))
			navigationAction := wvTypes.WKNavigationAction(getVal(2))
			windowFeatures := wvTypes.WKWindowFeatures(getVal(3))
			ret := cb(sender, configuration, navigationAction, windowFeatures)
			*(*wvTypes.WKWebView)(getPtr(getVal(4))) = ret
		},
	}
}

func makeTWKDecidePolicyForNavigationActionPreferences(cb TWKDecidePolicyForNavigationActionPreferences) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKDecidePolicyForNavigationActionPreferences",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(sender: TObject; navigationAction: WKNavigationAction; var actionPolicy: WKNavigationActionPolicy; var preferences: WKWebpagePreferences);
			sender := lcl.AsObject(getVal(0))
			navigationAction := wvTypes.WKNavigationAction(getVal(1))
			actionPolicy := (*wvTypes.WKNavigationActionPolicy)(getPtr(getVal(2)))
			preferences := (*wvTypes.WKWebpagePreferences)(getPtr(getVal(3)))
			cb(sender, navigationAction, actionPolicy, preferences)
		},
	}
}

func makeTWKDecidePolicyForNavigationResponse(cb TWKDecidePolicyForNavigationResponse) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKDecidePolicyForNavigationResponse",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(sender: TObject; navigationResponse: WKNavigationResponse; var responsePolicy: WKNavigationResponsePolicy);
			sender := lcl.AsObject(getVal(0))
			navigationResponse := wvTypes.WKNavigationResponse(getVal(1))
			responsePolicy := (*wvTypes.WKNavigationResponsePolicy)(getPtr(getVal(2)))
			cb(sender, navigationResponse, responsePolicy)
		},
	}
}

func makeTWKDownloadCancelCompletionHandler(cb TWKDownloadCancelCompletionHandler) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKDownloadCancelCompletionHandler",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(sender: TObject; download: WKDownload; aData: Pointer; aLength: Integer );
			sender := lcl.AsObject(getVal(0))
			download := wvTypes.WKDownload(getVal(1))
			data := uintptr(getVal(2))
			length := int32(getVal(3))
			cb(sender, download, data, length)
		},
	}
}

func makeTWKDownloadDecideDestinationUsingResponseSuggestedFilename(cb TWKDownloadDecideDestinationUsingResponseSuggestedFilename) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKDownloadDecideDestinationUsingResponseSuggestedFilename",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(sender: TObject; download: WKDownload; response: NSURLResponse; suggestedFilename: string);
			sender := lcl.AsObject(getVal(0))
			download := wvTypes.WKDownload(getVal(1))
			response := NSURLResponse(getVal(2))
			suggestedFilename := api.GoStr(getVal(3))
			cb(sender, download, response, suggestedFilename)
		},
	}
}

func makeTWKDownloadFailWithError(cb TWKDownloadFailWithError) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKDownloadFailWithError",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(sender: TObject; download: WKDownload; error: string; aData: Pointer; aLength: Integer );
			sender := lcl.AsObject(getVal(0))
			download := wvTypes.WKDownload(getVal(1))
			error_ := api.GoStr(getVal(2))
			data := uintptr(getVal(3))
			length := int32(getVal(4))
			cb(sender, download, error_, data, length)
		},
	}
}

func makeTWKDownloadFinish(cb TWKDownloadFinish) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKDownloadFinish",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(sender: TObject; download: WKDownload);
			sender := lcl.AsObject(getVal(0))
			download := wvTypes.WKDownload(getVal(1))
			cb(sender, download)
		},
	}
}

func makeTWKDownloadWillPerformHTTPRedirectionNewRequest(cb TWKDownloadWillPerformHTTPRedirectionNewRequest) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKDownloadWillPerformHTTPRedirectionNewRequest",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(sender: TObject; download: WKDownload; response: NSHTTPURLResponse; request: NSURLRequest);
			sender := lcl.AsObject(getVal(0))
			download := wvTypes.WKDownload(getVal(1))
			response := NSHTTPURLResponse(getVal(2))
			request := NSURLRequest(getVal(3))
			cb(sender, download, response, request)
		},
	}
}

func makeTWKEvaluateJavaScriptCallback(cb TWKEvaluateJavaScriptCallback) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKEvaluateJavaScriptCallback",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(data: Pointer; error_: string);
			data := uintptr(getVal(0))
			error_ := api.GoStr(getVal(1))
			cb(data, error_)
		},
	}
}

func makeTWKRunJavaScriptAlert(cb TWKRunJavaScriptAlert) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKRunJavaScriptAlert",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(sender: TObject; message_: string; frame: WKFrameInfo);
			sender := lcl.AsObject(getVal(0))
			message := api.GoStr(getVal(1))
			frame := wvTypes.WKFrameInfo(getVal(2))
			cb(sender, message, frame)
		},
	}
}

func makeTWKRunJavaScriptConfirmCompletion(cb TWKRunJavaScriptConfirmCompletion) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKRunJavaScriptConfirmCompletion",
		cb: func(getVal func(i int) uintptr) {
			// 3 : function(sender: TObject; message_: string; frame: WKFrameInfo): Boolean;
			sender := lcl.AsObject(getVal(0))
			message := api.GoStr(getVal(1))
			frame := wvTypes.WKFrameInfo(getVal(2))
			ret := cb(sender, message, frame)
			*(*bool)(getPtr(getVal(3))) = ret
		},
	}
}

func makeTWKRunJavaScriptTextInputCompletion(cb TWKRunJavaScriptTextInputCompletion) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKRunJavaScriptTextInputCompletion",
		cb: func(getVal func(i int) uintptr) {
			// 4 : function(sender: TObject; prompt: string; defaultText: string; frame: WKFrameInfo): string;
			sender := lcl.AsObject(getVal(0))
			prompt := api.GoStr(getVal(1))
			defaultText := api.GoStr(getVal(2))
			frame := wvTypes.WKFrameInfo(getVal(3))
			ret := cb(sender, prompt, defaultText, frame)
			*(*uintptr)(getPtr(getVal(4))) = api.PasStr(ret)
		},
	}
}

func makeTWKStartURLSchemeTask(cb TWKStartURLSchemeTask) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKStartURLSchemeTask",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(sender: TObject; urlSchemeTask: WKURLSchemeTask);
			sender := lcl.AsObject(getVal(0))
			urlSchemeTask := wvTypes.WKURLSchemeTask(getVal(1))
			cb(sender, urlSchemeTask)
		},
	}
}

func makeTWKStopURLSchemeTask(cb TWKStopURLSchemeTask) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKStopURLSchemeTask",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(sender: TObject; urlSchemeTask: WKURLSchemeTask);
			sender := lcl.AsObject(getVal(0))
			urlSchemeTask := wvTypes.WKURLSchemeTask(getVal(1))
			cb(sender, urlSchemeTask)
		},
	}
}

func makeTWKWebViewDidClose(cb TWKWebViewDidClose) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWKWebViewDidClose",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(sender: TObject);
			sender := lcl.AsObject(getVal(0))
			cb(sender)
		},
	}
}

func makeTWkCommitNavigation(cb TWkCommitNavigation) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkCommitNavigation",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(sender: TObject; navigation: WKNavigation);
			sender := lcl.AsObject(getVal(0))
			navigation := wvTypes.WKNavigation(getVal(1))
			cb(sender, navigation)
		},
	}
}

func makeTWkFailNavigationWithError(cb TWkFailNavigationWithError) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkFailNavigationWithError",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(sender: TObject; navigation: WKNavigation; error: string);
			sender := lcl.AsObject(getVal(0))
			navigation := wvTypes.WKNavigation(getVal(1))
			error_ := api.GoStr(getVal(2))
			cb(sender, navigation, error_)
		},
	}
}

func makeTWkFailProvisionalNavigationWithError(cb TWkFailProvisionalNavigationWithError) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkFailProvisionalNavigationWithError",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(sender: TObject; navigation: WKNavigation; error: string);
			sender := lcl.AsObject(getVal(0))
			navigation := wvTypes.WKNavigation(getVal(1))
			error_ := api.GoStr(getVal(2))
			cb(sender, navigation, error_)
		},
	}
}

func makeTWkFinishNavigation(cb TWkFinishNavigation) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkFinishNavigation",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(sender: TObject; navigation: WKNavigation);
			sender := lcl.AsObject(getVal(0))
			navigation := wvTypes.WKNavigation(getVal(1))
			cb(sender, navigation)
		},
	}
}

func makeTWkMouseEvent(cb TWkMouseEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkMouseEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; event: PWKButtonEvent);
			sender := lcl.AsObject(getVal(0))
			event := *(*TWKButtonEvent)(getPtr(getVal(1)))
			cb(sender, event)
		},
	}
}

func makeTWkNavigationActionDidBecomeDownload(cb TWkNavigationActionDidBecomeDownload) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkNavigationActionDidBecomeDownload",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(sender: TObject; navigationAction: WKNavigationAction; download: WKDownload);
			sender := lcl.AsObject(getVal(0))
			navigationAction := wvTypes.WKNavigationAction(getVal(1))
			download := wvTypes.WKDownload(getVal(2))
			cb(sender, navigationAction, download)
		},
	}
}

func makeTWkNavigationResponseDidBecomeDownload(cb TWkNavigationResponseDidBecomeDownload) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkNavigationResponseDidBecomeDownload",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(sender: TObject; navigationResponse: WKNavigationResponse; download: WKDownload);
			sender := lcl.AsObject(getVal(0))
			navigationResponse := wvTypes.WKNavigationResponse(getVal(1))
			download := wvTypes.WKDownload(getVal(2))
			cb(sender, navigationResponse, download)
		},
	}
}

func makeTWkProcessMessageEvent(cb TWkProcessMessageEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkProcessMessageEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(sender: TObject; userContentController: WKUserContentController; name, data: string);
			sender := lcl.AsObject(getVal(0))
			userContentController := wvTypes.WKUserContentController(getVal(1))
			name := api.GoStr(getVal(2))
			data := api.GoStr(getVal(3))
			cb(sender, userContentController, name, data)
		},
	}
}

func makeTWkReceiveServerRedirectForProvisionalNavigation(cb TWkReceiveServerRedirectForProvisionalNavigation) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkReceiveServerRedirectForProvisionalNavigation",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(sender: TObject; navigation: WKNavigation);
			sender := lcl.AsObject(getVal(0))
			navigation := wvTypes.WKNavigation(getVal(1))
			cb(sender, navigation)
		},
	}
}

func makeTWkStartProvisionalNavigation(cb TWkStartProvisionalNavigation) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkStartProvisionalNavigation",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(sender: TObject; navigation: WKNavigation);
			sender := lcl.AsObject(getVal(0))
			navigation := wvTypes.WKNavigation(getVal(1))
			cb(sender, navigation)
		},
	}
}

func makeTWkWebContentProcessDidTerminate(cb TWkWebContentProcessDidTerminate) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TWkWebContentProcessDidTerminate",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(sender: TObject);
			sender := lcl.AsObject(getVal(0))
			cb(sender)
		},
	}
}
