package goutils

import "testing"

func TestGetLocalNetworkInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetLocalNetworkInfo()
		})
	}
}
