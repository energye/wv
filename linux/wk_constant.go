//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import "github.com/energye/lcl/types"

type WebKitContextMenuItem uintptr
type WebKitContextMenu uintptr
type WebKitJavascriptResult uintptr
type WebKitNavigationAction uintptr
type WebKitPolicyDecision uintptr
type WebKitNavigationPolicyDecision = WebKitPolicyDecision
type WebKitResponsePolicyDecision = WebKitPolicyDecision
type WebKitSettings uintptr
type WebKitURIRequest uintptr
type WebKitURIResponse uintptr
type WebKitURISchemeRequest uintptr
type WebKitWebView uintptr
type WebKitURISchemeResponse uintptr
type WebKitWebsitePolicies uintptr
type WebKitWebContext uintptr
type WebKitCookieManager uintptr
type PWkAction uintptr
type PSoupDate uintptr
type PSoupMessageHeaders uintptr
type PInputStream uintptr
type PSoupCookie uintptr
type PList uintptr

type TScrolledWindow = uintptr

type TWkProcessId = int32

const (
	WkPID_BROWSER TWkProcessId = iota
	WkPID_RENDER
)

type WebKitContextMenuAction = int32

const (
	WEBKIT_CONTEXT_MENU_ACTION_NO_ACTION WebKitContextMenuAction = iota
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_LINK_IN_NEW_WINDOW
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_LINK_TO_DISK
	WEBKIT_CONTEXT_MENU_ACTION_COPY_LINK_TO_CLIPBOARD
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_IMAGE_IN_NEW_WINDOW
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_IMAGE_TO_DISK
	WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_TO_CLIPBOARD
	WEBKIT_CONTEXT_MENU_ACTION_COPY_IMAGE_URL_TO_CLIPBOARD
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_FRAME_IN_NEW_WINDOW
	WEBKIT_CONTEXT_MENU_ACTION_GO_BACK
	WEBKIT_CONTEXT_MENU_ACTION_GO_FORWARD
	WEBKIT_CONTEXT_MENU_ACTION_STOP
	WEBKIT_CONTEXT_MENU_ACTION_RELOAD
	WEBKIT_CONTEXT_MENU_ACTION_COPY
	WEBKIT_CONTEXT_MENU_ACTION_CUT
	WEBKIT_CONTEXT_MENU_ACTION_PASTE
	WEBKIT_CONTEXT_MENU_ACTION_DELETE
	WEBKIT_CONTEXT_MENU_ACTION_SELECT_ALL
	WEBKIT_CONTEXT_MENU_ACTION_INPUT_METHODS
	WEBKIT_CONTEXT_MENU_ACTION_UNICODE
	WEBKIT_CONTEXT_MENU_ACTION_SPELLING_GUESS
	WEBKIT_CONTEXT_MENU_ACTION_NO_GUESSES_FOUND
	WEBKIT_CONTEXT_MENU_ACTION_IGNORE_SPELLING
	WEBKIT_CONTEXT_MENU_ACTION_LEARN_SPELLING
	WEBKIT_CONTEXT_MENU_ACTION_IGNORE_GRAMMAR
	WEBKIT_CONTEXT_MENU_ACTION_FONT_MENU
	WEBKIT_CONTEXT_MENU_ACTION_BOLD
	WEBKIT_CONTEXT_MENU_ACTION_ITALIC
	WEBKIT_CONTEXT_MENU_ACTION_UNDERLINE
	WEBKIT_CONTEXT_MENU_ACTION_OUTLINE
	WEBKIT_CONTEXT_MENU_ACTION_INSPECT_ELEMENT
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_VIDEO_IN_NEW_WINDOW
	WEBKIT_CONTEXT_MENU_ACTION_OPEN_AUDIO_IN_NEW_WINDOW
	WEBKIT_CONTEXT_MENU_ACTION_COPY_VIDEO_LINK_TO_CLIPBOARD
	WEBKIT_CONTEXT_MENU_ACTION_COPY_AUDIO_LINK_TO_CLIPBOARD
	WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_CONTROLS
	WEBKIT_CONTEXT_MENU_ACTION_TOGGLE_MEDIA_LOOP
	WEBKIT_CONTEXT_MENU_ACTION_ENTER_VIDEO_FULLSCREEN
	WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PLAY
	WEBKIT_CONTEXT_MENU_ACTION_MEDIA_PAUSE
	WEBKIT_CONTEXT_MENU_ACTION_MEDIA_MUTE
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_VIDEO_TO_DISK
	WEBKIT_CONTEXT_MENU_ACTION_DOWNLOAD_AUDIO_TO_DISK
	WEBKIT_CONTEXT_MENU_ACTION_INSERT_EMOJI
	WEBKIT_CONTEXT_MENU_ACTION_PASTE_AS_PLAIN_TEXT
	WEBKIT_CONTEXT_MENU_ACTION_CUSTOM = 10000
)

