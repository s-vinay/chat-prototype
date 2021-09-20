# Manages Server Functionality of Chat Server

###### For local build run
1. _GO_POST_PROCESS_FILE="gofmt -s -w" java -jar openapi-generator-cli-5.2.0.jar generate -i apischema/chatobjects.json -g go -o pkg/apiobjects --additional-properties packageName=apiobjects --additional-properties isGoSubmodule=false --global-property models --enable-post-process-file_