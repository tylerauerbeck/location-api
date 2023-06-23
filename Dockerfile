FROM gcr.io/distroless/static

# Copy the binary that goreleaser built
COPY location-api /location-api

# Run the web service on container startup.
ENTRYPOINT ["/location-api"]
CMD ["serve"]
