package template

import (
	"fmt"
	"io"

	"github.com/icrowley/fake"
	"github.com/valyala/fasttemplate"
)

func Build(templateString string) string {
	t, err := fasttemplate.NewTemplate(templateString, "{", "}")
	if err != nil {
		fmt.Println("error when parseing template: %s", err)
	}
	s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		switch tag {
		case "firstname":
			return w.Write([]byte(fake.MaleFirstName()))
		case "lastname":
			return w.Write([]byte(fake.LastName()))
		case "age":
			return w.Write([]byte(fake.DigitsN(2)))
		default:
			return w.Write([]byte("unknown tag"))
		}
	})
	return s
}