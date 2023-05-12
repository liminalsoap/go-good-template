package main

import (
	"log"
	"todo/config"
	_ "todo/docs"
	"todo/internal/app"
)

//	@title			ToDo API
//	@version		1.0
//	@description	This is a swagger specification todo api.
//	@contact.name	fastwalker
//	@contact.url	http://www.github.com/iceforsik228
//	@contact.email	stchikhichin@gmail.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8080
//	@BasePath		/v1
func main() {
	log.Println("Start")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}
	app.Run(cfg)
}
