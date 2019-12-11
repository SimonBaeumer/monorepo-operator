FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y \
    ssh \
    git \
    iputils-ping \
    vim

WORKDIR /root

RUN mkdir .ssh && chmod 700 .ssh

# Copy private key
COPY id_rsa.private /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

# Copy known_hosts
COPY known_hosts /root/.ssh/known_hosts
RUN chmod 644 /root/.ssh/known_hosts

# Set git user configuration
RUN git config --global user.email "you@example.com"
RUN git config --global user.name "Your Name"

CMD [ "/bin/bash", "-c", " while :; do echo 'Hit CTRL+C'; sleep 1; done"]
