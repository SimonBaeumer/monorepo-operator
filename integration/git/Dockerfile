FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install -y \
    openssh-server \
    ssh \
    git \
    iputils-ping \
    vim

WORKDIR /root
RUN mkdir .ssh && chmod 700 .ssh
RUN touch .ssh/authorized_keys && chmod 600 .ssh/authorized_keys

COPY id_rsa.pub /root/id_rsa.pub
RUN cat /root/id_rsa.pub >> .ssh/authorized_keys

RUN mkdir -p /srv/git

RUN mkdir -p /srv/git/monorepo.git && cd /srv/git/monorepo.git && git init --bare
RUN mkdir -p /srv/git/subtree1.git && cd /srv/git/subtree1.git && git init --bare
RUN mkdir -p /srv/git/subtree2.git && cd /srv/git/subtree2.git && git init --bare
RUN mkdir -p /srv/git/subtree3.git && cd /srv/git/subtree3.git && git init --bare

EXPOSE 22

CMD [ "/etc/init.d/ssh", "start", "-D"]
