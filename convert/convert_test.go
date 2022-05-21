package convert

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		input    int
		expected string
	}{
		// Checking 0
		{0, "zero"},
		// Check single digit
		{1, "one"},
		{9, "nine"},
		// Check teens
		{11, "eleven"},
		{19, "nineteen"},
		// Check tens
		{10, "ten"},
		{90, "ninety"},
		// Check teens with "-""
		{21, "twenty-one"},
		{99, "ninety-nine"},
		// Check hundreds
		{100, "one hundred"},
		{900, "nine hundred"},
		// Check hundreds with "and"
		{101, "one hundred and one"},
		{111, "one hundred and eleven"},
		{120, "one hundred and twenty"},
		// Check hundreds with "and" and "-"
		{121, "one hundred and twenty-one"},
		{999, "nine hundred and ninety-nine"},
		// Check thousands
		{1000, "one thousand"},
		{9000, "nine thousand"},
		// Check thousands
		{10000, "ten thousand"},
		{90000, "ninety thousand"},
		// Check max value
		{100000, "one hundred thousand"},
		// Check combined chunks with "and"
		{1001, "one thousand and one"},
		{1011, "one thousand and eleven"},
		{10001, "ten thousand and one"},
		{10099, "ten thousand and ninety-nine"},
		// Check combined chunks with ","
		{1100, "one thousand, one hundred"},
		{9999, "nine thousand, nine hundred and ninety-nine"},
		{10100, "ten thousand, one hundred"},
		// Check all together
		{55555, "fifty-five thousand, five hundred and fifty-five"},

		// Finally, make sure, the one from example are passing
		{1000, "one thousand"},
		{101, "one hundred and one"},
		{352, "three hundred and fifty-two"},
		{12345, "twelve thousand, three hundred and forty-five"},
	}

	for _, table := range tables {
		result := Number2Words(table.input)
		if result != table.expected {
			t.Errorf("%d conversion was incorrect, got: '%s', expected: '%s'.", table.input, result, table.expected)
		}
	}
}
