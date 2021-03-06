package http

import (
	"strings"
	"testing"

	"github.com/go-martini/martini"

	"github.com/github/orchestrator/go/config"
	"github.com/outbrain/golib/log"
	test "github.com/outbrain/golib/tests"
)

func init() {
	config.Config.HostnameResolveMethod = "none"
	log.SetLevel(log.ERROR)
}

func TestGetSynonymPath(t *testing.T) {
	api := HttpAPI{}

	{
		path := "relocate-slaves"
		synonym := api.getSynonymPath(path)
		test.S(t).ExpectEquals(synonym, "relocate-replicas")
	}
	{
		path := "relocate-slaves/:host/:port"
		synonym := api.getSynonymPath(path)
		test.S(t).ExpectEquals(synonym, "relocate-replicas/:host/:port")
	}
}

func TestKnownPaths(t *testing.T) {
	m := martini.Classic()
	api := HttpAPI{}

	api.RegisterRequests(m)

	pathsMap := make(map[string]bool)
	for _, path := range registeredPaths {
		pathBase := strings.Split(path, "/")[0]
		pathsMap[pathBase] = true
	}
	test.S(t).ExpectTrue(pathsMap["health"])
	test.S(t).ExpectTrue(pathsMap["lb-check"])
	test.S(t).ExpectTrue(pathsMap["relocate"])
	test.S(t).ExpectTrue(pathsMap["relocate-slaves"])

	for path, synonym := range apiSynonyms {
		test.S(t).ExpectTrue(pathsMap[path])
		test.S(t).ExpectTrue(pathsMap[synonym])
	}
}
