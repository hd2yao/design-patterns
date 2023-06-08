package server

type Server interface {
	HandlerRequest(string, string) (int, string)
}
