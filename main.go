package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/prometheus/alertmanager/template"
	"net/http"
	"time"
)

type RenderRequest struct {
	Template string `json:"template"`
	Data     string `json:"data"`
}

func main() {
	t, err := template.FromGlobs()
	if err != nil {
		panic(err)
	}
	d := template.Data{
		Receiver: "test_receiver",
		Status:   "firing",
		Alerts: []template.Alert{{
			Status: "Firing",
			Labels: map[string]string{
				"service": "test",
				"env":     "prod",
				"other":   "value",
			},
			Annotations: map[string]string{
				"description":      "Error has happened",
				"extra_annotation": "My extra annotation",
			},
			StartsAt: time.Now(),
			EndsAt:   time.Now().Add(72 * time.Hour),
		}},
		GroupLabels: template.KV{
			"service": "test",
			"env":     "prod",
		},
		CommonLabels: template.KV{
			"service": "test",
			"env":     "prod",
			"other":   "value",
		},
		CommonAnnotations: template.KV{
			"description": "Error has happened",
		},
		ExternalURL: "http://google.com",
	}
	b, _ := json.MarshalIndent(d, "", "   ")
	fmt.Println(string(b))

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "/Users/achaplianka/Dvelop/Personal/Alertmanager-Template-Preview/web")
	app.Post("/render", func(c *fiber.Ctx) {
		rr := new(RenderRequest)
		if err := c.BodyParser(rr); err != nil {
			_ = c.JSON(fiber.Map{"Error": err.Error()})
			return
		}
		data := new(template.Data)
		err := json.Unmarshal([]byte(rr.Data), data)
		if err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
		}

		res, err := t.ExecuteTextString(rr.Template, data)
		if err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
		c.Send(res)
	})
	err = app.Listen("0.0.0.0:3000")
	if err != nil {
		panic(err)
	}
}
