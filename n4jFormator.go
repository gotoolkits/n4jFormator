package main

import (
	encoded "github.com/gotoolkits/n4jFormator/encoded"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

var (
	sHost = "8081"

	cypherSearchNodes         = `MATCH (n) RETURN n`
	cypherSearchRelationships = `MATCH p=()-->() RETURN p`
)

func main() {

	e := echo.New()
	e.HideBanner = true

	ne4j, err := encoded.InitNe4j()
	if err != nil {
		log.Fatalln("neo4j server connect failed.", err)
	}

	log.Println("⇨ neo4j server connect successfully.")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/api/getNodes", ne4j.FnFormatNodes)
	e.GET("/api/getPaths", ne4j.FnFormatRelationships)

	log.Println("⇨ http server starting on ", ":"+sHost)

	e.Logger.Fatal(e.Start(":" + sHost))

}
