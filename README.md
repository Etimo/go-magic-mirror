# Client/Server setup for learning

This is a base React/Golang setup intended for use in a collaborative IOT magic mirror project.
Anyone in need of an easy React/Redux and Golang/Gorilla setup can of course use this.

nodemon is used to auto build the go-project defined under the "server" folder, the
react client is built using webpack and served by webpack-dev-server.

## Initial setup

- Install go
- Clone the repo
- In the repo folder run 'npm install' or 'yarn install'
- In the repo folder run 'npm install-go' or 'yarn install-go'
- To run in development mode: 'npm run dev' or 'yarn dev'
This will run the frontend on localhost:3000 and the golang server on localhost:8080, both auto refreshing on code updates.

- Build using 'npm run build' or 'yarn build'
- Run only frontend 'npm run client' or 'yarn client'
- Run only server 'npm run server' or 'yarn server'

