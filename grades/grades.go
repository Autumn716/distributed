package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}
	return result / float32(len(s.Grades))
}

type Students []Student

var (
	students     Students
	studentMutex sync.Mutex //增加互斥锁，保证Web并发访问的安全
)

func (ss Students) GetByID(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}

	return nil, fmt.Errorf("Student With ID %d not found", id)
}

type GradeType string

// 定义考试类型
const (
	GradeQuiz = GradeType("Quiz")
	GradeText = GradeType("Text")
	GradeExam = GradeType("Exam")
)

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
