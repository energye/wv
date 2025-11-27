//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package linux

import "github.com/energye/lcl/base"

// AsWkContextMenu Convert a pointer object to an existing class object
func AsWkContextMenu(obj any) IWkContextMenu {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkContextMenu)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkContextMenuItem Convert a pointer object to an existing class object
func AsWkContextMenuItem(obj any) IWkContextMenuItem {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkContextMenuItem)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkCookieDate Convert a pointer object to an existing class object
func AsWkCookieDate(obj any) IWkCookieDate {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkCookieDate)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkHeaders Convert a pointer object to an existing class object
func AsWkHeaders(obj any) IWkHeaders {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkHeaders)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkInputStream Convert a pointer object to an existing class object
func AsWkInputStream(obj any) IWkInputStream {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkInputStream)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkJSValue Convert a pointer object to an existing class object
func AsWkJSValue(obj any) IWkJSValue {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkJSValue)
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

// AsWkPolicyDecision Convert a pointer object to an existing class object
func AsWkPolicyDecision(obj any) IWkPolicyDecision {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkPolicyDecision)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkNavigationPolicyDecision Convert a pointer object to an existing class object
func AsWkNavigationPolicyDecision(obj any) IWkNavigationPolicyDecision {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkNavigationPolicyDecision)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkResponsePolicyDecision Convert a pointer object to an existing class object
func AsWkResponsePolicyDecision(obj any) IWkResponsePolicyDecision {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkResponsePolicyDecision)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkSettings Convert a pointer object to an existing class object
func AsWkSettings(obj any) IWkSettings {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkSettings)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkURIRequest Convert a pointer object to an existing class object
func AsWkURIRequest(obj any) IWkURIRequest {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkURIRequest)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkURIResponse Convert a pointer object to an existing class object
func AsWkURIResponse(obj any) IWkURIResponse {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkURIResponse)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkURISchemeRequest Convert a pointer object to an existing class object
func AsWkURISchemeRequest(obj any) IWkURISchemeRequest {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkURISchemeRequest)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkWebContext Convert a pointer object to an existing class object
func AsWkWebContext(obj any) IWkWebContext {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkWebContext)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkWebsitePolicies Convert a pointer object to an existing class object
func AsWkWebsitePolicies(obj any) IWkWebsitePolicies {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkWebsitePolicies)
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

// AsWkCookieManager Convert a pointer object to an existing class object
func AsWkCookieManager(obj any) IWkCookieManager {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkCookieManager)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkCookie Convert a pointer object to an existing class object
func AsWkCookie(obj any) IWkCookie {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkCookie)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkCookieList Convert a pointer object to an existing class object
func AsWkCookieList(obj any) IWkCookieList {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkCookieList)
	base.SetObjectInstance(result, instance)
	return result
}

// AsWkLoader Convert a pointer object to an existing class object
func AsWkLoader(obj any) IWkLoader {
	instance := base.GetInstance(obj)
	if instance == nil {
		return nil
	}
	result := new(TWkLoader)
	base.SetObjectInstance(result, instance)
	return result
}
