package service

const (
	HelloServiceName = "my-hello-service"

	EchoOperationName  = "echo"
	HelloOperationName = "say-hello"

	EN Language = "en"
	FR Language = "fr"
	DE Language = "de"
	ES Language = "es"
	TR Language = "tr"
)

type (
	EchoInput struct {
		Message string
	}
	EchoOutput EchoInput

	Language   string
	HelloInput struct {
		Name     string
		Language Language
	}
	HelloOutput struct {
		Message string
	}
)
