@echo off
setlocal 
cd app

set APPNAME=webreboot
set VERSION=0.02

set LDFLAGS="-X main.appName=%APPNAME% -X main.appVer=%VERSION% -s -w"

set OUT=..\output

set arch=amd64 386 arm arm64
set os=linux windows

(for %%a in (%arch%) do (
    (for %%o in (%os%) do (call :build %%a %%o))
))

:build
set GOARCH=%1
set GOOS=%2

if "%GOARCH%"=="" goto :eof

set EXT=

if "%GOOS%"=="windows" (
    set EXT=.exe
)

echo %GOARCH%/%GOOS%
go build -ldflags %LDFLAGS% -o %OUT%\%VERSION%\%GOOS%-%GOARCH%\%APPNAME%%EXT%
goto :eof

:End
endlocal
