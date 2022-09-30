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
- [X] image support
- [X] video support
- [X] add a config and a config reader
- [ ] add a control pannel w/ login
- [ ] abllity to add new files using control pannel
- [ ] edit source code and add new lines using control pannel
- [ ] check for updates whilst using control pannel
- [ ] add mulit threading so the service doesn't get over run
- [ ] make https secure
- [ ] secure the program to patch any venruabiliyies that pop-up

# Update Info
the recent update info about go-web!  
. simple fixes to help speed up and use less proccesing when runing  
. simple typo fixes (there will lot's more to come)  
. upgraded all go-web functions to use http.ServFile instead of the method before  
. added more documentation to help developers use it  
. added the basic outline of the admin pannels  
. fixed a bug with css in admin html files  
