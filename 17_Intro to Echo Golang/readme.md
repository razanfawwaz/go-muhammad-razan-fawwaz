# (17) Intro to Echo Golang

## Echo
Echo merupakan salah satu framework dari bahasa GO yang populer untuk membuat API. Echo memiliki fitur yang cukup lengkap dan mudah digunakan. Echo juga memiliki dokumentasi yang lengkap dan mudah dipahami.

## Instalasi
Untuk menginstall Echo, kita bisa menggunakan perintah berikut:
```bash
go get -u github.com/labstack/echo/...
```

## Menjalankan Server
Untuk menjalankan server, kita bisa menggunakan perintah berikut:
```go
func main(){
	e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":8080"))
}
```
