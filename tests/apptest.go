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
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "ADMIN123")
	t.NewTestRequest(req).MakeRequest()
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestMenuAPIWithoutToken() {
	println("test menu ...")
	req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//req.Header.Set("Authorization", "ADMIN123")
	t.NewTestRequest(req).MakeRequest()
	t.AssertStatus(400)
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestMenuAPIWithUnauthorizedToken() {
	println("test menu ...")
	req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "AMONYMOUS123")
	t.NewTestRequest(req).MakeRequest()
	t.AssertStatus(401)
	//t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestMenuAPIWithWrongToken() {
	println("test menu ...")
	req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "A123")
	t.NewTestRequest(req).MakeRequest()
	t.AssertStatus(401)
	//t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestClientApiWorks() {
	println("test menu ...")
	t.Get("/clients")
	//req, _ := http.NewRequest("GET", t.BaseUrl()+"/menu", nil)
	////req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//req.Header.Set("Authorization", "A123")
	//t.NewTestRequest(req).MakeRequest()
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestCollaboratersApiWorks() {
	println("test menu ...")
	t.Get("/collaborateurs")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestCollaboratersRelocationApiWorks() {
	println("test menu ...")
	t.Get("/collaborateurs?k=Relocation")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) TestCollaboratersImmigrationApiWorks() {
	println("test menu ...")
	t.Get("/collaborateurs?k=Immigration")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}
