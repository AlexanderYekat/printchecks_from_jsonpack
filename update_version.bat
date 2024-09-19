@echo off
setlocal enabledelayedexpansion

:: Получаем текущую дату
for /f "tokens=2 delims==" %%I in ('wmic os get localdatetime /value') do set datetime=%%I
set year=%datetime:~0,4%
set month=%datetime:~4,2%
set day=%datetime:~6,2%

:: Формируем новую версию
set new_version=%year%.%month%.%day%.1

:: Обновляем файл jsontokkt.go
powershell -Command "(gc jsontokkt.go) -replace 'const Version_of_program = \".*\"', 'const Version_of_program = \"%new_version%\"' | Out-File -encoding ASCII jsontokkt.go"

:: Обновляем файл versioninfo.json
powershell -Command "(gc versioninfo.json) -replace '\"Major\": \d+', '\"Major\": %year%' | Out-File -encoding ASCII versioninfo.json"
powershell -Command "(gc versioninfo.json) -replace '\"Minor\": \d+', '\"Minor\": %month%' | Out-File -encoding ASCII versioninfo.json"
powershell -Command "(gc versioninfo.json) -replace '\"Patch\": \d+', '\"Patch\": %day%' | Out-File -encoding ASCII versioninfo.json"
powershell -Command "(gc versioninfo.json) -replace '\"ProductVersion\": \".*\"', '\"ProductVersion\": \"v%new_version%\"' | Out-File -encoding ASCII versioninfo.json"

:: Добавляем измененные файлы в индекс Git
git add jsontokkt.go versioninfo.json

echo Version updated to %new_version%