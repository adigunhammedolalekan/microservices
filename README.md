# Event Driven Microservices Test

### How to run
##### Requirements
* docker
* docker-compose
* Go 1.1+
* make

```shell script
$ make run
```

This will start all the services and their dependencies

### Testing through API gateway

An API gateway is running on `:2004`, you can test the services by `curling: http://localhost:2004`

#### Available endpoints
* `/api/event/new` -  to create events. `http://localhost:2004/api/event/new`
###### Body Sample
```json
{
    "id": "01EBP4DP4VECW8PHDJJFNEDVKE",
    "name": "targets.acquired",
    "data": [
    {
        "id": "01EBP52N9Q18YDXADTZH1ZEFT51",
        "message": "some message to send",
        "created_on": "2020-06-25T16:31:18.993Z",
        "updated_on": "2020-06-25T16:31:18.993Z"
    },
    {
        "id": "01EBP52N9QR44YZAK3WEPXPS121",
        "message": "some message to send",
        "created_on": "2020-06-25T16:31:18.993Z",
        "updated_on": "2020-06-25T16:31:18.993Z"
    },
    {
        "id": "01EBP52N9QTCMDRD51WDXKX7761",
        "message": "some message to send",
        "created_on": "2020-06-25T16:31:18.993Z",
        "updated_on": "2020-06-25T16:31:18.993Z"
        }
    ],
    "created_on": "2020-06-25T16:31:18.993Z"
}
```

* `/api/events` - list all events. `http://localhost:2004/api/events`
