package src

type BookNumberGenerator interface {
	GenerateNumber(categoryId int) int
	IncrementNumber(categoryId int)
}


