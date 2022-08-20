## 2FA Backend

Backend application for [2FA Mobile](https://github.com/peterdee/2fa-mobile)

Stack: [Golang](https://golang.org), [Fiber](https://gofiber.io), [GORM](https://gorm.io)

Development: http://localhost:2244

Heroku: https://backend2fa.herokuapp.com

### Deploy

```shell script
git clone https://github.com/peterdee/2fa-backend
cd ./2fa-backend
gvm use go1.18
go mod download
```

### Environment variables

The `.env` file is required for all of the environments except for Heroku, see [.env.example](.env.example) for details

### Launch

```shell script
go run ./
```

Can be launched with [AIR](https://github.com/cosmtrek/air)

### Heroku

The `release` branch is automatically deployed to Heroku

### License

[MIT](LICENSE.md)
