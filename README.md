# Soundcloud Watcher

Tool to watch for any added/removed likes from a profile.

## Usage

All commands must be run with `--user <your_username>`

`SOUNDCLOUD_CLIENT_ID` must be present in your environment. You can get this ID by viewing your network requests when using soundcloud.

### Initialize a user

`scwatch init`

### Check for new state

`scwatch run`

### Display last state

`scwatch stat`

### Get info about a song

`scwatch song <id>`
