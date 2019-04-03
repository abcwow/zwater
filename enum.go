package main

type EnumVar struct {
	to   CUP
	from CUP
}

type Enum interface {
	enum(before EnumVar) (after EnumVer, err error)
}

type EnumSetting struct {
	forms []Enums
}

var m_EnumSetting *EnumSetting

func init() {

}
