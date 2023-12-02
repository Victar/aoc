package util

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

var BaseDir = "/Users/vkad2506/AdventOfCode/java/src/test/resources/"
var GptApiKey = os.Getenv("GPT_API_KEY")
var AocSession = os.Getenv("ADVENT_OF_CODE_SESSION")

type SolverTask struct {
	AocTask  AocTask  `json:"aocTask"`
	AiSolver AiSolver `json:"aiSolver"`
	AocSolve AocSolve `json:"aocSolve"`
}
type AocTask struct {
	Text         string `json:"text,omitempty"`
	Input        string `json:"input,omitempty"`
	Year         string `json:"year,omitempty"`
	Day          string `json:"day,omitempty"`
	SilverSolved bool   `json:"silverSolved"`
	GoldSolved   bool   `json:"goldSolved"`
}

type AiSolver struct {
	AIModel              string `json:"aiModel,omitempty"`
	AIMaxTokens          int    `json:"aiMaxTokens,omitempty"`
	AIPrompt             string `json:"aiPrompt,omitempty"`
	AIResponse           string `json:"aiResponse,omitempty"`
	AIResponseCode       string `json:"aiResponseCode,omitempty"`
	AIResponseCodeOutput string `json:"aiResponseCodeOutput,omitempty"`
	AIAnswerFound        bool   `json:"aiAnswerFound"`
	AIAnswer             string `json:"aiAnswer,omitempty"`
}

type AocSolve struct {
	AttemptLevel       string `json:"level,omitempty"`
	AttemptAnswer      string `json:"answer,omitempty"`
	AttemptResponse    string `json:"attemptResponse,omitempty"`
	AttemptInfo        string `json:"attemptInfo,omitempty"`
	AttemptSuccess     bool   `json:"attemptSuccess,omitempty"`
	AttemptWait        int    `json:"attemptWait,omitempty"`
	AttemptAllowSubmit bool   `json:"attemptAllowSubmit"`
}

func NewSolverTask(year int, day int, aiModel string, aiMaxTokens int) SolverTask {
	return SolverTask{AocTask: AocTask{Year: strconv.Itoa(year), Day: strconv.Itoa(day)}, AiSolver: AiSolver{AIModel: aiModel, AIMaxTokens: aiMaxTokens}, AocSolve: AocSolve{AttemptAllowSubmit: true}}
}

func (task *SolverTask) InitOacTask() error {
	textInput, err := DownloadInput(task.AocTask.Year, task.AocTask.Day)
	if err != nil {
		return err
	}
	task.AocTask.Input = textInput
	_, err = SaveStringToFile(textInput, task.GetPath()+"input.txt")
	if err != nil {
		return err
	}
	textText, err := DownloadText(task.AocTask.Year, task.AocTask.Day)
	if err != nil {
		return err
	} else {
		task.AocTask.Text = textText
	}
	return nil
}

func (task *SolverTask) IsSolved() bool {
	return task.IsGoldSolved() || (task.IsSilverSolved() && task.AocSolve.AttemptSuccess)
}
func (task *SolverTask) IsSilverSolved() bool {
	if task.IsGoldSolved() {
		return true
	}
	silverStart := strings.Index(task.AocTask.Text, "The first half of this puzzle is complete! It provides one gold star:")
	return silverStart >= 0
}

func (task *SolverTask) IsGoldSolved() bool {
	goldStart := strings.Index(task.AocTask.Text, "Both parts of this puzzle are complete! They provide two gold stars:")
	return goldStart >= 0
}

func (task *SolverTask) LevelToSolve() string {
	if task.IsGoldSolved() {
		return ""
	}
	if task.IsSilverSolved() {
		return "2"
	}
	return "1"
}

func (task *SolverTask) GetPath() string {
	return "year" + task.AocTask.Year + "/day" + task.AocTask.Day + "/"
}

func (task *SolverTask) GetTaskText() string {
	plainText := task.AocTask.Text
	startString := "---"
	endString := "Answer:"
	startIndex := strings.Index(plainText, startString)
	endIndex := strings.Index(plainText, endString)
	if endIndex > 0 && startIndex > 0 && startIndex < endIndex {
		return plainText[startIndex:endIndex]
	}
	return plainText
}

func (task *SolverTask) GetJsonString() string {
	task.AocTask.SilverSolved = task.IsSilverSolved()
	task.AocTask.GoldSolved = task.IsGoldSolved()
	jsonString, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonString)
}
