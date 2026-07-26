package storage

import (
	"testing"
)

func TestParseMessageType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    MessageType
		wantErr bool
	}{
		{
			name:  "valid user type",
			input: "user",
			want:  MessageTypeUser,
		},
		{
			name:  "valid llm type",
			input: "llm",
			want:  MessageTypeLLM,
		},
		{
			name:    "invalid type",
			input:   "invalid",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "uppercase user",
			input:   "USER",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMessageType(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMessageType(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseMessageType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
