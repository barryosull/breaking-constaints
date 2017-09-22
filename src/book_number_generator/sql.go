package book_number_generator

type Sql struct {
	number int
}

func Make() *Sql {
	return &Sql{1}
}

func (s *Sql) GenerateNumber(categoryId int) int {
	return s.number;
}

func (s *Sql) IncrementNumber(categoryId int) {
	s.number = s.number +1;
}