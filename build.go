package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

var nop = struct{}{}

var availablePractices = map[string]struct{}{
	"authority/":     nop,
	"common-params/": nop,
	"configuration/": nop,
	"custom-server/": nop,
	"database-sql/":  nop,
	"middlewire/":    nop,
	"project-layer/": nop,
	"template-mvc/":  nop,
	"warming-up/":    nop,
}

func main() {
	targetPractice := getPracticeChosen()

	checkValid(targetPractice)

	content := formatWorkflowContent(targetPractice)
	overwriteClassroomYaml(content)

	commitAndPush()
}

func checkValid(targetPractice string) {
	if _, ok := availablePractices[targetPractice]; !ok {
		fmt.Println("Chapter not found")
		os.Exit(-1)
	}
}

func commitAndPush() {
	exec.Command("git", "add", ".github/workflows/classroom.yml").Run()
	exec.Command("git", "commit", "-m", "Update autograding").Run()
	exec.Command("git", "push").Run()
}

func getPracticeChosen() string {
	return os.Args[1]
}

func formatWorkflowContent(name string) string {
	if name[len(name)-1] != '/' {
		name += "/"
	}
	t, err := template.New("n").Parse(classroomTemplate)
	if err != nil {
		os.Exit(-1)
	}
	var b bytes.Buffer
	t.Execute(&b, name)
	return b.String()
}

func overwriteClassroomYaml(content string) {
	f, err := os.OpenFile(".github/workflows/classroom.yml", os.O_RDWR, 0644)
	defer f.Close()
	if err != nil {
		os.Exit(-1)
	}
	f.WriteString(content)
}

const classroomTemplate = `name: GitHub Classroom Workflow

on: [push]

permissions:
  checks: write
  actions: read
  contents: read

jobs:
  build:
    name: Autograding
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Setup check target
        run: >
          export CHECK_TARGET={{.}} &&
          source ~/.bashrc
        shell: bash

      - name: Setup Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.19
        
      - name: Build practice entity
        run: go build -o main.exe {{.}}

      - uses: education/autograding@v1
`
