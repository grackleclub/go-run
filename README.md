# go-run
`go-run` is a simple bash script that monitors for file change and reloads automatically.

While live-reload tools in the code editor work great for rendering plain `html`+`css`, Go projects utilizing templates require tedious manual reloading. `go-run` automates this.

[![Shellcheck](https://github.com/grackleclub/go-run/actions/workflows/shellcheck.yml/badge.svg)](https://github.com/grackleclub/go-run/actions/workflows/shellcheck.yml) [![Release](https://github.com/grackleclub/go-run/actions/workflows/release.yml/badge.svg)](https://github.com/grackleclub/go-run/actions/workflows/release.yml)

## Features
- accepts arbitrary arguments and passes them through
- lists detected file changes
- stops upon program termination or signal interrupt
- preserves exit codes in all scenarios
- builds outside the working tree, leaving no artifact behind

## Getting Started on Linux
Copy the file and give it permission to run:
```sh
curl -o go-run https://raw.githubusercontent.com/grackleclub/go-run/refs/heads/main/go-run
```

Make the script executable:
```sh
chmod +x go-run
```

Make the script globally executable (optional):
```sh
sudo mv go-run /usr/local/bin
```

Verify installation:
```sh
go-run version
```

> [!TIP]
> `go-run` can be scoped to the project, and optimally kept in a `bin` directory, consolidated with other tools, so that it may be run as `bin/go-run`. Use the optional step of moving the script to `/usr/local/bin` to make `go-run` directly executable from anywhere.

## Use with mise
[mise](https://mise.jdx.dev) can manage go-run as a tool.

```toml
[tools]
go = "1.25"
sqlc = "latest"
"ubi:grackleclub/go-run" = "1.2.0"

[tasks.sqlc]
description = "generate query code"
dir = "db"
run = "sqlc generate"

[tasks.dev]
description = "start dev server with db; reload on file changes"
depends = ["sqlc"]
run = "go-run ./cmd/cli -v"
```

`mise run dev` now regenerates, builds, serves, and restarts on every save.
Arguments after the target directory reach your program, so
`go-run . -port 9000` is the task equivalent of `go run . -port 9000`.

> [!NOTE]
> .gitignored files are also ignored by `go-run`

## Demo and Testing Options
Demo the project using the [example](./example/) module:
![example-demonstration](./gifs/example.gif)

## Updating
Run `go-run update` to update:
![example-update](./gifs/update.gif)

## Feedback
😎 Open a [pull request](https://github.com/grackleclub/go-run/pulls)!
