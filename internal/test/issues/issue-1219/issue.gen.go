// Package issue1219 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package issue1219

import (
	"encoding/json"
	"fmt"
)

// DefaultAdditional1 defines model for DefaultAdditional1.
type DefaultAdditional1 struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
}

// DefaultAdditional2 defines model for DefaultAdditional2.
type DefaultAdditional2 struct {
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeDefaultDefault defines model for MergeDefaultDefault.
type MergeDefaultDefault struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeDefaultWithAny defines model for MergeDefaultWithAny.
type MergeDefaultWithAny struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MergeDefaultWithString defines model for MergeDefaultWithString.
type MergeDefaultWithString struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeDefaultWithout defines model for MergeDefaultWithout.
type MergeDefaultWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithAnyDefault defines model for MergeWithAnyDefault.
type MergeWithAnyDefault struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MergeWithAnyWithAny defines model for MergeWithAnyWithAny.
type MergeWithAnyWithAny struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MergeWithAnyWithString defines model for MergeWithAnyWithString.
type MergeWithAnyWithString struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeWithAnyWithout defines model for MergeWithAnyWithout.
type MergeWithAnyWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithStringDefault defines model for MergeWithStringDefault.
type MergeWithStringDefault struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeWithStringWithAny defines model for MergeWithStringWithAny.
type MergeWithStringWithAny struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// MergeWithStringWithout defines model for MergeWithStringWithout.
type MergeWithStringWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutDefault defines model for MergeWithoutDefault.
type MergeWithoutDefault struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutWithAny defines model for MergeWithoutWithAny.
type MergeWithoutWithAny struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutWithString defines model for MergeWithoutWithString.
type MergeWithoutWithString struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// MergeWithoutWithout defines model for MergeWithoutWithout.
type MergeWithoutWithout struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// WithAnyAdditional1 defines model for WithAnyAdditional1.
type WithAnyAdditional1 struct {
	Field1               *int                   `json:"field1,omitempty"`
	Field2               *string                `json:"field2,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// WithAnyAdditional2 defines model for WithAnyAdditional2.
type WithAnyAdditional2 struct {
	FieldA               *int                   `json:"fieldA,omitempty"`
	FieldB               *string                `json:"fieldB,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// WithStringAdditional1 defines model for WithStringAdditional1.
type WithStringAdditional1 struct {
	Field1               *int              `json:"field1,omitempty"`
	Field2               *string           `json:"field2,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// WithStringAdditional2 defines model for WithStringAdditional2.
type WithStringAdditional2 struct {
	FieldA               *int              `json:"fieldA,omitempty"`
	FieldB               *string           `json:"fieldB,omitempty"`
	AdditionalProperties map[string]string `json:"-"`
}

// WithoutAdditional1 defines model for WithoutAdditional1.
type WithoutAdditional1 struct {
	Field1 *int    `json:"field1,omitempty"`
	Field2 *string `json:"field2,omitempty"`
}

// WithoutAdditional2 defines model for WithoutAdditional2.
type WithoutAdditional2 struct {
	FieldA *int    `json:"fieldA,omitempty"`
	FieldB *string `json:"fieldB,omitempty"`
}

// Getter for additional properties for MergeDefaultWithAny. Returns the specified
// element and whether it was found
func (a MergeDefaultWithAny) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeDefaultWithAny
func (a *MergeDefaultWithAny) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeDefaultWithAny to handle AdditionalProperties
func (a *MergeDefaultWithAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeDefaultWithAny to handle AdditionalProperties
func (a MergeDefaultWithAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeDefaultWithString. Returns the specified
// element and whether it was found
func (a MergeDefaultWithString) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeDefaultWithString
func (a *MergeDefaultWithString) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeDefaultWithString to handle AdditionalProperties
func (a *MergeDefaultWithString) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeDefaultWithString to handle AdditionalProperties
func (a MergeDefaultWithString) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithAnyDefault. Returns the specified
// element and whether it was found
func (a MergeWithAnyDefault) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithAnyDefault
func (a *MergeWithAnyDefault) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithAnyDefault to handle AdditionalProperties
func (a *MergeWithAnyDefault) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithAnyDefault to handle AdditionalProperties
func (a MergeWithAnyDefault) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithAnyWithAny. Returns the specified
// element and whether it was found
func (a MergeWithAnyWithAny) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithAnyWithAny
func (a *MergeWithAnyWithAny) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithAnyWithAny to handle AdditionalProperties
func (a *MergeWithAnyWithAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithAnyWithAny to handle AdditionalProperties
func (a MergeWithAnyWithAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithAnyWithString. Returns the specified
// element and whether it was found
func (a MergeWithAnyWithString) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithAnyWithString
func (a *MergeWithAnyWithString) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithAnyWithString to handle AdditionalProperties
func (a *MergeWithAnyWithString) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithAnyWithString to handle AdditionalProperties
func (a MergeWithAnyWithString) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithStringDefault. Returns the specified
// element and whether it was found
func (a MergeWithStringDefault) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithStringDefault
func (a *MergeWithStringDefault) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithStringDefault to handle AdditionalProperties
func (a *MergeWithStringDefault) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithStringDefault to handle AdditionalProperties
func (a MergeWithStringDefault) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for MergeWithStringWithAny. Returns the specified
// element and whether it was found
func (a MergeWithStringWithAny) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for MergeWithStringWithAny
func (a *MergeWithStringWithAny) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for MergeWithStringWithAny to handle AdditionalProperties
func (a *MergeWithStringWithAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for MergeWithStringWithAny to handle AdditionalProperties
func (a MergeWithStringWithAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithAnyAdditional1. Returns the specified
// element and whether it was found
func (a WithAnyAdditional1) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithAnyAdditional1
func (a *WithAnyAdditional1) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithAnyAdditional1 to handle AdditionalProperties
func (a *WithAnyAdditional1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithAnyAdditional1 to handle AdditionalProperties
func (a WithAnyAdditional1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithAnyAdditional2. Returns the specified
// element and whether it was found
func (a WithAnyAdditional2) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithAnyAdditional2
func (a *WithAnyAdditional2) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithAnyAdditional2 to handle AdditionalProperties
func (a *WithAnyAdditional2) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithAnyAdditional2 to handle AdditionalProperties
func (a WithAnyAdditional2) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithStringAdditional1. Returns the specified
// element and whether it was found
func (a WithStringAdditional1) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithStringAdditional1
func (a *WithStringAdditional1) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithStringAdditional1 to handle AdditionalProperties
func (a *WithStringAdditional1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["field1"]; found {
		err = json.Unmarshal(raw, &a.Field1)
		if err != nil {
			return fmt.Errorf("error reading 'field1': %w", err)
		}
		delete(object, "field1")
	}

	if raw, found := object["field2"]; found {
		err = json.Unmarshal(raw, &a.Field2)
		if err != nil {
			return fmt.Errorf("error reading 'field2': %w", err)
		}
		delete(object, "field2")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithStringAdditional1 to handle AdditionalProperties
func (a WithStringAdditional1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Field1 != nil {
		object["field1"], err = json.Marshal(a.Field1)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field1': %w", err)
		}
	}

	if a.Field2 != nil {
		object["field2"], err = json.Marshal(a.Field2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'field2': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for WithStringAdditional2. Returns the specified
// element and whether it was found
func (a WithStringAdditional2) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for WithStringAdditional2
func (a *WithStringAdditional2) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for WithStringAdditional2 to handle AdditionalProperties
func (a *WithStringAdditional2) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["fieldA"]; found {
		err = json.Unmarshal(raw, &a.FieldA)
		if err != nil {
			return fmt.Errorf("error reading 'fieldA': %w", err)
		}
		delete(object, "fieldA")
	}

	if raw, found := object["fieldB"]; found {
		err = json.Unmarshal(raw, &a.FieldB)
		if err != nil {
			return fmt.Errorf("error reading 'fieldB': %w", err)
		}
		delete(object, "fieldB")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for WithStringAdditional2 to handle AdditionalProperties
func (a WithStringAdditional2) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.FieldA != nil {
		object["fieldA"], err = json.Marshal(a.FieldA)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldA': %w", err)
		}
	}

	if a.FieldB != nil {
		object["fieldB"], err = json.Marshal(a.FieldB)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'fieldB': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
