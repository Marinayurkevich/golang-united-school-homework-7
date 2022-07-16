package coverage

import (
	"os"
	"reflect"
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
func TestPeopleLess(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
		want bool
	}{
		{
			name: "Person 1 > Person 2 by birthday (true)",
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
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
		{
			name: "Person 1 > Person 2 by birthday (false)",
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
			args: args{
				i: 0,
				j: 1,
			},
			want: false,
		},
		{
			name: "Person 1 > Person 2 by name (true)",
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
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
		{
			name: "Person 1 > Person 2 by name (false)",
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
			args: args{
				i: 0,
				j: 1,
			},
			want: false,
		},
		{
			name: "Person 1 > Person 2 by lastname (true)",
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
			args: args{
				i: 0,
				j: 1,
			},
			want: true,
		},
		{
			name: "Person 1 > Person 2 by lastname (false)",
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
			args: args{
				i: 0,
				j: 1,
			},
			want: false,
		},
		{
			name: "Person 1 == Person 2",
			p: People{
				{
					firstName: "John",
					lastName:  "Smith",
					birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
				}, {
					firstName: "John",
					lastName:  "Smith",
					birthDay:  time.Date(1988, time.March, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			args: args{
				i: 0,
				j: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleSwap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
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
			args: args{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p0, p1 := tt.p[0], tt.p[1]

			tt.p.Swap(tt.args.i, tt.args.j)

			if !((p0 == tt.p[1]) && (p1 == tt.p[0])) {
				t.Errorf("Fail to swap")
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			name: " ",
			args: args{
				str: "1 1 1\n2 2\n3",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: " ",
			args: args{
				str: "1 1 1\n2 2 2\n3 3 3",
			},
			want: &Matrix{
				rows: 3,
				cols: 3,
				data: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			},
			wantErr: false,
		},
		{
			name: " ",
			args: args{
				str: "1 1 1\n2 V 2\n3 3 3",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixRows(t *testing.T) {
	type fields struct {
		rows int
		cols int
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		// TODO: Add test cases.
		{
			name: " ",
			fields: fields{
				rows: 3,
				cols: 3,
				data: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			},
			want: [][]int{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			if got := m.Rows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixCols(t *testing.T) {
	type fields struct {
		rows int
		cols int
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		{
			name: " ",
			fields: fields{
				rows: 3,
				cols: 3,
				data: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			if got := m.Cols(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}
