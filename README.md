# Backend of TonWork service.
TonWork that's platform/marketplace where you can find work, team, and interesting project that rooted with IT(and not also).
If you need team, you will can create a find order on, for example, data sciences developer or another. 
And you can pay for work by crypto currency(in our plan start from TON coint and move to another SOL, ETH, BTC...).

<h3>Work with API</h3>
<code>
"/auth":
    "/register"    POST 
    "/login"       POST 
    "/profile"     POST 

"/api/v2/":
    "/work":
    	("/")      POST
    	("/")      GET
    	("/:id")   GET
    	("/:id")   PATCH
    	("/:id")   DELETE
    "/posts":
    	("/")      POST
    	("/")      GET
    	("/:id")   GET
    	("/:id")   PATCH
    	("/:id")   DELETE
    "/subscribe":
    	("/buy")   POST
    	("/cancel")POST
</code>
