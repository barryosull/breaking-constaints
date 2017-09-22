package tests

import (
	"github.com/barryosull/breaking-constaints/src"
	"github.com/barryosull/breaking-constaints/src/book_number_generator"
	"testing"
)

func makeBookNumberGenerator() src.BookNumberGenerator {

	return book_number_generator.Make()
}

func TestNumbersAreGeneratedSequentially(t *testing.T) {

	generator := makeBookNumberGenerator();

	categoryId := 1

	first := generator.GenerateNumber(categoryId)

	generator.IncrementNumber(categoryId)

	second := generator.GenerateNumber(categoryId)

	if (first != 1) {
		t.Error("Expected 1, got "+string(first)+", number not incremented correctly")
	}

	if (second != 2) {
		t.Error("Expected 2, got "+string(second)+", number not incremented correctly")
	}
}