package component

import (
	"testing"

	"github.com/vmware-tanzu/octant/internal/util/json"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware-tanzu/octant/internal/testutil"
)

//func TestFormFieldLayout_UnmarshalJSON(t *testing.T) {
//	textFieldOne := NewFormFieldText("Test Horizontal", "Test Horizontal", "")
//	textFieldTwo := NewFormFieldText("Cluster Name", "clusterName", "")
//	expected := NewFormFieldLayout(RowSizeSix, []FormField{textFieldOne, textFieldTwo})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormFieldLayout
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, &got)
//}

//func TestFormFieldCheckBox_UnmarshalJSON(t *testing.T) {
//	choices := []InputChoice{
//		{Label: "foo", Value: "foo", Checked: false},
//		{Label: "bar", Value: "bar", Checked: true},
//		{Label: "baz", Value: "baz", Checked: false},
//	}
//
//	expected := NewFormFieldCheckBox("label", "name", choices)
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//	assertFormFieldEqual(t, expected, got)
//}
//
//func TestFormFieldRadio_UnmarshalJSON(t *testing.T) {
//	choices := []InputChoice{
//		{Label: "foo", Value: "foo", Checked: false},
//		{Label: "bar", Value: "bar", Checked: true},
//		{Label: "baz", Value: "baz", Checked: false},
//	}
//
//	expected := NewFormFieldRadio("label", "name", choices)
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, got)
//}
//
//func TestFormFieldText_UnmarshalJSON(t *testing.T) {
//	expected := NewFormFieldText("label", "name", "text")
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, got)
//}

//
//func TestFormFieldPassword_UnmarshalJSON(t *testing.T) {
//	expected := NewFormFieldPassword("label", "name", "text")
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, got)
//}
//
//func TestFormFieldNumber_UnmarshalJSON(t *testing.T) {
//	expected := NewFormFieldNumber("label", "name", "999")
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, got)
//}
//
//func TestFormFieldSelect_UnmarshalJSON(t *testing.T) {
//	choices := []InputChoice{
//		{Label: "foo", Value: "foo", Checked: false},
//		{Label: "bar", Value: "bar", Checked: true},
//		{Label: "baz", Value: "baz", Checked: false},
//	}
//
//	expected := NewFormFieldSelect("label", "name", choices, true)
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, got)
//}
//
//func TestFormFieldTextarea_UnmarshalJSON(t *testing.T) {
//	expected := NewFormFieldTextarea("label", "name", "999")
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, got)
//}
//
//func TestFormFieldHidden_UnmarshalJSON(t *testing.T) {
//	expected := NewFormFieldHidden("label", "name")
//	expected.AddValidator("error", map[FormValidator]interface{}{})
//
//	data, err := json.Marshal(&expected)
//	require.NoError(t, err)
//
//	var got FormField
//
//	require.NoError(t, json.Unmarshal(data, &got))
//
//	assertFormFieldEqual(t, expected, got)
//}

//func TestFormField_MarshalJSON(t *testing.T) {
//	tests := []struct {
//		name         string
//		formField    FormField
//		expectedPath string
//	}{
//		{
//			name:         "text",
//			formField:    NewFormFieldText("label", "name", "value"),
//			expectedPath: "config_form_text.json",
//		},
//	}
//
//	for _, tc := range tests {
//		t.Run(tc.name, func(t *testing.T) {
//			actual, err := json.Marshal(tc.formField)
//			require.NoError(t, err)
//
//			assert.JSONEq(t, string(expected), string(actual))
//		})
//	}
//}

//TODO: TestFormField_AddValidators
//TODO: TestForm_Marshal

func TestForm_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		formField FormField
	}{
		{
			name:      "text field",
			formField: NewFormFieldText("label", "name", "value"),
		},
		{
			name: "check box field",
			formField: NewFormFieldCheckBox("label", "name", []InputChoice{
				{
					Label:   "foo",
					Value:   "foo",
					Checked: true,
				},
				{
					Label:   "bar",
					Value:   "bar",
					Checked: false,
				},
			}),
		},
		{
			name: "radio field",
			formField: NewFormFieldRadio("label", "name", []InputChoice{
				{
					Label:   "foo",
					Value:   "foo",
					Checked: true,
				},
			}),
		},
		{
			name: "select field",
			formField: NewFormFieldSelect("label", "name", []InputChoice{
				{
					Label: "baz",
					Value: "baz",
				},
			}, true),
		},
		{
			name:      "password field",
			formField: NewFormFieldPassword("label", "name", "value"),
		},
		{
			name:      "number field",
			formField: NewFormFieldNumber("label", "name", "7"),
		},
		{
			name:      "text area field",
			formField: NewFormFieldTextarea("label", "name", "7"),
		},
		{
			name:      "hidden field",
			formField: NewFormFieldHidden("name", "7"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			form := Form{
				Fields: []FormField{test.formField},
			}

			data, err := json.Marshal(&form)
			require.NoError(t, err)

			var got Form
			require.NoError(t, json.Unmarshal(data, &got))

			assert.Equal(t, form, got)

		})
	}
}

func TestCreateFormForObject(t *testing.T) {
	object := testutil.CreatePod("pod")
	got, err := CreateFormForObject("action", object,
		NewFormFieldNumber("number", "name", "0"))
	require.NoError(t, err)

	expected := Form{
		Fields: []FormField{
			NewFormFieldNumber("number", "name", "0"),
			NewFormFieldHidden("apiVersion", object.APIVersion),
			NewFormFieldHidden("kind", object.Kind),
			NewFormFieldHidden("name", object.Name),
			NewFormFieldHidden("namespace", object.Namespace),
			NewFormFieldHidden("action", "action"),
		},
	}
	require.Equal(t, expected, got)
}

//func assertFormFieldEqual(t *testing.T, expected, got FormField) {
//	assert.Equal(t, expected.Config.Value, got.Config.Value)
//	assert.Equal(t, expected.Config.Name, got.Config.Name)
//	assert.Equal(t, expected.Config.Type, got.Config.Type)
//	assert.Equal(t, expected.Config.Configuration, got.Config.Configuration)
//}
