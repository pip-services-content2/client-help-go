package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-help-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type HelpCommandableHttpClientV1 struct {
	client  *version1.HelpCommandableHttpClientV1
	fixture *HelpClientFixtureV1
}

func newHelpCommandableHttpClientV1() *HelpCommandableHttpClientV1 {
	return &HelpCommandableHttpClientV1{}
}

func (c *HelpCommandableHttpClientV1) setup(t *testing.T) *HelpClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewHelpCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewHelpClientFixtureV1(c.client)

	return c.fixture
}

func (c *HelpCommandableHttpClientV1) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newHelpCommandableHttpClientV1()

	fixture := c.setup(t)
	fixture.TestTopicsCrudOperations(t)
	c.teardown(t)

	fixture = c.setup(t)
	fixture.TestArticlesCrudOperations(t)
	c.teardown(t)
}