type TSoupDateFormat = int32

const (
	SOUP_DATE_HTTP            TSoupDateFormat = iota + 1
	SOUP_DATE_COOKIE          TSoupDateFormat = 2
	SOUP_DATE_RFC2822         TSoupDateFormat = 3
	SOUP_DATE_ISO8601_COMPACT TSoupDateFormat = 4
	SOUP_DATE_ISO8601_FULL    TSoupDateFormat = 5
	SOUP_DATE_ISO8601         TSoupDateFormat = 5
	SOUP_DATE_ISO8601_XMLRPC  TSoupDateFormat = 6
)

// WebKitLoadEvent 值用于表示在#WebKitWebView加载操作期间发生的不同事件
type WebKitLoadEvent = int32

const (
	WEBKIT_LOAD_STARTED    WebKitLoadEvent = iota // 已发出新的加载请求。尚未接收到数据，已分配空结构来执行加载;由于传输问题，例如无法解析名称或连接到端口，负载仍然可能失败
	WEBKIT_LOAD_REDIRECTED                        // 临时数据源收到服务器重定向
	WEBKIT_LOAD_COMMITTED                         // 内容开始到达页面加载。建立了必要的运输要求，并且正在执行装载
	WEBKIT_LOAD_FINISHED                          // 加载完成。所有资源都已完成加载，或者在加载操作期间出现错误
)

// WebKitHardwareAccelerationPolicy 用于确定硬件加速策略的值
type WebKitHardwareAccelerationPolicy = int32

const (
	WEBKIT_HARDWARE_ACCELERATION_POLICY_ON_DEMAND WebKitHardwareAccelerationPolicy = iota // 硬件加速根据web内容的请求启用/禁用。
	WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS                                            // 硬件加速总是启用的，即使对于没有请求它的网站也是如此
	WEBKIT_HARDWARE_ACCELERATION_POLICY_NEVER                                             // 硬件加速总是被禁用的，即使是对于请求它的网站
)

// WebKitCacheModel 用于确定#WebKitWebContext缓存模型的值
type WebKitCacheModel = int32

const (
	WEBKIT_CACHE_MODEL_DOCUMENT_VIEWER  WebKitCacheModel = iota // 完全禁用缓存，这将大大减少内存使用。对于只访问单个本地文件而不导航到其他页面的应用程序非常有用。不会缓存任何远程资源。
	WEBKIT_CACHE_MODEL_WEB_BROWSER                              // 为查看一系列本地文件而优化的缓存模型——例如，文档查看器或网站设计器。WebKit将缓存适度数量的资源。
	WEBKIT_CACHE_MODEL_DOCUMENT_BROWSER                         // 通过缓存大量资源和以前查看的内容，大大提高文档加载速度。
)

// WebKitCookiePersistentStorage 用于表示cookie持久存储类型的值
type WebKitCookiePersistentStorage = int32

const (
	WEBKIT_COOKIE_PERSISTENT_STORAGE_TEXT   WebKitCookiePersistentStorage = iota // cookie存储在Mozilla的" Cookies .txt"格式的文本文件中
	WEBKIT_COOKIE_PERSISTENT_STORAGE_SQLITE                                      // cookie以当前Mozilla格式存储在SQLite文件中
)

// WebKitCookieAcceptPolicy 用于表示cookie接受策略的值
type WebKitCookieAcceptPolicy = int32

const (
	WEBKIT_COOKIE_POLICY_ACCEPT_ALWAYS         WebKitCookieAcceptPolicy = iota // 无条件接受所有cookie
	WEBKIT_COOKIE_POLICY_ACCEPT_NEVER                                          // 无条件拒绝所有cookie
	WEBKIT_COOKIE_POLICY_ACCEPT_NO_THIRD_PARTY                                 // 仅接受由加载的主文档设置的cookie
)

// WebKitUserContentInjectedFrames 指定要在哪些框架中插入用户样式表
type WebKitUserContentInjectedFrames = int32

