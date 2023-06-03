# Backend of TonWork service.
TonWork that's platform/marketplace where you can find work, team, and interesting project that rooted with IT(and not also).
If you need team, you will can create a find order on, for example, data sciences developer or another. 
And you can pay for work by crypto currency(in our plan start from TON coint and move to another SOL, ETH, BTC...).

<h3>Work with API</h3>

1. "/auth":
      - "/register"   POST
      - "/login"      POST
      - "/profile"    POST
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
"/auth/profile"            POST<br>
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



