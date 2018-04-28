package steam

import "testing"

func TestFindGameName(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Doesn't Exist", args{0}, ""},
		{"Exist (Duck Game)", args{312530}, "Duck Game"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGameName(tt.args.id); got != tt.want {
				t.Errorf("FindGameName() = %v, want %v", got, tt.want)
			}
		})
	}
}
