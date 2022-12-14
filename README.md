# Udacity Golang Final Project

## Description

This project is a Golang api that does CRUD operations on Customer objects. It runs on port `3000`

## Install

* Install Go version 1.18 or higher
* Clone the project with `git clone https://github.com/drichards188/udacityCrm.git`
* go to folder and run `go get -d ./...`

## Run 

* Navigate to `/crmBackend/cmd/main` and run `go run main.go`
* Make requests to `localhost:3000`

## Usage

The Customer object you will pass is 
`{ Id: int, Name: string, Role: string, Active: bool }`

* Get all customers: `/customers` GET
* Get single customer: `/customers/:id` GET 
* Create a customer: `/customers` POST with Customer body
* Update a customer: `/customers/:id` PUT with Customer body
* Delete a customer: `/customers/:id` DELETE
