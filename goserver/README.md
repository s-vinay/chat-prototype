# Manages Server Functionality of Chat Server

_Runs on wss port: 8081_

###### For local build run(Without Docker)

1. _GO_POST_PROCESS_FILE="gofmt -s -w" java -jar openapi-generator-cli-5.2.0.jar generate -i apischema/chatobjects.json
   -g go -o pkg/apiobjects --additional-properties packageName=apiobjects --additional-properties isGoSubmodule=false
   --global-property models --enable-post-process-file_
2. go build .
3. ./goserver

###### For Docker Build

1. docker build -t gochatserver:1.0 --network=host ./

###### For Docker run

1. docker run -d -p 8081:8081 --name=chatserver gochatserver:1.0

###### Example wss client json templates

#### LastSeen

--->
`{
"service": "chat",
"messageType": "LastSeen",
"user_id": "1234",
"last_ts": 0,
"sync_id": "axpjy17661"
}`

<--- {
`"service": "chat",
"messageType": "LastSeen",
"user_id": "1234",
"last_ts": 123456789,
"sync_id": "axpjy17661"
}`