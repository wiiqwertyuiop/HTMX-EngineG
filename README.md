# Go Chi MVC + HTMX Mini-Server

Based off of ideas from NestJS and using MVC principles, this mini server able to act as a standalone front+backend. Featuring

1) Serving static files
2) Wrapper around go-chi for easy routing. You can now more directly use the idea of "controllers", with easy pathing.
3) Inadverntly fell into the MVC pattern, making things feel a bit more natural.
4) All in one front end and back end. HTMX is used to make request to the data endpoints.
5) Controllers are extensible for different types of middleware, specific to that route.
6) And finally... good old fashion Go templates of course!

WIP, but looking forward to experimenting and using this more in the future.

I also have a typescript version here: https://github.com/wiiqwertyuiop/Express-Engine

## How to run

Simply do a `go build src/main.go`, or if you have [modd](https://github.com/cortesi/modd) installed, you can do `modd` to start the server.
