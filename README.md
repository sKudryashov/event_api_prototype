# Event api prototype:
**Usage example**

There is actually four actions: create, get all, get by event type, get by 
time range. Just type appropriate commands in your console.

Get by type: `curl -XPOST -H 'Content-Type: application/json' -d '{"eventType":"Usual"}' http://localhost:3030/readbytype`

Create: `curl -XPOST -H 'Content-Type: application/json' -d '{"eventType":"Usual","sessionStart":"1476628565","sessionEnd":"1476628965","linkClicked":"https://blog.golang.org/c-go-cgo","timestamp":12039109203,"params":{"C":"c++","D":"D++","R":"R is not a real language"}}' http://localhost:3030/add`

Read all: `curl -XGET -H 'Content-Type: application/json' http://localhost:3030/read`

Read by time range: `curl -XGET -H 'Content-Type: application/json' http://localhost:3030/readbytimerange/:start/:end`

Note that each action dataset has it's own validation rules and some fields are required. If the query didn't pass validation, panic would be triggered. Of course you may change this behaviour easily.

**How to launch it ?**

Probably the simplest way to do it without dancing around mongoDB and glide installation it's just download Docker, get into project root where docker-compose.yml is placed and then type docker-compose up -d. The app will be available on 192.168.99.100:3030
