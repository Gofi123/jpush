package jpush

type Options struct {
	SendNo          *int    `json:"sendno,omitempty"`
	TimeLive        *int    `json:"time_to_live,omitempty"`
	OverrideMsgId   *int64  `json:"override_msg_id,omitempty"`
	ApnsProduction  *bool   `json:"apns_production"`
	ApnsCollapseId  *string `json:"apns_collapse_id,omitempty"`
	BigPushDuration *int    `json:"big_push_duration,omitempty"`
}

type OptionInterface interface {
	Value() interface{}
	IntValue() int
	Int64Value() int64
	StringValue() string
	BoolValue() bool
}

func SendNo(no int) *Option {
	return option("SendNo", no)
}
func TimeLive(no int) *Option {
	return option("TimeLive", no)
}
func OverrideMsgId(no int64) *Option {
	return option("OverrideMsgId", no)
}
func ApnsProduction(no bool) *Option {
	return option("ApnsProduction", no)
}
func ApnsCollapseId(no string) *Option {
	return option("ApnsCollapseId", no)
}
func BigPushDuration(no int) *Option {
	return option("BigPushDuration", no)
}

type Option struct {
	key   string
	value interface{}
}

func option(key string, value interface{}) *Option {
	return &Option{
		key:   key,
		value: value,
	}
}

func (p *Option) Value() interface{} {
	return p.value
}

func (p *Option) IntValue() int {
	if p.Value() != nil {
		if v, ok := p.Value().(int); ok {
			return v
		}
	}
	return 0
}

func (p *Option) Int64Value() int64 {
	if p.Value() != nil {
		if v, ok := p.Value().(int64); ok {
			return v
		}
	}
	return 0
}

func (p *Option) BoolValue() bool {
	if p.Value() != nil {
		if v, ok := p.Value().(bool); ok {
			return v
		}
	}
	return false
}

func (p *Option) StringValue() string {
	if p.Value() != nil {
		if v, ok := p.Value().(string); ok {
			return v
		}
	}
	return ""
}

func (p *Option) Key() string {
	return p.key
}

func NewOptions(options ...*Option) *Options {
	ops := &Options{}
	for _, op := range options {
		switch op.Key() {
		case "SendNo":
			tmp := op.IntValue()
			ops.SendNo = &tmp
		case "TimeLive":
			tmp := op.IntValue()
			ops.TimeLive = &tmp
		case "OverrideMsgId":
			tmp := op.Int64Value()
			ops.OverrideMsgId = &tmp
		case "ApnsProduction":
			tmp := op.BoolValue()
			ops.ApnsProduction = &tmp
		case "ApnsCollapseId":
			tmp := op.StringValue()
			ops.ApnsCollapseId = &tmp
		case "BigPushDuration":
			tmp := op.IntValue()
			ops.BigPushDuration = &tmp

		}
	}
	return ops
}
