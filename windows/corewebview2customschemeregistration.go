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
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// ICoreWebView2CustomSchemeRegistration Parent: IInterfacedObject
type ICoreWebView2CustomSchemeRegistration interface {
	IInterfacedObject
	// GetSchemeName
	//  The name of the custom scheme to register.
	GetSchemeName(outSchemeName *string) types.HRESULT // function
	// GetTreatAsSecure
	//  Whether the sites with this scheme will be treated as a
	//  [Secure Context](https://developer.mozilla.org/docs/Web/Security/Secure_Contexts)
	//  like an HTTPS site. This flag is only effective when HasAuthorityComponent
	//  is also set to `true`.
	//  `false` by default.
	GetTreatAsSecure(outTreatAsSecure *int32) types.HRESULT // function
	// SetTreatAsSecure
	//  Set if the scheme will be treated as a Secure Context.
	SetTreatAsSecure(treatAsSecure int32) types.HRESULT // function
	// GetAllowedOrigins
	//  List of origins that are allowed to issue requests with the custom
	//  scheme, such as XHRs and subresource requests that have an Origin header.
	//  The origin of any request (requests that have the
	//  [Origin header](https://developer.mozilla.org/docs/Web/HTTP/Headers/Origin))
	//  to the custom scheme URI needs to be in this list. No-origin requests
	//  are requests that do not have an Origin header, such as link
	//  navigations, embedded images and are always allowed.
	//  Note: POST requests always contain an Origin header, therefore
	//  AllowedOrigins must be set for even for same origin POST requests.
	//  Note that cross-origin restrictions still apply.
	//  From any opaque origin (Origin header is null), no cross-origin requests
	//  are allowed.
	//  If the list is empty, no cross-origin request to this scheme is
	//  allowed.
	//  Origins are specified as a string in the format of
	//  scheme://host:port.
	//  The origins are string pattern matched with `*` (matches 0 or more
	//  characters) and `?` (matches 0 or 1 character) wildcards just like
	//  the URI matching in the
	//  [AddWebResourceRequestedFilter API](/dotnet/api/microsoft.web.webview2.core.corewebview2.addwebresourcerequestedfilter).
	//  For example, "http://*.example.com:80".
	//  Here's a set of examples of what is allowed and not:
	//
	//  | Request URI | Originating URL | AllowedOrigins | Allowed |
	//  | -- | -- | -- | -- |
	//  | `custom-scheme:request` | `https://www.example.com` | {"https://www.example.com"} | Yes |
	//  | `custom-scheme:request` | `https://www.example.com` | {"https://*.example.com"} | Yes |
	//  | `custom-scheme:request` | `https://www.example.com` | {"https://www.example2.com"} | No |
	//  | `custom-scheme-with-authority://host/path` | `custom-scheme-with-authority://host2` | {""} | No |
	//  | `custom-scheme-with-authority://host/path` | `custom-scheme-with-authority2://host` | {"custom-scheme-with-authority2://*"} | Yes |
	//  | `custom-scheme-without-authority:path` | custom-scheme-without-authority:path2 | {"custom-scheme-without-authority:*"} | No |
	//  | `custom-scheme-without-authority:path` | custom-scheme-without-authority:path2 | {"*"} | Yes |
	//
	//  The returned strings and the array itself must be deallocated with
	//  CoTaskMemFree.
	//  * out allowedOrigins: PPWideChar1 --> out allowedOrigins: PPWideChar ************** WEBVIEW4DELPHI ************** *
	GetAllowedOrigins(outAllowedOriginsCount *uint32, outAllowedOrigins *types.PPWideChar) types.HRESULT // function
	// SetAllowedOrigins
	//  Set the array of origins that are allowed to use the scheme.
	//  * var allowedOrigins: PWideChar --> allowedOrigins: PPWideChar ************** WEBVIEW4DELPHI ************** *
	SetAllowedOrigins(allowedOriginsCount uint32, allowedOrigins types.PPWideChar) types.HRESULT // function
	// GetHasAuthorityComponent
	//  Set this property to `true` if the URIs with this custom
	//  scheme will have an authority component (a host for custom schemes).
	//  Specifically, if you have a URI of the following form you should set the
	//  `HasAuthorityComponent` value as listed.
	//
	//  | URI | Recommended HasAuthorityComponent value |
	//  | -- | -- |
	//  | `custom-scheme-with-authority://host/path` | `true` |
	//  | `custom-scheme-without-authority:path` | `false` |
	//
	//  When this property is set to `true`, the URIs with this scheme will be
	//  interpreted as having a
	//  [scheme and host](https://html.spec.whatwg.org/multipage/origin.html#concept-origin-tuple)
	//  origin similar to an http URI. Note that the port and user
	//  information are never included in the computation of origins for
	//  custom schemes.
	//  If this property is set to `false`, URIs with this scheme will have an
	//  [opaque origin](https://html.spec.whatwg.org/multipage/origin.html#concept-origin-opaque)
	//  similar to a data URI.
	//  This property is `false` by default.
	//
	//  Note: For custom schemes registered as having authority component,
	//  navigations to URIs without authority of such custom schemes will fail.
	//  However, if the content inside WebView2 references
	//  a subresource with a URI that does not have
	//  an authority component, but of a custom scheme that is registered as
	//  having authority component, the URI will be interpreted as a relative path
	//  as specified in [RFC3986](https://www.rfc-editor.org/rfc/rfc3986).
	//  For example, `custom-scheme-with-authority:path` will be interpreted
	//  as `custom-scheme-with-authority://host/path`.
	//  However, this behavior cannot be guaranteed to remain in future
	//  releases so it is recommended not to rely on this behavior.
	GetHasAuthorityComponent(outHasAuthorityComponent *int32) types.HRESULT // function
	// SetHasAuthorityComponent
	//  Get has authority component.
	SetHasAuthorityComponent(hasAuthorityComponent int32) types.HRESULT // function
	AsIntfCustomSchemeRegistration() uintptr
}

