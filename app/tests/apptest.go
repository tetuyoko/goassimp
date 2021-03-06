package tests

import (
	"github.com/revel/revel/testing"
	"net/url"
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

func (t *ApplicationTest) TestThatRedisAPIWorks() {
	t.Get("/redis")
	t.AssertOk()
	t.AssertContains("Success")
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *ApplicationTest) TestThatRedisPingAPIWorks() {
	t.Get("/redis/ping")
	t.AssertOk()
	t.AssertContains("Success")
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *ApplicationTest) TestThatRedisSetGetAPIWorks() {
	t.PostForm("/redis/set", url.Values{
		"key": {"hage"},
		"val": {"1"},
	})
	t.AssertOk()
	t.AssertContains("OK")
	t.AssertContentType("application/json; charset=utf-8")

	t.Get("/redis/get/hage")

	t.AssertOk()
	t.AssertContains("1")
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *ApplicationTest) TestThatRedisHSetHGetAPIWorks() {
	t.PostForm("/redis/hset", url.Values{
		"key":   {"hsetkey"},
		"field": {"hsetfield"},
		"val":   {"1"},
	})
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	t.Get("/redis/hget/hsetkey")

	t.AssertOk()
	t.AssertContains("hsetfield")
	t.AssertContains("1")
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *ApplicationTest) After() {
	println("Tear down")
}
