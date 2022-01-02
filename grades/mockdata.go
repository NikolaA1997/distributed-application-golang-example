package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nikola",
			LastName:  "Andrić",
			Grades: []Grade{
				{
					Title: "Kolokvij 1",
					Type:  GradeTest,
					Score: 59,
				},
				{
					Title: "Kviz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				{
					Title: "Kolokvij 2",
					Type:  GradeTest,
					Score: 76,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Marija",
			LastName:  "Šunjić",
			Grades: []Grade{
				{
					Title: "Kolokvij 1",
					Type:  GradeTest,
					Score: 87,
				},
				{
					Title: "Kviz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				{
					Title: "Kolokvij 2",
					Type:  GradeTest,
					Score: 78,
				},
			},
		},
	}
}
