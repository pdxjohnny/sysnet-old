FROM busybox
ADD ./default_app_name_linux-amd64 /app
CMD ["/app"]
