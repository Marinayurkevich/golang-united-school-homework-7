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

func TestPeople_Len(t *testing.T) {
	tests := []struct {
		name string
		p    People
		want int
	}{
		{
			name: "2 people",
			p: People{
				{
					firstName: "John",
					lastName:  "Smith",
					birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
				}, {
					firstName: "Roy",
					lastName:  "Brown",
					birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeople_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
	}{
		// TODO: Add test cases.
		{
			name: "2 people",
			p: People{
				{
					firstName: "John",
					lastName:  "Smith",
					birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
				}, {
					firstName: "Roy",
					lastName:  "Brown",
					birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			args: args{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}
