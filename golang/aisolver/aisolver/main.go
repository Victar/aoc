package main

import (
	"aisolver/util"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var model = openai.GPT4TurboPreview
var maxToken = 4096

var submitAllow = false // double check
var runSolved = true    // double check

var year = 2023
var dayStart = 25
var daysToSolve = 1
var attempts = 1

// Make sure
// BaseDir, GptApiKey, AocSession are defined correctly
// var BaseDir = "/Users/vkad2506/AdventOfCode/java/src/test/resources/"
// var GptApiKey = os.Getenv("GPT_API_KEY")
// var AocSession = os.Getenv("ADVENT_OF_CODE_SESSION")

func main() {
	for day := dayStart; day < dayStart+daysToSolve; day++ {
		fmt.Printf("Running task for year %d day %d...\n", year, day)
		startTime := time.Now()
		for attempt := 1; attempt <= attempts; attempt++ {
			task := util.NewSolverTask(year, day, model, maxToken)
			err := runAny(&task)
			if err != nil {
				fmt.Printf("Error on attempt %d: %v\n", attempt, err)
				continue
			}
			if task.IsSolved() {
				fmt.Printf("Task solved on attempt %d!\n", attempt)
				break
			}
			if !task.AocSolve.AttemptAllowSubmit {
				fmt.Printf("Submit is not allowed  %d!\n", attempt)
				continue
			}
			if task.AocSolve.AttemptWait > 0 {
				waitNSeconds(task.AocSolve.AttemptWait)
			}
		}
		elapsedTime := time.Since(startTime)
		fmt.Printf("Task for day %d finished in %s\n", day, elapsedTime)
	}
}

func runAny(task *util.SolverTask) error {
	defer func() {
		_, _ = util.SaveStringToFile(task.GetJsonString(), task.GetPath()+"task.json")
		util.BackSolution(task.GetPath())
	}()

	err := task.InitOacTask()
	if err != nil {
		return err
	}

	if task.IsSolved() && !runSolved {
		return nil
	} else {
		err := task.InitPrompt()
		if err != nil {
			return err
		}
		solution := util.GetGPTResponseChat(task.AiSolver.AIPrompt, task.AiSolver.AIModel, task.AiSolver.AIMaxTokens)
		task.AiSolver.AIResponse = solution
		goCode, err := extractGoCode(solution)
		if err != nil {
			return err
		}
		task.AiSolver.AIResponseCode = goCode
		_, err = util.SaveStringToFile(goCode, task.GetPath()+"main.go")
		if err != nil {
			return err
		}
		output, err := util.RunGoSolution(task.GetPath())
		if err != nil {
			return err
		}
		println(output)
		task.AiSolver.AIResponseCodeOutput = output
		answer := getLastWord(output)
		task.AiSolver.AIAnswer = answer
		if answer != "" {
			task.AiSolver.AIAnswerFound = true
		}
	}

	//Is answer found try to submit it
	if task.AiSolver.AIAnswerFound {
		level := task.LevelToSolve()
		task.AocSolve.AttemptLevel = level
		task.AocSolve.AttemptAnswer = task.AiSolver.AIAnswer

		isFull, _ := util.IsGlobalLeaderboardFull(task.AocTask.Year, task.AocTask.Day, level)
		if !isFull {
			task.AocSolve.AttemptInfo = "LeaderBoard is not full solution was not submitted"
			task.AocSolve.AttemptAllowSubmit = false
		} else if submitAllow {
			answerText, err := util.SubmitAnswer(task.AocTask.Year, task.AocTask.Day, level, task.AiSolver.AIAnswer)
			if err != nil {
				return err
			}
			task.AocSolve.AttemptResponse = answerText
			task.AocSolve.AttemptSuccess = isCorrectAnswer(answerText)
			secondsToWait, err := extractSeconds(answerText)
			if err == nil {
				task.AocSolve.AttemptWait = secondsToWait
			}
		}
	}
	return nil
}

func waitNSeconds(nSeconds int) {
	fmt.Printf("Waiting for %d seconds...\n", nSeconds)
	// Loop to print countdown every second
	for i := nSeconds; i > 0; i-- {
		fmt.Printf("%d seconds left...\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Waited for", nSeconds, "seconds. Continue with the next task.")
}

func extractSeconds(input string) (int, error) {
	// Define a regular expression to match the number of seconds
	re := regexp.MustCompile(`(\d+)s`)

	// Find the matches in the input string
	matches := re.FindStringSubmatch(input)

	// Check if there is a match
	if len(matches) < 2 {
		return 0, fmt.Errorf("no seconds found in the input string")
	}

	// Extract the matched number of seconds
	secondsStr := matches[1]

	// Convert the string to an integer
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert seconds to integer: %v", err)
	}

	return seconds, nil
}

func getLastWord(input string) string {
	// Split the input string by whitespace and line breaks, and trim spaces
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r == ' ' || r == '\n' || r == '\r'
	})

	// If there are no words, return an empty string
	if len(words) == 0 {
		return ""
	}

	// Return the last word
	return words[len(words)-1]
}
func isCorrectAnswer(str string) bool {
	answerIndex := strings.Index(str, "That's the right answer!")
	return answerIndex >= 0
}

func extractGoCode(str string) (string, error) {
	start := "```golang"
	start2 := "```go"
	start3 := "```"

	end := "```"
	startIndex := strings.Index(str, start)
	if startIndex == -1 {
		startIndex = strings.Index(str, start2)
		start = start2
		if startIndex == -1 {
			startIndex = strings.Index(str, start3)
			start = start3
			if startIndex == -1 {
				return "", fmt.Errorf("Start string not found in the input string")
			}
		}
	}

	startIndex += len(start)

	endIndex := strings.Index(str[startIndex:], end)
	if endIndex == -1 {
		return "", fmt.Errorf("End string not found in the input string")
	}
	return str[startIndex : startIndex+endIndex], nil
}
