package people

import "github.com/sgash708/GO_Interface/implements"

// student 生徒構造体
type student struct {
	Name string
}

// NewStudent 生徒初期化
func NewStudent(name string) implements.Features {
	return &student{Name: name}
}

// GetName 名前取得
func (s *student) GetName() string {
	return s.Name
}
