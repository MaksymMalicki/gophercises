package quizgame

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type QuizData struct {
	Question string
	Answer   string
}

func RunQuizgame() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	score := 0
	stringifiedQuizData, err := csvReader.ReadAll()
	if err != nil {
		panic("Could't read the data")
	}

	var quizData []QuizData
	for _, data := range stringifiedQuizData {
		quizData = append(quizData, QuizData{data[0], data[1]})
	}
	randomGame := true
	if randomGame {
		rand.Shuffle(len(quizData), func(i, j int) { quizData[i], quizData[j] = quizData[j], quizData[i] })
	}
	numberOfQuestions := len(quizData)
	timer := time.NewTimer(30 * time.Second)
	for index, data := range quizData {
		select {
		case <-timer.C:
			fmt.Printf("\n\nTimeout, our score: %d/%d", score, numberOfQuestions)
			return
		default:
			fmt.Printf("Problem #%d: %s= ", index, data.Question)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			providedAnswer := scanner.Text()
			if providedAnswer == data.Answer {
				score++
			}
		}
	}
	fmt.Printf("Our score: %d/%d", score, numberOfQuestions)
}
