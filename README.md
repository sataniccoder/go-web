# go-web
a website hosting service in golang
so far you can't add any pages that reqiure server based action sush as login pages, this doesn't mean you can't add javascrpit to your page to make it reactive
to start it jusr edit the index.html and the main.css to what you would like and then run the program!

# run
simple run `go run go-web.go` or to compile ans run it use `go build go-web.go && ./go-web`

# TO-DO
- [X] basic hosting service
- [X] support multiple files and live updates
- [ ] image support
- [ ] video support
- [ ] add a config and a config reader
- [ ] add support for scripting so you can add login pages etc...
- [ ] add mulit threading so the service doesn't get over run
- [ ] secure the program to patch any venruabiliyies that pop-up
