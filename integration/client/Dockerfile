FROM alpine/git:1.0.7

COPY start.sh /bin/start.sh
RUN chmod +x /bin/start.sh

RUN apk update
RUN apk add git

COPY ssh .ssh
RUN chmod 700 .ssh
RUN chmod 600 .ssh/id_rsa
RUN chmod 644 .ssh/id_rsa.pub

RUN git config --global user.email "test@example.com"
RUN git config --global user.name "John Doe"

ENTRYPOINT [ "start.sh" ]