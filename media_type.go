package mediatype

import (
	"mime"
	"strings"
)

// Parse parses a string as [MediaType].
//
// If it is a valid media type, returns concrete MediaType value.
func Parse(s string) (*MediaType, error) {
	mt, params, err := mime.ParseMediaType(s)
	if err != nil {
		return nil, err
	}
	primary, sub, _ := strings.Cut(mt, "/")
	return &MediaType{
		Type:       primary,
		SubType:    SubType(sub),
		Parameters: params,
	}, nil
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

// Equal returns whether two media types are equal.
func (mt *MediaType) Equal(other *MediaType) bool {
	return mt.Type == other.Type && mt.SubType == other.SubType
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
