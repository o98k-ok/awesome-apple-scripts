tell application "System Events"
  tell process "Chrome"
    set frontmost to true
    click menu item "JavaScript 控制台" of menu "开发者" of menu item "开发者" of menu "视图" of menu bar 1
  end tell
end tell