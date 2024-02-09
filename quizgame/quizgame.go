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
	gameChan := make(chan int)
	go func() {
		for index, data := range quizData {
			fmt.Printf("Problem #%d: %s= ", index, data.Question)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			providedAnswer := scanner.Text()
			if providedAnswer == data.Answer {
				score++
			}
		}
		gameChan <- score
	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Printf("\n\nTimeout, our score: %d/%d", score, numberOfQuestions)
	case <-gameChan:
		fmt.Printf("\n\nYour score: %d/%d", score, numberOfQuestions)
	}

}
