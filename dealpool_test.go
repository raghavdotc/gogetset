package gogetset

import "testing"

func Test_isLastStep(t *testing.T) {
	type args struct {
		stepIndex int
		steps     []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isLastStep",
			args: args{
				stepIndex: 1,
				steps:     []string{"hello", "world"},
			},
			want: true,
		},
		{
			name: "isNotLastStep",
			args: args{
				stepIndex: 0,
				steps:     []string{"hello", "world"},
			},
			want: false,
		},
		{
			name: "isBeyondLastStep",
			args: args{
				stepIndex: 3,
				steps:     []string{"hello", "world"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLastStep(tt.args.stepIndex, tt.args.steps); got != tt.want {
				t.Errorf("isLastStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextStepIsArray(t *testing.T) {
	type args struct {
		step string
	}
	tests := []struct {
		name         string
		args         args
		wantArrayKey string
		wantIdx      int
		wantIsArray  bool
		wantErr      bool
	}{
		{
			args: args{
				"hello[10]",
			},
			wantArrayKey: "hello",
			wantIdx:      10,
			wantIsArray:  true,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArrayKey, gotIdx, gotIsArray, err := nextStepIsArray(tt.args.step)
			if (err != nil) != tt.wantErr {
				t.Errorf("stepIsArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotArrayKey != tt.wantArrayKey {
				t.Errorf("stepIsArray() gotArrayKey = %v, want %v", gotArrayKey, tt.wantArrayKey)
			}
			if gotIdx != tt.wantIdx {
				t.Errorf("stepIsArray() gotIdx = %v, want %v", gotIdx, tt.wantIdx)
			}
			if gotIsArray != tt.wantIsArray {
				t.Errorf("stepIsArray() gotIsArray = %v, want %v", gotIsArray, tt.wantIsArray)
			}
		})
	}
}
