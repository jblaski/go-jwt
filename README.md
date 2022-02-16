### Credit
This repository was initially started by following the great tutorial at the below URL:
https://betterprogramming.pub/hands-on-with-jwt-in-golang-8c986d1bb4c0

### JWT
We have some server running, to which we can make requests. Some requests should on respond if the client sending them is authenticated.

Basic steps:
 - client authenticates with server using username and password
 - if valid, the server signs a JWT, and returns it to the client
 - client can then include an Authorization header in subsequent requests to protected resources
 - the API validates the JWT when recieving these requests