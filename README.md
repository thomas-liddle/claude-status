Sets your Slack status to a random activity word (e.g. "Reticulating...") with the `:claude:` emoji. Intended for use as a Claude Code `PreToolUse` hook so your status updates in real time while Claude is working.

To use it you'll need the following env variable:
```
SLACK_TOKEN     (xoxp-... User OAuth Token)
```

To get a token:
 - go to https://api.slack.com/apps and go to the "Status Updater" app
 - install the app to your workspace
 - copy the "User OAuth Token" (starts with `xoxp-`)
 - add `export SLACK_TOKEN=xoxp-...` to your `~/.zshrc` or `~/.zprofile`

The binary is installed via `make install` from the module root. To wire it up, add the following to `~/.claude/settings.json`:

```json
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "",
        "hooks": [
          {
            "type": "command",
            "command": "/Users/YOUR_USERNAME/go/bin/claude-status",
            "async": true
          }
        ]
      }
    ]
  }
}
```
