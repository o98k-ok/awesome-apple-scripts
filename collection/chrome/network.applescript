#!/usr/bin/env osascript

tell application "System Events"
  tell process "Chrome"
    set frontmost to true
    delay 0.2
    click menu item "开发者工具" of menu "开发者" of menu item "开发者" of menu "视图" of menu bar 1
  end tell
end tell
