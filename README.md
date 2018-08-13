# goload
A quick and simple load tester using templates to populate random data

## Usage
```
goload get 'http://localhost:8080/encode/{firstname}/{lastname}' -n 10000 -c 10
```
Create 10 clients to send 10000 GET requests each to /encode/{firstname}/{lastname} 
where firstname and lastname are randomly generated for each request
