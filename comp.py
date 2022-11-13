import os

os.system("sudo docker build -t go-web .")
os.system("sudo docker run -p 8080:8080 -t go-web")