tell application "System Events"
  tell process "Chrome"
    set frontmost to true
    click menu item "移到“内建视网膜显示器”" of menu "窗口" of menu bar 1
  end tell
end tell