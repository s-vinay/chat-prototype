FROM ubuntu:bionic-20210827

RUN apt-get update
RUN apt install -y default-jre

RUN wget https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz \
	&& tar -xvf go1.15.2.linux-amd64.tar.gz \
	&& mv go /usr/local \
	&& wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.5/protoc-3.15.5-linux-x86_64.zip \
	&& unzip protoc-3.15.5-linux-x86_64.zip -d $HOME/.local

ADD . /appsrc/

RUN wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.2.0/openapi-generator-cli-5.2.0.jar \
      -O /appsrc/openapi-generator-cli.jar

ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH
ENV GO111MODULE=on
ENV PATH=$PATH:/root/go/bin:/root/.local/bin

#delve for debug support
RUN git clone https://github.com/go-delve/delve \
    && cd delve \
	&& go install github.com/go-delve/delve/cmd/dlv

RUN cd /appsrc/ \
    && GO_POST_PROCESS_FILE="gofmt -s -w" java -jar openapi-generator-cli.jar generate -i apischema/chatobjects.json -g go -o pkg/apiobjects --additional-properties packageName=apiobjects --additional-properties isGoSubmodule=false --global-property models --enable-post-process-file
    && go build .

CMD ["/appsrc/goserver"]