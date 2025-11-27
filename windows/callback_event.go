//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	wvTypes "github.com/energye/wv/types/windows"
)

type callback struct {
	name string
	cb   func(getVal func(i int) uintptr)
}

func getPtr(val uintptr) base.UnsafePointer {
	return base.UnsafePointer(val)
}

func makeTDragDropEvent(cb TDragDropEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDragDropEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender, Source: TObject; X,Y: Integer);
			sender := lcl.AsObject(getVal(0))
			source := lcl.AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, source, X, Y)
		},
	}
}

func makeTDragOverEvent(cb TDragOverEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TDragOverEvent",
		cb: func(getVal func(i int) uintptr) {
			// 6 : procedure(Sender, Source: TObject; X,Y: Integer; State: TDragState; var Accept: Boolean);
			sender := lcl.AsObject(getVal(0))
			source := lcl.AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			state := types.TDragState(getVal(4))
			accept := (*bool)(getPtr(getVal(5)))
			cb(sender, source, X, Y, state, accept)
		},
	}
}

func makeTEndDragEvent(cb TEndDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TEndDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender, Target: TObject; X,Y: Integer);
			sender := lcl.AsObject(getVal(0))
			target := lcl.AsObject(getVal(1))
			X := int32(getVal(2))
			Y := int32(getVal(3))
			cb(sender, target, X, Y)
		},
	}
}

func makeTLoaderBrowserProcessExitedEvent(cb TLoaderBrowserProcessExitedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLoaderBrowserProcessExitedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aEnvironment: ICoreWebView2Environment; const aArgs: ICoreWebView2BrowserProcessExitedEventArgs);
			sender := lcl.AsObject(getVal(0))
			environment := AsCoreWebView2Environment(getVal(1))
			args := AsCoreWebView2BrowserProcessExitedEventArgs(getVal(2))
			cb(sender, environment, args)
		},
	}
}

func makeTLoaderGetCustomSchemesEvent(cb TLoaderGetCustomSchemesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLoaderGetCustomSchemesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var aCustomSchemes: TWVCustomSchemeInfoArray);
			sender := lcl.AsObject(getVal(0))
			var customSchemes IWVCustomSchemeInfoArrayWrap
			customSchemes = AsWVCustomSchemeInfoArrayWrap(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &customSchemes)
			if customSchemes != nil {
				*(*uintptr)(getPtr(getVal(1))) = customSchemes.Instance()
			}
		},
	}
}

func makeTLoaderNewBrowserVersionAvailableEvent(cb TLoaderNewBrowserVersionAvailableEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLoaderNewBrowserVersionAvailableEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aEnvironment: ICoreWebView2Environment);
			sender := lcl.AsObject(getVal(0))
			environment := AsCoreWebView2Environment(getVal(1))
			cb(sender, environment)
		},
	}
}

func makeTLoaderNotifyEvent(cb TLoaderNotifyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLoaderNotifyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(Sender: TObject);
			sender := lcl.AsObject(getVal(0))
			cb(sender)
		},
	}
}

func makeTLoaderProcessInfosChangedEvent(cb TLoaderProcessInfosChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TLoaderProcessInfosChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aEnvironment: ICoreWebView2Environment);
			sender := lcl.AsObject(getVal(0))
			environment := AsCoreWebView2Environment(getVal(1))
			cb(sender, environment)
		},
	}
}

func makeTNotifyEvent(cb TNotifyEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TNotifyEvent",
		cb: func(getVal func(i int) uintptr) {
			// 1 : procedure(Sender: TObject);
			sender := lcl.AsObject(getVal(0))
			cb(sender)
		},
	}
}

func makeTOnAcceleratorKeyPressedEvent(cb TOnAcceleratorKeyPressedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAcceleratorKeyPressedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aController: ICoreWebView2Controller; const aArgs: ICoreWebView2AcceleratorKeyPressedEventArgs);
			sender := lcl.AsObject(getVal(0))
			controller := AsCoreWebView2Controller(getVal(1))
			args := AsCoreWebView2AcceleratorKeyPressedEventArgs(getVal(2))
			cb(sender, controller, args)
		},
	}
}

