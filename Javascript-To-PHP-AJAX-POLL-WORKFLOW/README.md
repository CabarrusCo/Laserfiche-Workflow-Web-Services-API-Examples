### Javascript to PHP, Poll Workflow for Completion Example

This example shows how to Start a Workflow with a Javascript to PHP AJAX call and then get a response to the application that the workflow has finished.

### Why?

The Workflow web services, whenever started, returns a response back immediately that a workflow has started. That's great for MOST cases. But
what if for some reason you want your application to know when the Workflow has FINISHED and with what STATUS and not just started? That
is when you have to implement a polling mechanism. An example of when this would be useful is if you want to use Workflow to do quick Database
CRUD actions and then relay it back to the end user that the action has been completed.

### How it works

1. Javascript uses axios to POST to the backend web services and passes two parameters. 

### Things to consider

+ This example assumes that you have to authenticate to IIS. It uses NTLM for this setup(Again, test server to test server). If your server is setup for anonymous auth, you may not NEED the NTLM auth part, or may have to tweak it depending on Auth Method
+ The credentials are stored in script. **THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script.** Please store them somewhere server side encrypted.
+ This example is test server to test server, special tweaks may have to put in place for PHP cURL to work with HTTPS, depending on your server setup

This is a minimalistic approach for calling the Workflow Web Services, we realized that really all you need to call the Workflow Web Services is the Input parameters(if needed). The other stuff that Laserfiche talks about in documentation is nice to have BUT NOT required.
