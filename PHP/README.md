### PHP Example

This example is a straight PHP example. It shows how to call the Workflow Web Services from PHP using cURL.

### Things to consider

+ This example assumes that you have to authenticate to IIS. It uses NTLM for this setup(Again, test server to test server). If your server is setup for anonymous auth, you may not NEED the NTLM auth part, or may have to tweak it depending on Auth Method
+ The credentials are stored in script. **THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script.** Please store them somewhere server side encrypted.
+ This example is test server to test server, special tweaks may have to put in place for PHP cURL to work with HTTPS, depending on your server setup

This is a minimalistic approach for calling the Workflow Web Services, we realized that really all you need to call the Workflow Web Services is the Input parameters(if needed). The other stuff that Laserfiche talks about in documentation is nice to have BUT NOT required.


### Workflow PHP Client

If you are looking to integrate Workflow Web Services with PHP, we have a released a Workflow Client written in PHP that will handle most of the heavy lifting for you. Please visit our Github Repository located [Here](https://github.com/CabarrusCo/Laserfiche-Workflow-PHP-Client) for more information.
