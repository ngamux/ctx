# ctx
Request-response utility for every HTTP context  

# Usage

```go
mux.Post("/", func(rw http.ResponseWriter, r *http.Request) error {
  c := ctx.New(rw, r)
  
  var data interface{}
  if err := c.GetJSON(&data); err != nil {
    return c.JSON(http.StatusBadRequest, ngamux.Map{
      "error": err.Error(),
    })
  }
  
  return c.JSON(http.StatusOK, data)
})
```

or use handler adaptor

```go
mux.Post("/", ctx.Handler(func(c *ctx.Context) error {  
  var data interface{}
  if err := c.GetJSON(&data); err != nil {
    return c.JSON(http.StatusBadRequest, ngamux.Map{
      "error": err.Error(),
    })
  }
  
  return c.JSON(http.StatusOK, data)
}))
```
