package main

import (
	"fmt"
	"net/http"
)

// "fmt"
// "strings"

// "github.com/gofiber/fiber/v2"

func main() {
	fmt.Println("From main package")

	res, err := http.Get("http://users:7070/users")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)

	// app := fiber.New()

	// app.Use("/", func(c *fiber.Ctx) error {

	// 	URI := c.Context().URI()

	// 	uriSplit := strings.Split(string(URI.FullURI()), "/")

	// 	// method := c.Context().Method()

	// 	// headers := c.Context().Request.Header

	// 	// body = c.Context().Request.Body()

	// 	fmt.Println(uriSplit[3])

	// 	type service struct {
	// 		Name   string `json:"name"`
	// 		Domain string `json:"domain"`
	// 	}

	// 	var services []service
	// 	services = append(services, service{Name: "users", Domain: "http://127:0.0.1:5050"})
	// 	services = append(services, service{Name: "home", Domain: "http://127:0.0.1:6000"})
	// 	services = append(services, service{Name: "cart", Domain: "http://127:0.0.1:6060"})
	// 	services = append(services, service{Name: "services", Domain: "http://127:0.0.1:7000"})
	// 	services = append(services, service{Name: "products", Domain: "http://127:0.0.1:7070"})
	// 	services = append(services, service{Name: "order-management", Domain: "http://127:0.0.1:8000"})

	// 	var requestedService service

	// 	for _, item := range services {
	// 		if item.Name == uriSplit[3] {
	// 			requestedService = item
	// 		}
	// 	}

	// 	if requestedService.Name == "" {
	// 		return c.Status(fiber.StatusBadRequest).JSON("service not available")
	// 	}

	// 	fmt.Println(requestedService)
	// 	return c.JSON(requestedService)

	// })

	// app.Listen(":3000")
}
