package analyze

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/ItakawaM/go-cryptotool/ciphers"
)

const (
	chiSquaredThreshold        float64 = 200.0
	englishDictionaryThreshold float64 = 0.5
)

type letterFrequency struct {
	letter    byte
	frequency float64
}

//go:embed english_words.txt
var englishWordsFile []byte

var englishFrequencies = [26]letterFrequency{
	{'a', 8.167},
	{'b', 1.492},
	{'c', 2.782},
	{'d', 4.253},
	{'e', 12.702},
	{'f', 2.228},
	{'g', 2.015},
	{'h', 6.094},
	{'i', 6.966},
	{'j', 0.153},
	{'k', 0.772},
	{'l', 4.025},
	{'m', 2.406},
	{'n', 6.749},
	{'o', 7.507},
	{'p', 1.929},
	{'q', 0.095},
	{'r', 5.987},
	{'s', 6.327},
	{'t', 9.056},
	{'u', 2.758},
	{'v', 0.978},
	{'w', 2.360},
	{'x', 0.150},
	{'y', 1.974},
	{'z', 0.074},
}

type AnalysisResult struct {
	Key          byte
	ChiScore     float64
	EnglishScore float64
}

func (ar AnalysisResult) String() string {
	return fmt.Sprintf("[%02d]: %.3f | %.3f", ar.Key, ar.ChiScore, ar.EnglishScore)
}

func AnalyzeCaesarFile(inputFilepath string) ([]AnalysisResult, error) {
	inFile, err := os.Open(inputFilepath)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	buffer := make([]byte, 16*1024) // Read 16KB or less
	n, err := io.ReadFull(inFile, buffer)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	buffer = buffer[:n]

	return AnalyzeCaesarBuffer(buffer)
}

func AnalyzeCaesarBuffer(buffer []byte) ([]AnalysisResult, error) {
	var results []AnalysisResult

	dst := make([]byte, len(buffer))
	frequencies := calculateLetterFrequency(buffer)

	sortedFrequencies := frequencies
	sort.Slice(sortedFrequencies[:], func(i, j int) bool {
		return sortedFrequencies[i].frequency > sortedFrequencies[j].frequency
	})

	englishDictionary, err := loadDictionary()
	if err != nil {
		return nil, err
	}

	englishMax := byte('e')
	for _, candidate := range sortedFrequencies {
		key := (candidate.letter - englishMax + 26) % 26

		caesarCipher, err := ciphers.NewCaesarCipher(int(key))
		if err != nil {
			return nil, err
		}
		caesarCipher.DecryptBlock(dst, buffer)

		newFrequencies := calculateLetterFrequency(dst)
		decryptedScore := calculateChiSquared(newFrequencies)

		englishScore := calculateEnglish(dst, englishDictionary)

		results = append(results, AnalysisResult{
			Key:          key,
			ChiScore:     decryptedScore,
			EnglishScore: englishScore,
		})

		if decryptedScore <= chiSquaredThreshold && englishScore >= englishDictionaryThreshold {
			break
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].ChiScore < results[j].ChiScore
	})

	return results, nil
}

func calculateLetterFrequency(buffer []byte) [26]letterFrequency {
	var frequencies [26]letterFrequency
	for i := range 26 {
		frequencies[i].letter = byte('a' + i)
	}

	var total float64
	for _, char := range buffer {
		if char >= 'a' && char <= 'z' {
			frequencies[char-'a'].frequency++
			total++
		} else if char >= 'A' && char <= 'Z' {
			frequencies[char-'A'].frequency++
			total++
		}
	}

	if total == 0 {
		return frequencies
	}

	for i := range frequencies {
		frequencies[i].frequency = (frequencies[i].frequency / total) * 100
	}

	return frequencies
}

func calculateChiSquared(frequencies [26]letterFrequency) float64 {
	score := 0.0

	for i := range 26 {
		expected := englishFrequencies[i].frequency
		difference := frequencies[i].frequency - expected
		score += (difference * difference) / expected
	}

	return score
}

func calculateEnglish(buffer []byte, dictionary map[string]struct{}) float64 {
	words := strings.Fields(strings.ToLower(string(buffer)))
	if len(words) == 0 {
		return 0.0
	}

	matches := 0
	for _, word := range words {
		word = strings.Trim(word, ".,!?;:\"'()-")
		if _, ok := dictionary[word]; ok {
			matches++
		}
	}

	return float64(matches) / float64(len(words))
}

func loadDictionary() (map[string]struct{}, error) {
	dictionary := make(map[string]struct{})
	for word := range strings.FieldsSeq(strings.ToLower(string(englishWordsFile))) {
		dictionary[word] = struct{}{}
	}

	return dictionary, nil
}
