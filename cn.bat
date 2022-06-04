@echo off

:: set MONGO_HOST=<hostname>
:: set MONGO_USERNAME=<name>
:: set MONGO_PASSWORD=<password>
:: set MONGO_DATABASE=cn

powershell.exe Set-ExecutionPolicy RemoteSigned -Scope LocalMachine
powershell.exe .\cn.ps1
powershell.exe Set-ExecutionPolicy Undefined -Scope LocalMachine
