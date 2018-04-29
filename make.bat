@echo off
set BINARY=stream-set
set VERSION=v0.0.3
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
del .\release\versioninfo.json
IF "%1"=="dev" (
    echo "Dev Build"
    REM rsrc -manifest "main.manifest" -ico "img/icon.ico" -o "rsrc.syso"
    go generate
    go build -i -o release/%BINARY%.exe
) 
IF "%1"=="pro" (
    echo "Production Build"
    REM rsrc -manifest "main.manifest" -ico "img/icon.ico" -o "rsrc.syso"
    go generate
    go build -i -ldflags="-H windowsgui" -o release/%BINARY%.exe
    7z a .\release\%BINARY%.zip .\release\*
)