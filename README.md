# Go/Gin Quote Service

Simple Go/Gin REST service for storing and retrieving quotes


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://www.heroku.com/deploy?template=https://github.com/heroku-examples/go-gin)

This repository goes along with the blog post [Deploying a Simple Go/Gin Appliction on Heroku](https://www.heroku.com/blog/deploying-simple-go-gin-application-on-heroku/).

### Manual Deployment

#### Download and install Go

On macOS you can use Homebrew, on other systems you can download the installer from the [official Go website](https://golang.org/dl/).

```bash
brew install go
```

#### Create your workspace
```bash
go mod initi gihub.com/YOUR-USERNAME/YOUR-REPO-NAME
```

#### Add the Gin Framework
```bash
go get github.com/gin-gonic/gin
```

#### Run the server
```bash
go run main.go
```

### Create a Heroku App
#### Login & Create app
```bash
heroku login

heroku apps:create my-go-gin-api
```

#### Show Git remotes
```bash
git remote show heroku
```

#### Tidy up your Go project
```bash
go mod tidy
```

### Deploying Your Application
Assuming you're in the project directory, `git add .` works great. Otherwise you might need `git add -All`.
```bash
git add .
git commit -m "Prepare app for Heroku deployment"
git push heroku main 
```

### Test the live application
```bash
curl -s -X GET https://<<APP-URL>>/quote | jq
```
