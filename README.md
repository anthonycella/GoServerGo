# GoServerGo

A basic go server app that promps the user to enter a positive integer between 0 and 20 inclusive. Once the user inputs the integer the application will send the integer to the server as a string to compute the factorial of said integer. The application is terminated when the user inputs "x". All other input besides positive integers between 0 and 20 inclusive and x is considered invalid and will print an error message to the user.

![GoServerGo demo](./goserverdemo.gif)


## Tech Stack
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Getting Started
1. If you have not done so already, install Go on your local computer. Installation instructions can be found [here](https://go.dev/doc/install)
2. Open up a terminal window and clone down the repository to your local computer
```
git clone https://github.com/anthonycella/GoServerGo
```
3. Once you are in the repository navigate into the go-server file
```
cd go-server
```
4. Once in the go-server file start up the server with the following command
```
go run server.go
```
5. Open up a second terminal window and navigate into the GoServerGo repository and then into the go-client file
```
cd go-client
```
6. Start up the client and follow the prompts!
```
go run client.go
```

#### Author: Anthony Cella