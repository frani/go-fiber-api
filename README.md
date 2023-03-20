# Go Fiber API with Mongo

Current boilerplate is a migration from and old NodeJS template that I used to use.

# Folder structure

```sh
| Folder       | Suffix Extention | description                                                                            |
| ------------ | ---------------- | -------------------------------------------------------------------------------------- |
| /routes      | \*.route.js      | Fiber Routes                                                                           |
| /validations | \*.validation.js | Request Validation Schemas                                                             |
| /middlewares | \*.mid.js        | custom Fiber middlewares                                                               |
| /controllers | \*.controller.js | controller (a.k.a. handlers )                                                          |
| /services    | \*.service.js    | business logic, incluye intertal or externaa servicies as SDK (service layer)          |
| /database    | \*.db.js         | Function to manage Mongoose schemas                                                    |
| /models      | \*.model.js      | Mongoose models/models (data layer)                                                    |
| /config      | \*.conf.js       | Environment variables and configuration related things                                 |
| /helpers     | \*.helper.js     | Helpful functions that can be used in many places but is not necesary a business logic |
| /tests       | \*.test.js       | well... just tests for your API                                                        |
```
