package design

import . "goa.design/goa/v3/dsl"

var _ = API("hello-world", func() {
	Title("The hello-world API")
	Description("A simple API that says hello")
	Version("0.1.0")

	Server("hello-world", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:5000")
		})
	})
})

var _ = Service("hello", func() {
	Description("The hello service, it says hello")

	Method("say-hello", func() {
		Payload(func() {
			Field(1, "name", String, "The name to say hello to")
			Required("name")
		})

		Result(String)

		HTTP(func() {
			GET("/hello/{name}")
		})

		GRPC(func() {
			// for now empty
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
