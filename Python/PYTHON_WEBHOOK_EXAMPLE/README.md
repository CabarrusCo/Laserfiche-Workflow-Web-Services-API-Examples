### Python Webhook Example

This example shows how to setup Python as Webhook. The Webhook then GETs parameters from the URL. The parameters are then passed along to the Workflow API.

### Things to consider

+ This example assumes that you have to authenticate to IIS. It uses requests_nltm for this setup(test server to test server). If your server is setup for anonymous auth, you may not NEED the requests_ntlm auth part or may have to tweak it depending on Auth Method.
+ The credentials are stored in script. **THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script.** Please store them somewhere server side encrypted.

### Other documentation

+ This examples uses a small cherrypy wrapper for Python called "Webhook-Listener". "Webhook-Listener" is a minimal Python Package that lets you setup Webhooks and pass the data along to other functions. More information can be found here. https://pypi.org/project/Webhook-Listener/
