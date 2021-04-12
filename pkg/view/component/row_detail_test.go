package component

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware-tanzu/octant/internal/util/json"
)

func TestTableRow_AddExpandableDetail(t *testing.T) {
	cases := []struct {
		name     string
		row      TableRow
		details  *ExpandableRowDetail
		expected *ExpandableRowDetail
	}{
		{
			name: "one detail",
			row: TableRow{
				"abc": NewText("124"),
			},
			details:  NewExpandableRowDetail(NewText("detail")),
			expected: NewExpandableRowDetail(NewText("detail")),
		},
		{
			name: "multiple details",
			row: TableRow{
				"abc": NewText("124"),
			},
			details: NewExpandableRowDetail([]Component{
				NewText("detail"),
				NewText("detail 2"),
			}...),
			expected: NewExpandableRowDetail(NewText("detail"), NewText("detail 2")),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.row.AddExpandableDetail(tc.details)
			require.Equal(t, tc.expected, tc.row[ExpandableRowKey])
		})
	}
}

func TestTableTow_ExpandableDetail_Marshal(t *testing.T) {
	cases := []struct {
		name         string
		input        *ExpandableRowDetail
		expectedPath string
		isErr        bool
	}{
		{
			name: "in general",
			input: &ExpandableRowDetail{
				Base: newBase(TypeExpandableRowDetail, nil),
				Config: ExpandableDetailConfig{
					Body: []Component{
						NewText("test"),
					},
				},
			},
			expectedPath: "expandable_row.json",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := json.Marshal(tc.input)
			isErr := err != nil
			if isErr != tc.isErr {
				t.Fatalf("Unexpected error: %v", err)
			}
			expected, err := ioutil.ReadFile(path.Join("testdata", tc.expectedPath))
			require.NoError(t, err)
			assert.JSONEq(t, string(expected), string(actual))
		})
	}
}
