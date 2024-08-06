# {{ .appName }}

{{ .appDescription }}

# Demo Application

```sh
$> {{ .appName }} -h
{{ .appName }} project template demo application

Usage:
  {{ .appName }} [flags]
  {{ .appName }} [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  example     example subcommand which adds or multiplies two given integers
  help        Help about any command
  version     golang-cli-template version

Flags:
  -h, --help   help for golang-cli-template

Use "{{ .appName }} [command] --help" for more information about a command.
```

```sh
$> {{ .appName }} example 2 5 --add
7

$> {{ .appName }} example 2 5 --multiply
10
```

# Makefile Targets
```sh
$> make
bootstrap                      install build deps
build                          build golang binary
clean                          clean up environment
cover                          display test coverage
docker-build                   dockerize golang application
fmt                            format go files
help                           list makefile targets
install                        install golang binary
lint                           lint go files
pre-commit                     run pre-commit hooks
run                            run the app
test                           display test coverage
```

# Contribute
If you find issues in that setup or have some nice features / improvements, I would welcome an issue or a PR :)

# Original Credits
This template is based on the boilerplate from https://github.com/FalcoSuessgott/golang-cli-template