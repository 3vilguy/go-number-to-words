package convert

var UNITS = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var TEENS = []string{"", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var TENS = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func Number2Words(number int) string {
	// Quick 0 check
	if number == 0 {
		return UNITS[0]
	}

	var chunkIndex = 0
	var combined = ""

	for number > 0 {
		// Take 3-digit chunk of the number
		chunkOfNumber := number % 1000

		if chunkOfNumber > 0 {
			var hundreds = chunkOfNumber / 100 % 10
			var tens = chunkOfNumber / 10 % 10
			var units = chunkOfNumber % 10

			currentChunkToWord := ""

			// Combine with previous chunks ?
			if number > 1000 && hundreds == 0 && chunkIndex == 0 {
				currentChunkToWord = " and "
			} else {
				if number > 1000 {
					currentChunkToWord = ", "
				}
			}

			// Add Hundreds
			if hundreds > 0 {
				currentChunkToWord += UNITS[hundreds] + " hundred"
				if tens > 0 || units > 0 {
					currentChunkToWord += " and "
				}
			}

			// Add Tens
			if tens > 0 {
				if tens == 1 && units > 0 {
					currentChunkToWord += TEENS[units]
				} else {
					currentChunkToWord += TENS[tens]
					if units > 0 {
						currentChunkToWord += "-"
					}
				}
			}

			// Add Units
			if tens != 1 && units > 0 {
				currentChunkToWord += UNITS[units]
			}

			// Assign first chunk or combine with previously combined words
			if chunkIndex == 0 {
				combined = currentChunkToWord
			} else {
				combined = currentChunkToWord + " thousand" + combined
			}
		}

		number = number / 1000
		chunkIndex++
	}

	return combined
}
