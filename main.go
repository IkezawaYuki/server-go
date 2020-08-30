package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"net/http"
	"os"
)

func main(){

	_ = godotenv.Load()

	gomniauth.SetSecurityKey(os.Getenv("SECURITY_KEY"))
	gomniauth.WithProviders(
		google.New(
			os.Getenv("CLIENT_ID"),
			os.Getenv("SECRET_VALUE"),
			os.Getenv("REDIRECT_URL"),
			),
		)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())


	e.GET("/auth/callback", callback)
	e.GET("/v1/auth/login", login)

	// Routes
	e.GET("/", hello)
	e.GET("/clear", clear)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}


func login(c echo.Context) error{
	provider, err := gomniauth.Provider("google")
	if err != nil{
		panic(err)
	}
	loginURL, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil{
		panic(err)
	}
	return c.JSON(200, loginURL)
}

func callback(c echo.Context) error{
	provider, err := gomniauth.Provider("google")
	if err != nil{
		panic(err)
	}

	cred, err := provider.CompleteAuth(objx.MustFromURLQuery(c.QueryString()))
	if err != nil{
		panic(err)
	}

	user, err := provider.GetUser(cred)
	if err != nil{
		panic(err)
	}

	fmt.Println(user)
	return c.JSON(200, "YES!!!")
}

func clear(c echo.Context) error{
	return c.JSON(http.StatusOK, "clear!!!!!")
}