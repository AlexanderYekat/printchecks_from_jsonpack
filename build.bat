@echo off
setlocal enabledelayedexpansion
cd /d "%~dp0"

echo [1/3] go generate
go generate ./...
if errorlevel 1 goto :error

echo [2/3] go build -o jsontokkt.exe
go build -o jsontokkt.exe .
if errorlevel 1 goto :error

echo [3/3] cleanup old default binary name (if any)
if exist clientrabbit.exe del /q clientrabbit.exe >nul 2>&1

echo Done. Output: jsontokkt.exe
exit /b 0

:error
echo Build failed. See messages above.
exit /b 1