func makeTOnAddScriptToExecuteOnDocumentCreatedCompletedEvent(cb TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: wvstring);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := api.GoStr(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnBasicAuthenticationRequestedEvent(cb TOnBasicAuthenticationRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBasicAuthenticationRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2BasicAuthenticationRequestedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2BasicAuthenticationRequestedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnBrowserExtensionEnableCompletedEvent(cb TOnBrowserExtensionEnableCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserExtensionEnableCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aExtensionID: wvstring);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			extensionID := api.GoStr(getVal(2))
			cb(sender, errorCode, extensionID)
		},
	}
}

func makeTOnBrowserExtensionRemoveCompletedEvent(cb TOnBrowserExtensionRemoveCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserExtensionRemoveCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aExtensionID: wvstring);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			extensionID := api.GoStr(getVal(2))
			cb(sender, errorCode, extensionID)
		},
	}
}

func makeTOnBrowserProcessExitedEvent(cb TOnBrowserProcessExitedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBrowserProcessExitedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aEnvironment: ICoreWebView2Environment; const aArgs: ICoreWebView2BrowserProcessExitedEventArgs);
			sender := lcl.AsObject(getVal(0))
			environment := AsCoreWebView2Environment(getVal(1))
			args := AsCoreWebView2BrowserProcessExitedEventArgs(getVal(2))
			cb(sender, environment, args)
		},
	}
}

func makeTOnBytesReceivedChangedEvent(cb TOnBytesReceivedChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnBytesReceivedChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aDownloadOperation: ICoreWebView2DownloadOperation; aDownloadID: integer);
			sender := lcl.AsObject(getVal(0))
			downloadOperation := AsCoreWebView2DownloadOperation(getVal(1))
			downloadID := int32(getVal(2))
			cb(sender, downloadOperation, downloadID)
		},
	}
}

func makeTOnCallDevToolsProtocolMethodCompletedEvent(cb TOnCallDevToolsProtocolMethodCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCallDevToolsProtocolMethodCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: wvstring; aExecutionID: integer);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := api.GoStr(getVal(2))
			executionID := int32(getVal(3))
			cb(sender, errorCode, result, executionID)
		},
	}
}

func makeTOnCapturePreviewCompletedEvent(cb TOnCapturePreviewCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCapturePreviewCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aErrorCode: HRESULT);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			cb(sender, errorCode)
		},
	}
}

func makeTOnClearBrowsingDataCompletedEvent(cb TOnClearBrowsingDataCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClearBrowsingDataCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aErrorCode: HRESULT);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			cb(sender, errorCode)
		},
	}
}

func makeTOnClearCacheCompletedEvent(cb TOnClearCacheCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClearCacheCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			cb(sender, result)
		},
	}
}

func makeTOnClearDataForOriginCompletedEvent(cb TOnClearDataForOriginCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClearDataForOriginCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			cb(sender, result)
		},
	}
}

func makeTOnClientCertificateRequestedEvent(cb TOnClientCertificateRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnClientCertificateRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2ClientCertificateRequestedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ClientCertificateRequestedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnCompMsgEvent(cb TOnCompMsgEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCompMsgEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; var aMessage: TMessage; var aHandled: Boolean);
			sender := lcl.AsObject(getVal(0))
			message := (*types.TMessage)(getPtr(getVal(1)))
			handled := (*bool)(getPtr(getVal(2)))
			cb(sender, message, handled)
		},
	}
}

func makeTOnContentLoadingEvent(cb TOnContentLoadingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContentLoadingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2ContentLoadingEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ContentLoadingEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnContextMenuRequestedEvent(cb TOnContextMenuRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnContextMenuRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2ContextMenuRequestedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ContextMenuRequestedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnCustomItemSelectedEvent(cb TOnCustomItemSelectedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnCustomItemSelectedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aMenuItem: ICoreWebView2ContextMenuItem);
			sender := lcl.AsObject(getVal(0))
			menuItem := AsCoreWebView2ContextMenuItem(getVal(1))
			cb(sender, menuItem)
		},
	}
}

func makeTOnDOMContentLoadedEvent(cb TOnDOMContentLoadedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDOMContentLoadedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2DOMContentLoadedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2DOMContentLoadedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnDevToolsProtocolEventReceivedEvent(cb TOnDevToolsProtocolEventReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDevToolsProtocolEventReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 5 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2DevToolsProtocolEventReceivedEventArgs; const aEventName : wvstring; aEventID : integer);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(getVal(2))
			eventName := api.GoStr(getVal(3))
			eventID := int32(getVal(4))
			cb(sender, webView, args, eventName, eventID)
		},
	}
}

