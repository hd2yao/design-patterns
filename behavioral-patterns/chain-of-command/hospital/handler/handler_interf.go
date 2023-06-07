package handler

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
