@echo off
set BINARY=stream-set
set VERSION=v0.0.21
set GOARCH=amd64
set GOOS=windows
REM go build -o release/%BINARY%-%VERSION%-windows-amd64.exe
setlocal
set "list=.json .ini"
(
 for %%i in (%list%) do (
  echo Copying: %%i
  copy *%%i release
 )
)
set "list=img server"
(
 for %%i in (%list%) do (
  echo Copying: %%i
  echo D|xcopy /s /Y /EXCLUDE:exclude.txt %%i release\%%i
 )
)
IF "%1"=="dev" (
    echo "Dev Build"
    rsrc -manifest main.manifest -o release/rsrc.syso
    go build -o release/%BINARY%-%VERSION%-windows-%GOARCH%.exe
) 
IF "%1"=="pro" (
    echo "Production Build"
    rsrc -manifest main.manifest -o release/rsrc.syso
    go build -ldflags="-H windowsgui" -o release/%BINARY%-%VERSION%-windows-amd64.exe
)