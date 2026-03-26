# VHS demo tape for mrglab
# Run: vhs demo.tape

Output docs/assets/demo.gif

Set Shell "bash"
Set FontSize 14
Set Width 1400
Set Height 800
Set Padding 10

# Launch in dev mode (mocked data, no API calls)
Type "mrglab -dev"
Enter
Sleep 2s

# Select a project
Sleep 500ms
Enter
Sleep 2s

# Browse merge requests
Sleep 500ms
Down
Sleep 500ms
Down
Sleep 500ms

# View MR details
Enter
Sleep 2s

# Scroll details
Down
Sleep 300ms
Down
Sleep 300ms
Down
Sleep 300ms

# Navigate discussions
Type "n"
Sleep 1s
Type "n"
Sleep 1s

# Close details panel
Escape
Sleep 1s

# Toggle help
Type "?"
Sleep 2s
Escape
Sleep 1s

# Quit
Ctrl+c
