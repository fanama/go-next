# Backend in go + frontend Next.js

## prerequise

- nodejs :
  - > sudo apt-get install build-essential nodejs

- go

 - linux

    - > curl -O https://storage.googleapis.com/golang/go1.15.4.linux-amd64.tar.gz

    - > sudo tar -C /usr/local -xzf go1.15.4.linux-amd64.tar.gz

    - > export PATH=$PATH:/usr/local/go/bin


## Frontend

### START
 - > cd frontend
 - > npm install

### DEV

- run the app
    > npm run dev

### BUILD
  - linux :
    - > npm run build


## Server

### START
 get dependencies 
    > go get 

### DEV



- run the app
    > go run main.go

### BUILD
  - linux :
    - > go build -o server
