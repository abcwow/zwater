package main

/////////////////////////////////

type CUP struct {
	id       int
	capacity int
	current  int
}

type CupsSetting struct {
	cups []CUP
}

var m_CupsSetting *CupsSetting

func init() {

}
