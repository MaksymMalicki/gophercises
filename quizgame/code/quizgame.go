package quizgame

import (
	"bufio"
	"encoding/csv"
	"flag"
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

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffleQuestions := flag.Bool("shuffle", false, "shuffle the quiz questions")

	flag.Parse()

	print(*timeLimit, *shuffleQuestions)

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	stringifiedQuizData, err := csvReader.ReadAll()
	if err != nil {
		panic("Could't read the data")
	}

	var quizData []QuizData
	for _, data := range stringifiedQuizData {
		quizData = append(quizData, QuizData{data[0], data[1]})
	}

	if *shuffleQuestions {
		rand.Shuffle(len(quizData), func(i, j int) { quizData[i], quizData[j] = quizData[j], quizData[i] })
	}

	numberOfQuestions := len(quizData)
	score := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
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
