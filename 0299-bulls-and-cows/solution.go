package bullsandcows

import (
	"fmt"
)

func getHint(secret string, guess string) string {
	bullSlice := []rune(secret)
	cowCount := map[rune]int{}

	for _, d := range secret {
		count := cowCount[d]
		cowCount[d] = count + 1
	}

	var bulls, cows int
	cowMap := map[rune]int{}
	for i, d := range guess {
		if bullSlice[i] == d {
			bulls++
			count := cowCount[d]
			cowCount[d] = count - 1
			continue
		}

		if _, exists := cowCount[d]; exists {
			cowMap[d] = cowMap[d] + 1
		}
	}

	for k, v := range cowMap {
		count := cowCount[k]

		if count < v {
			cows += count
		} else {
			cows += v
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func getHint1(secret, guess string) string {
	var bulls, cows int
	freq := [10]int{}
	gfreq := [10]int{}

	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
			continue
		}

		freq[secret[i]-'0'] = freq[secret[i]-'0'] + 1
		gfreq[guess[i]-'0'] = gfreq[guess[i]-'0'] + 1
	}

	for i := range freq {
		if freq[i] > gfreq[i] {
			cows = cows + gfreq[i]
		} else {
			cows = cows + freq[i]
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
