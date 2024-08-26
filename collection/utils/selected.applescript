#!/usr/bin/env osascript

tell application "System Events" to keystroke "c" using {command down}
delay 0.5
set selectedText to the clipboard
log selectedText
return selectedText
