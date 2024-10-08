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
	"github.com/energye/lcl/types"
	"unsafe"
)

// getParam 从指定索引和地址获取事件中的参数
// 不再使用FreePascal导出的了，直接在这处理
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
		case TNotifyEvent:
			fn.(TNotifyEvent)(lcl.AsObject(getPtr(0)))

		case TOnExecuteScriptCompletedEvent:
			fn.(TOnExecuteScriptCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)), GoStr(getVal(2)), int32(getVal(3)))

		case TOnCapturePreviewCompletedEvent:
			fn.(TOnCapturePreviewCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)))

		case TOnWebResourceResponseViewGetContentCompletedEvent:
			fn.(TOnWebResourceResponseViewGetContentCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)), AsStream(getVal(2)), int32(getVal(3)))

		case TOnGetCookiesCompletedEvent:
			cookieList := AsCoreWebView2CookieList(getVal(2))
			fn.(TOnGetCookiesCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)), cookieList)

		case TOnTrySuspendCompletedEvent:
			fn.(TOnTrySuspendCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)), GoBool(getVal(2)))

		case TOnPrintToPdfCompletedEvent:
			fn.(TOnPrintToPdfCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)), GoBool(getVal(2)))

		case TOnCallDevToolsProtocolMethodCompletedEvent:
			fn.(TOnCallDevToolsProtocolMethodCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)), GoStr(getVal(2)), int32(getVal(3)))

		case TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent:
			fn.(TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent)(lcl.AsObject(getPtr(0)), int32(getVal(1)), GoStr(getVal(2)))

		case TOnMoveFocusRequestedEvent:
			controller := AsCoreWebView2Controller(getVal(1))
			args := AsCoreWebView2MoveFocusRequestedEventArgs(getVal(2))
			fn.(TOnMoveFocusRequestedEvent)(lcl.AsObject(getPtr(0)), controller, args)

		case TOnAcceleratorKeyPressedEvent:
			controller := AsCoreWebView2Controller(getVal(1))
			args := AsCoreWebView2AcceleratorKeyPressedEventArgs(getVal(2))
			fn.(TOnAcceleratorKeyPressedEvent)(lcl.AsObject(getPtr(0)), controller, args)

		case TOnBrowserProcessExitedEvent:
			environment := AsCoreWebView2Environment(getVal(1))
			args := AsCoreWebView2BrowserProcessExitedEventArgs(getVal(2))
			fn.(TOnBrowserProcessExitedEvent)(lcl.AsObject(getPtr(0)), environment, args)

		case TOnNavigationStartingEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2NavigationStartingEventArgs(getVal(2))
			fn.(TOnNavigationStartingEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnNavigationCompletedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2NavigationCompletedEventArgs(getVal(2))
			fn.(TOnNavigationCompletedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnSourceChangedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2SourceChangedEventArgs(getVal(2))
			fn.(TOnSourceChangedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnContentLoadingEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ContentLoadingEventArgs(getVal(2))
			fn.(TOnContentLoadingEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnNewWindowRequestedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2NewWindowRequestedEventArgs(getVal(2))
			fn.(TOnNewWindowRequestedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnWebResourceRequestedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2WebResourceRequestedEventArgs(getVal(2))
			fn.(TOnWebResourceRequestedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnScriptDialogOpeningEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ScriptDialogOpeningEventArgs(getVal(2))
			fn.(TOnScriptDialogOpeningEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnPermissionRequestedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2PermissionRequestedEventArgs(getVal(2))
			fn.(TOnPermissionRequestedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnProcessFailedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ProcessFailedEventArgs(getVal(2))
			fn.(TOnProcessFailedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnWebMessageReceivedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2WebMessageReceivedEventArgs(getVal(2))
			fn.(TOnWebMessageReceivedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnDevToolsProtocolEventReceivedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(getVal(2))
			fn.(TOnDevToolsProtocolEventReceivedEvent)(lcl.AsObject(getPtr(0)), webview, args, GoStr(getVal(3)), int32(getVal(4)))

		case TOnWebResourceResponseReceivedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2WebResourceResponseReceivedEventArgs(getVal(2))
			fn.(TOnWebResourceResponseReceivedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnDOMContentLoadedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2DOMContentLoadedEventArgs(getVal(2))
			fn.(TOnDOMContentLoadedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnFrameCreatedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2FrameCreatedEventArgs(getVal(2))
			fn.(TOnFrameCreatedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnDownloadStartingEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2DownloadStartingEventArgs(getVal(2))
			fn.(TOnDownloadStartingEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnClientCertificateRequestedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ClientCertificateRequestedEventArgs(getVal(2))
			fn.(TOnClientCertificateRequestedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnBytesReceivedChangedEvent:
			downloadID := int32(getVal(2))
			downloadOperation := AsCoreWebView2DownloadOperation(getVal(1))
			fn.(TOnBytesReceivedChangedEvent)(lcl.AsObject(getPtr(0)), downloadOperation, downloadID)

		case TOnEstimatedEndTimeChangedEvent:
			downloadID := int32(getVal(2))
			downloadOperation := AsCoreWebView2DownloadOperation(getVal(1))
			fn.(TOnEstimatedEndTimeChangedEvent)(lcl.AsObject(getPtr(0)), downloadOperation, downloadID)

		case TOnDownloadStateChangedEvent:
			downloadID := int32(getVal(2))
			downloadOperation := AsCoreWebView2DownloadOperation(getVal(1))
			fn.(TOnDownloadStateChangedEvent)(lcl.AsObject(getPtr(0)), downloadOperation, downloadID)

		case TOnFrameNameChangedEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			frameID := uint32(getVal(2))
			fn.(TOnFrameNameChangedEvent)(lcl.AsObject(getPtr(0)), frame, frameID)

		case TOnFrameDestroyedEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			frameID := uint32(getVal(2))
			fn.(TOnFrameDestroyedEvent)(lcl.AsObject(getPtr(0)), frame, frameID)

		case TOnInitializationErrorEvent:
			errorCode := int32(getVal(1))
			errorMessage := GoStr(getVal(2))
			fn.(TOnInitializationErrorEvent)(lcl.AsObject(getPtr(0)), errorCode, errorMessage)

		case TOnPrintCompletedEvent:
			errorCode := int32(getVal(1))
			printStatus := TWVPrintStatus(getVal(2))
			fn.(TOnPrintCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, printStatus)

		case TOnRefreshIgnoreCacheCompletedEvent:
			errorCode := int32(getVal(1))
			resulIObjectAsJson := GoStr(getVal(2))
			fn.(TOnRefreshIgnoreCacheCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, resulIObjectAsJson)

		case TOnRetrieveHTMLCompletedEvent:
			result := GoBool(getVal(1))
			html := GoStr(getVal(2))
			fn.(TOnRetrieveHTMLCompletedEvent)(lcl.AsObject(getPtr(0)), result, html)

		case TOnRetrieveTextCompletedEvent:
			result := GoBool(getVal(1))
			text := GoStr(getVal(2))
			fn.(TOnRetrieveTextCompletedEvent)(lcl.AsObject(getPtr(0)), result, text)

		case TOnRetrieveMHTMLCompletedEvent:
			result := GoBool(getVal(1))
			mHtml := GoStr(getVal(2))
			fn.(TOnRetrieveMHTMLCompletedEvent)(lcl.AsObject(getPtr(0)), result, mHtml)

		case TOnClearCacheCompletedEvent:
			result := GoBool(getVal(1))
			fn.(TOnClearCacheCompletedEvent)(lcl.AsObject(getPtr(0)), result)

		case TOnClearDataForOriginCompletedEvent:
			result := GoBool(getVal(1))
			fn.(TOnClearDataForOriginCompletedEvent)(lcl.AsObject(getPtr(0)), result)

		case TOnOfflineCompletedEvent:
			result := GoBool(getVal(1))
			fn.(TOnOfflineCompletedEvent)(lcl.AsObject(getPtr(0)), result)

		case TOnIgnoreCertificateErrorsCompletedEvent:
			result := GoBool(getVal(1))
			fn.(TOnIgnoreCertificateErrorsCompletedEvent)(lcl.AsObject(getPtr(0)), result)

		case TOnSimulateKeyEventCompletedEvent:
			result := GoBool(getVal(1))
			fn.(TOnSimulateKeyEventCompletedEvent)(lcl.AsObject(getPtr(0)), result)

		case TOnIsMutedChangedEvent:
			webview := AsCoreWebView2(getVal(1))
			fn.(TOnIsMutedChangedEvent)(lcl.AsObject(getPtr(0)), webview)

		case TOnIsDocumentPlayingAudioChangedEvent:
			webview := AsCoreWebView2(getVal(1))
			fn.(TOnIsDocumentPlayingAudioChangedEvent)(lcl.AsObject(getPtr(0)), webview)

		case TOnIsDefaultDownloadDialogOpenChangedEvent:
			webview := AsCoreWebView2(getVal(1))
			fn.(TOnIsDefaultDownloadDialogOpenChangedEvent)(lcl.AsObject(getPtr(0)), webview)

		case TOnProcessInfosChangedEvent:
			environment := AsCoreWebView2Environment(getVal(1))
			fn.(TOnProcessInfosChangedEvent)(lcl.AsObject(getPtr(0)), environment)

		case TOnFrameNavigationStartingEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2NavigationStartingEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			fn.(TOnFrameNavigationStartingEvent)(lcl.AsObject(getPtr(0)), frame, args, frameID)

		case TOnFrameNavigationCompletedEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2NavigationCompletedEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			fn.(TOnFrameNavigationCompletedEvent)(lcl.AsObject(getPtr(0)), frame, args, frameID)

		case TOnFrameContentLoadingEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2ContentLoadingEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			fn.(TOnFrameContentLoadingEvent)(lcl.AsObject(getPtr(0)), frame, args, frameID)

		case TOnFrameDOMContentLoadedEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2DOMContentLoadedEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			fn.(TOnFrameDOMContentLoadedEvent)(lcl.AsObject(getPtr(0)), frame, args, frameID)

		case TOnFrameWebMessageReceivedEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2WebMessageReceivedEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			fn.(TOnFrameWebMessageReceivedEvent)(lcl.AsObject(getPtr(0)), frame, args, frameID)

		case TOnBasicAuthenticationRequestedEvent:
			frame := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2BasicAuthenticationRequestedEventArgs(getVal(2))
			fn.(TOnBasicAuthenticationRequestedEvent)(lcl.AsObject(getPtr(0)), frame, args)

		case TOnContextMenuRequestedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ContextMenuRequestedEventArgs(getVal(2))
			fn.(TOnContextMenuRequestedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnCustomItemSelectedEvent:
			menuItem := AsCoreWebView2ContextMenuItem(getVal(1))
			fn.(TOnCustomItemSelectedEvent)(lcl.AsObject(getPtr(0)), menuItem)

		case TOnStatusBarTextChangedEvent:
			webview := AsCoreWebView2(getVal(1))
			fn.(TOnStatusBarTextChangedEvent)(lcl.AsObject(getPtr(0)), webview)

		case TOnFramePermissionRequestedEvent:
			frame := AsCoreWebView2Frame(getVal(1))
			args := AsCoreWebView2PermissionRequestedEventArgs(getVal(2))
			frameID := uint32(getVal(3))
			fn.(TOnFramePermissionRequestedEvent)(lcl.AsObject(getPtr(0)), frame, args, frameID)

		case TOnClearBrowsingDataCompletedEvent:
			errorCode := int32(getVal(1))
			fn.(TOnClearBrowsingDataCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode)

		case TOnServerCertificateErrorActionsCompletedEvent:
			errorCode := int32(getVal(1))
			fn.(TOnServerCertificateErrorActionsCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode)

		case TOnServerCertificateErrorDetectedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2ServerCertificateErrorDetectedEventArgs(getVal(2))
			fn.(TOnServerCertificateErrorDetectedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnFaviconChangedEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsUnknown(getVal(2))
			fn.(TOnFaviconChangedEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnGetFaviconCompletedEvent:
			errorCode := int32(getVal(1))
			faviconStream := AsStream(getVal(2))
			fn.(TOnGetFaviconCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, faviconStream)

		case TOnPrintToPdfStreamCompletedEvent:
			errorCode := int32(getVal(1))
			pdfStream := AsStream(getVal(2))
			fn.(TOnPrintToPdfStreamCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, pdfStream)

		case TOnGetCustomSchemesEvent:
			var data = (*uintptr)(getPtr(1))
			var size = (*int32)(getPtr(2))
			var customSchemes TWVCustomSchemeInfoArray
			fn.(TOnGetCustomSchemesEvent)(lcl.AsObject(getPtr(0)), customSchemes)
			*size = int32(len(customSchemes))
			if *size > 0 {
				dataPtr := make([]*tWVCustomSchemeInfoPtr, *size, *size)
				for i, scheme := range customSchemes {
					dataPtr[i] = &tWVCustomSchemeInfoPtr{
						SchemeName:            PascalStr(scheme.SchemeName),
						TreatAsSecure:         PascalBool(scheme.TreatAsSecure),
						AllowedDomains:        PascalStr(scheme.AllowedDomains),
						HasAuthorityComponent: PascalBool(scheme.HasAuthorityComponent),
					}
				}
				*data = uintptr(unsafe.Pointer(dataPtr[0]))
			}

		case TOnGetNonDefaultPermissionSettingsCompletedEvent:
			errorCode := int32(getVal(1))
			collectionView := AsCoreWebView2PermissionSettingCollectionView(getVal(2))
			fn.(TOnGetNonDefaultPermissionSettingsCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, collectionView)

		case TOnSetPermissionStateCompletedEvent:
			errorCode := int32(getVal(1))
			fn.(TOnSetPermissionStateCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode)

		case TOnLaunchingExternalUriSchemeEvent:
			webview := AsCoreWebView2(getVal(1))
			args := AsCoreWebView2LaunchingExternalUriSchemeEventArgs(getVal(2))
			fn.(TOnLaunchingExternalUriSchemeEvent)(lcl.AsObject(getPtr(0)), webview, args)

		case TOnGetProcessExtendedInfosCompletedEvent:
			errorCode := int32(getVal(1))
			value := AsCoreWebView2ProcessExtendedInfoCollection(getVal(2))
			fn.(TOnGetProcessExtendedInfosCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, value)

		case TOnBrowserExtensionRemoveCompletedEvent:
			errorCode := int32(getVal(1))
			extensionID := GoStr(getVal(2))
			fn.(TOnBrowserExtensionRemoveCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, extensionID)

		case TOnBrowserExtensionEnableCompletedEvent:
			errorCode := int32(getVal(1))
			extensionID := GoStr(getVal(2))
			fn.(TOnBrowserExtensionEnableCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, extensionID)

		case TOnProfileAddBrowserExtensionCompletedEvent:
			errorCode := int32(getVal(1))
			extension := AsCoreWebView2BrowserExtension(getVal(2))
			fn.(TOnProfileAddBrowserExtensionCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, extension)

		case TOnProfileGetBrowserExtensionsCompletedEvent:
			errorCode := int32(getVal(1))
			extensionList := AsCoreWebView2BrowserExtensionList(getVal(2))
			fn.(TOnProfileGetBrowserExtensionsCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, extensionList)

		case TOnProfileDeletedEvent:
			profile := AsCoreWebView2Profile(getVal(1))
			fn.(TOnProfileDeletedEvent)(lcl.AsObject(getPtr(0)), profile)

		case TOnExecuteScriptWithResultCompletedEvent:
			errorCode := int32(getVal(1))
			result := AsCoreWebView2ExecuteScriptResult(getVal(2))
			executionID := int32(getVal(3))
			fn.(TOnExecuteScriptWithResultCompletedEvent)(lcl.AsObject(getPtr(0)), errorCode, result, executionID)

		case TOnCompMsgEvent:
			message := (*types.TMessage)(getPtr(1))
			handled := (*bool)(getPtr(2))
			fn.(TOnCompMsgEvent)(lcl.AsObject(getPtr(0)), message, handled)

		case TDragDropEvent:
			fn.(TDragDropEvent)(lcl.AsObject(getVal(0)), lcl.AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		case TDragOverEvent:
			fn.(TDragOverEvent)(lcl.AsObject(getVal(0)), lcl.AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)), types.TDragState(getVal(4)), (*bool)(getPtr(5)))

		case TStartDragEvent:
			obj := lcl.AsDragObject(getVal(1))
			fn.(TStartDragEvent)(lcl.AsObject(getVal(0)), &obj)
			if obj != nil {
				*(*uintptr)(unsafePointer(getVal(1))) = obj.Instance()
			}

		case TEndDragEvent:
			fn.(TEndDragEvent)(lcl.AsObject(getVal(0)), lcl.AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// WVLoader
		case TOnLoaderNotifyEvent:
			fn.(TOnLoaderNotifyEvent)(lcl.AsObject(getPtr(0)))

		case TOnLoaderGetCustomSchemesEvent:
			var data = (*uintptr)(getPtr(1))
			var size = (*int32)(getPtr(2))
			var customSchemes TWVCustomSchemeInfoArray
			fn.(TOnLoaderGetCustomSchemesEvent)(lcl.AsObject(getPtr(0)), &customSchemes)
			*size = int32(len(customSchemes))
			if *size > 0 {
				// 数组不能使用 []*tWVCustomSchemeInfoPtr 定义
				dataPtr := make([]tWVCustomSchemeInfoPtr, *size, *size)
				for i, scheme := range customSchemes {
					dataPtr[i] = tWVCustomSchemeInfoPtr{
						SchemeName:            PascalStr(scheme.SchemeName),
						TreatAsSecure:         PascalBool(scheme.TreatAsSecure),
						AllowedDomains:        PascalStr(scheme.AllowedDomains),
						HasAuthorityComponent: PascalBool(scheme.HasAuthorityComponent),
					}
				}
				*data = uintptr(unsafe.Pointer(&dataPtr[0]))
			}
		case TOnLoaderNewBrowserVersionAvailableEvent:
			env := AsCoreWebView2Environment(getVal(1))
			fn.(TOnLoaderNewBrowserVersionAvailableEvent)(lcl.AsObject(getPtr(0)), env)

		case TOnLoaderBrowserProcessExitedEvent:
			env := AsCoreWebView2Environment(getVal(1))
			args := AsCoreWebView2BrowserProcessExitedEventArgs(getVal(2))
			fn.(TOnLoaderBrowserProcessExitedEvent)(lcl.AsObject(getPtr(0)), env, args)

		case TOnLoaderProcessInfosChangedEvent:
			env := AsCoreWebView2Environment(getVal(1))
			fn.(TOnLoaderProcessInfosChangedEvent)(lcl.AsObject(getPtr(0)), env)

		default:
		}
	}
	return 0
}

// Init
//
//	Webview2初始化
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	inits.Init(libs, resources)
	SetWVEventCallback(eventCallback)
	SetWVRemoveEventCallback(removeEventCallback)
}
