package frontend

type FrontEnd = Client

func GetFrontend() *FrontEnd {
	return new(FrontEnd)
}
