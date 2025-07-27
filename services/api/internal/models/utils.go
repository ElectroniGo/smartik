package models

func GetAllModels() []interface{} {
	return []interface{}{
		&Student{},
		&Subject{},
		&Exam{},
		&AnswerScript{},
	}
}
