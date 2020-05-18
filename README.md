# Client/Server setup for learning

React/Golang setup intended for in a IOT magic mirror project.
Anyone in need of an easy React/Redux and Golang/Gorilla setup can of course use this.

nodemon is used to auto build the go-project defined under the "server" folder, the
react client is built using webpack and served by webpack-dev-server.

## Initial setup

- Install go
- Get the project using `go get github.com/etimo/go-magic-mirror`
- Go to the right folder in your gopath i.e. `$GOPATH/src/github.com/etimo/go-magic-mirror`
- Run 'npm install' or 'yarn install' to install project development and frontend dependencies.
- In same-folder run 'npm run install-go' or 'yarn install-go' to install all Go dependencies.
- To run in development mode: 'npm run dev' or 'yarn dev'
This will run the frontend using "webpack-dev-server" on localhost:3000 and the golang server on localhost:8080, both auto refreshing on code updates.

- Build using 'npm run build' or 'yarn build'
  The go-magic mirror binary will also serve the generate JS bundle from folder `./dist`
- Run only frontend 'npm run client' or 'yarn client'
- Run only server 'npm run server' or 'yarn server'
- The server part can also be built using `go build`

## Building for ARMv7
To build an ARMv7 compatible binary and dist folder for the Raspberry Pi 3 and up. Run the following command:
`export GO_ARM=7; export GO_ARCH=arm; yarn build`
## Architecture:

The frontend opens a websocket to the backend, currently only a single connection to the backend is allowed.


