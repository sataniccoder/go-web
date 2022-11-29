import os


print("""
[!!] This is a simple script to setup a go web server [!!]


1) compile and run go-web
2) setup config file
3) setup https (optional)
4) Build the docker image && run the container
5) exit
""")

while True:
    while True:
        path = input("Enter the path to the go-web directory: ")
        # verify that main.go is in the directory
        if os.path.isfile(path + "/main.go"):
            break
        else:
            print("[!!] main.go not found in the directory [!!]")

    op = input("> ")
    if op == "1":
        os.system("go build "+path+"main.go && ./"+path+"main")
    elif op == "2":
        
        port = input("Enter the port to run the server on: ")
        ver = input("enter verbose mode (true/false): ")
        if ver == "true":
            ver = "1"
        else:
            ver = "0"

        verb = """
# these comments will be ignored by the program, WARNING: DO NOT EDIT THE WAY THIS FILE IS SET OUT!
# if you add another config header you will have too edit the list size of the veriable 'data_list' on config.go
 

# port the server will be served on, (default: 8080) non-root ports are recomended :^)
port="""+port+"""


# diff templates folder, (default: templates/) can be: /home/user/folders_maby/templates/
folder=templates/

# verbose, this is set if the user want's too see what pages are being accses (not recmoneded for preduction due to how bugg the output could get)
# it would also hide errors, as they will be printed no matter what, this helps with secure development
# 0=false, 1=true, default: 1)
verbose="""+ver+"""
        
        """
        with open(path+"config.conf", "w") as f:
            f.write(verb)
            f.close()
            
        print("[!!] config file created [!!]")


    elif op == "3":
        c = input("enter country code (eg: GB, US): ")
        s = input("enter state (eg: London, California): ")
        l = input("enter location (eg: London, San Francisco): ")
        o = input("enter organisation (eg: Google, Apple): ")
        u = input("enter unit (eg: IT, Security): ")
        cn = input("enter common name (eg: google.com, apple.com): ")

        dns_1 = input("enter dns 1 (eg: localhost, 127.0.0.1): ")
        dns_2 = input("enter dns 2 (eg: localhost, 127.0.0.1): ")
        ip = input("enter ip (eg: localhost, 127.0.0.1): ")
        cet = """
[ req ]
prompt = no
default_bits = 4096
distinguished_name = req_distinguished_name
req_extensions = req_ext

[ req_distinguished_name ]
C="""+c+"""
ST="""+s+"""
L="""+l+"""
O="""+o+"""
OU="""+u+"""
CN="""+cn+"""

[ req_ext ]
subjectAltName = @alt_names

[alt_names]
DNS.1 = """+dns_1+"""
DNS.2 = """+dns_2+"""
IP.1 = """+ip+"""

        """
        with open(path+"certs/cet.conf", "w") as f:
            f.write(cet)
        print("[!!] cet.conf created [!!]")
        os.system("openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout "+path+"certs/key.pem -out "+path+"certs/cert.pem -config "+path+"certs/cet.conf")
        print("[!!] key.pem and cert.pem created [!!]")
        
    elif op == "4":
        os.system("sudo docker build -t go-web .")
        os.system("sudo docker run -p 8080:8080 -t go-web")
    elif op == "5":
        break
    else:
        print("[!!] invalid option [!!]")