package mediatype_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/aereal/mediatype"
)

func TestMediaType_marshaling(t *testing.T) {
	t.Parallel()

	repr := "text/svg+xml; charset=utf-8"
	mt, err := mediatype.Parse(repr)
	if err != nil {
		t.Fatal(err)
	}
	jv, err := json.Marshal(mt)
	if err != nil {
		t.Fatal(err)
	}
	wantQuoted := strconv.Quote(repr)
	if s := string(jv); s != wantQuoted {
		t.Errorf("marshaled text differs:\n\twant: %q\n\t got: %q", wantQuoted, s)
	}
}

func TestMediaType_unmarshaling(t *testing.T) {
	t.Parallel()

	got := new(mediatype.MediaType)
	if err := json.Unmarshal([]byte(`"text/svg+xml; charset=utf-8"`), got); err != nil {
		t.Fatal(err)
	}
	want := &mediatype.MediaType{
		Type:       "text",
		SubType:    "svg+xml",
		Parameters: map[string]string{"charset": "utf-8"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("unmarshal failed:\n\twant: %#v\n\t got: %#v", want, got)
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input   string
		want    *mediatype.MediaType
		wantErr error
	}{
		{
			input: "text/svg+",
			want: &mediatype.MediaType{
				Type:       "text",
				SubType:    mediatype.SubType("svg+"),
				Parameters: map[string]string{},
			},
			wantErr: nil,
		},
		{
			input:   "",
			want:    nil,
			wantErr: literalError("mime: no media type"),
		},
		{
			input: "text/plain",
			want: &mediatype.MediaType{
				Type:       "text",
				SubType:    mediatype.SubType("plain"),
				Parameters: map[string]string{},
			},
			wantErr: nil,
		},
		{
			input: "text/plain;charset=utf-8",
			want: &mediatype.MediaType{
				Type:    "text",
				SubType: mediatype.SubType("plain"),
				Parameters: map[string]string{
					"charset": "utf-8",
				},
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			got, gotErr := mediatype.Parse(tc.input)
			if !errors.Is(tc.wantErr, gotErr) {
				t.Errorf("error:\n\twant: type=%T msg=%s\n\t got: type=%T msg=%s", tc.wantErr, tc.wantErr, gotErr, gotErr)
			}
			if gotErr != nil {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MediaType:\n\twant: %#v\n\t got: %#v", tc.want, got)
			}
		})
	}
}

type literalError string

func (e literalError) Error() string { return string(e) }

func (e literalError) Is(other error) bool {
	return string(e) == other.Error()
}
