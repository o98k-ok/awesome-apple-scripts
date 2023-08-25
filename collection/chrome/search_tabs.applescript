#!/usr/bin/env osascript

tell application "System Events"
  tell process "Chrome"
    set frontmost to true
    delay 0.5
    click menu item "搜索标签页..." of menu "标签页" of menu bar 1
  end tell
end tell
