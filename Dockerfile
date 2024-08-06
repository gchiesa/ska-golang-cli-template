FROM alpine:3.20
COPY {{ .appName }} /usr/bin/{{ .appName }}
ENTRYPOINT ["/usr/bin/{{ .appName }}"]