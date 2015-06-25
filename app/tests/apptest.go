package tests

import (
	"github.com/revel/revel/testing"
)

type ApplicationTest struct {
	testing.TestSuite
}

func (t *ApplicationTest) Before() {
	println("Set up")
}

func (t *ApplicationTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *ApplicationTest) TestThatRedisPingAPIWorks() {
	t.Get("/redis/ping")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *ApplicationTest) TestThatRedisAPIWorks() {
	t.Get("/redis")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *ApplicationTest) After() {
	println("Tear down")
}
