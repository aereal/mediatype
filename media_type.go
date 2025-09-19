package mediatype

import (
	"encoding"
	"fmt"
	"mime"
	"strings"
)

// Parse parses a string as [MediaType].
//
// If it is a valid media type, returns concrete MediaType value.
func Parse(s string) (*MediaType, error) {
	mt := new(MediaType)
	if err := unmarshal(mt, s); err != nil {
		return nil, err
	}
	return mt, nil
}

func unmarshal(mt *MediaType, s string) error {
	parsed, params, err := mime.ParseMediaType(s)
	if err != nil {
		return err
	}
	primary, sub, _ := strings.Cut(parsed, "/")
	*mt = MediaType{
		Type:       primary,
		SubType:    SubType(sub),
		Parameters: params,
	}
	return nil
}

// MediaType is an identifier of resource types as described in [RFC 2046].
//
// [RFC 2046]: https://datatracker.ietf.org/doc/html/rfc2046
type MediaType struct {
	// Type is also called the top-level type, the part before the `/`.
	//
	// refs. https://datatracker.ietf.org/doc/html/rfc2046#section-2
	Type string
	// SubType is the part after the `/`.
	SubType SubType
	// Parameters are modifiers of the media subtype.
	//
	// refs. https://datatracker.ietf.org/doc/html/rfc2046#section-1
	Parameters map[string]string
}

var (
	_ fmt.Stringer             = (*MediaType)(nil)
	_ encoding.TextMarshaler   = (*MediaType)(nil)
	_ encoding.TextUnmarshaler = (*MediaType)(nil)
)

// Equal returns whether two media types are equal.
func (mt *MediaType) Equal(other *MediaType) bool {
	return mt.Type == other.Type && mt.SubType == other.SubType
}

func (mt *MediaType) String() string {
	return mime.FormatMediaType(mt.Type+"/"+mt.SubType.String(), mt.Parameters)
}

func (mt *MediaType) MarshalText() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (mt *MediaType) UnmarshalText(b []byte) error {
	return unmarshal(mt, string(b))
}

// SubType is a subtype part of the [MediaType].
type SubType string

// Base returns a base subtype name of the subtype.
//
// refs. https://datatracker.ietf.org/doc/html/rfc6838#section-4.2.8
func (st SubType) Base() string {
	base, _, _ := strings.Cut(string(st), "+")
	return base
}

// Suffix returns a structured syntax name suffix of the subtype and its existence boolean.
//
// If the subtype has no structured syntax name suffix, returns empty string and false value.
//
// refs. https://datatracker.ietf.org/doc/html/rfc6838#section-4.2.8
func (st SubType) Suffix() (string, bool) {
	_, suffix, hasSuffix := strings.Cut(string(st), "+")
	return suffix, hasSuffix && suffix != ""
}

func (st SubType) String() string { return string(st) }
