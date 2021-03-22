from golang

label maintainer="mengxianfan" 
label version="1.0"

#run pip install flask

copy app.go /app/

workdir /app

expose 5000

cmd go run app.go
