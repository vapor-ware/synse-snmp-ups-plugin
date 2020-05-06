#
# Builder Image
#
FROM vaporio/foundation:bionic as builder

RUN apt-get update \
 && apt-get install -y --no-install-recommends ca-certificates

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
