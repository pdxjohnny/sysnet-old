FROM busybox
ADD ./build/default_app_name_linux-amd64 /app
CMD ["/app"]

