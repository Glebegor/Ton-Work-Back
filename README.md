# Backend of TonWork service.
TonWork that's platform/marketplace where you can find work, team, and interesting project that rooted with IT(and not also).
If you need team, you will can create a find order on, for example, data sciences developer or another. 
And you can pay for work by crypto currency(in our plan start from TON coint and move to another SOL, ETH, BTC...).

<h3>Start project by cmd</h3>
<h4>1-Start database:</h4>
You need to download a docker image(postgres).<br>
And after create database:<br>
<code>docker run --name=name-of-database -e POSTGRES_PASSWORD="password" -p 5436:5432 -d --rm postgres</code>
<h4>2-Create migrations:</h4>
<code>migrate -path ./schema -database "postgres://postgres:password@localhost:5436/postgres?sslmode=disable" up</code>
<h4>3-Create .env file:</h4>
DB_PASSWORD='password'<br>
Secret_Key='SECRET-KEY'
<h4>4-Run project:</h4>

<h3>Run project by Makefile:</h3>
<h4>installdb:</h4>
This command will install docker image to your computer.<br>
<code>docker pull postgres</code>
<h4>build:</h4>
This command will run docker database(PostgresSQL). <br>
<code>docker run --name=name-of-database -e POSTGRES_PASSWORD='password' -p 5436:5432 -d --rm postgres</code>
<h4>migrate:</h4>
This command will create migrations for your database.<br>
<code>migrate -path ./schema -database "postgres://postgres:password@localhost:5436/postgres?sslmode=disable" up</code>
<h4>run:</h4>
This command will run project.<br>
<code>go run cmd/main.go</code>

<h3>Loading by docker compose</h3>
<code>docker compose build</code><br>
<code>docker compose up</code>

<h3>Work with API</h3>

1. "/auth":
      - "/register"   POST
      - "/login"      POST
      - "/profile/:username"    GET
      - "/updateProfile"    PUT
2. "/api/v2/":
      - "/work":
          - "/"       POST
          - "/"       GET
          - "/:id"    GET
          - "/:id"    PATCH
          - "/:id"    DELETE
      - "/posts":
          - "/"       POST
          - "/"       GET
          - "/:id"    GET
          - "/:id"    PATCH
          - "/:id"    DELETE
      - "/subscribe":
          - "/buy"    POST
          - "/cancel" POST
          - "/timetoend" GET
      - "/chat":
          - "/CreateRoom" POST
          - "/GetRooms" GET
          - "/JoinRoom/:roomId" WebSocket
          - "/GetClients/:roomId" Get
         
<h4>Or you can see it in this format(Handler is realizated if has "+" on the right side):</h4><br>
"/auth":<br>
"/auth/updateProfile"             PUT<br>
"/auth/register"                  POST+<br>
"/auth/login"                     POST+<br>
"/auth/profile"                   GET+<br>
<br>
"/api/v2/":<br>
<br>
"/api/v2/work":<br>
"/api/v2/work/"                   POST+<br>
"/api/v2/work/"                   GET+<br>
"/api/v2/work/:id"                GET+<br>
"/api/v2/work/:id"                PATCH+<br>
"/api/v2/work/:id"                DELETE+<br>
<br>
"/api/v2/posts":<br>
"/api/v2/posts/"                  POST+<br>
"/api/v2/posts/"                  GET+<br>
"/api/v2/posts/:id"               GET+<br>
"/api/v2/posts/:id"               PATCH+<br>
"/api/v2/posts/:id"               DELETE+<br>
<br>
"/api/v2/subscribe":<br>
"/api/v2/subscribe/buy"           POST+<br>
"/api/v2/subscribe/cancel"        POST+<br>
"/api/v2/subscribe/timetoend"     GET+<br>
<br>
"/api/v2/chat:<br>
"/api/v2/chat/CreateRoom"         POST+<br>
"/api/v2/chat/GetRooms"           GET+<br>
"/api/v2/chat/JoinRoom/:roomId"   WebSocket+<br>
"/api/v2/chat/GetClients/:roomId" GET+<br>
<h3>JWT token and his structure</h3>
<hr>
JWT token looks like this:
<code>qweioucu34ioslk1j23lkjds.dkowjrlekwjriodfslvvldkwsjqr.dsfouqweopriuoiu3oi3o2uadflsk</code>
And has three parts: header.payload.signature<br>
Claims of our JWT token: userId, userUsername, userName, userSurname.<br>
Please, save in cookies or session storage. <br>
Header for token: Authorization.<br> 

<h3>Requests and responses on every link</h3>
<h3>AUTH</h3>
<hr>
<h4>"/auth/register", method:POST.</h4>

Type | JSON 
--- | ---
Request | { "username": "User Name", "password_hash": "123456789", "email": "nickname@gmail.com" }
Response | { "Status": "OK" } 
Error Response | { "message": "Some text" } 

<h4>"/auth/login", method:POST.</h4>

