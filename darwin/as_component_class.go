//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package darwin

import "github.com/energye/lcl/base"

// AsWkNSURL Convert a pointer object to an existing class object
func AsWkNSURL(obj any) IWkNSURL {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSURL)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNSURLAuthenticationChallenge Convert a pointer object to an existing class object
func AsWkNSURLAuthenticationChallenge(obj any) IWkNSURLAuthenticationChallenge {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSURLAuthenticationChallenge)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNSURLCredential Convert a pointer object to an existing class object
func AsWkNSURLCredential(obj any) IWkNSURLCredential {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSURLCredential)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNSURLRequest Convert a pointer object to an existing class object
func AsWkNSURLRequest(obj any) IWkNSURLRequest {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSURLRequest)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNSMutableURLRequest Convert a pointer object to an existing class object
func AsWkNSMutableURLRequest(obj any) IWkNSMutableURLRequest {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSMutableURLRequest)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNSURLResponse Convert a pointer object to an existing class object
func AsWkNSURLResponse(obj any) IWkNSURLResponse {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSURLResponse)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNSHTTPURLResponse Convert a pointer object to an existing class object
func AsWkNSHTTPURLResponse(obj any) IWkNSHTTPURLResponse {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSHTTPURLResponse)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkDownload Convert a pointer object to an existing class object
func AsWkDownload(obj any) IWkDownload {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkDownload)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkDownloadDelegate Convert a pointer object to an existing class object
func AsWkDownloadDelegate(obj any) IWkDownloadDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkDownloadDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkFrameInfo Convert a pointer object to an existing class object
func AsWkFrameInfo(obj any) IWkFrameInfo {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkFrameInfo)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNavigation Convert a pointer object to an existing class object
func AsWkNavigation(obj any) IWkNavigation {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNavigation)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNavigationAction Convert a pointer object to an existing class object
func AsWkNavigationAction(obj any) IWkNavigationAction {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNavigationAction)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNavigationDelegate Convert a pointer object to an existing class object
func AsWkNavigationDelegate(obj any) IWkNavigationDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNavigationDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNavigationResponse Convert a pointer object to an existing class object
func AsWkNavigationResponse(obj any) IWkNavigationResponse {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNavigationResponse)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkPreferences Convert a pointer object to an existing class object
func AsWkPreferences(obj any) IWkPreferences {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkPreferences)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkScriptMessageHandler Convert a pointer object to an existing class object
func AsWkScriptMessageHandler(obj any) IWkScriptMessageHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkScriptMessageHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkUIDelegate Convert a pointer object to an existing class object
func AsWkUIDelegate(obj any) IWkUIDelegate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkUIDelegate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkURLSchemeHandler Convert a pointer object to an existing class object
func AsWkURLSchemeHandler(obj any) IWkURLSchemeHandler {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkURLSchemeHandler)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkURLSchemeTask Convert a pointer object to an existing class object
func AsWkURLSchemeTask(obj any) IWkURLSchemeTask {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkURLSchemeTask)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkUserContentController Convert a pointer object to an existing class object
func AsWkUserContentController(obj any) IWkUserContentController {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkUserContentController)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkUserScript Convert a pointer object to an existing class object
func AsWkUserScript(obj any) IWkUserScript {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkUserScript)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkWebpagePreferences Convert a pointer object to an existing class object
func AsWkWebpagePreferences(obj any) IWkWebpagePreferences {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkWebpagePreferences)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkWebview Convert a pointer object to an existing class object
func AsWkWebview(obj any) IWkWebview {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkWebview)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkWebViewConfiguration Convert a pointer object to an existing class object
func AsWkWebViewConfiguration(obj any) IWkWebViewConfiguration {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkWebViewConfiguration)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkWebviewParent Convert a pointer object to an existing class object
func AsWkWebviewParent(obj any) IWkWebviewParent {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkWebviewParent)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkWindowFeatures Convert a pointer object to an existing class object
func AsWkWindowFeatures(obj any) IWkWindowFeatures {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkWindowFeatures)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNSProgress Convert a pointer object to an existing class object
func AsWkNSProgress(obj any) IWkNSProgress {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNSProgress)
	base.SetObjectInstance(result, instance)
	return result
}
