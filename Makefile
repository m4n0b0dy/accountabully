build-win:
	rsrc -ico ./application/ui/icon/logo.ico
	go build -ldflags "-H=windowsgui" -o accountabully.exe
	if not exist "C:\\Program Files\\Accountabully" mkdir "C:\\Program Files\\Accountabully"
	copy accountabully.exe "C:\\Program Files\\Accountabully\\accountabully.exe"