func makeTOnDownloadStartingEvent(cb TOnDownloadStartingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadStartingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2DownloadStartingEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2DownloadStartingEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnDownloadStateChangedEvent(cb TOnDownloadStateChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnDownloadStateChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aDownloadOperation: ICoreWebView2DownloadOperation; aDownloadID: integer);
			sender := lcl.AsObject(getVal(0))
			downloadOperation := AsCoreWebView2DownloadOperation(getVal(1))
			downloadID := int32(getVal(2))
			cb(sender, downloadOperation, downloadID)
		},
	}
}

func makeTOnEstimatedEndTimeChangedEvent(cb TOnEstimatedEndTimeChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnEstimatedEndTimeChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aDownloadOperation: ICoreWebView2DownloadOperation; aDownloadID: integer);
			sender := lcl.AsObject(getVal(0))
			downloadOperation := AsCoreWebView2DownloadOperation(getVal(1))
			downloadID := int32(getVal(2))
			cb(sender, downloadOperation, downloadID)
		},
	}
}

func makeTOnExecuteScriptCompletedEvent(cb TOnExecuteScriptCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnExecuteScriptCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResultObjectAsJson: wvstring; aExecutionID: integer);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			resultObjectAsJson := api.GoStr(getVal(2))
			executionID := int32(getVal(3))
			cb(sender, errorCode, resultObjectAsJson, executionID)
		},
	}
}

func makeTOnExecuteScriptWithResultCompletedEvent(cb TOnExecuteScriptWithResultCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnExecuteScriptWithResultCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; errorCode: HResult; const aResult: ICoreWebView2ExecuteScriptResult; aExecutionID : integer);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := AsCoreWebView2ExecuteScriptResult(getVal(2))
			executionID := int32(getVal(3))
			cb(sender, errorCode, result, executionID)
		},
	}
}

func makeTOnFaviconChangedEvent(cb TOnFaviconChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFaviconChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: IUnknown);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := lcl.AsUnknown(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnFrameContentLoadingEvent(cb TOnFrameContentLoadingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameContentLoadingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; const aArgs: ICoreWebView2ContentLoadingEventArgs; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2ContentLoadingEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			cb(sender, frame, args, frameID)
		},
	}
}

func makeTOnFrameCreatedEvent(cb TOnFrameCreatedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameCreatedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2FrameCreatedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2FrameCreatedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnFrameDOMContentLoadedEvent(cb TOnFrameDOMContentLoadedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameDOMContentLoadedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; const aArgs: ICoreWebView2DOMContentLoadedEventArgs; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2DOMContentLoadedEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			cb(sender, frame, args, frameID)
		},
	}
}

func makeTOnFrameDestroyedEvent(cb TOnFrameDestroyedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameDestroyedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			frameID := uint32(getVal(2))
			cb(sender, frame, frameID)
		},
	}
}

func makeTOnFrameNameChangedEvent(cb TOnFrameNameChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameNameChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			frameID := uint32(getVal(2))
			cb(sender, frame, frameID)
		},
	}
}

func makeTOnFrameNavigationCompletedEvent(cb TOnFrameNavigationCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameNavigationCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; const aArgs: ICoreWebView2NavigationCompletedEventArgs; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2NavigationCompletedEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			cb(sender, frame, args, frameID)
		},
	}
}

func makeTOnFrameNavigationStartingEvent(cb TOnFrameNavigationStartingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameNavigationStartingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; const aArgs: ICoreWebView2NavigationStartingEventArgs; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2NavigationStartingEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			cb(sender, frame, args, frameID)
		},
	}
}

func makeTOnFramePermissionRequestedEvent(cb TOnFramePermissionRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFramePermissionRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; const aArgs: ICoreWebView2PermissionRequestedEventArgs2; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			args := ICoreWebView2PermissionRequestedEventArgs2(getVal(2))
			frameID := uint32(getVal(3))
			cb(sender, frame, args, frameID)
		},
	}
}

func makeTOnFrameScreenCaptureStartingEvent(cb TOnFrameScreenCaptureStartingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameScreenCaptureStartingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; const aArgs: ICoreWebView2ScreenCaptureStartingEventArgs; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2ScreenCaptureStartingEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			cb(sender, frame, args, frameID)
		},
	}
}

