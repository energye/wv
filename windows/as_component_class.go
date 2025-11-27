//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package windows

import "github.com/energye/lcl/base"

// AsWVLoader Convert a pointer object to an existing class object
func AsWVLoader(obj any) IWVLoader {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWVLoader)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWVProxySettings Convert a pointer object to an existing class object
func AsWVProxySettings(obj any) IWVProxySettings {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWVProxySettings)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWVWinControl Convert a pointer object to an existing class object
func AsWVWinControl(obj any) IWVWinControl {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWVWinControl)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWVWindowParent Convert a pointer object to an existing class object
func AsWVWindowParent(obj any) IWVWindowParent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWVWindowParent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWVBrowserBase Convert a pointer object to an existing class object
func AsWVBrowserBase(obj any) IWVBrowserBase {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWVBrowserBase)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWVBrowser Convert a pointer object to an existing class object
func AsWVBrowser(obj any) IWVBrowser {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWVBrowser)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2 Convert a pointer object to an existing class object
func AsCoreWebView2(obj any) ICoreWebView2 {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PrintSettings Convert a pointer object to an existing class object
func AsCoreWebView2PrintSettings(obj any) ICoreWebView2PrintSettings {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PrintSettings)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Settings Convert a pointer object to an existing class object
func AsCoreWebView2Settings(obj any) ICoreWebView2Settings {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Settings)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Environment Convert a pointer object to an existing class object
func AsCoreWebView2Environment(obj any) ICoreWebView2Environment {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Environment)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Controller Convert a pointer object to an existing class object
func AsCoreWebView2Controller(obj any) ICoreWebView2Controller {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Controller)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2CompositionController Convert a pointer object to an existing class object
func AsCoreWebView2CompositionController(obj any) ICoreWebView2CompositionController {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2CompositionController)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Profile Convert a pointer object to an existing class object
func AsCoreWebView2Profile(obj any) ICoreWebView2Profile {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Profile)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2AcceleratorKeyPressedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2AcceleratorKeyPressedEventArgs(obj any) ICoreWebView2AcceleratorKeyPressedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2AcceleratorKeyPressedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2MoveFocusRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2MoveFocusRequestedEventArgs(obj any) ICoreWebView2MoveFocusRequestedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2MoveFocusRequestedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2BrowserProcessExitedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2BrowserProcessExitedEventArgs(obj any) ICoreWebView2BrowserProcessExitedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2BrowserProcessExitedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2NavigationStartingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NavigationStartingEventArgs(obj any) ICoreWebView2NavigationStartingEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2NavigationStartingEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2NavigationCompletedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NavigationCompletedEventArgs(obj any) ICoreWebView2NavigationCompletedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2NavigationCompletedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2SourceChangedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2SourceChangedEventArgs(obj any) ICoreWebView2SourceChangedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2SourceChangedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ContentLoadingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ContentLoadingEventArgs(obj any) ICoreWebView2ContentLoadingEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ContentLoadingEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2NewWindowRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NewWindowRequestedEventArgs(obj any) ICoreWebView2NewWindowRequestedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2NewWindowRequestedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebResourceRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceRequestedEventArgs(obj any) ICoreWebView2WebResourceRequestedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebResourceRequestedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ScriptDialogOpeningEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ScriptDialogOpeningEventArgs(obj any) ICoreWebView2ScriptDialogOpeningEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ScriptDialogOpeningEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PermissionRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2PermissionRequestedEventArgs(obj any) ICoreWebView2PermissionRequestedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PermissionRequestedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ProcessFailedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ProcessFailedEventArgs(obj any) ICoreWebView2ProcessFailedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ProcessFailedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebMessageReceivedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2WebMessageReceivedEventArgs(obj any) ICoreWebView2WebMessageReceivedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebMessageReceivedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2DevToolsProtocolEventReceivedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2DevToolsProtocolEventReceivedEventArgs(obj any) ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2DevToolsProtocolEventReceivedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebResourceResponseReceivedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponseReceivedEventArgs(obj any) ICoreWebView2WebResourceResponseReceivedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebResourceResponseReceivedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2DOMContentLoadedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2DOMContentLoadedEventArgs(obj any) ICoreWebView2DOMContentLoadedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2DOMContentLoadedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2FrameCreatedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2FrameCreatedEventArgs(obj any) ICoreWebView2FrameCreatedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2FrameCreatedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2DownloadStartingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2DownloadStartingEventArgs(obj any) ICoreWebView2DownloadStartingEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2DownloadStartingEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ClientCertificateRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ClientCertificateRequestedEventArgs(obj any) ICoreWebView2ClientCertificateRequestedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ClientCertificateRequestedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2BasicAuthenticationRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2BasicAuthenticationRequestedEventArgs(obj any) ICoreWebView2BasicAuthenticationRequestedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2BasicAuthenticationRequestedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ContextMenuRequestedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuRequestedEventArgs(obj any) ICoreWebView2ContextMenuRequestedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ContextMenuRequestedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ServerCertificateErrorDetectedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ServerCertificateErrorDetectedEventArgs(obj any) ICoreWebView2ServerCertificateErrorDetectedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ServerCertificateErrorDetectedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2LaunchingExternalUriSchemeEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2LaunchingExternalUriSchemeEventArgs(obj any) ICoreWebView2LaunchingExternalUriSchemeEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2LaunchingExternalUriSchemeEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ClientCertificate Convert a pointer object to an existing class object
func AsCoreWebView2ClientCertificate(obj any) ICoreWebView2ClientCertificate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ClientCertificate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ClientCertificateCollection Convert a pointer object to an existing class object
func AsCoreWebView2ClientCertificateCollection(obj any) ICoreWebView2ClientCertificateCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ClientCertificateCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2StringCollection Convert a pointer object to an existing class object
func AsCoreWebView2StringCollection(obj any) ICoreWebView2StringCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2StringCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Deferral Convert a pointer object to an existing class object
func AsCoreWebView2Deferral(obj any) ICoreWebView2Deferral {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Deferral)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2BasicAuthenticationResponse Convert a pointer object to an existing class object
func AsCoreWebView2BasicAuthenticationResponse(obj any) ICoreWebView2BasicAuthenticationResponse {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2BasicAuthenticationResponse)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ExecuteScriptResult Convert a pointer object to an existing class object
func AsCoreWebView2ExecuteScriptResult(obj any) ICoreWebView2ExecuteScriptResult {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ExecuteScriptResult)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ProcessExtendedInfoCollection Convert a pointer object to an existing class object
func AsCoreWebView2ProcessExtendedInfoCollection(obj any) ICoreWebView2ProcessExtendedInfoCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ProcessExtendedInfoCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2DownloadOperation Convert a pointer object to an existing class object
func AsCoreWebView2DownloadOperation(obj any) ICoreWebView2DownloadOperation {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2DownloadOperation)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Frame Convert a pointer object to an existing class object
func AsCoreWebView2Frame(obj any) ICoreWebView2Frame {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Frame)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PermissionSetting Convert a pointer object to an existing class object
func AsCoreWebView2PermissionSetting(obj any) ICoreWebView2PermissionSetting {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PermissionSetting)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2HttpHeadersCollectionIterator Convert a pointer object to an existing class object
func AsCoreWebView2HttpHeadersCollectionIterator(obj any) ICoreWebView2HttpHeadersCollectionIterator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2HttpHeadersCollectionIterator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2BrowserExtension Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtension(obj any) ICoreWebView2BrowserExtension {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2BrowserExtension)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2HttpResponseHeaders Convert a pointer object to an existing class object
func AsCoreWebView2HttpResponseHeaders(obj any) ICoreWebView2HttpResponseHeaders {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2HttpResponseHeaders)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2HttpRequestHeaders Convert a pointer object to an existing class object
func AsCoreWebView2HttpRequestHeaders(obj any) ICoreWebView2HttpRequestHeaders {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2HttpRequestHeaders)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2BrowserExtensionList Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtensionList(obj any) ICoreWebView2BrowserExtensionList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2BrowserExtensionList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ProcessInfo Convert a pointer object to an existing class object
func AsCoreWebView2ProcessInfo(obj any) ICoreWebView2ProcessInfo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ProcessInfo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PermissionSettingCollectionView Convert a pointer object to an existing class object
func AsCoreWebView2PermissionSettingCollectionView(obj any) ICoreWebView2PermissionSettingCollectionView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PermissionSettingCollectionView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2File Convert a pointer object to an existing class object
func AsCoreWebView2File(obj any) ICoreWebView2File {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2File)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2CookieList Convert a pointer object to an existing class object
func AsCoreWebView2CookieList(obj any) ICoreWebView2CookieList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2CookieList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ContextMenuItemCollection Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuItemCollection(obj any) ICoreWebView2ContextMenuItemCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ContextMenuItemCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ControllerOptions Convert a pointer object to an existing class object
func AsCoreWebView2ControllerOptions(obj any) ICoreWebView2ControllerOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ControllerOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ContextMenuItem Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuItem(obj any) ICoreWebView2ContextMenuItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ContextMenuItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebResourceResponse Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponse(obj any) ICoreWebView2WebResourceResponse {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebResourceResponse)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2SharedBuffer Convert a pointer object to an existing class object
func AsCoreWebView2SharedBuffer(obj any) ICoreWebView2SharedBuffer {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2SharedBuffer)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PointerInfo Convert a pointer object to an existing class object
func AsCoreWebView2PointerInfo(obj any) ICoreWebView2PointerInfo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PointerInfo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Cookie Convert a pointer object to an existing class object
func AsCoreWebView2Cookie(obj any) ICoreWebView2Cookie {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Cookie)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebResourceRequestRef Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceRequestRef(obj any) ICoreWebView2WebResourceRequestRef {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebResourceRequestRef)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ProcessInfoCollection Convert a pointer object to an existing class object
func AsCoreWebView2ProcessInfoCollection(obj any) ICoreWebView2ProcessInfoCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ProcessInfoCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2CookieManager Convert a pointer object to an existing class object
func AsCoreWebView2CookieManager(obj any) ICoreWebView2CookieManager {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2CookieManager)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ContextMenuTarget Convert a pointer object to an existing class object
func AsCoreWebView2ContextMenuTarget(obj any) ICoreWebView2ContextMenuTarget {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ContextMenuTarget)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ScriptException Convert a pointer object to an existing class object
func AsCoreWebView2ScriptException(obj any) ICoreWebView2ScriptException {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ScriptException)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WindowFeatures Convert a pointer object to an existing class object
func AsCoreWebView2WindowFeatures(obj any) ICoreWebView2WindowFeatures {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WindowFeatures)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2FrameInfo Convert a pointer object to an existing class object
func AsCoreWebView2FrameInfo(obj any) ICoreWebView2FrameInfo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2FrameInfo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ProcessExtendedInfo Convert a pointer object to an existing class object
func AsCoreWebView2ProcessExtendedInfo(obj any) ICoreWebView2ProcessExtendedInfo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ProcessExtendedInfo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2FrameInfoCollection Convert a pointer object to an existing class object
func AsCoreWebView2FrameInfoCollection(obj any) ICoreWebView2FrameInfoCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2FrameInfoCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2FrameInfoCollectionIterator Convert a pointer object to an existing class object
func AsCoreWebView2FrameInfoCollectionIterator(obj any) ICoreWebView2FrameInfoCollectionIterator {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2FrameInfoCollectionIterator)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Certificate Convert a pointer object to an existing class object
func AsCoreWebView2Certificate(obj any) ICoreWebView2Certificate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Certificate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ObjectCollectionView Convert a pointer object to an existing class object
func AsCoreWebView2ObjectCollectionView(obj any) ICoreWebView2ObjectCollectionView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ObjectCollectionView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ObjectCollection Convert a pointer object to an existing class object
func AsCoreWebView2ObjectCollection(obj any) ICoreWebView2ObjectCollection {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ObjectCollection)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebResourceResponseView Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponseView(obj any) ICoreWebView2WebResourceResponseView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebResourceResponseView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2RegionRectCollectionView Convert a pointer object to an existing class object
func AsCoreWebView2RegionRectCollectionView(obj any) ICoreWebView2RegionRectCollectionView {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2RegionRectCollectionView)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2FileSystemHandle Convert a pointer object to an existing class object
func AsCoreWebView2FileSystemHandle(obj any) ICoreWebView2FileSystemHandle {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2FileSystemHandle)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2Notification Convert a pointer object to an existing class object
func AsCoreWebView2Notification(obj any) ICoreWebView2Notification {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2Notification)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ScreenCaptureStartingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2ScreenCaptureStartingEventArgs(obj any) ICoreWebView2ScreenCaptureStartingEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ScreenCaptureStartingEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2NonClientRegionChangedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NonClientRegionChangedEventArgs(obj any) ICoreWebView2NonClientRegionChangedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2NonClientRegionChangedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2NotificationReceivedEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2NotificationReceivedEventArgs(obj any) ICoreWebView2NotificationReceivedEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2NotificationReceivedEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2SaveAsUIShowingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2SaveAsUIShowingEventArgs(obj any) ICoreWebView2SaveAsUIShowingEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2SaveAsUIShowingEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2SaveFileSecurityCheckStartingEventArgs Convert a pointer object to an existing class object
func AsCoreWebView2SaveFileSecurityCheckStartingEventArgs(obj any) ICoreWebView2SaveFileSecurityCheckStartingEventArgs {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2SaveFileSecurityCheckStartingEventArgs)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWVCustomSchemeInfoArrayWrap Convert a pointer object to an existing class object
func AsWVCustomSchemeInfoArrayWrap(obj any) IWVCustomSchemeInfoArrayWrap {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWVCustomSchemeInfoArrayWrap)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2EnvironmentOptions Convert a pointer object to an existing class object
func AsCoreWebView2EnvironmentOptions(obj any) ICoreWebView2EnvironmentOptions {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2EnvironmentOptions)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebResourceRequestOwn Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceRequestOwn(obj any) ICoreWebView2WebResourceRequestOwn {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebResourceRequestOwn)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2CustomSchemeRegistration Convert a pointer object to an existing class object
func AsCoreWebView2CustomSchemeRegistration(obj any) ICoreWebView2CustomSchemeRegistration {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2CustomSchemeRegistration)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2BrowserExtensionEnableCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtensionEnableCompletedHandler(obj any) ICoreWebView2BrowserExtensionEnableCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2BrowserExtensionEnableCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2BrowserExtensionRemoveCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2BrowserExtensionRemoveCompletedHandler(obj any) ICoreWebView2BrowserExtensionRemoveCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2BrowserExtensionRemoveCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ClearBrowsingDataCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2ClearBrowsingDataCompletedHandler(obj any) ICoreWebView2ClearBrowsingDataCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ClearBrowsingDataCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2CustomItemSelectedEventHandler Convert a pointer object to an existing class object
func AsCoreWebView2CustomItemSelectedEventHandler(obj any) ICoreWebView2CustomItemSelectedEventHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2CustomItemSelectedEventHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2GetCookiesCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2GetCookiesCompletedHandler(obj any) ICoreWebView2GetCookiesCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2GetCookiesCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler(obj any) ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2GetNonDefaultPermissionSettingsCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PrintCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2PrintCompletedHandler(obj any) ICoreWebView2PrintCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PrintCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PrintToPdfCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2PrintToPdfCompletedHandler(obj any) ICoreWebView2PrintToPdfCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PrintToPdfCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2PrintToPdfStreamCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2PrintToPdfStreamCompletedHandler(obj any) ICoreWebView2PrintToPdfStreamCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2PrintToPdfStreamCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ProfileAddBrowserExtensionCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2ProfileAddBrowserExtensionCompletedHandler(obj any) ICoreWebView2ProfileAddBrowserExtensionCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ProfileAddBrowserExtensionCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2ProfileGetBrowserExtensionsCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2ProfileGetBrowserExtensionsCompletedHandler(obj any) ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2ProfileGetBrowserExtensionsCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2SetPermissionStateCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2SetPermissionStateCompletedHandler(obj any) ICoreWebView2SetPermissionStateCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2SetPermissionStateCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2TrySuspendCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2TrySuspendCompletedHandler(obj any) ICoreWebView2TrySuspendCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2TrySuspendCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsCoreWebView2WebResourceResponseViewGetContentCompletedHandler Convert a pointer object to an existing class object
func AsCoreWebView2WebResourceResponseViewGetContentCompletedHandler(obj any) ICoreWebView2WebResourceResponseViewGetContentCompletedHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TCoreWebView2WebResourceResponseViewGetContentCompletedHandler)
	base.SetObjectInstance(result, instance)
	return result
}
