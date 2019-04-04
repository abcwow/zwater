package main

import (
	"fmt"
)

type EnumVar struct {
	to   CUP
	from CUP
}

type Enum interface {
	enum(before EnumVar) (after EnumVar, err error)
}

type EnumSetting struct {
	forms []Enum
}

/////////////////////////////

type FullSelf struct {
}

func (m *FullSelf) enum(before EnumVar) (after EnumVar, err error) {
	if before.to.id != before.from.id {
		return after, fmt.Errorf("op1 should on same cup")
	}

	if before.to.current == before.to.capacity {
		return after, fmt.Errorf("op1 should on a cup not full")
	}

	after = before
	after.to.current = after.to.capacity

	return after, nil

}

/////////////////////////////

type EmptySelf struct {
}

func (m *EmptySelf) enum(before EnumVar) (after EnumVar, err error) {

	if before.to.id != before.from.id {
		return after, fmt.Errorf("op2 should on same cup")
	}

	if before.to.current == 0 {
		return after, fmt.Errorf("op1 should on a cup not empty")
	}

	after = before
	after.to.current = 0

	return after, nil
}

/////////////////////////////

type FullOther struct {
}

func (m *FullOther) enum(before EnumVar) (after EnumVar, err error) {

	if before.to.id == before.from.id {
		return after, fmt.Errorf("op3 should on two cups")
	}

	if before.from.current+before.to.current <= before.to.capacity {
		return after, fmt.Errorf("op3 should make cup full")
	}

	after = before
	after.to.current = after.to.capacity
	after.from.current = after.from.current + after.to.current - before.to.capacity

	return after, nil
}

/////////////////////////////

type ToOtherSelfEmpty struct {
}

func (m *ToOtherSelfEmpty) enum(before EnumVar) (after EnumVar, err error) {

	if before.to.id == before.from.id {
		return after, fmt.Errorf("op4 should on two cups")
	}

	if before.to.current+before.from.current > before.to.capacity {

		return after, fmt.Errorf("op4 should make cup b full & cup a empty")
	}

	after = before
	after.from.current = 0
	if after.to.current+after.from.current > after.to.capacity {
		after.to.current = after.to.capacity
	} else {
		after.to.current = after.to.current + after.from.current
	}

	return after, nil

}

/////////////////////////////
var m_EnumSetting EnumSetting

func init() {

	m_EnumSetting.forms = append(m_EnumSetting.forms, new(FullSelf))
	m_EnumSetting.forms = append(m_EnumSetting.forms, new(EmptySelf))
	m_EnumSetting.forms = append(m_EnumSetting.forms, new(FullOther))
	m_EnumSetting.forms = append(m_EnumSetting.forms, new(ToOtherSelfEmpty))

}
