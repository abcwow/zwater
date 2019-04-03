package main

type JudgeItem struct {
}

//////////////////////////

type JudgeItem struct {
	enumidx int //current op idx
	op      OP  //prev op info
}

func (d *JudgeItem) ToString() string {

	opident := op.Identity()

	str := fmt.Sprintf("op%d", d.enumidx) + "_" + opident

	return str

}

/////////////////////////////////////

const (
	REVERSE int = 0
	NEXT    int = 1
)

type JudgeTable struct {
	ops map[string]bool
}

func NewJudgeTalbe() *JudgeTable {

}

func (m *JudgeTable) Add(item *JudgeItem) {

}

func (m *JudgeTable) Find(item *JudgeItem) bool {

}

func (m *JudgeTable) Judge(cur, prev OP) int {
	item := JudgeItem{cur.enumidx, prev}

	if m.Find(&item) == true {
		return REVERSE
	}

	m.Add(&item)

	return NEXT

}

var m_JudgeTable *JudgeTable

func init() {

}
