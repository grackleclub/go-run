# example

This example program runs for `$GORUN_TIMEOUT` seconds, and is used for testing [`go-run`](../go-run). It prints the contents of [static/test.txt](../static/test.txt) every second, but is relaunced if the contents of the file changes.

A simple way to trigger a file change is by overwriting or appending data to a file:
```sh
echo "I was updated at $(date)" > ./example/static/test.txt
```

## ffmpeg
Screen capture converted to gif with `ffmpeg`:
```
ffmpeg -i ~/Videos/Screencasts/example.webm -vf "fps=30,scale=1080:-1:flags=lanczos" -c:v gif example.gif
```
