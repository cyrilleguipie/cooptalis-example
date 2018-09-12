package tests

import (
	"github.com/revel/revel/testing"
	"net/http"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestMenuAPIWorksWithAuthorizedToken() {
	println("test menu ...")
	req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	req.Header.Set("Authorization", "ADMIN123")
	t.NewTestRequest(req).MakeRequest()
	t.AssertOk()
}

func (t *AppTest) TestMenuAPIWithoutToken() {
	req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	t.NewTestRequest(req).MakeRequest()
	t.AssertStatus(400)
}

func (t *AppTest) TestMenuAPIWithUnauthorizedToken() {
	req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	req.Header.Set("Authorization", "ANONYMOUS123")
	t.NewTestRequest(req).MakeRequest()
	t.AssertStatus(401)
}

func (t *AppTest) TestMenuAPIWithWrongToken() {
	req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	req.Header.Set("Authorization", "A123")
	t.NewTestRequest(req).MakeRequest()
	t.AssertStatus(401)
}

func (t *AppTest) TestClientApiWorks() {
	t.Get("/clients")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestCollaboratersApiWorks() {
	t.Get("/collaborateurs")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestCollaboratersRelocationApiWorks() {
	t.Get("/collaborateurs?k=Relocation")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestCollaboratersImmigrationApiWorks() {
	t.Get("/collaborateurs?k=Immigration")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}
