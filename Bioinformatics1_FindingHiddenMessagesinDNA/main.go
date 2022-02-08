//package Bioinformatics1_FindingHiddenMessagesinDNA
package main

import "fmt"

//def patternCount(self, text, pattern):
//	count = 0
//	for i in range(0,(len(text)-len(pattern))+1):
//		subtext = text[i:len(pattern)+i]
//		if subtext == pattern:
//			count +=1
//	return count

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

func main() {
	fmt.Println("Here to expect Bioinformatics Code in Golang")
	text := "GCGCG"
	pattern := "GCG"
	fmt.Println(patternCount(text, pattern))
}
