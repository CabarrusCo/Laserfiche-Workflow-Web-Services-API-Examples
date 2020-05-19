### Python Example

This example is a straight Python example. It shows how to call the Workflow Web Service from Python using the Requests library.

### Things to consider

+ This example assumes that you have to authenticate to IIS. It uses requests_nltm for this setup(test server to test server). If your server is setup for anonymous auth, you may not NEED the requests_ntlm auth part or may have to tweak it depending on Auth Method
+ The credentials are stored in script. **THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script.** Please store them somewhere server side encrypted.

This is a minimalistic approach for calling the Workflow Web Services, we realized that really all you need to call the Workflow Web Services is the Input parameters(if needed). The other stuff that Laserfiche talks about in documentation is nice to have BUT NOT required.
