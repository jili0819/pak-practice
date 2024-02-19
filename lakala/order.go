package lakala

import (
	"context"
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/core"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"strings"
	"time"
)

const (
	AccountTypeAliPay = "ALIPAY"
	AccountTypeWeChat = "WECHAT"
)

type (
	// PayRequest 支付请求
	PayRequest struct {
		MerchantNo    string      `json:"merchant_no"`  // 拉卡拉分配的商户号
		TermNo        string      `json:"term_no"`      // 拉卡拉分配的业务终端号
		OutTradeNo    string      `json:"out_trade_no"` // 商户系统唯一，对应数据库表中外部请求流水号。
		AccountType   string      `json:"account_type"` // 微信：WECHAT 支付宝：ALIPAY 银联：UQRCODEPAY 翼支付: BESTPAY 苏宁易付宝: SUNING 拉卡拉支付账户：LKLACC 网联小钱包：NUCSPAY
		TransType     string      `json:"trans_type"`   // 41:NATIVE（（ALIPAY，云闪付支持） 51:JSAPI（微信公众号支付，支付宝服务窗支付，银联JS支付，翼支付JS支付、拉卡拉钱包支付） 71:微信小程序支付
		TotalAmount   string      `json:"total_amount"`
		LocationInfo  Location    `json:"location_info"`
		BusiMode      string      `json:"busi_mode"`    // 业务模式： ACQ-收单 不填，默认为“ACQ-收单”
		Subject       string      `json:"subject"`      // 标题，用于简单描述订单或商品主题，会传递给账户端 （账户端控制，实际最多42个字符），微信支付必送。
		PayOrderNo    string      `json:"pay_order_no"` // 拉卡拉订单系统订单号，以拉卡拉支付业务订单号为驱动的支付行为，需上传该字段。
		NotifyUrl     string      `json:"notify_url"`   // 商户通知地址，如果上传，且 pay_order_no 不存在情况下，则按此地址通知商户(详见“[交易通知]”接口)
		SettleType    string      `json:"settle_type"`  // “0”或者空，常规结算方式，如需接拉卡拉分账通需传“1”，商户未开通分账之前切记不用上送此参数。；
		Remark        string      `json:"remark"`
		PromoInfo     string      `json:"promo_info"` // 优惠相关信息，JSON格式
		AccBusiFields interface{} `json:"acc_busi_fields"`
	}

	Location struct {
		RequestIp string `json:"request_ip"` // 请求方的IP地址，存在必填，格式如36.45.36.95
		//BaseStation string `json:"base_station"` // 客户端设备的基站信息（主扫时基站信息使用该字段）
		//Location    string `json:"location"`     // 商户终端的地理位置，整体格式：纬度,经度，+表示北纬、东经，-表示南纬、 西经。 经度格式：1位正负号+3位整数+1位小数点+5位小数；纬度格式：1位正负号+2位整数+1位小数点+6位小数；举例：+31.221345,+121.12345
	}

	// ZhiFuBaoAccBusiFields 支付宝支付请求参数
	ZhiFuBaoAccBusiFields struct {
		UserId         string `json:"user_id"`         // 买家在支付宝的用户id,支付宝用户ID,支付宝的buyer_user_id ,trans_type为41-NATIVE情况下不需要传，为51情况下必须传入
		TimeoutExpress string `json:"timeout_express"` // 预下单有效时间,预下单的订单的有效时间，以分钟为单位。如果在有效时间内没有完成付款，则在账户端该订单失效。如果不上送，以账户端订单失效时间为准。 建议不超过15分钟。不传值则默认5分钟。
		ExtendParams   struct {
			SysServiceProviderId string `json:"sys_service_provider_id"` // 服务商的PID,系统商编号，该参数作为系统商返佣数据提取的依据，请填写系统商签约协议的 PID
			HbFqNum              string `json:"hb_fq_num"`               // 花呗分期期数,支付宝花呗分期必送字段: 花呗分期数 3：3期 6：6期 12：12期
			HbFqSellerPercent    string `json:"hb_fq_seller_percent"`    //卖家承担手续费比例,
			FoodOrderType        string `json:"food_order_type"`         // 点餐场景类型,qr_order（店内扫码点餐），pre_order（预点到店自提），home_delivery （外送到家），direct_payment（直接付款），other（其它）
		} `json:"extend_params"` // 业务扩展参数
		// GoodsDetail        string `json:"goods_detail"`         // 商品详情,订单包含的商品列表信息，Json数组。
		QuitUrl            string `json:"quit_url"`             // 用户付款中途退出返回商户网站的地址,transType=81时（即支付宝H5支付），此参数必传
		StoreId            string `json:"store_id"`             // 商户门店编号,支付宝收单上送
		DisablePayChannels string `json:"disable_pay_channels"` // “credit_group”表示禁用信用支付类（包含信用卡,花呗，花呗分期） “pcredit”表示禁用花呗,“pcreditpayInstallment”表示禁用花呗分期,“creditCard“表示禁用信用卡,如果想禁用多个可在枚举间加,隔开
		BusinessParams     string `json:"business_params"`      // 商户传入业务信息,商户传入业务信息，应用于安全，营销等参数直传场景，格式为 json 格式。 示例：{“enable_thirdpar ty_subsidy”:”N”,”source”:”xxxx”},source送值与吱口令生成接口中上送的source对应，值内容为“inderict_wx_吱口令申领接口中source值”
		MinAge             string `json:"min_age"`              // 允许的最小买家年龄
	}

	// WeChatAccBusiFields 微信支付请求参数
	WeChatAccBusiFields struct {
		UserId         string `json:"user_id"`         // 买家在支付宝的用户id,支付宝用户ID,支付宝的buyer_user_id ,trans_type为41-NATIVE情况下不需要传，为51情况下必须传入
		TimeoutExpress string `json:"timeout_express"` // 预下单有效时间,预下单的订单的有效时间，以分钟为单位。如果在有效时间内没有完成付款，则在账户端该订单失效。如果不上送，以账户端订单失效时间为准。 建议不超过15分钟。不传值则默认5分钟。
		DeviceInfo     string `json:"device_info"`     // 终端设备号(门店号或收银设备ID)，注意：PC网页或JSAPI支付请传”WEB”
	}

	// PayResponse 返回信息
	PayResponse struct {
		MerchantNo       string      `json:"merchant_no"`  // 商户号
		TermNo           string      `json:"term_no"`      // 商户请求流水号
		OutTradeNo       string      `json:"out_trade_no"` // 拉卡拉交易流水号
		LogNo            string      `json:"log_no"`       // 拉卡拉对账单流水号
		SettleMerchantNo string      `json:"settle_merchant_no"`
		SettleTermNo     string      `json:"settle_term_no"`
		AccRespFields    interface{} `json:"acc_resp_fields"`
	}
	// AliPayAccResp 支付宝返回
	AliPayAccResp struct {
		PrepayId string `json:"prepay_id"` // 预下单Id
	}

	// WeChatAccResp 微信返回
	WeChatAccResp struct {
		PrepayId  string `json:"prepay_id"`  // 预下单Id
		PaySign   string `json:"pay_sign"`   // 支付签名信息
		AppId     string `json:"app_id"`     // 小程序id
		TimeStamp string `json:"time_stamp"` // 时间戳
		NonceStr  string `json:"nonce_str"`  // 随机字符串
		Package   string `json:"package"`    // 订单详情扩展字符串
		SignType  string `json:"sign_type"`  // 签名方式
	}
)

