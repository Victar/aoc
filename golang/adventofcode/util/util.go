package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func ReadInput(path string) ([][]rune, error) {
	file, err := os.Open(BaseDir + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}
	return lines, scanner.Err()
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(BaseDir + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func ReadFileSingle(path string) (string, error) {
	filePath := BaseDir + path
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func SaveStringToFile(fileContent string, path string) (string, error) {
	// Create or open the file for writing
	if err := os.MkdirAll(filepath.Dir(BaseDir+path), os.ModePerm); err != nil {
		return "", err
	}
	file, err := os.Create(BaseDir + path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write the string to the file
	_, err = file.WriteString(fileContent)
	if err != nil {
		return "", err
	}

	return path, nil
}

func RunGoSolution(path string) (string, error) {
	// Change the working directory
	scriptPath := BaseDir + path
	err := os.Chdir(scriptPath)
	if err != nil {
		//fmt.Println("Error changing directory:", err)
		return "", err
	}

	// Initialize a Go module
	initCmd := exec.Command("go", "mod", "init", "ai")
	initCmd.Stdout = os.Stdout
	initCmd.Stderr = os.Stderr
	if err := initCmd.Run(); err != nil {
		//fmt.Println("Error running 'go mod init':", err)
		return "", err
	}

	// Build the Go application
	buildCmd := exec.Command("go", "build")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return "", err
	}

	// Execute the compiled binary and capture output
	exeCmd := exec.Command("./ai")
	output, err := exeCmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func BackSolution(path string) {
	scriptPath := BaseDir + path
	filesToCopy := []string{"ai", "main.go", "go.mod", "task.json"}

	err := copyFilesToTimestampedFolder(scriptPath, filesToCopy)
	if err != nil {
		fmt.Println(err)
	}
}

func copyFilesToTimestampedFolder(path string, filesToCopy []string) error {
	// Create a timestamped folder name
	timestamp := time.Now().Format("20060102_150405")
	folderName := fmt.Sprintf("ai_%s", timestamp)

	// Create the subfolder
	err := os.Mkdir(path+folderName, 0755)
	if err != nil {
		return err
	}

	// Copy files to the subfolder
	for _, file := range filesToCopy {
		sourcePath := path + file
		destinationPath := filepath.Join(path+folderName, file)

		err := copyFile(sourcePath, destinationPath)
		if err != nil {
			fmt.Println(err)
		}

		// Delete the file from the source after copying
		err = os.Remove(sourcePath)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Files copied to %s\n\n", folderName)
	return nil
}

func copyFile(sourcePath, destinationPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}
