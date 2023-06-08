package main

import (
	"design-patterns/structural-patterns/proxy/servers/proxy"
	"fmt"
)

func main() {
	nginxServer := proxy.NewNginxServer()
	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCode, body := nginxServer.Application.HandlerRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.Application.HandlerRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.Application.HandlerRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.Application.HandlerRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = nginxServer.Application.HandlerRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
}
