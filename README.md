# Discord-predictions

Script to parse out predicted gender and age from Discord `analytics/events.json` file

# Usage

- Request you data from Discord in _User Settings > Privacy & Safety > Request all of my Data_
- Wait for an email from Discord and download/unzip the `package.zip` file
- `./discord-predictions-${arch} -f path/to/package/activity/analytics/events.{xyz}.json`

# Pre-built platforms

- Windows 10: `.\discord-predictions.exe -f [file]`
- macOS (apple silicon): `./discord-predictions-macOS-arm64 -f [file]`
- macOS (intel): `./discord-predictions-macOS-amd64 -f [file]`
- Linux: `./discord-predictions-linux-amd64 -f [file]`
