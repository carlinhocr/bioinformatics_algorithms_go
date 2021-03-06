//package Bioinformatics1_FindingHiddenMessagesinDNA
package main

import "fmt"

func patternCount(text string, pattern string) int {
	count := 0
	n := len(text)
	k := len(pattern)
	for i := 0; i < (n-k)+1; i++ {
		subtext := text[i : i+k]
		if subtext == pattern {
			count++
		}
	}
	return count
}

//def frequencyTable(self, text="hola", k=0):
//	freqMap = {}
//	n = len(text)
//	for i in range(0, n - k):
//		pattern = text[i:k+i]
//		if not pattern in freqMap:
//			freqMap[pattern]=1
//		else:
//			freqMap[pattern] += 1
//	return freqMap

func frequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)
	for i := 0; i < (n - k); i++ {
		pattern := text[i : i+k]
		if _, ok := freqMap[pattern]; ok == false {
			freqMap[pattern] = 1
		} else {
			freqMap[pattern]++
		}
	}
	return freqMap
}

//def MaxMap(self,freqMap):
//	maximum = 0
//	for element in freqMap:
//		if abs(freqMap[element]) > maximum:
//			maximum = abs(freqMap[element])
//	return maximum

func maxMap(freqMap map[string]int) int {
	maximum := 0
	for _, element := range freqMap {
		if element > maximum {
			maximum = element
		}
	}
	return maximum
}

//def BetterFrequentWords(self,text,k):
//	frequentPatterns = []
//	freqMap = self.frequencyTable(text,k)
//	max = self.MaxMap(freqMap)
//	for pattern in freqMap:
//		# the patern is the key in the dictionary, the value it is how often it is found in text
//		if freqMap[pattern] == max:
//			frequentPatterns.append(pattern)
//	return frequentPatterns

func main() {
	fmt.Println("Here to expect Bioinformatics Code in Golang")
	// Exercise 1 PatternCount
	text := "GCGCG"
	pattern := "GCG"
	fmt.Println(patternCount(text, pattern))
	// Exercise 2 BetterFrequentWords
	text = "ACGTTGCATGTCGCATGATGCATGAGAGCT"
	k := 4
	table := frequencyTable(text, k)
	fmt.Println(table)
	fmt.Println(maxMap(table))
}
