## Skeleton

## Arsitektur

![Architecture](assets/imgs/architecture.png)

## Request Flow

![Request Flow](assets/imgs/flow.png)

![Request Explaination](assets/imgs/explain.png)

## Flow Pagination

![Pagination](assets/imgs/paginated.png)

## Flow Create/Update/Delete

![Create/Update/Delete](assets/imgs/create.png)

## Flow Get

![Get](assets/imgs/get.png)

## Tool

Saat ini masih menggunakan protoc-gen-go versi yang lama (v1.4.3)

```
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    github.com/golang/protobuf/protoc-gen-go
```

## Perintah

- Build Dependency Graph

```
go run cmd/dic/main.go
```

- Application

```
go run cmd/app/main.go
```

- Generator

```
go run cmd/generator/main.go
```

## Testing

```
$ go test ./... [-v]
```

## Contoh

- [Todo Application](https://github.com/crowdeco/skeleton-todo)

## Dokumentasi

- [Instalasi](docs/install.md)

- [Awal Memulai](docs/basic_usage.md)