const (
	WEBKIT_USER_CONTENT_INJECT_ALL_FRAMES WebKitUserContentInjectedFrames = iota // 将用户样式表插入到web视图加载的所有框架中，包括嵌套框架。这是默认值
	WEBKIT_USER_CONTENT_INJECT_TOP_FRAME                                         // 在web视图加载的顶层框架中*只*插入用户样式表，而在嵌套框架中*不插入
)

// WebKitUserStyleLevel 指定如何处理用户样式表
type WebKitUserStyleLevel = int32

const (
	WEBKIT_USER_STYLE_LEVEL_USER   WebKitUserStyleLevel = iota // 样式表是用户样式表，其内容总是覆盖其他样式表。这是默认值
	WEBKIT_USER_STYLE_LEVEL_AUTHOR                             // 样式表将被视为由加载的文档提供的样式表。这意味着其他用户样式表可能仍然会覆盖它
)

// WebKitUserScriptInjectionTime 指定将在文档的哪个位置插入用户脚本
type WebKitUserScriptInjectionTime = int32

const (
	WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_START WebKitUserScriptInjectionTime = iota // 在加载文档的开头插入用户脚本的代码。这是默认值
	WEBKIT_USER_SCRIPT_INJECT_AT_DOCUMENT_END                                        // 在加载文档的末尾插入用户脚本的代码
)

// WebKitPolicyDecisionType 用于确定期间的策略决策类型的值
type WebKitPolicyDecisionType = int32

const (
	// 这种类型的政策决定是在WebKit即将在主框架或子框架中导航到新页面时请求的
	// 可接受的政策决定是webkit_policy_decision_use()或webkit_policy_decision_ignore()
	// 这种类型的政策决定始终是 WebKitNavigationPolicyDecision
	WEBKIT_POLICY_DECISION_TYPE_NAVIGATION_ACTION WebKitPolicyDecisionType = iota
	// 当WebKit要创建一个新窗口时，需要这种类型的策略决策。可接受的策略决策是webkit_policy_decision_use()或webkit_policy_decision_ignore()
	// 这种类型的策略决策总是一个 WebKitNavigationPolicyDecision
	// 这些决定对于为新窗口实现特殊操作非常有用，例如，当键盘修饰符处于活动状态时，强制在选项卡中打开新窗口，或者处理<a>元素上的特殊目标属性
	WEBKIT_POLICY_DECISION_TYPE_NEW_WINDOW_ACTION
	// 当WebKit收到网络资源的响应并准备开始加载时，使用这种类型的决策
	// 注意，这些资源包括页面的所有子资源，例如图像和样式表以及主文档
	// 对该决策的适当策略响应是webkit_policy_decision_use()、webkit_policy_decision_ignore()或webkit_policy_decision_download()
	// 这种类型的策略决策总是一个 WebKitResponsePolicyDecision
	// 这个决定对于强制下载某些类型的资源而不是在WebView中呈现或完全阻止资源的传输非常有用
	WEBKIT_POLICY_DECISION_TYPE_RESPONSE
)

// WebKitAutoplayPolicy 用于指定自动播放策略的值
type WebKitAutoplayPolicy = int32

const (
	WEBKIT_AUTOPLAY_ALLOW               WebKitAutoplayPolicy = iota // 不要限制自动播放
	WEBKIT_AUTOPLAY_ALLOW_WITHOUT_SOUND                             // 允许视频自动播放，如果他们没有音轨，或者如果他们的音轨是静音的
	WEBKIT_AUTOPLAY_DENY                                            // 永远不要允许自动播放
)

// WebKitWebProcessTerminationReason 用于指定web进程异常终止原因的值
type WebKitWebProcessTerminationReason = int32

const (
	WEBKIT_WEB_PROCESS_CRASHED               WebKitWebProcessTerminationReason = iota // web进程崩溃
	WEBKIT_WEB_PROCESS_EXCEEDED_MEMORY_LIMIT                                          // web进程超过内存限制。处理步骤
	WEBKIT_WEB_PROCESS_TERMINATED_BY_API                                              // web进程终止由API调用请求。自:2.34
)

type JSCValuePropertyFlags = int32

const (
	JSC_VALUE_PROPERTY_CONFIGURABLE JSCValuePropertyFlags = 1 << 0
	JSC_VALUE_PROPERTY_ENUMERABLE   JSCValuePropertyFlags = 1 << 1
	JSC_VALUE_PROPERTY_WRITABLE     JSCValuePropertyFlags = 1 << 2
)

