package main

import (
	"fmt"

	people "github.com/sgash708/GO_Interface/People"
)

func main() {
	s := people.NewStudent("taro")
	studentName := s.GetName()

	t := people.NewTeacher("George")
	teacherName := t.GetName()

	fmt.Println(studentName, teacherName)
}
