FROM busybox
ADD ./sysnet_linux-amd64 /app
CMD ["/app"]
