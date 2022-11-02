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
- [X] add a control pannel w/ login
- [X] add mulit threading so the service doesn't get over run
- [ ] abllity to add new files using control pannel
- [ ] check for updates whilst using control pannel
- [ ] make https secure
- [ ] secure the program to patch any venruabiliyies that pop-up

# Update Info
the recent update info about go-web!  
. simple fixes to help speed up and use less proccesing when runing  
. simple typo fixes (there will lot's more to come)  
. you can now download html files
. the ability to genrate a new admin pin (terminal warns you that another one has been made)

# Coming up
. ablility to upload files  
. ability to restart the program from the admin pannel (if you make new changes too it, it would recompile and excute it)  
. ability to easly create post pages (trying to make a scriping language for it [PROBABLY WILL CHANGE AND WON'T BE USED])


