# ska-start: managed-dokerfile + ska-replace-match:(?s).*
FROM ubuntu:latest

COPY {{ .appName }} /usr/bin/{{ .appName }}
ENTRYPOINT ["/usr/bin/{{ .appName }}"]

LABEL version="1.0" maintainer="{{ .maintainerName }} <{{ .maintainerEmail }}>"
# ska-end