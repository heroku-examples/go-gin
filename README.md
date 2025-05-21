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
~/project$ go mod initi gihub.com/YOUR-USERNAME/YOUR-REPO-NAME
```

#### Add the Gin Framework
```bash
~/project$ go get github.com/gin-gonic/gin
```

#### Run the server
```bash
~/project$ go run main.go
```

### Create a Heroku App
#### Login & Create app
```bash
~/project$ heroku login

~/project$ heroku apps:create my-go-gin-api
```

#### Show Git remotes
```bash
~/project$ git remote show heroku
```

#### Tidy up your Go project
```bash
~/project$ go mod tidy
```

### Deploying Your Application
```bash
~/project$ git add .
~/project$ git commit -m "Prepare app for Heroku deployment"
~/project$ git push heroku main 
```

### Test the live application
```bash
~/project$ curl -s -X GET https://<<APP-URL>>/quote | jq
```