FROM scratch
COPY shruti-cron /
ENTRYPOINT ["/shruti-cron"]
