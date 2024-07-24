package options

import (
	"testing"
)

func TestLine(t *testing.T) {
	input := `overwrite=true
source="C:\Users\atvfo\AppData\Local\WSJT-X\wsjtx_log.adi"
dest="F:\filebackups"`

	tests := []struct {
		expectedCount        int
		expectedLiteralKey   string
		expectedLiteralValue string
	}{
		{0, "overwrite", "true"},
		{1, "source", `"C:\Users\atvfo\AppData\Local\WSJT-X\wsjtx_log.adi"`},
		{2, "dest", `"F:\filebackups"`},
	}

	o := NewFromString(input)

	for i, tt := range tests {
		currentOption := o.options[i]

		if currentOption.Key != tt.expectedLiteralKey {
			t.Fatalf("tests[%d] - Key wrong. expected=%q, got=%q",
				i, tt.expectedLiteralKey, currentOption.Key)
		}

		if currentOption.Value != tt.expectedLiteralValue {
			t.Fatalf("tests[%d] - Value wrong. expected=%q, got=%q",
				i, tt.expectedLiteralValue, currentOption.Value)
		}
	}
}

func TestGetValueByKey(t *testing.T) {
	input := `overwrite=true
source="C:\Users\atvfo\AppData\Local\WSJT-X\wsjtx_log.adi"
dest="F:\filebackups"`

	tests := []struct {
		expectedCount        int
		expectedLiteralKey   string
		expectedLiteralValue string
	}{
		{0, "overwrite", "true"},
		{1, "source", `"C:\Users\atvfo\AppData\Local\WSJT-X\wsjtx_log.adi"`},
		{2, "dest", `"F:\filebackups"`},
	}

	o := NewFromString(input)

	for i, tt := range tests {
		search := o.GetValueByKey(tt.expectedLiteralKey)

		for _, yy := range search {

			if yy != tt.expectedLiteralValue {
				t.Fatalf("tests[%d] - Value wrong. expected=%q, got=%q",
					i, tt.expectedLiteralValue, yy)
			}

		}
	}
}
