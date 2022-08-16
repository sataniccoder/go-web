# go-web
a static website hosting service in golang, this is mainly used for remote deployment, this means once you have set the program up with a user and password
you can accses the control pannel from your browser, from there you can add new pages, update the programs source code and much more!

# INFO
so far you can't add any pages that reqiure server based action sush as login pages, this doesn't mean you can't add javascrpit to your page to make it
reactive
to start it jusr edit the index.html and the main.css to what you would like and then run the program!

# run
simple run `go run go-web.go` or to compile ans run it use `go build go-web.go && ./go-web`

# TO-DO
- [X] basic hosting service
- [X] support multiple files and live updates
- [ ] image support
- [ ] video support
- [ ] add a config and a config reader
- [ ] add a control pannel w/ login
- [ ] add mulit threading so the service doesn't get over run
- [ ] secure the program to patch any venruabiliyies that pop-up
