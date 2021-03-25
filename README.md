
## go-middleware

Middleware for Go functions.

Works perfectly with `socket.io` event handlers – 
https://github.com/googollee/go-socket.io

### Installation

`go get github.com/fakundo/go-middleware`

### Usage

Create middleware

```go
import (
  middleware "github.com/fakundo/go-middleware"
  socketio "github.com/googollee/go-socket.io"
)

var requireAuth = middleware.Create(func(s socketio.Conn, next func()) {
  if authorized(s) {
    next()
  } else {
    s.emit("error", AuthError)
  }
})
```

Use it

```go
import socketio "github.com/googollee/go-socket.io"

io, _ = socketio.NewServer(nil)

io.OnEvent("/", "some-event", requireAuth(func(s socketio.Conn) {
  // some event handler code
}))


io.OnEvent("/", "another-event", requireAuth(func(s socketio.Conn, arg string) {
  // another event handler code
}))
```

### Middleware composition

```go
import middleware "github.com/fakundo/go-middleware"

...

io.OnEvent("/", "event", middleware.Use(someMiddleware, requireAuth, func(s socketio.Conn) {
  // event handler code
}))
```

### About middleware behaviour

- *currently* the decorated function type must be `interface` (exactly like `socket.io` event handlers)
- the origin function and middleware can have different number of arguments, middleware will only take required arguments, and the last argument is always `next`
- `next()` calls the origin function
- the origin function may return something, so `next()` will return that too

### License

MIT
