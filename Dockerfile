FROM iron/go
WORKDIR /app
# Now just add the binary
ADD app /app/
EXPOSE 7777 
ENTRYPOINT ["./app"]
