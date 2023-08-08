<h1 align="center">SATURNUS</h1>
<img src="https://i.pinimg.com/originals/02/01/1e/02011ec8554277b8c70bf22fb192123c.gif" />

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif"></p>


## Description
Example and Template about Transcoding HTTP to GRPC with Golang and Mongodb. Fyi, GRPC is modern open source high performance Remote Procedure (RPC) framework that can run in any environtment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services. gRPC is based from protocol buffer which is an input schema that will be carried out by the client or user.

In this project, I provide examples or templates about http to gRPC transcoding. Actually, gRPC itself is based on HTTP v2 but is not published. So, Google Cloud Endpoint itself supports transcoding http to grpc, more or less the schema is like this.

<div>
    <img src="https://poornimanayar.co.uk/media/4wjjytij/image.png" />
</div>

## Endpoint
- GRPC
<table>
    <tr>
        <th>Endpoint</th>
        <th>Description</th>
        <th>Method</th>
    </tr>
    <tr>
        <th>TodoService/CreateProduct</th>
        <th>To create todo</th>
        <th>POST</th>
    </tr>
    <tr>
        <th>TodoService/FindById</th>
        <th>To get todo by id</th>
        <th>GET</th>
    </tr>
    <tr>
        <th>TodoService/Update</th>
        <th>To update todo by id</th>
        <th>PUT</th>
    </tr>
    <tr>
        <th>TodoService/Delete</th>
        <th>To delete todo by id</th>
        <th>DELETE</th>
    </tr>
    <tr>
        <th>TodoService/FindAll</th>
        <th>To sell all todo</th>
        <th>GET</th>
    </tr>
</table>

- REST
<table>
    <tr>
        <th>Endpoint</th>
        <th>Description</th>
        <th>Method</th>
    </tr>
    <tr>
        <th>/v1/todo</th>
        <th>To create todo</th>
        <th>POST</th>
    </tr>
    <tr>
        <th>/v1/todo/{id}</th>
        <th>To get todo by id</th>
        <th>GET</th>
    </tr>
    <tr>
        <th>/v1/todo/{id}</th>
        <th>To update todo by id</th>
        <th>PUT</th>
    </tr>
    <tr>
        <th>/v1/todo/{id}</th>
        <th>To delete todo by id</th>
        <th>DELETE</th>
    </tr>
    <tr>
        <th>/v1/todo</th>
        <th>To sell all todo</th>
        <th>GET</th>
    </tr>
</table>

## Usage
### 1.1 Usage with Golang
Run this command to running download all module
```go
go mod tidy
```
then running file `server.go` to run the server
```go
go run server.go
```
and the output will be like this
```
2023/08/08 14:26:20 GRPC Server running on localhost:40000
2023/08/08 14:26:20 HTTP Server running on localhost:8080
```

### 1.2 Usage with Docker
Run this command to download docker-compose (this command for linux)
```bash
sudo apt install docker-compose -y
```
then running this command
```
docker-compose up -d
```
and service will be running on port `40000 to GRPC` and `8080 to REST`

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif"></p>

> For example schema to input in GRPC you can see in [`example`](https://github.com/rulanugrh/saturnus/tree/master/example) folder