func makeTOnFrameWebMessageReceivedEvent(cb TOnFrameWebMessageReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnFrameWebMessageReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; const aFrame: ICoreWebView2Frame; const aArgs: ICoreWebView2WebMessageReceivedEventArgs; aFrameID: cardinal);
			sender := lcl.AsObject(getVal(0))
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2WebMessageReceivedEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			cb(sender, frame, args, frameID)
		},
	}
}

func makeTOnGetCookiesCompletedEvent(cb TOnGetCookiesCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetCookiesCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: ICoreWebView2CookieList);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := AsCoreWebView2CookieList(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnGetCustomSchemesEvent(cb TOnGetCustomSchemesEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetCustomSchemesEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var aCustomSchemes: TWVCustomSchemeInfoArray);
			sender := lcl.AsObject(getVal(0))
			var customSchemes IWVCustomSchemeInfoArrayWrap
			customSchemes = AsWVCustomSchemeInfoArrayWrap(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &customSchemes)
			if customSchemes != nil {
				*(*uintptr)(getPtr(getVal(1))) = customSchemes.Instance()
			}
		},
	}
}

func makeTOnGetFaviconCompletedEvent(cb TOnGetFaviconCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetFaviconCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: IStream);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := lcl.AsStreamAdapter(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnGetNonDefaultPermissionSettingsCompletedEvent(cb TOnGetNonDefaultPermissionSettingsCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetNonDefaultPermissionSettingsCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: ICoreWebView2PermissionSettingCollectionView);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := AsCoreWebView2PermissionSettingCollectionView(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnGetProcessExtendedInfosCompletedEvent(cb TOnGetProcessExtendedInfosCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnGetProcessExtendedInfosCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: ICoreWebView2ProcessExtendedInfoCollection);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := AsCoreWebView2ProcessExtendedInfoCollection(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnIgnoreCertificateErrorsCompletedEvent(cb TOnIgnoreCertificateErrorsCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIgnoreCertificateErrorsCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			cb(sender, result)
		},
	}
}

func makeTOnInitializationErrorEvent(cb TOnInitializationErrorEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnInitializationErrorEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aErrorMessage: wvstring);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			errorMessage := api.GoStr(getVal(2))
			cb(sender, errorCode, errorMessage)
		},
	}
}

func makeTOnIsDefaultDownloadDialogOpenChangedEvent(cb TOnIsDefaultDownloadDialogOpenChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsDefaultDownloadDialogOpenChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aWebView: ICoreWebView2);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			cb(sender, webView)
		},
	}
}

func makeTOnIsDocumentPlayingAudioChangedEvent(cb TOnIsDocumentPlayingAudioChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsDocumentPlayingAudioChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aWebView: ICoreWebView2);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			cb(sender, webView)
		},
	}
}

func makeTOnIsMutedChangedEvent(cb TOnIsMutedChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnIsMutedChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aWebView: ICoreWebView2);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			cb(sender, webView)
		},
	}
}

func makeTOnLaunchingExternalUriSchemeEvent(cb TOnLaunchingExternalUriSchemeEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnLaunchingExternalUriSchemeEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2LaunchingExternalUriSchemeEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2LaunchingExternalUriSchemeEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnMoveFocusRequestedEvent(cb TOnMoveFocusRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnMoveFocusRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aController: ICoreWebView2Controller; const aArgs: ICoreWebView2MoveFocusRequestedEventArgs);
			sender := lcl.AsObject(getVal(0))
			controller := AsCoreWebView2Controller(getVal(1))
			args := AsCoreWebView2MoveFocusRequestedEventArgs(getVal(2))
			cb(sender, controller, args)
		},
	}
}

func makeTOnNavigationCompletedEvent(cb TOnNavigationCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNavigationCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2NavigationCompletedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2NavigationCompletedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnNavigationStartingEvent(cb TOnNavigationStartingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNavigationStartingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2NavigationStartingEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2NavigationStartingEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnNewWindowRequestedEvent(cb TOnNewWindowRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNewWindowRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2NewWindowRequestedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2NewWindowRequestedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnNonClientRegionChangedEvent(cb TOnNonClientRegionChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNonClientRegionChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aController : ICoreWebView2CompositionController; const aArgs : ICoreWebView2NonClientRegionChangedEventArgs);
			sender := lcl.AsObject(getVal(0))
			controller := AsCoreWebView2CompositionController(getVal(1))
			args := AsCoreWebView2NonClientRegionChangedEventArgs(getVal(2))
			cb(sender, controller, args)
		},
	}
}

