//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

// AsWkContextMenu Convert a pointer object to an existing class object
func AsWkContextMenu(obj interface{}) IWkContextMenu {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkContextMenu := new(TWkContextMenu)
	SetObjectInstance(wkContextMenu, instance)
	return wkContextMenu
}

// AsWkContextMenuItem Convert a pointer object to an existing class object
func AsWkContextMenuItem(obj interface{}) IWkContextMenuItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkContextMenuItem := new(TWkContextMenuItem)
	SetObjectInstance(wkContextMenuItem, instance)
	return wkContextMenuItem
}

// AsWkCookie Convert a pointer object to an existing class object
func AsWkCookie(obj interface{}) IWkCookie {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkCookie := new(TWkCookie)
	SetObjectInstance(wkCookie, instance)
	return wkCookie
}

// AsWkCookieDate Convert a pointer object to an existing class object
func AsWkCookieDate(obj interface{}) IWkCookieDate {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkCookieDate := new(TWkCookieDate)
	SetObjectInstance(wkCookieDate, instance)
	return wkCookieDate
}

// AsWkCookieList Convert a pointer object to an existing class object
func AsWkCookieList(obj interface{}) IWkCookieList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkCookieList := new(TWkCookieList)
	SetObjectInstance(wkCookieList, instance)
	return wkCookieList
}

// AsWkCookieManager Convert a pointer object to an existing class object
func AsWkCookieManager(obj interface{}) IWkCookieManager {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkCookieManager := new(TWkCookieManager)
	SetObjectInstance(wkCookieManager, instance)
	return wkCookieManager
}

// AsWkHeaders Convert a pointer object to an existing class object
func AsWkHeaders(obj interface{}) IWkHeaders {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkHeaders := new(TWkHeaders)
	SetObjectInstance(wkHeaders, instance)
	return wkHeaders
}

// AsWkInputStream Convert a pointer object to an existing class object
func AsWkInputStream(obj interface{}) IWkInputStream {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkInputStream := new(TWkInputStream)
	SetObjectInstance(wkInputStream, instance)
	return wkInputStream
}

// AsWkJSValue Convert a pointer object to an existing class object
func AsWkJSValue(obj interface{}) IWkJSValue {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkJSValue := new(TWkJSValue)
	SetObjectInstance(wkJSValue, instance)
	return wkJSValue
}

// AsWkNavigationAction Convert a pointer object to an existing class object
func AsWkNavigationAction(obj interface{}) IWkNavigationAction {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkNavigationAction := new(TWkNavigationAction)
	SetObjectInstance(wkNavigationAction, instance)
	return wkNavigationAction
}

// AsWkNavigationPolicyDecision Convert a pointer object to an existing class object
func AsWkNavigationPolicyDecision(obj interface{}) IWkNavigationPolicyDecision {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkNavigationPolicyDecision := new(TWkNavigationPolicyDecision)
	SetObjectInstance(wkNavigationPolicyDecision, instance)
	return wkNavigationPolicyDecision
}

// AsWkPolicyDecision Convert a pointer object to an existing class object
func AsWkPolicyDecision(obj interface{}) IWkPolicyDecision {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkPolicyDecision := new(TWkPolicyDecision)
	SetObjectInstance(wkPolicyDecision, instance)
	return wkPolicyDecision
}

// AsWkResponsePolicyDecision Convert a pointer object to an existing class object
func AsWkResponsePolicyDecision(obj interface{}) IWkResponsePolicyDecision {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkResponsePolicyDecision := new(TWkResponsePolicyDecision)
	SetObjectInstance(wkResponsePolicyDecision, instance)
	return wkResponsePolicyDecision
}

// AsWkSettings Convert a pointer object to an existing class object
func AsWkSettings(obj interface{}) IWkSettings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkSettings := new(TWkSettings)
	SetObjectInstance(wkSettings, instance)
	return wkSettings
}

// AsWkURIRequest Convert a pointer object to an existing class object
func AsWkURIRequest(obj interface{}) IWkURIRequest {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkURIRequest := new(TWkURIRequest)
	SetObjectInstance(wkURIRequest, instance)
	return wkURIRequest
}

// AsWkURIResponse Convert a pointer object to an existing class object
func AsWkURIResponse(obj interface{}) IWkURIResponse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkURIResponse := new(TWkURIResponse)
	SetObjectInstance(wkURIResponse, instance)
	return wkURIResponse
}

// AsWkURISchemeRequest Convert a pointer object to an existing class object
func AsWkURISchemeRequest(obj interface{}) IWkURISchemeRequest {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkURISchemeRequest := new(TWkURISchemeRequest)
	SetObjectInstance(wkURISchemeRequest, instance)
	return wkURISchemeRequest
}

// AsWkWebContext Convert a pointer object to an existing class object
func AsWkWebContext(obj interface{}) IWkWebContext {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkWebContext := new(TWkWebContext)
	SetObjectInstance(wkWebContext, instance)
	return wkWebContext
}

// AsWkWebsitePolicies Convert a pointer object to an existing class object
func AsWkWebsitePolicies(obj interface{}) IWkWebsitePolicies {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	wkWebsitePolicies := new(TWkWebsitePolicies)
	SetObjectInstance(wkWebsitePolicies, instance)
	return wkWebsitePolicies
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
