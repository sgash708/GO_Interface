package main

import (
	"fmt"

	people "github.com/sgash708/GO_Interface/People"
)

func main() {
	s := people.NewStudent("taro")
	studentName := s.GetName()

	fmt.Println(studentName)
}
