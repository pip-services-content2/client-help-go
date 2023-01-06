package build

import (
	clients1 "github.com/pip-services-content2/client-help-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type HelpServiceFactory struct {
	*cbuild.Factory
}

func NewHelpServiceFactory() *HelpServiceFactory {
	c := &HelpServiceFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-help", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-help", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-help", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewHelpNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewHelpMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewHelpCommandableHttpClientV1)

	return c
}
