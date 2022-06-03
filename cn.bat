set id = "123424"
powershell.exe Rename-Computer -NewName (Invoke-WebRequest -URI http://localhost:3000/getname?id=%id).Content
