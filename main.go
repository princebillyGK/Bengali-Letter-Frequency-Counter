package Bengali_Letter_Frequency_Calculator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type CharFreq struct {
	characterInt int32
	freq         int
}

type ShortedCharFreqList []CharFreq

func (s ShortedCharFreqList) Len() int           { return len(s) }
func (s ShortedCharFreqList) Less(i, j int) bool { return s[i].freq < s[j].freq }
func (s ShortedCharFreqList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func describeError(description string, err error) {
	fmt.Println(description)
	log.Fatal(err)
	return
}

func main() {

	outputFile, outputFileErr := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputFileErr != nil {
		describeError("Failed to create file", outputFileErr)
		return
	}

	bytes, readFileErr := ioutil.ReadFile("content.txt")
	if readFileErr != nil {
		describeError("Failed to input file", readFileErr)
		return
	}

	content := string(bytes)
	w := bufio.NewWriter(outputFile)

	characterFrequencies := make(map[int32]int)

	for _, v := range content {
		characterFrequencies[v]++
	}

	shortedList := make(ShortedCharFreqList, len(characterFrequencies))
	index := 0
	for charInt, freq := range characterFrequencies {
		shortedList[index] = CharFreq{charInt, freq}
		index++
	}

	sort.Sort(sort.Reverse(shortedList))

	for _, charFreq := range shortedList {
		if _, err := fmt.Fprintf(w, "%v (%v) : %v \n", string(charFreq.characterInt), charFreq.characterInt, charFreq.freq);
			err != nil {
			describeError("Something went wrong", err)
		}
	}

	if flushError := w.Flush(); flushError != nil {
		describeError("Failed to write output file", flushError)
	}
}
