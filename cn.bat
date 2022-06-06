@echo off

powershell.exe Rename-Computer -NewName $(Invoke-WebRequest -URI http://localhost:3001/getname?id=$((Get-ComputerInfo).WindowsProductId)).Content
