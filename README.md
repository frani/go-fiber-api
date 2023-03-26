# Go Fiber API with Mongo

Current boilerplate is a migration from and old NodeJS template that I used to use.

# Folder structure

```sh
| Folder       | Suffix Extention | description                                                                            |
| ------------ | ---------------- | -------------------------------------------------------------------------------------- |
| /routes      | \*_route.go      | Fiber Routes                                                                           |
| /validations | \*_validation.go | Request Validation Schemas                                                             |
| /middlewares | \*_mid.go        | custom Fiber middlewares                                                               |
| /handlers    | \*_handler.go    | handler (a.k.a. controllers )                                                          |
| /services    | \*_service.go    | business logic, incluye intertal or externaa servicies as SDK (service layer)          |
| /database    | \*_db.go         | Function to manage Mongoose schemas                                                    |
| /models      | \*_model.go      | Mongoose models/models (data layer)                                                    |
| /config      | \*_conf.go       | Environment variables and configuration related things                                 |
| /helpers     | \*_helper.go     | Helpful functions that can be used in many places but is not necesary a business logic |
| /tests       | \*_test.go       | well... just tests for your API                                                        |
```
