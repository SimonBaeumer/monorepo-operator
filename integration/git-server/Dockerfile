FROM jkarlos/git-server-docker

COPY create-repo.sh /bin/create-repo.sh
RUN chmod +x /bin/create-repo.sh

RUN cd /git-server/repos && create-repo.sh monorepo.git
RUN cd /git-server/repos && create-repo.sh repo01.git
RUN cd /git-server/repos && create-repo.sh repo02.git

EXPOSE 22

CMD ["sh", "start.sh"]