func (l *Lakala) Pay(ctx context.Context, body *PayRequest) (resp *PayResponse, err error) {
	resp = &PayResponse{}
	switch body.AccountType {
	case AccountTypeAliPay:
		resp.AccRespFields = &AliPayAccResp{}
	case AccountTypeWeChat:
		resp.AccRespFields = &WeChatAccResp{}
	}

	if err = l.Post(ctx, true, "", "", nil, nil, body, resp); err != nil {
		return nil, err
	}
	return
}

// Post url 为空时，使用默认域名
func (l *Lakala) Post(ctx context.Context, needSign bool, url, uri string, header, query, body, resp interface{}) (err error) {
	if url == "" {
		url = l.Config.BaseUrl
	}
	goutInstance := gout.POST(url + uri)
	if needSign {
		timestamp := time.Now().Unix()
		goutInstance.SetHeader(core.H{
			"Authorization": strings.Join([]string{
				fmt.Sprintf("LKLAPI-SHA256withRSA appid=\"%s\"", l.Config.AppId),
				fmt.Sprintf("serial_no=\"%s\"", l.Config.SerialNo),
				fmt.Sprintf("timestamp=\"%d\"", timestamp),
				fmt.Sprintf("LKLAPI-SHA256withRSA appid=\"%s\"", l.Config.AppId),
				fmt.Sprintf("LKLAPI-SHA256withRSA appid=\"%s\"", l.Config.AppId),
				fmt.Sprintf("LKLAPI-SHA256withRSA appid=\"%s\"", l.Config.AppId),
			}, ","),
		})
	}

	if header != nil {
		goutInstance.SetHeader(header)
	}
	if query != nil {
		goutInstance.SetQuery(query)
	}
	if body != nil {
		goutInstance.SetJSON(body)
	}

	if err = goutInstance.Debug(l.Debug).BindJSON(resp).Do(); err != nil {
		logrus.Error("拉卡拉支付请求失败", zap.Any("body", body), zap.Error(err))
		return
	}
	return
}
