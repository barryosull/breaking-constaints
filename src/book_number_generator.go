package src

type BookNumberGenerator interface {
	GenerateNumber(categoryId int) int
	NumberUsed(categoryId int, number int) error
	Reset()
}

type BookNumberNotUnique struct {

}

func (e *BookNumberNotUnique) Error() string {
	return "Book number not unique in category"
}


