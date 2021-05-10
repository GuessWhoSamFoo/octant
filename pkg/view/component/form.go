/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package component

import (
	"fmt"
	"github.com/vmware-tanzu/octant/internal/util/json"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

type FormType string

const (
	FieldTypeCheckBox FormType = "checkbox"
	FieldTypeRadio    FormType = "radio"
	FieldTypeText     FormType = "text"
	FieldTypePassword FormType = "password"
	FieldTypeNumber   FormType = "number"
	FieldTypeSelect   FormType = "select"
	FieldTypeTextarea FormType = "textarea"
	FieldTypeHidden   FormType = "hidden"
	FieldTypeLayout   FormType = "layout"
)

type FormValidator string

const (
	FormValidatorMin           FormValidator = "min"
	FormValidatorMax           FormValidator = "max"
	FormValidatorRequired      FormValidator = "required"
	FormValidatorRequiredTrue  FormValidator = "requiredTrue"
	FormValidatorEmail         FormValidator = "email"
	FormValidatorMinLength     FormValidator = "minLength"
	FormValidatorMaxLength     FormValidator = "maxLength"
	FormValidatorPattern       FormValidator = "pattern"
	FormValidatorNullValidator FormValidator = "nullValidator"
)

type InputChoice struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Checked bool   `json:"checked"`
}

// FormFieldOptions provides additional configuration for form fields with multiple inputs
type FormFieldOptions struct {
	Choices  []InputChoice `json:"choices,omitempty"`
	Multiple bool          `json:"multiple,omitempty"`
	Fields   []FormField   `json:"fields,omitempty"`
	// Layout InputLayout `json:"layout,omitempty"`
	// vertical, vertical-inline, horizontal, horizontal-inline, compact
}

// FormField is a component for fields within a Form
// +octant:component
type FormField struct {
	Base
	Config FormFieldConfig `json:"config"`
}

var _ Component = (*FormField)(nil)

type formFieldMarshal FormField

type FormFieldConfig struct {
	Type          FormType                      `json:"type"`
	Label         string                        `json:"label"`
	Name          string                        `json:"name"`
	Value         interface{}                   `json:"value"`
	Configuration *FormFieldOptions             `json:"configuration,omitempty"`
	Placeholder   string                        `json:"placeholder,omitempty"`
	Error         string                        `json:"error,omitempty"`
	Validators    map[FormValidator]interface{} `json:"validators,omitempty"`
}

func (ffc *FormFieldConfig) UnmarshalJSON(data []byte) error {
	x := struct {
		Type          FormType                      `json:"type"`
		Label         string                        `json:"label"`
		Name          string                        `json:"name"`
		Value         interface{}                   `json:"value"`
		Configuration *FormFieldOptions             `json:"configuration,omitempty"`
		Placeholder   string                        `json:"placeholder,omitempty"`
		Error         string                        `json:"error,omitempty"`
		Validators    map[FormValidator]interface{} `json:"validators,omitempty"`
	}{}
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	ffc.Type = x.Type
	ffc.Label = x.Label
	ffc.Name = x.Name
	ffc.Value = x.Value
	if x.Configuration != nil {
		ffc.Configuration = x.Configuration
	}
	ffc.Placeholder = x.Placeholder
	ffc.Error = x.Error
	ffc.Validators = x.Validators
	return nil
}

func (ff *FormField) MarshalJSON() ([]byte, error) {
	m := formFieldMarshal(*ff)
	m.Metadata.Type = TypeFormField
	return json.Marshal(&m)
}

// NewFormFieldLayout creates a group of form fields
func NewFormFieldLayout(size RowSize) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeLayout,
			Value: size,
			Configuration: &FormFieldOptions{
				Fields: []FormField{},
			},
		},
	}
}

type RowSize string

