# Welcome to Accountabully!

<div align="center">
  <img src="application/ui/icon/logo.ico" alt="Accountabully Logo">
</div>

## Overview
Accountabully is your personal productivity enforcer. Tired of getting distracted by YouTube, Instagram, or Reddit? Accountabully helps you stay focused and accountable by closing or minimizing the apps you specify.

True to its name, it acts like a "bully" for your productivity, keeping distractions at bay until you're ready to focus again.

**Warning**: Accountabully is not for the faint of heart. It's a no-nonsense app that will keep you on track! Distracting apps are shut down without mercy or saving, so use it wisely. Or just stay focused and avoid the bully altogether!

## Features
- **Customizable Block List**: Add any number of apps to your block list.
- **Real-Time Monitoring**: Checks every second to enforce your rules by closing or minimizing specified apps.
- **Background Operation**: Runs quietly in the system tray, so it doesn’t disrupt your workflow.
- **Runs Locally Only**: No internet connection required and no data ever leaves your pc, ensuring your privacy.
- **System Tray Controls**:
    - Start/Stop the bully.
    - Add or modify blocked apps and actions.
    - Lock the bully to prevent tampering (bully must be running).
    - Quit the app when you're done.

## Getting Started
1. **Run Accountabully**: Launch the application for the first time.
2. **Configure Rules**:
    - Click on the system tray icon, then "Show Rules" to open the settings.
    - Add the names of apps you want to block along with the desired action (close or minimize).
    - (note) Because web browsers like Firefox and Chrome run as one Pid, Accountabully will close the entire browser instead of a single tab.
3. **Activate the Bully**:
    - Check the system tray icon to ensure the bully is running.
    - Once activated, it will enforce the rules in real time.
4. **Lock the Bully**:
    - Use the "Lock" option from the system tray to prevent yourself from stopping it.

## Pro Tips
1. Locking the bully is the ultimate step to stay accountable—it ensures you can't easily bypass your own productivity safeguards!
2. By adding Accountabully to startup programs and Window's Task Scheduler (Taskschd.msc) with a recurring time trigger ([guide](https://www.backup4all.com/how-to-create-a-new-task-using-windows-task-scheduler-kb.html)), you can ensure it runs automatically at all times.
3. There's a boolean, hardCoreMode in main.go in ui. It turns bully on at startup, locks app on startup, and auto restarts bully after a set time. Closing accountabully just makes it angrier.

## Setup
**Prerequisites**: Install GCC
- [GCC Install Guide](https://dev.to/gamegods3/how-to-install-gcc-in-windows-10-the-easier-way-422j)

**Download Exe**: Download accountabully.exe from the releases page

**Or Build From Scratch**: golang, rsrc, fork this repo
- [Go Install Guide](https://go.dev/doc/tutorial/getting-started)
- `go install -v github.com/akavel/rsrc@latest`
- ```
  git clone git@github.com:helpful-ml/accountabully.git
  cd accountabully
  rsrc -ico ./application/ui/icon/logo.ico
  go build -ldflags "-H=windowsgui" -o accountabully.exe
  if not exist "C:\\Program Files\\Accountabully" mkdir "C:\\Program Files\\Accountabully"
  copy accountabully.exe "C:\\Program Files\\Accountabully\\accountabully.exe"
  ```
---
Enjoy focused, distraction-free sessions with Accountabully!

Currently only supports Windows OS. (could be forked for other OS)