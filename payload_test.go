package jpush

import (
	"testing"
	"bytes"
)

func TestPayload(t *testing.T) {
	platform := NewPlatform()
	platform.IsAll = true
	audience := NewAudience()
	audience.SetTag("深圳", "北京")
	androidAlert := "Hi, JPush!"
	extra := map[string]interface{}{
		"newsid": 321,
	}
	notification := &Notification{
		Android: &AndroidNotification{
			Alert:     &androidAlert,
			Title:     "Send to Android",
			BuilderId: 1,
			Extras:    extra,
		},
		Ios: &IosNotification{
			Alert:  &androidAlert,
			Sound:  "default",
			Badge:  "+1",
			Extras: extra,
		},
	}
	message := NewMessage()
	message.Title = "msg"
	message.Extras = map[string]interface{}{
		"key": "value",
	}
	message.ContentType = "text"
	message.Content = androidAlert
	p := &Payload{
		Cid:          "8103a4c628a0b98974ec1949-711261d4-5f17-4d2f-a855-5e5a8909b26e",
		Platform:     platform,
		Audience:     audience,
		Notification: notification,
		Message:      message,
		SmsMessage: &SmsMessage{
			TempId: 1250,
			TempPara: map[string]interface{}{
				"code": "123456",
			},
			DelayTime: 3600,
		},
		Options: NewOptions(TimeLive(60), ApnsProduction(false), ApnsCollapseId("jiguang_test_201706011100")),
	}

	b, err := p.MarshalJSON()
	if err != nil {
		t.Errorf("marashalJSON err: %+v", err)
	}
	expect := `{"platform":"all","audience":{"tag":["深圳","北京"]},"notification":{"android":{"alert":"Hi, JPush!","title":"Send to Android","builder_id":1,"extras":{"newsid":321}},"ios":{"alert":"Hi, JPush!","sound":"default","badge":"+1","extras":{"newsid":321}}},"message":{"msg_content":"Hi, JPush!","title":"msg","content_type":"text","extras":{"key":"value"}},"sms_message":{"delay_time":3600,"temp_id":1250,"temp_para":{"code":"123456"}},"options":{"time_to_live":60,"apns_production":false,"apns_collapse_id":"jiguang_test_201706011100"},"cid":"8103a4c628a0b98974ec1949-711261d4-5f17-4d2f-a855-5e5a8909b26e"}`
	if !bytes.Equal(b, []byte(expect)) {
		t.Errorf("不相等，expect: %s, got: %s", b, expect)
		println(string(b))
	}
}

func TestNewOptions(t *testing.T) {
	ops := NewOptions(SendNo(1), TimeLive(12),
		OverrideMsgId(123),
		ApnsProduction(true),
		ApnsCollapseId("1234"),
		BigPushDuration(123456))
	if *ops.SendNo != 1 || *ops.TimeLive != 12 ||
		*ops.OverrideMsgId != 123 || *ops.ApnsProduction != true ||
		*ops.ApnsCollapseId != "1234" || *ops.BigPushDuration != 123456 {
			t.Error("NewOptions 失败")
	}

	if _, ok :=SendNo(1).Value().(int); !ok {
		t.Error("Option.Value 失败")
	}
}