func makeTOnNotificationCloseRequestedEvent(cb TOnNotificationCloseRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNotificationCloseRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aNotification: ICoreWebView2Notification; const aArgs: IUnknown);
			sender := lcl.AsObject(getVal(0))
			notification := AsCoreWebView2Notification(getVal(1))
			args := lcl.AsUnknown(getVal(2))
			cb(sender, notification, args)
		},
	}
}

func makeTOnNotificationReceivedEvent(cb TOnNotificationReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnNotificationReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView : ICoreWebView2; const aArgs : ICoreWebView2NotificationReceivedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2NotificationReceivedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnOfflineCompletedEvent(cb TOnOfflineCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnOfflineCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			cb(sender, result)
		},
	}
}

func makeTOnPermissionRequestedEvent(cb TOnPermissionRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPermissionRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2PermissionRequestedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2PermissionRequestedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnPrintCompletedEvent(cb TOnPrintCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; aResult: TWVPrintStatus);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := wvTypes.TWVPrintStatus(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnPrintToPdfCompletedEvent(cb TOnPrintToPdfCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintToPdfCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := api.GoBool(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnPrintToPdfStreamCompletedEvent(cb TOnPrintToPdfStreamCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnPrintToPdfStreamCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: IStream);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := lcl.AsStreamAdapter(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnProcessFailedEvent(cb TOnProcessFailedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProcessFailedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2ProcessFailedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ProcessFailedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnProcessInfosChangedEvent(cb TOnProcessInfosChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProcessInfosChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aEnvironment: ICoreWebView2Environment);
			sender := lcl.AsObject(getVal(0))
			environment := AsCoreWebView2Environment(getVal(1))
			cb(sender, environment)
		},
	}
}

func makeTOnProfileAddBrowserExtensionCompletedEvent(cb TOnProfileAddBrowserExtensionCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProfileAddBrowserExtensionCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: ICoreWebView2BrowserExtension);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := AsCoreWebView2BrowserExtension(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnProfileDeletedEvent(cb TOnProfileDeletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProfileDeletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aProfile: ICoreWebView2Profile);
			sender := lcl.AsObject(getVal(0))
			profile := AsCoreWebView2Profile(getVal(1))
			cb(sender, profile)
		},
	}
}

func makeTOnProfileGetBrowserExtensionsCompletedEvent(cb TOnProfileGetBrowserExtensionsCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnProfileGetBrowserExtensionsCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: ICoreWebView2BrowserExtensionList);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := AsCoreWebView2BrowserExtensionList(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnRefreshIgnoreCacheCompletedEvent(cb TOnRefreshIgnoreCacheCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRefreshIgnoreCacheCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResultObjectAsJson: wvstring);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			resultObjectAsJson := api.GoStr(getVal(2))
			cb(sender, errorCode, resultObjectAsJson)
		},
	}
}

func makeTOnRetrieveHTMLCompletedEvent(cb TOnRetrieveHTMLCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRetrieveHTMLCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aResult: boolean; const aHTML: wvstring);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			hTML := api.GoStr(getVal(2))
			cb(sender, result, hTML)
		},
	}
}

func makeTOnRetrieveMHTMLCompletedEvent(cb TOnRetrieveMHTMLCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRetrieveMHTMLCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aResult: boolean; const aMHTML: wvstring);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			mHTML := api.GoStr(getVal(2))
			cb(sender, result, mHTML)
		},
	}
}

func makeTOnRetrieveTextCompletedEvent(cb TOnRetrieveTextCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnRetrieveTextCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aResult: boolean; const aText: wvstring);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			text := api.GoStr(getVal(2))
			cb(sender, result, text)
		},
	}
}

