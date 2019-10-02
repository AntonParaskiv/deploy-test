package main

import "testing"

func Test_printMessage(t *testing.T) {
	tests := []struct {
		name        string
		wantMessage string
	}{
		{
			name:        "Success",
			wantMessage: Message,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMessage := printMessage(); gotMessage != tt.wantMessage {
				t.Errorf("printMessage() = %v, want %v", gotMessage, tt.wantMessage)
			}
		})
	}
}
