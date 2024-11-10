//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

// An object that conveys ongoing progress to the user for a specified task.
// https://developer.apple.com/documentation/foundation/nsprogress?language=objc
type NSProgress uintptr

// An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
// https://developer.apple.com/documentation/foundation/nsurl?language=objc
type NSURL uintptr

// The metadata associated with the response to an HTTP protocol URL load request.
// https://developer.apple.com/documentation/foundation/nshttpurlresponse?language=objc
type NSHTTPURLResponse uintptr

// A mutable URL load request that is independent of protocol or URL scheme.
// https://developer.apple.com/documentation/foundation/nsmutableurlrequest?language=objc
type NSMutableURLRequest uintptr

// A challenge from a server requiring authentication from the client.
// https://developer.apple.com/documentation/foundation/nsurlauthenticationchallenge?language=objc
type NSURLAuthenticationChallenge uintptr

// An authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
// https://developer.apple.com/documentation/foundation/nsurlcredential?language=objc
type NSURLCredential uintptr

// A URL load request that is independent of protocol or URL scheme.
// https://developer.apple.com/documentation/foundation/nsurlrequest?language=objc
type NSURLRequest uintptr

// The metadata associated with the response to a URL load request, independent of protocol and URL scheme.
// https://developer.apple.com/documentation/foundation/nsurlresponse?language=objc
type NSURLResponse uintptr

// An object that represents the download of a web resource.
// https://developer.apple.com/documentation/webkit/wkdownload?language=objc
type WKDownload uintptr

// A protocol you implement to track download progress and handle redirects, authentication challenges, and failures.
// https://developer.apple.com/documentation/webkit/wkdownloaddelegate?language=objc
type WKDownloadDelegateProtocol uintptr

// An object that contains information about a frame on a webpage.
// https://developer.apple.com/documentation/webkit/wkframeinfo?language=objc
type WKFrameInfo uintptr

// An object that tracks the loading progress of a webpage.
// https://developer.apple.com/documentation/webkit/wknavigation?language=objc
type WKNavigation uintptr

// An object that contains information about an action that causes navigation to occur.
// https://developer.apple.com/documentation/webkit/wknavigationaction?language=objc
type WKNavigationAction uintptr

// https://developer.apple.com/documentation/webkit/wknavigationresponse?language=objc
// An object that contains the response to a navigation request, and which you use to make navigation-related policy decisions.
type WKNavigationResponse uintptr

// An object that encapsulates the standard behaviors to apply to websites.
// https://developer.apple.com/documentation/webkit/wkpreferences?language=objc
type WKPreferences uintptr

// An interface for receiving messages from JavaScript code running in a webpage.
// https://developer.apple.com/documentation/webkit/wkscriptmessagehandler?language=objc
type WKScriptMessageHandlerProtocol uintptr
type WKScriptMessageHandler = WKScriptMessageHandlerProtocol

// An interface that WebKit uses to request custom resources from your app.
// https://developer.apple.com/documentation/webkit/wkurlschemetask?language=objc
type WKURLSchemeTask uintptr

// An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view.
// https://developer.apple.com/documentation/webkit/wkusercontentcontroller?language=objc
type WKUserContentController uintptr

// A script that the web view injects into a webpage.
// https://developer.apple.com/documentation/webkit/wkuserscript?language=objc
type WKUserScript uintptr

// An object that specifies the behaviors to use when loading and rendering page content.
// https://developer.apple.com/documentation/webkit/wkwebpagepreferences?language=objc
type WKWebpagePreferences uintptr

// An object that displays interactive web content, such as for an in-app browser.
// https://developer.apple.com/documentation/webkit/wkwebview?language=objc
type WkWebview uintptr

// A collection of properties that you use to initialize a web view.
// https://developer.apple.com/documentation/webkit/wkwebviewconfiguration?language=objc
type WKWebViewConfiguration uintptr

// A protocol for loading resources with URL schemes that WebKit doesn't handle.
// https://developer.apple.com/documentation/webkit/wkurlschemehandler?language=objc
type WKURLSchemeHandlerProtocol uintptr
type WKURLSchemeHandler = WKURLSchemeHandlerProtocol

// Display-related attributes that a webpage requests for its window.
// https://developer.apple.com/documentation/webkit/wkwindowfeatures?language=objc
type WKWindowFeatures uintptr

// https://developer.apple.com/documentation/webkit/wkuidelegate?language=objc
// The methods for presenting native user interface elements on behalf of a webpage.
type WKUIDelegateProtocol uintptr
type WKUIDelegate = WKUIDelegateProtocol

// Methods for accepting or rejecting navigation changes, and for tracking the progress of navigation requests.
// https://developer.apple.com/documentation/webkit/wknavigationdelegate?language=objc
type WKNavigationDelegateProtocol uintptr
type WKNavigationDelegate = WKNavigationDelegateProtocol

type NSEventModifierFlags = int32
type NSTimeInterval = float64

// Constants that specify how long the credential will be kept.
// https://developer.apple.com/documentation/foundation/nsurlcredentialpersistence?language=objc
type NSURLCredentialPersistence = int32

