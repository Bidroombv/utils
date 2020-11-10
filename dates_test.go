package utils

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestDateISO_String(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "It should be stringify",
			fields: fields{
				Time: time.Date(2020, 11, 11, 0, 0, 0, 0, time.Local),
			},
			want: "2020-11-11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DateISO{
				Time: tt.fields.Time,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testStruct struct {
	ID   string  `json:"id"`
	Date DateISO `json:"date"`
}

func TestDateISO_UnmarshalJSON(t *testing.T) {
	t.Run("It should be correctly unmarshaled", func(t *testing.T) {
		var res testStruct

		exampleJSON := `{
						"id": "lorem-ipsum",
						"date": "2020-11-11"
					}`

		err := json.Unmarshal([]byte(exampleJSON), &res)
		if err != nil {
			t.Error("Couldn't unmarshal JSON")
		}

		if res.Date.Day() != 11 || res.Date.Month() != 11 || res.Date.Year() != 2020 {
			t.Errorf("Wrong date, res: %s, expected 2020-11-11", res.Date)
		}
	})
}

func TestDateISO_MarshalJSON(t *testing.T) {
	t.Run("It should be correctly marshaled", func(t *testing.T) {
		given := testStruct{
			ID: "lorem-ipsum",
			Date: DateISO{
				Time: time.Date(2020, 11, 11, 0, 0, 0, 0, time.Local),
			},
		}

		expectedJson := `{
						"id": "lorem-ipsum",
						"date": "2020-11-11"
					}`

		resJson, err := json.Marshal(given)
		if err != nil {
			t.Errorf("Couldn't marshal JSON, err: %v", err)
		}

		r, err := jsonBytesEqual([]byte(expectedJson), resJson)
		if err != nil {
			t.Error("Couldn't unmarshal JSON")
		}

		if !r {
			t.Errorf("Wrong JSON, given: %v, expected: %v", string(resJson), expectedJson)
		}

	})
}

// JSONBytesEqual compares the JSON in two byte slices.
func jsonBytesEqual(a, b []byte) (bool, error) {
	var j, j2 testStruct
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