type JSCTypedArrayType = int32

const (
	JSC_TYPED_ARRAY_NONE JSCTypedArrayType = 0
	JSC_TYPED_ARRAY_INT8
	JSC_TYPED_ARRAY_INT16
	JSC_TYPED_ARRAY_INT32
	JSC_TYPED_ARRAY_INT64
	JSC_TYPED_ARRAY_UINT8
	JSC_TYPED_ARRAY_UINT8_CLAMPED
	JSC_TYPED_ARRAY_UINT16
	JSC_TYPED_ARRAY_UINT32
	JSC_TYPED_ARRAY_UINT64
	JSC_TYPED_ARRAY_FLOAT32
	JSC_TYPED_ARRAY_FLOAT64
)

type JSCCheckSyntaxMode = int32

const (
	JSC_CHECK_SYNTAX_MODE_SCRIPT JSCCheckSyntaxMode = iota
	JSC_CHECK_SYNTAX_MODE_MODULE
)

type JSCCheckSyntaxResult = int32

const (
	JSC_CHECK_SYNTAX_RESULT_SUCCESS JSCCheckSyntaxResult = iota
	JSC_CHECK_SYNTAX_RESULT_RECOVERABLE_ERROR
	JSC_CHECK_SYNTAX_RESULT_IRRECOVERABLE_ERROR
	JSC_CHECK_SYNTAX_RESULT_UNTERMINATED_LITERAL_ERROR
	JSC_CHECK_SYNTAX_RESULT_OUT_OF_MEMORY_ERROR
	JSC_CHECK_SYNTAX_RESULT_STACK_OVERFLOW_ERROR
)

type JSType = int32

const (
	JtInvalid JSType = iota
	JtNotSupported
	JtUndefined
	JtNull
	JtString
	JtNumber
	JtInteger
	JtBoolean
	JtException
)

type TSoupMessageHeadersType = int32

const (
	SOUP_MESSAGE_HEADERS_REQUEST TSoupMessageHeadersType = iota
	SOUP_MESSAGE_HEADERS_RESPONSE
	SOUP_MESSAGE_HEADERS_MULTIPART
)

type TGdkEventType = int32

const (
	TGdkEventTypeMinValue   TGdkEventType = -0x7FFFFFFF
	GDK_NOTHING             TGdkEventType = -1
	GDK_DELETE              TGdkEventType = 0
	GDK_DESTROY             TGdkEventType = 1
	GDK_EXPOSE              TGdkEventType = 2
	GDK_MOTION_NOTIFY       TGdkEventType = 3
	GDK_BUTTON_PRESS        TGdkEventType = 4
	GDK_DOUBLE_BUTTON_PRESS TGdkEventType = 5
	GDK_2BUTTON_PRESS       TGdkEventType = 5
	GDK_TRIPLE_BUTTON_PRESS TGdkEventType = 6
	GDK_3BUTTON_PRESS       TGdkEventType = 6
	GDK_BUTTON_RELEASE      TGdkEventType = 7
	GDK_KEY_PRESS           TGdkEventType = 8
	GDK_KEY_RELEASE         TGdkEventType = 9
	GDK_ENTER_NOTIFY        TGdkEventType = 10
	GDK_LEAVE_NOTIFY        TGdkEventType = 11
	GDK_FOCUS_CHANGE        TGdkEventType = 12
	GDK_CONFIGURE           TGdkEventType = 13
	GDK_MAP                 TGdkEventType = 14
	GDK_UNMAP               TGdkEventType = 15
	GDK_PROPERTY_NOTIFY     TGdkEventType = 16
	GDK_SELECTION_CLEAR     TGdkEventType = 17
	GDK_SELECTION_REQUEST   TGdkEventType = 18
	GDK_SELECTION_NOTIFY    TGdkEventType = 19
	GDK_PROXIMITY_IN        TGdkEventType = 20
	GDK_PROXIMITY_OUT       TGdkEventType = 21
	GDK_DRAG_ENTER          TGdkEventType = 22
	GDK_DRAG_LEAVE          TGdkEventType = 23
	GDK_DRAG_MOTION_        TGdkEventType = 24
	GDK_DRAG_STATUS_        TGdkEventType = 25
	GDK_DROP_START          TGdkEventType = 26
	GDK_DROP_FINISHED       TGdkEventType = 27
	GDK_CLIENT_EVENT        TGdkEventType = 28
	GDK_VISIBILITY_NOTIFY   TGdkEventType = 29
	GDK_SCROLL              TGdkEventType = 31
	GDK_WINDOW_STATE        TGdkEventType = 32
	GDK_SETTING             TGdkEventType = 33
	GDK_OWNER_CHANGE        TGdkEventType = 34
	GDK_GRAB_BROKEN         TGdkEventType = 35
	GDK_DAMAGE              TGdkEventType = 36
	GDK_TOUCH_BEGIN         TGdkEventType = 37
	GDK_TOUCH_UPDATE        TGdkEventType = 38
	GDK_TOUCH_END           TGdkEventType = 39
	GDK_TOUCH_CANCEL        TGdkEventType = 40
	GDK_TOUCHPAD_SWIPE      TGdkEventType = 41
	GDK_TOUCHPAD_PINCH      TGdkEventType = 42
	GDK_PAD_BUTTON_PRESS    TGdkEventType = 43
	GDK_PAD_BUTTON_RELEASE  TGdkEventType = 44
	GDK_PAD_RING            TGdkEventType = 45
	GDK_PAD_STRIP           TGdkEventType = 46
	GDK_PAD_GROUP_MODE      TGdkEventType = 47
	GDK_EVENT_LAST          TGdkEventType = 48
	TGdkEventTypeMaxValue   TGdkEventType = 0x7FFFFFFF
)

