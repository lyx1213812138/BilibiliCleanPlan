package data

type Vgroup interface {
	GetVideo() ([]Video, error)
	GetLabel() Label
	SetLabel(Label)
	GetStr() (string, error)
	GetUpID() int
}

type VgType int

const (
	IsSeason VgType = iota
	IsUp
)

type Label int

const (
	NoLab      Label = iota
	Ignore           // 1
	Normal           // 2
	Prefer           // 3
	VeryPrefer       // 4
)
