import os, strutils

proc getCurModuleContent(pkgName: string): string =
  result =
    "module " & pkgName & "\n\n" & "go 1.19\n"
    
proc getCurReadmeContent(dirName: string): string = 
  result =
    "# " & dirName & "\n\n" 

var exceptDirs = @[".git", "misc"]

proc addExceptDir(dirName: string) {.used.} =
  exceptDirs.add(dirName)

proc genBatchGoModFile() =
  let curDir = getCurrentDir()
  for (dirType, path) in walkDir(curDir):
    if dirType != pcDir: continue

    let dirName = path.split(r"\")[^1]
    if dirName in exceptDirs: continue

    echo "write in " & dirName
    let gomod = open(dirName&"/go.mod", mode = fmWrite)
    gomod.write(getCurModuleContent(dirName))

proc genBatchReadmeFile() =
  let curDir = getCurrentDir()
  for (dirType, path) in walkDir(curDir):
    if dirType != pcDir: continue

    let dirName = path.split(r"\")[^1]
    if dirName in exceptDirs: continue

    echo "write in " & dirName
    let gomod = open(dirName&"/README.md", mode = fmWrite)
    gomod.write(getCurReadmeContent(dirName))

if isMainModule:
  genBatchReadmeFile(); genBatchGoModFile()
