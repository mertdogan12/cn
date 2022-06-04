$ID = (Get-ComputerInfo).WindowsProductId
Rename-Computer -NewName (Invoke-WebRequest -URI http://localhost:3000/getname?id=$ID).Content
