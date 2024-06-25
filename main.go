package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Block struct {
	Id int
}

type Count struct {
	Count int
}

type Contact struct {
	Name  string
	Email string
}

func NewContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func NewData() Data {
	return Data{
		Contacts: []Contact{
			NewContact("Juan", "juan@example.com"),
			NewContact("Pedro", "pedro@example.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: NewData(),
		Form: NewFormData(),
	}
}

type Blocks struct {
	Start  int
	Next   int
	More   bool
	Blocks []Block
}

func main() {
	e := echo.New()
	e.Renderer = NewTemplates()
	e.Use(middleware.Logger())

	count := Count{Count: 0}
	// data := NewData()
	page := newPage()

	e.GET("/", func(c echo.Context) error {
		// count.Count++
		return c.Render(200, "index", page)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.hasEmail(email) {
			formData := NewFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email

			formData.Errors["email"] = "Email already exists"

			return c.Render(400, "form", formData)
		}

		page.Data.Contacts = append(page.Data.Contacts, NewContact(name, email))
		return c.Render(200, "display", page)
	})

	e.GET("/blocks", func(c echo.Context) error {
		startStr := c.QueryParam("start")
		start, err := strconv.Atoi(startStr)
		if err != nil {
			start = 0
		}

		blocks := []Block{}
		for i := start; i < start+10; i++ {
			blocks = append(blocks, Block{Id: i})
		}

		template := "blocks"
		if start == 0 {
			template = "blocks-index"
		}
		return c.Render(http.StatusOK, template, Blocks{
			Start:  start,
			Next:   start + 10,
			More:   start+10 < 100,
			Blocks: blocks,
		})
	})

	e.GET("/index", func(c echo.Context) error {
		startStr := c.QueryParam("start")
		start, err := strconv.Atoi(startStr)
		if err != nil {
			start = 0
		}

		blocks := []Block{}
		for i := start; i < start+10; i++ {
			blocks = append(blocks, Block{Id: i})
		}

		template := "index"
		if start == 0 {
			template = "index"
		}
		return c.Render(http.StatusOK, template, Blocks{
			Start:  start,
			Next:   start + 10,
			More:   start+10 < 100,
			Blocks: blocks,
		})
	})

	e.Logger.Fatal(e.Start(":42069"))
}
