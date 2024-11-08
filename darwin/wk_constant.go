//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

type NSURL uintptr
type NSHTTPURLResponse uintptr
type NSMutableURLRequest uintptr
type NSURLAuthenticationChallenge uintptr
type NSURLCredential uintptr
type NSURLRequest uintptr
type NSURLResponse uintptr
type WKDownload uintptr
type WKDownloadDelegateProtocol uintptr
type WKFrameInfo uintptr
type WKNavigation uintptr
type WKNavigationAction uintptr
type WKNavigationDelegateProtocol uintptr
type WKNavigationResponse uintptr
type WKPreferences uintptr
type WKScriptMessageHandlerProtocol uintptr
type WKUIDelegateProtocol uintptr
type WKURLSchemeHandlerProtocol uintptr
type WKURLSchemeTask uintptr
type WKUserContentController uintptr
type WKUserScript uintptr
type WKWebpagePreferences uintptr
type WkWebview uintptr
type WKWebViewConfiguration uintptr
type WKURLSchemeHandler = WKURLSchemeHandlerProtocol
type WKWindowFeatures uintptr
type WKUIDelegate = WKUIDelegateProtocol
type WKNavigationDelegate = WKNavigationDelegateProtocol
type NSProgress uintptr

type NSURLCredentialPersistence = int32

const (
	NSURLCredentialPersistenceNone NSURLCredentialPersistence = iota
	NSURLCredentialPersistenceForSession
	NSURLCredentialPersistencePermanent
	NSURLCredentialPersistenceSynchronizable // available in 10_8, 6_0
)

type NSURLRequestCachePolicy = int32

const (
	NSURLRequestUseProtocolCachePolicy NSURLRequestCachePolicy = iota
	NSURLRequestReloadIgnoringLocalCacheData
	NSURLRequestReturnCacheDataElseLoad
	NSURLRequestReturnCacheDataDontLoad
	NSURLRequestReloadIgnoringLocalAndRemoteCacheData
	NSURLRequestReloadRevalidatingCacheData
	NSURLRequestReloadIgnoringCacheData NSURLRequestCachePolicy = NSURLRequestReloadIgnoringLocalCacheData
)

type NSURLRequestNetworkServiceType = int32

const (
	NSURLNetworkServiceTypeDefault NSURLRequestNetworkServiceType = iota
	NSURLNetworkServiceTypeVoIP
	NSURLNetworkServiceTypeVideo
	NSURLNetworkServiceTypeBackground
	NSURLNetworkServiceTypeVoice
)

type WKContentMode = int32
type WKNavigationType = int32
type NSEventModifierFlags = int32
type WKNavigationActionPolicy = int32
type WKNavigationResponsePolicy = int32
type NSTimeInterval = float64
