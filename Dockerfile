#
# Builder Image
#
FROM vaporio/golang:1.13 as builder

#
# Final Image
#
FROM vaporio/scratch-ish:1.0.0

LABEL org.label-schema.schema-version="1.0" \
      org.label-schema.name="vaporio/snmp-ups-plugin" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/synse-snmp-ups-plugin" \
      org.label-schema.vendor="Vapor IO"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy the executable.
COPY synse-snmp-ups-plugin ./plugin

ENTRYPOINT ["./plugin"]
