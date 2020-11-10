package utils

import (
	"encoding/json"
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

func TestDateISO_UnmarshalJSON(t *testing.T) {
	t.Run("It should be correctly unmarshaled", func(t *testing.T) {
		type testStruct struct {
			ID   string  `json:"id"`
			Date DateISO `json:"date"`
		}

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
