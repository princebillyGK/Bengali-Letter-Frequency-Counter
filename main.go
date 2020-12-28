package Bengali_Letter_Frequency_Calculator

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const InputDirectory = "./inputs"
const OutputDirectory = "./outputs"

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

type InputItem struct {
	Info os.FileInfo
	path string
}



func traverseFilesAndRun(directory string, cb func(string) error) error {
	traversingError := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir(){
			if analyzingError := cb(path); analyzingError != nil {
				return analyzingError
			}
		}
		return nil
	})
	if traversingError != nil {
		return traversingError
	}
	return nil
}

func main() {

	outputFile, outputFileErr := os.OpenFile(OutputDirectory, os.O_WRONLY|os.O_CREATE, 0666)
	if outputFileErr != nil {
		describeError("Failed to create file", outputFileErr)
		return
	}

	characterFrequencies := make(map[int32]int)
	w := bufio.NewWriter(outputFile)

	err1 := traverseFilesAndRun(InputDirectory, func(path string) error {
		bytes, readFileErr := ioutil.ReadFile(path)
		if readFileErr != nil {
			return readFileErr
		}
		content := string(bytes)
		for _, v := range content {
			characterFrequencies[v]++
		}
		return nil
	})
	if err1 != nil {
		log.Fatal(err1)
		return
	}

	index := 0
	shortedList := make(ShortedCharFreqList, len(characterFrequencies))
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
