package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Nick",
			LastName:  "Dicker",
			Grades: []Grade{
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 11,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 100,
				},
				{
					Title: "Quiz 3",
					Type:  GradeQuiz,
					Score: 60,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Fuck",
			LastName:  "You",
			Grades: []Grade{
				{
					Title: "Text 1",
					Type:  GradeText,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 77,
				},
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 66,
				},
			},
		},
		{
			ID:        4,
			FirstName: "Lucss",
			LastName:  "Flank",
			Grades: []Grade{
				{
					Title: "Exam 1",
					Type:  GradeExam,
					Score: 22,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Text 2",
					Type:  GradeText,
					Score: 23,
				},
			},
		},
		{
			ID:        5,
			FirstName: "Jack",
			LastName:  "Sluofuck",
			Grades: []Grade{
				{
					Title: "Text 2",
					Type:  GradeText,
					Score: 45,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 56,
				},
				{
					Title: "Text 2",
					Type:  GradeText,
					Score: 67,
				},
			},
		},
	}
}
