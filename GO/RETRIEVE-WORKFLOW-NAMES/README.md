### Go Grab All Workflow Names Example

This example shows how to call the endpoint on Laserfiche server that houses all the workflow data. It will pull back all Workflow Names.

### Things to consider

+ This example assumes that you have to authenticate to IIS. It uses the go-ntlmssp package for handling auth for this setup(test server to test server). If your server is setup for anonymous auth, you may not NEED this package.
+ The credentials are stored in script. **THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script.** Please store them somewhere server side encrypted.

### Other documentation

+ This example is our first attempt at interacting with Laserfiche with Go. We will upload an example of how to Start a workflow with Go and pass parameters soon.
