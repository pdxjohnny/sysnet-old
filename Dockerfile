FROM busybox
ADD ./default_app_name_linux-386 /app
CMD ["/app"]
