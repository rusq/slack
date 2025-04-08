package slack

import (
	"reflect"
	"testing"
)

func attachmentIDPtr(s string) *AttachmentID {
	id := AttachmentID(s)
	return &id
}

func TestAttachmentID_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		a       *AttachmentID
		args    args
		want    *AttachmentID
		wantErr bool
	}{
		{
			name: "empty",
			a:    new(AttachmentID),
			args: args{
				data: []byte(""),
			},
			want:    new(AttachmentID),
			wantErr: false,
		},
		{
			name: "string",
			a:    new(AttachmentID),
			args: args{
				data: []byte(`"1234567890"`),
			},
			want:    attachmentIDPtr("1234567890"),
			wantErr: false,
		},
		{
			name: "integer",
			a:    new(AttachmentID),
			args: args{
				data: []byte(`1234567890`),
			},
			want:    attachmentIDPtr("1234567890$"),
			wantErr: false,
		},
		{
			name: "invalid",
			a:    new(AttachmentID),
			args: args{
				data: []byte(`{}`),
			},
			want:    new(AttachmentID),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AttachmentID.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.a, tt.want) {
				t.Errorf("AttachmentID.UnmarshalJSON() = %v, want %v", tt.a, tt.want)
			}
		})
	}
}

func TestAttachmentID_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       *AttachmentID
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			a:       new(AttachmentID),
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name:    "string",
			a:       attachmentIDPtr("1234567890"),
			want:    []byte(`"1234567890"`),
			wantErr: false,
		},
		{
			name:    "integer",
			a:       attachmentIDPtr("1234567890$"),
			want:    []byte(`1234567890`),
			wantErr: false,
		},
		{
			name:    "just $",
			a:       attachmentIDPtr("$"),
			want:    []byte(`"$"`),
			wantErr: false,
		},
		{
			name:    "a string ending with $",
			a:       attachmentIDPtr("have some $"),
			want:    []byte(`"have some $"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("AttachmentID.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttachmentID.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