type TCoreWebView2CustomSchemeRegistration struct {
	TInterfacedObject
}

func (m *TCoreWebView2CustomSchemeRegistration) GetSchemeName(outSchemeName *string) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var schemeNamePtr uintptr
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(2, m.Instance(), uintptr(base.UnsafePointer(&schemeNamePtr)))
	*outSchemeName = api.GoStr(schemeNamePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomSchemeRegistration) GetTreatAsSecure(outTreatAsSecure *int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var treatAsSecurePtr uintptr
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(3, m.Instance(), uintptr(base.UnsafePointer(&treatAsSecurePtr)))
	*outTreatAsSecure = int32(treatAsSecurePtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomSchemeRegistration) SetTreatAsSecure(treatAsSecure int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(4, m.Instance(), uintptr(treatAsSecure))
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomSchemeRegistration) GetAllowedOrigins(outAllowedOriginsCount *uint32, outAllowedOrigins *types.PPWideChar) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var allowedOriginsCountPtr uintptr
	var allowedOriginsPtr uintptr
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(5, m.Instance(), uintptr(base.UnsafePointer(&allowedOriginsCountPtr)), uintptr(base.UnsafePointer(&allowedOriginsPtr)))
	*outAllowedOriginsCount = uint32(allowedOriginsCountPtr)
	*outAllowedOrigins = types.PPWideChar(allowedOriginsPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomSchemeRegistration) SetAllowedOrigins(allowedOriginsCount uint32, allowedOrigins types.PPWideChar) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(6, m.Instance(), uintptr(allowedOriginsCount), uintptr(allowedOrigins))
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomSchemeRegistration) GetHasAuthorityComponent(outHasAuthorityComponent *int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	var hasAuthorityComponentPtr uintptr
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(7, m.Instance(), uintptr(base.UnsafePointer(&hasAuthorityComponentPtr)))
	*outHasAuthorityComponent = int32(hasAuthorityComponentPtr)
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomSchemeRegistration) SetHasAuthorityComponent(hasAuthorityComponent int32) types.HRESULT {
	if !m.IsValid() {
		return 0
	}
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(8, m.Instance(), uintptr(hasAuthorityComponent))
	return types.HRESULT(r)
}

func (m *TCoreWebView2CustomSchemeRegistration) AsIntfCustomSchemeRegistration() uintptr {
	return m.GetIntfPointer(0)
}

// NewCoreWebView2CustomSchemeRegistration class constructor
func NewCoreWebView2CustomSchemeRegistration(customSchemeInfo TWVCustomSchemeInfo) ICoreWebView2CustomSchemeRegistration {
	customSchemeInfoPtr := customSchemeInfo.ToPas()
	var customSchemeRegistrationPtr uintptr // ICoreWebView2CustomSchemeRegistration
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(0, uintptr(base.UnsafePointer(customSchemeInfoPtr)), uintptr(base.UnsafePointer(&customSchemeRegistrationPtr)))
	ret := AsCoreWebView2CustomSchemeRegistration(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, customSchemeRegistrationPtr)
	}
	return ret
}

// NewCoreWebView2CustomSchemeRegistrationWithStrX2BoolX2 class constructor
func NewCoreWebView2CustomSchemeRegistrationWithStrX2BoolX2(schemeName string, allowedDomains string, treatAsSecure bool, hasAuthorityComponent bool) ICoreWebView2CustomSchemeRegistration {
	var customSchemeRegistrationPtr uintptr // ICoreWebView2CustomSchemeRegistration
	r := coreWebView2CustomSchemeRegistrationAPI().SysCallN(1, api.PasStr(schemeName), api.PasStr(allowedDomains), api.PasBool(treatAsSecure), api.PasBool(hasAuthorityComponent), uintptr(base.UnsafePointer(&customSchemeRegistrationPtr)))
	ret := AsCoreWebView2CustomSchemeRegistration(r)
	if intf, ok := ret.(base.IIntfs); ok {
		intf.Create(1)
		intf.SetIntfPointer(0, customSchemeRegistrationPtr)
	}
	return ret
}

var (
	coreWebView2CustomSchemeRegistrationOnce   base.Once
	coreWebView2CustomSchemeRegistrationImport *imports.Imports = nil
)

func coreWebView2CustomSchemeRegistrationAPI() *imports.Imports {
	coreWebView2CustomSchemeRegistrationOnce.Do(func() {
		coreWebView2CustomSchemeRegistrationImport = api.NewDefaultImports()
		coreWebView2CustomSchemeRegistrationImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_Create", 0), // constructor NewCoreWebView2CustomSchemeRegistration
			/* 1 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_CreateWithStrX2BoolX2", 0), // constructor NewCoreWebView2CustomSchemeRegistrationWithStrX2BoolX2
			/* 2 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_Get_SchemeName", 0), // function GetSchemeName
			/* 3 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_Get_TreatAsSecure", 0), // function GetTreatAsSecure
			/* 4 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_Set_TreatAsSecure", 0), // function SetTreatAsSecure
			/* 5 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_GetAllowedOrigins", 0), // function GetAllowedOrigins
			/* 6 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_SetAllowedOrigins", 0), // function SetAllowedOrigins
			/* 7 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_Get_HasAuthorityComponent", 0), // function GetHasAuthorityComponent
			/* 8 */ imports.NewTable("TCoreWebView2CustomSchemeRegistration_Set_HasAuthorityComponent", 0), // function SetHasAuthorityComponent
		}
	})
	return coreWebView2CustomSchemeRegistrationImport
}
