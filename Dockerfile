# ska-start:new-base + ska-replace-match:FROM .*
FROM ubuntu:latest
# ska-end

COPY {{ .appName }} /usr/bin/{{ .appName }}
ENTRYPOINT ["/usr/bin/{{ .appName }}"]

# ska-start:new-central-managed-labels + ska-inject-after:@end
LABEL version="1.0" maintainer="{{ .maintainerName }} <{{ .maintainerEmail }}>"
# ska-end