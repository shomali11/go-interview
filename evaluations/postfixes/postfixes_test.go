package postfixes

import "testing"

func TestPostfix(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			"Test valid postfix",
			args{"2 3 +"},
			5.0,
			false,
		},
		{
			"Test valid postfix with plenty of whitespace",
			args{" 2  3  + "},
			5.0,
			false,
		},
		{
			"Test invalid empty postfix",
			args{""},
			0,
			true,
		},
		{
			"Test invalid postfix of one number",
			args{"3 +"},
			0,
			true,
		},
		{
			"Test invalid postfix of too many numbers",
			args{"1 2 3 +"},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
