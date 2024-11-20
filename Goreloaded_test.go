package main

import "testing"

func TestGoreloaded(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{{
		input: "1E (hex) files were added",
		want:  "30 files were added",
	}, {
		input: "It has been 10 (bin) years",
		want:  "It has been 2 years",
	}, {
		input: "Ready, set, go (up) !",
		want:  "Ready, set, GO!",
	}, {
		input: "I should stop SHOUTING (low)",
		want:  "I should stop shouting",
	}, {
		input: "Welcome to the Brooklyn bridge (cap)",
		want:  "Welcome to the Brooklyn Bridge",
	}, {
		input: "This is so exciting (up, 2)",
		want:  "This is SO EXCITING",
	}, {
		input: "I was sitting over there ,and then BAMM !!",
		want:  "I was sitting over there, and then BAMM!!",
	}, {
		input: "I was thinking ... You were right",
		want:  "I was thinking... You were right",
	}, {
		input: "I am exactly how they describe me: ' awesome '",
		want:  "I am exactly how they describe me: 'awesome'",
	}, {
		input: "As Elton John said: ' I am the most well-known homosexual in the world '",
		want:  "As Elton John said: 'I am the most well-known homosexual in the world'",
	}, {
		input: "There it was. A amazing rock!",
		want:  "There it was. An amazing rock!",
	}, {
		input: "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
		want:  "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
	}, {
		input: "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
		want:  "I have to pack 5 outfits. Packed 26 just to be sure",
	}, {
		input: "Don not be sad ,because sad backwards is das . And das not good",
		want:  "Don not be sad, because sad backwards is das. And das not good",
	}, {
		input: "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
		want:  "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
	},
	}

	// Iterate through each test case
	for _, tc := range testCases {

		// Run the function with the test input
		result := Goreloaded(tc.input)

		// Compare the result with the expected output
		if result != tc.want {
			t.Errorf("Input: %s\n%s\n%s (Expected output)", tc.input, result, tc.want)
		}

	}
}
