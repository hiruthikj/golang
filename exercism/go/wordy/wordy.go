package wordy

import (
	"strconv"
	"strings"
)

func operate(splitQuestion []string) (int, bool) {
	currentResult := 0
	i := 0
	
	for i<len(splitQuestion) {
		number, err := strconv.Atoi(splitQuestion[i])
		if err != nil {
			return 0, false
		}
		i++
		
		if 

		
	}
	
}

func cleanQuestion (question string) (string, bool) {
	question, found := strings.CutPrefix(question, "What is")
	if !found {
		return "", false
	}

	question, found = strings.CutSuffix(question, "?")
	if !found {
		return "", false
	}
	return question, true
}

func Answer(question string) (int, bool) {
	cleanedQuestion, ok := cleanQuestion(question)
	if !ok {
		return 0, false
	}

	splitQuestion := strings.Split(cleanedQuestion, " ")
	result, ok := operate(splitQuestion)

	return result, ok
}
