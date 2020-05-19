### Javascript to PHP AJAX example
This example is a Javascript AJAX to PHP backend example.

Javascript passes the parameter to a PHP script, then PHP echos back up to the front end. The Javascript uses two libraries for
handling the AJAX call, AXIOS and qs.

## Things to consider
+ This example assumes that you have to authenticate to IIS. It uses NTLM for this setup(Test server to test server). If your server is setup for anonymous auth, you may not NEED the NTLM auth part or may have to tweak it depending on Auth Method.
+ The credentials are stored in script. THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script. Please store them somewhere server side encrypted.
+ This example is test server to test server, special tweaks may have to put in place for PHP cURL to work with HTTPS, depending on your server setup
