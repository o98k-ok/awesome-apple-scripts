tell application "System Events"
  tell process "Chrome"
    set frontmost to true
    click menu item "关闭其他标签页" of menu "标签页" of menu bar 1
  end tell
end tell