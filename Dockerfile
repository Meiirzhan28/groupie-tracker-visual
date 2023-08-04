# Start the build step with the golang:alpine As builder image
# and name the step "builder"
# By default the stages have no names, 
# and you access them by their integer number, starting with 0 for the first FROM instruction
FROM golang:alpine As builder
# define the working directory
# the WORKDIR command specifies the working directory 
# if the directory we have specified does not exist, it will automatically create it
LABEL name="Groupie-tracker"
LABEL description="alem school project"
LABEL authors="raqazhet,zadilbay"
LABEL release-date="24.01.2022"
WORKDIR /app

# copy all files to the working directory
COPY . .

# create the .eh file of our product
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

# Start a new build step with the alpine image as the base
FROM alpine

# define the working directory
WORKDIR /app
# The COPY --from=0 line copies only the assembled artifact from the previous stage to this new stage. 
# Go SDK and any intermediate artifacts are left behind and not stored in the final image.
COPY --from=builder /app .
# this is port
EXPOSE 3000

CMD ["./app"]