const (
	RowSizeOne    RowSize = "grid cols@xs:1 gap:lg"
	RowSizeTwo    RowSize = "grid cols@xs:2 gap:lg"
	RowSizeThree  RowSize = "grid cols@xs:3 gap:lg"
	RowSizeFour   RowSize = "grid cols@xs:4 gap:lg"
	RowSizeFive   RowSize = "grid cols@xs:5 gap:lg"
	RowSizeSix    RowSize = "grid cols@xs:6 gap:lg"
	RowSizeSeven  RowSize = "grid cols@xs:7 gap:lg"
	RowSizeEight  RowSize = "grid cols@xs:8 gap:lg"
	RowSizeNine   RowSize = "grid cols@xs:9 gap:lg"
	RowSizeTen    RowSize = "grid cols@xs:10 gap:lg"
	RowSizeEleven RowSize = "grid cols@xs:11 gap:lg"
	RowSizeTwelve RowSize = "grid cols@xs:12 gap:lg"
)

func NewFormFieldCheckBox(label, name string, choices []InputChoice) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeCheckBox,
			Label: label,
			Name:  name,
			Configuration: &FormFieldOptions{
				Choices:  choices,
				Multiple: true,
			},
		},
	}
}

// AddValidator adds validator(s)
func (ff *FormField) AddValidator(errorMessage string, validators map[FormValidator]interface{}) {
	ff.Config.Error = errorMessage
	ff.Config.Validators = validators
}

func NewFormFieldRadio(label, name string, choices []InputChoice) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeRadio,
			Label: label,
			Name:  name,
			Configuration: &FormFieldOptions{
				Choices:  choices,
				Multiple: false,
			},
		},
	}
}

func NewFormFieldText(label, name, value string) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeText,
			Label: label,
			Name:  name,
			Value: value,
		},
	}
}

func NewFormFieldPassword(label, name, value string) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypePassword,
			Label: label,
			Name:  name,
			Value: value,
		},
	}
}

func NewFormFieldNumber(label, name, value string) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeNumber,
			Label: label,
			Name:  name,
			Value: value,
		},
	}
}

func NewFormFieldSelect(label, name string, choices []InputChoice, multiple bool) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeSelect,
			Label: label,
			Name:  name,
			Configuration: &FormFieldOptions{
				Choices:  choices,
				Multiple: multiple,
			},
		},
	}
}

func NewFormFieldTextarea(label, name, value string) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeTextarea,
			Label: label,
			Name:  name,
			Value: value,
		},
	}
}

func NewFormFieldHidden(name, value string) FormField {
	return FormField{
		Base: newBase(TypeFormField, nil),
		Config: FormFieldConfig{
			Type:  FieldTypeHidden,
			Name:  name,
			Value: value,
		},
	}
}

type Form struct {
	Fields []FormField `json:"fields"`
	Action string      `json:"action,omitempty"`
}

type formMarshal Form

func (f *Form) MarshalJSON() ([]byte, error) {
	m := formMarshal{
		Fields: f.Fields,
		Action: f.Action,
	}
	return json.Marshal(&m)
}

func (f *Form) UnmarshalJSON(data []byte) error {
	x := struct {
		Fields []TypedObject `json:"fields"`
		Action string        `json:"action,omitempty"`
	}{}

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	for _, typedObject := range x.Fields {
		component, err := typedObject.ToComponent()
		if err != nil {
			return err
		}

		field, ok := component.(*FormField)
		if !ok {
			fmt.Errorf("item was not a form field")
		}

		f.Fields = append(f.Fields, *field)
	}
	f.Action = x.Action
	return nil
}

// CreateFormForObject creates a form for an object with additional fields.
func CreateFormForObject(actionName string, object runtime.Object, fields ...FormField) (Form, error) {
	if object == nil {
		return Form{}, errors.New("object is nil")
	}

	apiVersion, kind := object.GetObjectKind().GroupVersionKind().ToAPIVersionAndKind()
	accessor, err := meta.Accessor(object)
	if err != nil {
		return Form{}, err
	}

	fields = append(fields,
		NewFormFieldHidden("apiVersion", apiVersion),
		NewFormFieldHidden("kind", kind),
		NewFormFieldHidden("name", accessor.GetName()),
		NewFormFieldHidden("namespace", accessor.GetNamespace()),
		NewFormFieldHidden("action", actionName),
	)

	return Form{Fields: fields}, nil
}
