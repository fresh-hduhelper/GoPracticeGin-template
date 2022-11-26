package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var availablePractices = `authority/common-params/configuration/custom-server/database-sql/middlewire/project-layer/template-mvc/warming-up/`

func main() {
	targetPractice := getPracticeChosen()

	checkValid(targetPractice)

	content := formatWorkflowContent(targetPractice)
	overwriteClassroomYaml(content)

	commitAndPush("finished")
}

func checkValid(targetPractice string) {
	if !strings.Contains(availablePractices, targetPractice) {
		os.Exit(-1)
	}
}

func commitAndPush(msg string) {
	exec.Command("git", "add", ".github/workflows/classroom.yml").Run()
	exec.Command("git", "commit", "-m", "Update autograding").Run()
	exec.Command("git", "push").Run()
	fmt.Printf("msg: %v\n", msg)
}

func getPracticeChosen() (pracName string) {
	flag.StringVar(&pracName,
		"Chosen Practice's Name", "warming-up",
		"pointing out which practice to check")
	flag.Parse()
	return
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

      - name: Setup Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.19
        
      - name: Build practice entity
        run: go build -o main.exe {{.}}

      - uses: education/autograding@v1
`
