package people

import "github.com/sgash708/GO_Interface/implements"

// teacher 先生構造体
type teacher struct {
	Name string
}

// NewTeacher 先生初期化
func NewTeacher(name string) implements.Features {
	return &teacher{Name: name}
}

// GetName 名前取得
func (t *teacher) GetName() string {
	return t.Name
}
