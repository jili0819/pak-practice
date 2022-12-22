package main

import "reflect"

// ApiInter 简单工厂模式
const (
	SpeakImplsName   = "SpeakImpls"
	WrangleImplsName = "WrangleImpls"
)

type (
	ApiInter interface {
		Say(string) string
		Name() string
	}
	SpeakImpls   struct{}
	WrangleImpls struct{}
)

func (s *SpeakImpls) Say(msg string) string {
	return msg
}

func (s *SpeakImpls) Name() string {
	return SpeakImplsName
}

func (a *WrangleImpls) Say(msg string) string {
	return msg
}

func (a *WrangleImpls) Name() string {
	return WrangleImplsName
}

func NewApiImpls(facType interface{}) ApiInter {
	name := reflect.TypeOf(facType).Name()
	switch name {
	case SpeakImplsName:
		return &SpeakImpls{}
	case WrangleImplsName:
		return &WrangleImpls{}
	}
	return &WrangleImpls{}
}

func main() {
	NewApiImpls(new(SpeakImpls))
}
