//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

type IWKDownloadEvent interface {
}

type IWKNavigationDelegateEvent interface {
}

type IReceiveScriptMessageEvent interface {
}

type IWKUIDelegateEvent interface {
}

type IWKURLSchemeHandlerEvent interface {
}

// WKScriptMessageHandlerProtocol

type TWkProcessMessageEvent func(sender TWKUserContentController, name, data string)

// WKNavigationDelegateProtocol

type TWKDecidePolicyForNavigationActionPreferences func(webView IWkWebview, navigationAction TWKNavigationAction, actionPolicy *WKNavigationActionPolicy, preferences *TWKWebpagePreferences)
type TWKDecidePolicyForNavigationResponse func(webView IWkWebview, navigationResponse TWKNavigationResponse, responsePolicy *WKNavigationResponsePolicy)
type TWkStartProvisionalNavigation func(webView IWkWebview, navigation TWKNavigation)
type TWkReceiveServerRedirectForProvisionalNavigation func(webView IWkWebview, navigation TWKNavigation)
type TWkFailProvisionalNavigationWithError func(webView IWkWebview, navigation TWKNavigation, error string)
type TWkCommitNavigation func(webView IWkWebview, navigation TWKNavigation)
type TWkFinishNavigation func(webView IWkWebview, navigation TWKNavigation)
type TWkFailNavigationWithError func(webView IWkWebview, navigation TWKNavigation, error string)

// TWkReceiveAuthenticationChallenge func(webView IWkWebview, challenge NSURLAuthenticationChallenge, var disposition NSURLSessionAuthChallengeDisposition, var credential INSURLCredential)

type TWkWebContentProcessDidTerminate func(webView IWkWebview)

// TWkAuthenticationChallengeShouldAllowDeprecatedTLS func(webView IWkWebview, challenge NSURLAuthenticationChallenge) Boolean

type TWkNavigationActionDidBecomeDownload func(webView IWkWebview, navigationAction TWKNavigationAction, download TWKDownload)
type TWkNavigationResponseDidBecomeDownload func(webView IWkWebview, navigationResponse TWKNavigationResponse, download TWKDownload)

// WKURLSchemeHandlerProtocol

type TWKStartURLSchemeTask func(webView IWkWebview, urlSchemeTask TWKURLSchemeTask)
type TWKStopURLSchemeTask func(webView IWkWebview, urlSchemeTask TWKURLSchemeTask)

// WKUIDelegateProtocol

type TWKCreateWebView func(webView IWkWebview, configuration_ TWKWebViewConfiguration, navigationAction TWKNavigationAction, windowFeatures TWKWindowFeatures) IWkWebview
type TWKRunJavaScriptAlert func(webView IWkWebview, message_ string, frame TWKFrameInfo)
type TWKRunJavaScriptConfirmCompletion func(webView IWkWebview, message_ string, frame TWKFrameInfo) bool
type TWKRunJavaScriptTextInputCompletion func(webView IWkWebview, prompt string, defaultText string, frame TWKFrameInfo) string
type TWKWebViewDidClose func(webView IWkWebview)

// WKDownload

type TWKDownloadCancelCompletionHandler func(webView IWkWebview, download TWKDownload, aData uintptr, aLength int32 /*NSData*/)
type TWKDownloadDecideDestinationUsingResponseSuggestedFilename func(webView IWkWebview, download TWKDownload, response TNSURLResponse, suggestedFilename string)
type TWKDownloadWillPerformHTTPRedirectionNewRequest func(webView IWkWebview, download TWKDownload, response TNSHTTPURLResponse, request TNSURLRequest)

// TWKDownloadReceiveAuthenticationChallenge func(webView IWkWebview, download IWKDownload, challenge NSURLAuthenticationChallenge)

type TWKDownloadFinish func(webView IWkWebview, download TWKDownload)
type TWKDownloadFailWithError func(webView IWkWebview, download TWKDownload, error string, aData uintptr, aLength int32 /*NSData*/)
