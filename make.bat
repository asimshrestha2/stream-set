set BINARY=stream-set
set VERSION=v0.0.1
REM set GOARCH=amd64
REM set GOOS=windows
REM go build -o release/%BINARY%-%VERSION%-windows-amd64.exe

IF "%1"=="dev" (
    echo "Dev Build"
    rsrc -manifest main.manifest -o release/rsrc.syso
    go build -o release/%BINARY%-%VERSION%-windows-amd64.exe
) 
IF "%1"=="pro" (
    echo "Production Build"
    rsrc -manifest main.manifest -o release/rsrc.syso
    go build -ldflags="-H windowsgui" -o release/%BINARY%-%VERSION%-windows-amd64.exe
)