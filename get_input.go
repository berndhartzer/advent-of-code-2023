package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type aocInput struct {
	buf *bytes.Buffer
}

func (a *aocInput) asIntSlice() []int {
	var contents []int
	scanner := bufio.NewScanner(a.buf)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("failed to transform input")
		}
		contents = append(contents, num)
	}

	return contents
}

func (a *aocInput) asStringSlice() []string {
	var contents []string
	scanner := bufio.NewScanner(a.buf)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents
}

func (a *aocInput) asString() string {
	s := a.buf.String()
	s = strings.TrimSuffix(s, "\n")
	return s
}

func (a *aocInput) asCommaSeparatedInts() []int {
	stringVal := a.asString()
	split := strings.Split(stringVal, ",")

	intVals := []int{}
	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic("failed to transform input")
		}

		intVals = append(intVals, n)
	}

	return intVals
}

func getInput(day int) (*aocInput, error) {
	filePath := fmt.Sprintf("./input/%d.txt", day)

	// Attempt to retrieve input from file on disk first
	data, err := getFileOnDisk(filePath)
	if err != nil {
		fmt.Printf("Error getting file on disk: %v\ncontinuing...\n", err)
	} else {
		return &aocInput{
			buf: data,
		}, nil
	}

	session, err := os.ReadFile("./input/.session")
	if err != nil {
		return nil, fmt.Errorf("failed to read session: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.AddCookie(&http.Cookie{
		Name: "session",
		Value: strings.TrimRight(string(session), "\r\n"),
	})

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 status code: %v", resp.StatusCode)
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create new file %s: %w", filePath, err)
	}
	defer newFile.Close()

	fileBuf := new(bytes.Buffer)

	w := io.MultiWriter(newFile, fileBuf)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to write to multiwriter: %w", err)
	}

	return &aocInput{
		buf: fileBuf,
	}, nil
}

func getFileOnDisk(filePath string) (*bytes.Buffer, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file %s is empty", filePath)
	}

	return bytes.NewBuffer(data), nil
}
