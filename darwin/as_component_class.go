//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

// AsNSHTTPURLResponse Convert a pointer object to an existing class object
func AsNSHTTPURLResponse(obj interface{}) INSHTTPURLResponse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSHTTPURLResponse := new(TNSHTTPURLResponse)
	SetObjectInstance(nSHTTPURLResponse, instance)
	return nSHTTPURLResponse
}

// AsNSMutableURLRequest Convert a pointer object to an existing class object
func AsNSMutableURLRequest(obj interface{}) INSMutableURLRequest {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSMutableURLRequest := new(TNSMutableURLRequest)
	SetObjectInstance(nSMutableURLRequest, instance)
	return nSMutableURLRequest
}

// AsNSProgress Convert a pointer object to an existing class object
func AsNSProgress(obj interface{}) INSProgress {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSProgress := new(TNSProgress)
	SetObjectInstance(nSProgress, instance)
	return nSProgress
}

// AsNSURL Convert a pointer object to an existing class object
func AsNSURL(obj interface{}) INSURL {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSURL := new(TNSURL)
	SetObjectInstance(nSURL, instance)
	return nSURL
}

// AsNSURLAuthenticationChallenge Convert a pointer object to an existing class object
func AsNSURLAuthenticationChallenge(obj interface{}) INSURLAuthenticationChallenge {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSURLAuthenticationChallenge := new(TNSURLAuthenticationChallenge)
	SetObjectInstance(nSURLAuthenticationChallenge, instance)
	return nSURLAuthenticationChallenge
}

// AsNSURLCredential Convert a pointer object to an existing class object
func AsNSURLCredential(obj interface{}) INSURLCredential {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSURLCredential := new(TNSURLCredential)
	SetObjectInstance(nSURLCredential, instance)
	return nSURLCredential
}

// AsNSURLRequest Convert a pointer object to an existing class object
func AsNSURLRequest(obj interface{}) INSURLRequest {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSURLRequest := new(TNSURLRequest)
	SetObjectInstance(nSURLRequest, instance)
	return nSURLRequest
}

// AsNSURLResponse Convert a pointer object to an existing class object
func AsNSURLResponse(obj interface{}) INSURLResponse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	nSURLResponse := new(TNSURLResponse)
	SetObjectInstance(nSURLResponse, instance)
	return nSURLResponse
}

// AsWKDownload Convert a pointer object to an existing class object
func AsWKDownload(obj interface{}) IWKDownload {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKDownload := new(TWKDownload)
	SetObjectInstance(wKDownload, instance)
	return wKDownload
}

// AsWKDownloadDelegate Convert a pointer object to an existing class object
func AsWKDownloadDelegate(obj interface{}) IWKDownloadDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKDownloadDelegate := new(TWKDownloadDelegate)
	SetObjectInstance(wKDownloadDelegate, instance)
	return wKDownloadDelegate
}

// AsWKFrameInfo Convert a pointer object to an existing class object
func AsWKFrameInfo(obj interface{}) IWKFrameInfo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKFrameInfo := new(TWKFrameInfo)
	SetObjectInstance(wKFrameInfo, instance)
	return wKFrameInfo
}

// AsWKNavigation Convert a pointer object to an existing class object
func AsWKNavigation(obj interface{}) IWKNavigation {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKNavigation := new(TWKNavigation)
	SetObjectInstance(wKNavigation, instance)
	return wKNavigation
}

// AsWKNavigationAction Convert a pointer object to an existing class object
func AsWKNavigationAction(obj interface{}) IWKNavigationAction {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKNavigationAction := new(TWKNavigationAction)
	SetObjectInstance(wKNavigationAction, instance)
	return wKNavigationAction
}

// AsWKNavigationDelegate Convert a pointer object to an existing class object
func AsWKNavigationDelegate(obj interface{}) IWKNavigationDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKNavigationDelegate := new(TWKNavigationDelegate)
	SetObjectInstance(wKNavigationDelegate, instance)
	return wKNavigationDelegate
}

