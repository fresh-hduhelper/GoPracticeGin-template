package main

import (
	"log"
	"os"
	"os/exec"
)

var workflow = `name: GitHub Classroom Workflow

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

    - uses: actions/setup-go@v3.2.1
      with:
          go-version: 1.19

    - name: Collect Dependencies
      run: pip3 install requests
      
    - name: Build
      run: |
      
    - name: Run after build
      run: 

    - uses: education/autograding@v1
`

func main() {
	//TODO:
	// 1. hash 所有文件夹及其内容，有修改的话，在启动本地检查的时候再检查
	// 2. go run build.go 时能够覆写 workflow
	//
}

var checkScriptContent = `//go:build check
//  +build check
package main

func main() {

}
`

func genCheckScript() {
	checkScript, err := os.Create("check.go")
	if err != nil {
		log.Fatal(err)
	}
	_, err = checkScript.Write([]byte(checkScriptContent))
	if err != nil {
		log.Fatal(err)
	}
}

type CommandBatch []*exec.Cmd

func (cb *CommandBatch) lazyInit() {
	if cb == nil {
		*cb = make([]*exec.Cmd, 16)
	}
}

func (cb *CommandBatch) AddBatch(cmdBatch ...*exec.Cmd) {
	cb.lazyInit()
	*cb = append(*cb, cmdBatch...)
}

func (cb *CommandBatch) ExecAll() {
	if cb == nil || len(*cb) <= 0 {
		return
	}
	for _, cmd := range *cb {
		cmd.Start()
	}
}

func updateWorkflowOverwrite() {

}

func updateWorkflowPush() {
	var cmds CommandBatch
	cmds.AddBatch(
		exec.Command("git", "add", ".github/workflows/classroom.yml"),
		exec.Command("git", "commit", "-m", "Update autograding"),
		exec.Command("git", "push"),
	)
	cmds.ExecAll()
}
