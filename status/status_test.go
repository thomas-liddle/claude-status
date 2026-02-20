package status

import "testing"

func TestSetSlackStatusUntilEOD(t *testing.T) {
	type args struct {
		emoji string
		text  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"simple", args{"best-album", "TØP"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSlackStatusUntilEOD(tt.args.emoji, tt.args.text)
		})
	}
}