// AsWKNavigationResponse Convert a pointer object to an existing class object
func AsWKNavigationResponse(obj interface{}) IWKNavigationResponse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKNavigationResponse := new(TWKNavigationResponse)
	SetObjectInstance(wKNavigationResponse, instance)
	return wKNavigationResponse
}

// AsWKPreferences Convert a pointer object to an existing class object
func AsWKPreferences(obj interface{}) IWKPreferences {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKPreferences := new(TWKPreferences)
	SetObjectInstance(wKPreferences, instance)
	return wKPreferences
}

// AsWKScriptMessageHandler Convert a pointer object to an existing class object
func AsWKScriptMessageHandler(obj interface{}) IWKScriptMessageHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKScriptMessageHandler := new(TWKScriptMessageHandler)
	SetObjectInstance(wKScriptMessageHandler, instance)
	return wKScriptMessageHandler
}

// AsWKUIDelegate Convert a pointer object to an existing class object
func AsWKUIDelegate(obj interface{}) IWKUIDelegate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKUIDelegate := new(TWKUIDelegate)
	SetObjectInstance(wKUIDelegate, instance)
	return wKUIDelegate
}

// AsWKURLSchemeHandler Convert a pointer object to an existing class object
func AsWKURLSchemeHandler(obj interface{}) IWKURLSchemeHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKURLSchemeHandler := new(TWKURLSchemeHandler)
	SetObjectInstance(wKURLSchemeHandler, instance)
	return wKURLSchemeHandler
}

// AsWKURLSchemeTask Convert a pointer object to an existing class object
func AsWKURLSchemeTask(obj interface{}) IWKURLSchemeTask {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKURLSchemeTask := new(TWKURLSchemeTask)
	SetObjectInstance(wKURLSchemeTask, instance)
	return wKURLSchemeTask
}

// AsWKUserContentController Convert a pointer object to an existing class object
func AsWKUserContentController(obj interface{}) IWKUserContentController {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKUserContentController := new(TWKUserContentController)
	SetObjectInstance(wKUserContentController, instance)
	return wKUserContentController
}

// AsWKUserScript Convert a pointer object to an existing class object
func AsWKUserScript(obj interface{}) IWKUserScript {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKUserScript := new(TWKUserScript)
	SetObjectInstance(wKUserScript, instance)
	return wKUserScript
}

// AsWKWebViewConfiguration Convert a pointer object to an existing class object
func AsWKWebViewConfiguration(obj interface{}) IWKWebViewConfiguration {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKWebViewConfiguration := new(TWKWebViewConfiguration)
	SetObjectInstance(wKWebViewConfiguration, instance)
	return wKWebViewConfiguration
}

// AsWKWebpagePreferences Convert a pointer object to an existing class object
func AsWKWebpagePreferences(obj interface{}) IWKWebpagePreferences {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKWebpagePreferences := new(TWKWebpagePreferences)
	SetObjectInstance(wKWebpagePreferences, instance)
	return wKWebpagePreferences
}

// AsWKWindowFeatures Convert a pointer object to an existing class object
func AsWKWindowFeatures(obj interface{}) IWKWindowFeatures {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wKWindowFeatures := new(TWKWindowFeatures)
	SetObjectInstance(wKWindowFeatures, instance)
	return wKWindowFeatures
}

// AsWkWebview Convert a pointer object to an existing class object
func AsWkWebview(obj interface{}) IWkWebview {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkWebview := new(TWkWebview)
	SetObjectInstance(wkWebview, instance)
	return wkWebview
}

// AsWkWebviewParent Convert a pointer object to an existing class object
func AsWkWebviewParent(obj interface{}) IWkWebviewParent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkWebviewParent := new(TWkWebviewParent)
	SetObjectInstance(wkWebviewParent, instance)
	return wkWebviewParent
}
