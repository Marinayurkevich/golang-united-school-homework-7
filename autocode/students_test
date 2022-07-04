package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeopleLen(t *testing.T) {
	var p = People{Person{
		firstName: "John",
		lastName:  "Smith",
		birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
	}, Person{
		firstName: "Roy",
		lastName:  "Brown",
		birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
	}}
	if len(p) != p.Len() {
		t.Errorf("len(p) is not correct: %d, got %d", p.Len(), len(p))
	}

}
