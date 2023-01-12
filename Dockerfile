FROM c4stus/raspberrypi:blecron

WORKDIR /data

RUN apt-get update -y && apt-get upgrade -y && apt-get install -yq --no-install-recommends \
    locales \
    systemd
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

RUN echo "en_US.UTF-8 UTF-8" > /etc/locale.gen && locale-gen
ENV LANG en_US.utf8
ENV TZ=Europe/Warsaw
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN wget https://dl.google.com/go/go1.19.4.linux-armv6l.tar.gz
RUN tar -C /usr/local/ -xzf go1.19.4.linux-armv6l.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" >> .profile

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

COPY ./src /data

CMD sh run.sh