func makeTOnSaveAsUIShowingEvent(cb TOnSaveAsUIShowingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSaveAsUIShowingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2SaveAsUIShowingEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2SaveAsUIShowingEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnSaveFileSecurityCheckStartingEvent(cb TOnSaveFileSecurityCheckStartingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSaveFileSecurityCheckStartingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2SaveFileSecurityCheckStartingEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2SaveFileSecurityCheckStartingEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnScreenCaptureStartingEvent(cb TOnScreenCaptureStartingEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnScreenCaptureStartingEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2ScreenCaptureStartingEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ScreenCaptureStartingEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnScriptDialogOpeningEvent(cb TOnScriptDialogOpeningEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnScriptDialogOpeningEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2ScriptDialogOpeningEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ScriptDialogOpeningEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnServerCertificateErrorActionsCompletedEvent(cb TOnServerCertificateErrorActionsCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerCertificateErrorActionsCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aErrorCode: HRESULT);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			cb(sender, errorCode)
		},
	}
}

func makeTOnServerCertificateErrorDetectedEvent(cb TOnServerCertificateErrorDetectedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnServerCertificateErrorDetectedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2ServerCertificateErrorDetectedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ServerCertificateErrorDetectedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnSetPermissionStateCompletedEvent(cb TOnSetPermissionStateCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSetPermissionStateCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aErrorCode: HRESULT);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			cb(sender, errorCode)
		},
	}
}

func makeTOnShowSaveAsUICompletedEvent(cb TOnShowSaveAsUICompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnShowSaveAsUICompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HResult; aResult: TWVSaveAsUIResult);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := wvTypes.TWVSaveAsUIResult(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnSimulateKeyEventCompletedEvent(cb TOnSimulateKeyEventCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSimulateKeyEventCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			result := api.GoBool(getVal(1))
			cb(sender, result)
		},
	}
}

func makeTOnSourceChangedEvent(cb TOnSourceChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnSourceChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2SourceChangedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2SourceChangedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnStatusBarTextChangedEvent(cb TOnStatusBarTextChangedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnStatusBarTextChangedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; const aWebView: ICoreWebView2);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			cb(sender, webView)
		},
	}
}

func makeTOnTrySuspendCompletedEvent(cb TOnTrySuspendCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnTrySuspendCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; aErrorCode: HRESULT; aResult: boolean);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := api.GoBool(getVal(2))
			cb(sender, errorCode, result)
		},
	}
}

func makeTOnWebMessageReceivedEvent(cb TOnWebMessageReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebMessageReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2WebMessageReceivedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2WebMessageReceivedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnWebResourceRequestedEvent(cb TOnWebResourceRequestedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebResourceRequestedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2WebResourceRequestedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2WebResourceRequestedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnWebResourceResponseReceivedEvent(cb TOnWebResourceResponseReceivedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebResourceResponseReceivedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 3 : procedure(Sender: TObject; const aWebView: ICoreWebView2; const aArgs: ICoreWebView2WebResourceResponseReceivedEventArgs);
			sender := lcl.AsObject(getVal(0))
			webView := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2WebResourceResponseReceivedEventArgs(getVal(2))
			cb(sender, webView, args)
		},
	}
}

func makeTOnWebResourceResponseViewGetContentCompletedEvent(cb TOnWebResourceResponseViewGetContentCompletedEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TOnWebResourceResponseViewGetContentCompletedEvent",
		cb: func(getVal func(i int) uintptr) {
			// 4 : procedure(Sender: TObject; aErrorCode: HRESULT; const aResult: IStream; aResourceID : integer);
			sender := lcl.AsObject(getVal(0))
			errorCode := types.HRESULT(getVal(1))
			result := lcl.AsStreamAdapter(getVal(2))
			resourceID := int32(getVal(3))
			cb(sender, errorCode, result, resourceID)
		},
	}
}

func makeTStartDragEvent(cb TStartDragEvent) *callback {
	if cb == nil {
		return nil
	}
	return &callback{
		name: "TStartDragEvent",
		cb: func(getVal func(i int) uintptr) {
			// 2 : procedure(Sender: TObject; var DragObject: TDragObject);
			sender := lcl.AsObject(getVal(0))
			var dragObject lcl.IDragObject
			dragObject = lcl.AsDragObject(*(*uintptr)(getPtr(getVal(1))))
			cb(sender, &dragObject)
			if dragObject != nil {
				*(*uintptr)(getPtr(getVal(1))) = dragObject.Instance()
			}
		},
	}
}
