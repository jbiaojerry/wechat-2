// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

const (
	// 微信服务器推送过来的消息类型
	MSG_TYPE_TEXT     = "text"     // 文本消息
	MSG_TYPE_IMAGE    = "image"    // 图片消息
	MSG_TYPE_VOICE    = "voice"    // 语音消息
	MSG_TYPE_VIDEO    = "video"    // 视频消息
	MSG_TYPE_LOCATION = "location" // 地理位置消息
	MSG_TYPE_LINK     = "link"     // 链接消息
	MSG_TYPE_EVENT    = "event"    // 事件推送
)

const (
	// 微信服务器推送过来的事件类型
	EVENT_TYPE_SUBSCRIBE         = "subscribe"         // 订阅, 包括点击订阅和扫描二维码
	EVENT_TYPE_UNSUBSCRIBE       = "unsubscribe"       // 取消订阅
	EVENT_TYPE_SCAN              = "SCAN"              // 已经订阅用户扫描二维码事件
	EVENT_TYPE_LOCATION          = "LOCATION"          // 上报地理位置事件
	EVENT_TYPE_CLICK             = "CLICK"             // 点击菜单拉取消息时的事件推送
	EVENT_TYPE_VIEW              = "VIEW"              // 点击菜单跳转链接时的事件推送
	EVENT_TYPE_MASSSENDJOBFINISH = "MASSSENDJOBFINISH" // 微信服务器推送群发结果
	EVENT_TYPE_MERCHANTORDER     = "merchant_order"    // 订单付款通知
)
