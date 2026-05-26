// Copyright (c) 2019 Ashley Jeffs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package gabs implements a wrapper around creating and parsing unknown or
// dynamic map structures resulting from JSON parsing.
package gabs

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

//------------------------------------------------------------------------------

var (
	// ErrOutOfBounds indicates an index was out of bounds.
	ErrOutOfBounds = errors.New("out of bounds")

	// ErrNotObjOrArray is returned when a target is not an object or array type
	// but needs to be for the intended operation.
	ErrNotObjOrArray = errors.New("not an object or array")

	// ErrNotObj is returned when a target is not an object but needs to be for
	// the intended operation.
	ErrNotObj = errors.New("not an object")

	// ErrInvalidQuery is returned when a seach query was not valid.
	ErrInvalidQuery = errors.New("invalid search query")

	// ErrNotArray is returned when a target is not an array but needs to be for
	// the intended operation.
	ErrNotArray = errors.New("not an array")

	// ErrPathCollision is returned when creating a path failed because an
	// element collided with an existing value.
	ErrPathCollision = errors.New("encountered value collision whilst building path")

	// ErrInvalidInputObj is returned when the input value was not a
	// map[string]interface{}.
	ErrInvalidInputObj = errors.New("invalid input object")

	// ErrInvalidInputText is returned when the input data could not be parsed.
	ErrInvalidInputText = errors.New("input text could not be parsed")

	// ErrNotFound is returned when a query leaf is not found.
	ErrNotFound = errors.New("field not found")

	// ErrInvalidPath is returned when the filepath was not valid.
	ErrInvalidPath = errors.New("invalid file path")

	// ErrInvalidBuffer is returned when the input buffer contained an invalid
	// JSON string.
	ErrInvalidBuffer = errors.New("input buffer contained invalid JSON")
)

var (
	r1 *strings.Replacer
	r2 *strings.Replacer
)

func init() {
	r1 = strings.NewReplacer("~1", "/", "~0", "~")
	r2 = strings.NewReplacer("~1", ".", "~0", "~")
}

//------------------------------------------------------------------------------