type TGdkModifierTypeIdx = int32

const (
	TGdkModifierTypeIdxMinValue   TGdkModifierTypeIdx = 0
	GDK_SHIFT_MASK                TGdkModifierTypeIdx = 0
	GDK_LOCK_MASK                 TGdkModifierTypeIdx = 1
	GDK_CONTROL_MASK              TGdkModifierTypeIdx = 2
	GDK_MOD1_MASK                 TGdkModifierTypeIdx = 3
	GDK_MOD2_MASK                 TGdkModifierTypeIdx = 4
	GDK_MOD3_MASK                 TGdkModifierTypeIdx = 5
	GDK_MOD4_MASK                 TGdkModifierTypeIdx = 6
	GDK_MOD5_MASK                 TGdkModifierTypeIdx = 7
	GDK_BUTTON1_MASK              TGdkModifierTypeIdx = 8
	GDK_BUTTON2_MASK              TGdkModifierTypeIdx = 9
	GDK_BUTTON3_MASK              TGdkModifierTypeIdx = 10
	GDK_BUTTON4_MASK              TGdkModifierTypeIdx = 11
	GDK_BUTTON5_MASK              TGdkModifierTypeIdx = 12
	GDK_MODIFIER_RESERVED_13_MASK TGdkModifierTypeIdx = 13
	GDK_MODIFIER_RESERVED_14_MASK TGdkModifierTypeIdx = 14
	GDK_MODIFIER_RESERVED_15_MASK TGdkModifierTypeIdx = 15
	GDK_MODIFIER_RESERVED_16_MASK TGdkModifierTypeIdx = 16
	GDK_MODIFIER_RESERVED_17_MASK TGdkModifierTypeIdx = 17
	GDK_MODIFIER_RESERVED_18_MASK TGdkModifierTypeIdx = 18
	GDK_MODIFIER_RESERVED_19_MASK TGdkModifierTypeIdx = 19
	GDK_MODIFIER_RESERVED_20_MASK TGdkModifierTypeIdx = 20
	GDK_MODIFIER_RESERVED_21_MASK TGdkModifierTypeIdx = 21
	GDK_MODIFIER_RESERVED_22_MASK TGdkModifierTypeIdx = 22
	GDK_MODIFIER_RESERVED_23_MASK TGdkModifierTypeIdx = 23
	GDK_MODIFIER_RESERVED_24_MASK TGdkModifierTypeIdx = 24
	GDK_MODIFIER_RESERVED_25_MASK TGdkModifierTypeIdx = 25
	GDK_SUPER_MASK                TGdkModifierTypeIdx = 26
	GDK_HYPER_MASK                TGdkModifierTypeIdx = 27
	GDK_META_MASK                 TGdkModifierTypeIdx = 28
	GDK_MODIFIER_RESERVED_29_MASK TGdkModifierTypeIdx = 29
	GDK_RELEASE_MASK              TGdkModifierTypeIdx = 30
	TGdkModifierTypeIdxMaxValue   TGdkModifierTypeIdx = 31
)

// TGdkModifierType SET
type TGdkModifierType = types.TSet