const (
	// The credential should not be stored.
	NSURLCredentialPersistenceNone NSURLCredentialPersistence = iota
	// The credential should be stored only for this session.
	NSURLCredentialPersistenceForSession
	// The credential should be stored in the keychain.
	NSURLCredentialPersistencePermanent
	// The credential should be stored permanently in the keychain, and in addition should be distributed to other devices based on the owning Apple ID.
	NSURLCredentialPersistenceSynchronizable // available in 10_8, 6_0
)

// The constants used to specify interaction with the cached responses.
// https://developer.apple.com/documentation/foundation/nsurlrequestcachepolicy?language=objc
type NSURLRequestCachePolicy = int32

const (
	// Use the caching logic defined in the protocol implementation, if any, for a particular URL load request.
	NSURLRequestUseProtocolCachePolicy NSURLRequestCachePolicy = iota
	// The URL load should be loaded only from the originating source.
	NSURLRequestReloadIgnoringLocalCacheData
	// Use existing cache data, regardless or age or expiration date, loading from originating source only if there is no cached data.
	NSURLRequestReturnCacheDataElseLoad
	// Use existing cache data, regardless or age or expiration date, and fail if no cached data is available.
	NSURLRequestReturnCacheDataDontLoad
	// Ignore local cache data, and instruct proxies and other intermediates to disregard their caches so far as the protocol allows.
	NSURLRequestReloadIgnoringLocalAndRemoteCacheData
	// Use cache data if the origin source can validate it; otherwise, load from the origin.
	NSURLRequestReloadRevalidatingCacheData
	NSURLRequestReloadIgnoringCacheData NSURLRequestCachePolicy = NSURLRequestReloadIgnoringLocalCacheData
)

// Constants that specify how a request uses network resources.
// https://developer.apple.com/documentation/foundation/nsurlrequestnetworkservicetype?language=objc
type NSURLRequestNetworkServiceType = int32

const (
	//A service type for standard network traffic.
	NSURLNetworkServiceTypeDefaultNSURLRequestNetworkServiceType NSURLRequestNetworkServiceType = iota
	//A service type for VoIP traffic.
	NSURLNetworkServiceTypeVoIP
	//	A service type for low-delay tolerant, very low-loss tolerant, inelastic flow, and constant packet rate connections.
	NSURLNetworkServiceTypeVideo
	//A service type for high-delay tolerant, high-loss tolerant, elastic flow, and variable size connections.
	NSURLNetworkServiceTypeBackground
	//A service type for low-delay tolerant, very low-loss tolerant, inelastic flow, and constant packet rate connections.
	NSURLNetworkServiceTypeVoice
	//A service for low-loss tolerant, inelastic flow, jitter tolerant, short but bursty rate, and variable size connections.
	NSURLNetworkServiceTypeCallSignaling NSURLRequestNetworkServiceType = 11
	//A service type for medium-delay tolerant, elastic and inelastic flow, bursty, and long-lived connections.
	NSURLNetworkServiceTypeResponsiveData NSURLRequestNetworkServiceType = 6
	//A service type for medium-delay tolerant, low-medium-loss tolerant, elastic flow, constant packet interval, and variable rate and size connections.
	NSURLNetworkServiceTypeAVStreaming NSURLRequestNetworkServiceType = 8
	//A service type for low-delay tolerant, low-to-medium-loss tolerant, elastic flow, variable packet interval, rate, size responsive and time-sensitive connections.
	NSURLNetworkServiceTypeResponsiveAV NSURLRequestNetworkServiceType = 9
)

// Constants that indicate how to render web view content.
// https://developer.apple.com/documentation/webkit/wkcontentmode?language=objc
type WKContentMode = int32

const (
	//The content mode that is appropriate for the current device.
	WKContentModeRecommended WKContentMode = iota
	//The content mode that represents a desktop experience.
	WKContentModeDesktop
	//The content mode that represents a mobile experience.
	WKContentModeMobile
)

// The type of action that triggered the navigation.
// https://developer.apple.com/documentation/webkit/wknavigationtype?language=objc
type WKNavigationType = int32

const (
	// A link activation.
	WKNavigationTypeLinkActivated WKNavigationType = iota
	// A request to submit a form.
	WKNavigationTypeFormSubmitted
	// A request for the frame’s next or previous item.
	WKNavigationTypeBackForward
	// A request to reload the webpage.
	WKNavigationTypeReload
	// A request to resubmit a form.
	WKNavigationTypeFormResubmitted
	// A navigation request that originates for some other reason.
	WKNavigationTypeOther
)

// Constants that indicate whether to allow or cancel navigation to a webpage from an action.
// https://developer.apple.com/documentation/webkit/wknavigationactionpolicy?language=objc
type WKNavigationActionPolicy = int32

const (
	// Cancel the navigation.
	WKNavigationActionPolicyCancel WKNavigationActionPolicy = iota
	// Allow the navigation to continue.
	WKNavigationActionPolicyAllow
	// Allow the download to proceed.
	WKNavigationActionPolicyDownload
)

// Constants that indicate whether to allow or cancel navigation to a webpage from a response.
// https://developer.apple.com/documentation/webkit/wknavigationresponsepolicy?language=objc
type WKNavigationResponsePolicy = int32

const (
	// Cancel the navigation.
	WKNavigationResponsePolicyCancel WKNavigationResponsePolicy = iota
	// Allow the navigation to continue.
	WKNavigationResponsePolicyAllow
	// Allow the download to proceed.
	WKNavigationResponsePolicyDownload
)
