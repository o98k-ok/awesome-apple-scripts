#!/usr/bin/env osascript

tell application "System Events"
  tell process "Chrome"
    set frontmost to true
    delay 0.5
    click menu item "移到“内建视网膜显示器”" of menu "窗口" of menu bar 1
  end tell
end tell