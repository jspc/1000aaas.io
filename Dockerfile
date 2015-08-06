FROM scratch
COPY main /main
EXPOSE 80
CMD ["/main", "-bind=:80"]