Type | JSON 
--- | ---
Request | { "username": "User Name", "password_hash": "123456789" }
Response | { "token": "wqewqeqwr123o1kepo2k-c439!(#_$I(#$.@)#@!O)$K@J)!$.!@(#JWJDWADISIADOUI" }
Error Response | { "message": "Some text" } 

<h4>"/auth/profile/${Username}", method:GET.</h4>

Type | JSON | Headers
--- | --- | ---
Request | write param to url, example: "/auth/profile/Glebegor" | Headers
Response | { "username": "User name", "email": "email@gmail.com", "telefon": "+3242 3242 432", "position": "Position", "description": "Description text", "subscribe": "Subscribe", "companies": "Companies", "name": "name", "surname": "Surname", "id": "Id" } | Headers
Error Response | { "message": "Some text" } | Headers

<h3>POSTS</h3>
<hr>
<h4>"/api/v2/posts/", method:GET.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | --- 
Response | {data: [{ "title": "Title", "description": "Title", "text": "aqweqweqwesd", "tags": "['asdasd','asdasd','asdasd']", "rating": 1, "id": 2 }... ]}| --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/posts/", method:POST.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | { "title": "Title", "description": "Title", "text": "aqweqweqwesd", "tags": "['asdasd','asdasd','asdasd']", "rating": 1} | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" }| --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/posts/:id", method:GET.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | ---
Response | { "title": "Title", "description": "Title", "text": "aqweqweqwesd", "tags": "['asdasd','asdasd','asdasd']", "rating": 1}| --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/posts/:id", method:PUT.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | { "title": "Title", "description": "Title", "text": "aqweqweqwesd", "tags": "['asdasd','asdasd','asdasd']", "rating": 1} | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" } | --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/posts/:id", method:DELETE.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" } | --- 
Error Response | { "message": "Some text" } | --- 

<h3>WORKS</h3>
<hr>
<h4>"/api/v2/posts/", method:GET.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | --- 
Response | { "data": [ { "title": "title", "description": "description", "text": "text", "tags": "['asdasd', 'asdad']", "technologies": "['qweqw', 'qweqe']", "company": "qwec rewqrq", "price": 2, "experienceLevel": "qweq ", "type_of_job": "qwe ", "invites": 0, "rating": 2, "id": 1 } ] }| --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/work/", method:POST.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | { "title": "title", "description": "description", "text": "text", "tags": "['asdasd', 'asdad']", "technologies": "['qweqw', 'qweqe']", "company": "qwec rewqrq", "price": 2, "experienceLevel": "qweq ", "type_of_job": "qwe ", "invites": 0, "rating": 2 } | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" }| --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/work/:id", method:GET.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | ---
Response | { "title": "title", "description": "description", "text": "text", "tags": "['asdasd', 'asdad']", "technologies": "['qweqw', 'qweqe']", "company": "qwec rewqrq", "price": 2, "experienceLevel": "qweq ", "type_of_job": "qwe ", "invites": 0, "rating": 2 }| --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/work/:id", method:PUT.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | { "title": "title", "description": "description", "text": "text", "tags": "['asdasd', 'asdad']", "technologies": "['qweqw', 'qweqe']", "company": "qwec rewqrq", "price": 2, "experienceLevel": "qweq ", "type_of_job": "qwe "} | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" } | --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/work/:id", method:DELETE.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" } | --- 
Error Response | { "message": "Some text" } | --- 

<h3>SUBSCRIBES</h3>
<hr>
<h4>"/api/v2/subscribe/buy", method:POST.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" } | --- 
Error Response | { "message": "Some text" } | --- 
{ "name":"qweqwe", "id":"qweqew" }
<h4>"/api/v2/subscribe/cancel", method:POST.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "Status": "OK" } | --- 
Error Response | { "message": "Some text" } | --- 

<h4>"/api/v2/subscribe/timetoend", method:GET.</h4>

Type | JSON | Headers 
--- | --- | --- 
Request | --- | Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a" 
Response | { "TimeToEnd": -1, "Username": "username"} | --- 
Error Response | { "message": "Some text" } | --- 

<h3>CHAT</h3>
<hr>
<h4>"/api/v2/chat/CreateRoom", method:POST</h4>

Type | JSON | Headers
--- | --- | ---
Request | { "name":"qweqwe", "id":"qweqew" } |  Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a"
Response | { "name":"qweqwe", "id":"qweqew" } | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v2/chat/JoinRoom/:roomId", method:WEBSOCKET</h4>

Type | JSON | Headers
--- | --- | ---
Request | --- |  Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a"
Response | { --- } | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v2/chat/GetRooms", method:GET</h4>

Type | JSON | Headers
--- | --- | ---
Request | --- |  Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a"
Response | [ { "id": "room1", "name": "room1" }...] | ---
Error Response | { "message": "Some text" } | ---

<h4>"/api/v2/chat/GetClients/:roomId", method:GET</h4>

Type | JSON | Headers
--- | --- | ---
Request | --- |  Authorization: "Bearer tokenqw.qweqweqe.qwesaid0@OI#U!sf09a"
Response | [ { "id": "3", "username":"username" }...] | ---
Error Response | { "message": "Some text" } | ---