// JSONPointerToSlice parses a JSON pointer path
// (https://tools.ietf.org/html/rfc6901) and returns the path segments as a
// slice.
//
// Because the characters '~' (%x7E) and '/' (%x2F) have special meanings in
// gabs paths, '~' needs to be encoded as '~0' and '/' needs to be encoded as
// '~1' when these characters appear in a reference key.
func JSONPointerToSlice(path string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// DotPathToSlice returns a slice of path segments parsed out of a dot path.
//
// Because '.' (%x2E) is the segment separator, it must be encoded as '~1'
// if it appears in the reference key. Likewise, '~' (%x7E) must be encoded
// as '~0' since it is the escape character for encoding '.'.
func DotPathToSlice(path string) []string { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------------------------------

// Container references a specific element within a wrapped structure.
type Container struct {
	object interface{}
}

// Data returns the underlying value of the target element in the wrapped
// structure.
func (g *Container) Data() interface{} { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------------------------------

func (g *Container) searchStrict(allowWildcard bool, hierarchy ...string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Search attempts to find and return an object within the wrapped structure by
// following a provided hierarchy of field names to locate the target.
//
// If the search encounters an array then the next hierarchy field name must be
// either a an integer which is interpreted as the index of the target, or the
// character '*', in which case all elements are searched with the remaining
// search hierarchy and the results returned within an array.
func (g *Container) Search(hierarchy ...string) *Container { _ = "STUB: not implemented"; return nil }

// Path searches the wrapped structure following a path in dot notation,
// segments of this path are searched according to the same rules as Search.
//
// Because the characters '~' (%x7E) and '.' (%x2E) have special meanings in
// gabs paths, '~' needs to be encoded as '~0' and '.' needs to be encoded as
// '~1' when these characters appear in a reference key.
func (g *Container) Path(path string) *Container { _ = "STUB: not implemented"; return nil }

// JSONPointer parses a JSON pointer path (https://tools.ietf.org/html/rfc6901)
// and either returns a *gabs.Container containing the result or an error if the
// referenced item could not be found.
//
// Because the characters '~' (%x7E) and '/' (%x2F) have special meanings in
// gabs paths, '~' needs to be encoded as '~0' and '/' needs to be encoded as
// '~1' when these characters appear in a reference key.
func (g *Container) JSONPointer(path string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// S is a shorthand alias for Search.
func (g *Container) S(hierarchy ...string) *Container { _ = "STUB: not implemented"; return nil }

// Exists checks whether a field exists within the hierarchy.
func (g *Container) Exists(hierarchy ...string) bool { _ = "STUB: not implemented"; return false }

// ExistsP checks whether a dot notation path exists.
func (g *Container) ExistsP(path string) bool { _ = "STUB: not implemented"; return false }

// Index attempts to find and return an element within a JSON array by an index.
func (g *Container) Index(index int) *Container { _ = "STUB: not implemented"; return nil }

// Children returns a slice of all children of an array element. This also works
// for objects, however, the children returned for an object will be in a random
// order and you lose the names of the returned objects this way. If the
// underlying container value isn't an array or map nil is returned.
func (g *Container) Children() []*Container { _ = "STUB: not implemented"; return nil }

// ChildrenMap returns a map of all the children of an object element. IF the
// underlying value isn't a object then an empty map is returned.
func (g *Container) ChildrenMap() map[string]*Container { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------------------------------

// Set attempts to set the value of a field located by a hierarchy of field
// names. If the search encounters an array then the next hierarchy field name
// is interpreted as an integer index of an existing element, or the character
// '-', which indicates a new element appended to the end of the array.
//
// Any parts of the hierarchy that do not exist will be constructed as objects.
// This includes parts that could be interpreted as array indexes.
//
// Returns a container of the new value or an error.
func (g *Container) Set(value interface{}, hierarchy ...string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetP sets the value of a field at a path using dot notation, any parts
// of the path that do not exist will be constructed, and if a collision occurs
// with a non object type whilst iterating the path an error is returned.
func (g *Container) SetP(value interface{}, path string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetIndex attempts to set a value of an array element based on an index.
func (g *Container) SetIndex(value interface{}, index int) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetJSONPointer parses a JSON pointer path
// (https://tools.ietf.org/html/rfc6901) and sets the leaf to a value. Returns
// an error if the pointer could not be resolved due to missing fields.
func (g *Container) SetJSONPointer(value interface{}, path string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Object creates a new JSON object at a target path. Returns an error if the
// path contains a collision with a non object type.
func (g *Container) Object(hierarchy ...string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ObjectP creates a new JSON object at a target path using dot notation.
// Returns an error if the path contains a collision with a non object type.
func (g *Container) ObjectP(path string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ObjectI creates a new JSON object at an array index. Returns an error if the
// object is not an array or the index is out of bounds.
func (g *Container) ObjectI(index int) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Array creates a new JSON array at a path. Returns an error if the path
// contains a collision with a non object type.
func (g *Container) Array(hierarchy ...string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayP creates a new JSON array at a path using dot notation. Returns an
// error if the path contains a collision with a non object type.
func (g *Container) ArrayP(path string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayI creates a new JSON array within an array at an index. Returns an error
// if the element is not an array or the index is out of bounds.
func (g *Container) ArrayI(index int) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayOfSize creates a new JSON array of a particular size at a path. Returns
// an error if the path contains a collision with a non object type.
func (g *Container) ArrayOfSize(size int, hierarchy ...string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayOfSizeP creates a new JSON array of a particular size at a path using
// dot notation. Returns an error if the path contains a collision with a non
// object type.
func (g *Container) ArrayOfSizeP(size int, path string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayOfSizeI create a new JSON array of a particular size within an array at
// an index. Returns an error if the element is not an array or the index is out
// of bounds.
func (g *Container) ArrayOfSizeI(size, index int) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete an element at a path, an error is returned if the element does not
// exist or is not an object. In order to remove an array element please use
// ArrayRemove.
func (g *Container) Delete(hierarchy ...string) error { _ = "STUB: not implemented"; return nil }

// DeleteP deletes an element at a path using dot notation, an error is returned
// if the element does not exist.
func (g *Container) DeleteP(path string) error { _ = "STUB: not implemented"; return nil }

// MergeFn merges two objects using a provided function to resolve collisions.
//
// The collision function receives two interface{} arguments, destination (the
// original object) and source (the object being merged into the destination).
// Which ever value is returned becomes the new value in the destination object
// at the location of the collision.
func (g *Container) MergeFn(source *Container, collisionFn func(destination, source interface{}) interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// path doesn't exist. So set the value

// Merge a source object into an existing destination object. When a collision
// is found within the merged structures (both a source and destination object
// contain the same non-object keys) the result will be an array containing both
// values, where values that are already arrays will be expanded into the
// resulting array.
//
// It is possible to merge structures will different collision behaviours with
// MergeFn.
func (g *Container) Merge(source *Container) error { _ = "STUB: not implemented"; return nil }

//------------------------------------------------------------------------------

/*
Array modification/search - Keeping these options simple right now, no need for
anything more complicated since you can just cast to []interface{}, modify and
then reassign with Set.
*/

// ArrayAppend attempts to append a value onto a JSON array at a path. If the
// target is not a JSON array then it will be converted into one, with its
// original contents set to the first element of the array.
func (g *Container) ArrayAppend(value interface{}, hierarchy ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// ArrayAppendP attempts to append a value onto a JSON array at a path using dot
// notation. If the target is not a JSON array then it will be converted into
// one, with its original contents set to the first element of the array.
func (g *Container) ArrayAppendP(value interface{}, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// ArrayConcat attempts to append a value onto a JSON array at a path. If the
// target is not a JSON array then it will be converted into one, with its
// original contents set to the first element of the array.
//
// ArrayConcat differs from ArrayAppend in that it will expand a value type
// []interface{} during the append operation, resulting in concatenation of each
// element, rather than append as a single element of []interface{}.
func (g *Container) ArrayConcat(value interface{}, hierarchy ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the data exists, and it is not a slice of interface,
// append it as the first element of our new array.

// If the data exists, and it is a slice of interface,
// assign it to our variable.

// If we have been given a slice of interface, expand it when appending.

// ArrayConcatP attempts to append a value onto a JSON array at a path using dot
// notation. If the target is not a JSON array then it will be converted into one,
// with its original contents set to the first element of the array.
//
// ArrayConcatP differs from ArrayAppendP in that it will expand a value type
// []interface{} during the append operation, resulting in concatenation of each
// element, rather than append as a single element of []interface{}.
func (g *Container) ArrayConcatP(value interface{}, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// ArrayRemove attempts to remove an element identified by an index from a JSON
// array at a path.
func (g *Container) ArrayRemove(index int, hierarchy ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// ArrayRemoveP attempts to remove an element identified by an index from a JSON
// array at a path using dot notation.
func (g *Container) ArrayRemoveP(index int, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// ArrayElement attempts to access an element by an index from a JSON array at a
// path.
func (g *Container) ArrayElement(index int, hierarchy ...string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayElementP attempts to access an element by an index from a JSON array at
// a path using dot notation.
func (g *Container) ArrayElementP(index int, path string) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayCount counts the number of elements in a JSON array at a path.
func (g *Container) ArrayCount(hierarchy ...string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ArrayCountP counts the number of elements in a JSON array at a path using dot
// notation.
func (g *Container) ArrayCountP(path string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//------------------------------------------------------------------------------

func walkObject(path string, obj, flat map[string]interface{}, includeEmpty bool) {
	_ = "STUB: not implemented"
	return
}

func walkArray(path string, arr []interface{}, flat map[string]interface{}, includeEmpty bool) {
	_ = "STUB: not implemented"
	return
}

// Flatten a JSON array or object into an object of key/value pairs for each
// field, where the key is the full path of the structured field in dot path
// notation matching the spec for the method Path.
//
// E.g. the structure `{"foo":[{"bar":"1"},{"bar":"2"}]}` would flatten into the
// object: `{"foo.0.bar":"1","foo.1.bar":"2"}`. `{"foo": [{"bar":[]},{"bar":{}}]}`
// would flatten into the object `{}`
//
// Returns an error if the target is not a JSON object or array.
func (g *Container) Flatten() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// FlattenIncludeEmpty a JSON array or object into an object of key/value pairs
		// for each field, just as Flatten, but includes empty arrays and objects, where
		// the key is the full path of the structured field in dot path notation matching
		// the spec for the method Path.
		//
		// E.g. the structure `{"foo": [{"bar":[]},{"bar":{}}]}` would flatten into the
		// object: `{"foo.0.bar":[],"foo.1.bar":{}}`.
		//
		// Returns an error if the target is not a JSON object or array.
		nil
}

func (g *Container) FlattenIncludeEmpty() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Container) flatten(includeEmpty bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//------------------------------------------------------------------------------

// Bytes marshals an element to a JSON []byte blob.
func (g *Container) Bytes() []byte { _ = "STUB: not implemented"; return nil }

// BytesIndent marshals an element to a JSON []byte blob formatted with a prefix
// and indent string.
func (g *Container) BytesIndent(prefix, indent string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// String marshals an element to a JSON formatted string.
func (g *Container) String() string { _ = "STUB: not implemented"; return "" }

// StringIndent marshals an element to a JSON string formatted with a prefix and
// indent string.
func (g *Container) StringIndent(prefix, indent string) string {
	_ = "STUB: not implemented"
	return ""
}

// EncodeOpt is a functional option for the EncodeJSON method.
type EncodeOpt func(e *json.Encoder)

// EncodeOptHTMLEscape sets the encoder to escape the JSON for html.
func EncodeOptHTMLEscape(doEscape bool) EncodeOpt {
	_ = "STUB: not implemented"
	return *new(EncodeOpt)
}

// EncodeOptIndent sets the encoder to indent the JSON output.
func EncodeOptIndent(prefix, indent string) EncodeOpt {
	_ = "STUB: not implemented"
	return *new(EncodeOpt)
}

// EncodeJSON marshals an element to a JSON formatted []byte using a variant
// list of modifier functions for the encoder being used. Functions for
// modifying the output are prefixed with EncodeOpt, e.g. EncodeOptHTMLEscape.
func (g *Container) EncodeJSON(encodeOpts ...EncodeOpt) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Do not escape by default.

// New creates a new gabs JSON object.
func New() *Container { _ = "STUB: not implemented"; return nil }

// Wrap an already unmarshalled JSON object (or a new map[string]interface{})
// into a *Container.
func Wrap(root interface{}) *Container { _ = "STUB: not implemented"; return nil }

// ParseJSON unmarshals a JSON byte slice into a *Container.
func ParseJSON(sample []byte) (*Container, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseJSONDecoder applies a json.Decoder to a *Container.
func ParseJSONDecoder(decoder *json.Decoder) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseJSONFile reads a file and unmarshals the contents into a *Container.
func ParseJSONFile(path string) (*Container, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseJSONBuffer reads a buffer and unmarshals the contents into a *Container.
func ParseJSONBuffer(buffer io.Reader) (*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalJSON returns the JSON encoding of this container. This allows
// structs which contain Container instances to be marshaled using
// json.Marshal().
func (g *Container) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//------------------------------------------------------------------------------
