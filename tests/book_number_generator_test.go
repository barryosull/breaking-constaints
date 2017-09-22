package tests

import (
	"github.com/barryosull/breaking-constaints/src"
	"github.com/barryosull/breaking-constaints/src/book_number_generator"
	"testing"
)

func makeBookNumberGenerator() src.BookNumberGenerator {

	generator := book_number_generator.Make()
	generator.Reset()
	return generator
}

func TestNumbersAreGeneratedSequentially(t *testing.T) {

	generator := makeBookNumberGenerator();

	categoryId := 1

	first := generator.GenerateNumber(categoryId)

	err := generator.NumberUsed(categoryId, first)

	if (err != nil) {
		t.Error("Got error")
		t.Error(err)
	}

	second := generator.GenerateNumber(categoryId)

	if (first != 1) {
		t.Error("Expected 1, got "+string(first)+", number not incremented correctly")
	}

	if (second != 2) {
		t.Error("Expected 2, got "+string(second)+", number not incremented correctly")
	}
}

func TestCannotHaveDuplicateNumersInCategory(t *testing.T) {

	generator := makeBookNumberGenerator();

	categoryId := 1

	number := generator.GenerateNumber(categoryId)

	err := generator.NumberUsed(categoryId, number)

	if (err != nil) {
		t.Error("Got error")
		t.Error(err)
	}

	err = generator.NumberUsed(categoryId, number)

	if (err == nil) {
		t.Error("Expected error, not nothing")
		t.Error(err)
	}
}