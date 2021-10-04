package people

// Student 生徒インターフェース
type Student interface {
	GetName() string
}

// student 生徒構造体
type student struct {
	Name string
}

// NewStudent 生徒初期化
func NewStudent(name string) Student {
	return &student{Name: name}
}

func (s *student) GetName() string {
	return s.Name
}
