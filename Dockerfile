FROM alpine

LABEL MAINTAINER=PerwiraMedia

# Working Directory
WORKDIR /home/pm-boiler-code

# Copy in the source
COPY _bin/pm-boiler-code ./pm-boiler-code

# Make shell scripts executable
RUN chmod +x ./pm-boiler-code

# Running
CMD ["./pm-boiler-code"]