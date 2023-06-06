# Backend of TonWork service.
TonWork that's platform/marketplace where you can find work, team, and interesting project that rooted with IT(and not also).
If you need team, you will can create a find order on, for example, data sciences developer or another. 
And you can pay for work by crypto currency(in our plan start from TON coint and move to another SOL, ETH, BTC...).

<h3>Work with API</h3>

1. "/auth":
      - "/register"   POST
      - "/login"      POST
      - "/profile"    GET
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
         
<h4>Or you can see it in this format:</h4><br>
"/auth":<br>
"/auth/register"           POST<br>
"/auth/login"              POST<br>
"/auth/profile"            GET<br>
<br>
"/api/v2/":<br>
<br>
"/api/v2/work":<br>
"/api/v2/work/"            POST<br>
"/api/v2/work/"            GET<br>
"/api/v2/work/:id"         GET<br>
"/api/v2/work/:id"         PATCH<br>
"/api/v2/work/:id"         DELETE<br>
<br>
"/api/v2/posts":<br>
"/api/v2/posts/"           POST<br>
"/api/v2/posts/"           GET<br>
"/api/v2/posts/:id"        GET<br>
"/api/v2/posts/:id"        PATCH<br>
"/api/v2/posts/:id"        DELETE<br>
<br>
"/api/v2/subscribe":<br>
"/api/v2/subscribe/buy"    POST<br>
"/api/v2/subscribe/cancel" POST<br>

<h3>JWT token and his structure</h3>
JWT token looks like this:
<code>qweioucu34ioslk1j23lkjds.dkowjrlekwjriodfslvvldkwsjqr.dsfouqweopriuoiu3oi3o2uadflsk</code>
And has three parts: header.payload.signature<br>
claims of our JWT token: userId, userUsername, userName, userSurname.<br>
Please, save in cookies or session storage.  

<h3>Requests and responses on every link</h3>
<hr>
<h4>"/auth/register", method:POST.</h4>

Type | JSON 
--- | ---
Request | { "username": "User Name", "password": "123456789", "email": "nickname@gmail.com" }
Response | { "Status": "OK" } 
Error Response | { "message": "Some text" } 

<hr>
<h4>"/auth/login", method:POST.</h4>

Type | JSON 
--- | ---
Request | { "username": "User Name", "password": "123456789" }
Response | { "token": "wqewqeqwr123o1kepo2k-c439!(#_$I(#$.@)#@!O)$K@J)!$.!@(#JWJDWADISIADOUI" }
Error Response | { "message": "Some text" } 

<hr>
<h4>"/auth/profile/${Username}", method:GET.</h4>

Type | JSON 
--- | ---
Request | write param to url, example: "/auth/profile/Glebegor"
Response | { "username": "User name", "email": "email@gmail.com", "telefon": "+3242 3242 432", "position": "Position", "description": "Description text", "subscribe": "Subscribe", "companies": "Companies", "name": "name", "surname": "Surname", "id": "Id" }
Error Response | { "message": "Some text" } 

