FROM scratch

COPY . .

ENTRYPOINT #!/bin/sh

CMD ["/slackBot"]
