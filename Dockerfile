FROM iron/go
WORKDIR /app
ADD app /app/
EXPOSE 7777 
ENTRYPOINT ["./app"]
