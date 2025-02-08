package model

type DataHelper interface {
	Examine() error
}

func ExamineData(data DataHelper) error {
	return data.Examine()
}