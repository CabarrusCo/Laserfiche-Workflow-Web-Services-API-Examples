### Javascript to PHP, Poll Workflow for Completion With Javascript Example

This example shows how to Start a Workflow with a Javascript to PHP AJAX call and then use Javascript to poll the Workflow every 30 seconds.

### Why?

The Workflow web services, whenever started, returns a response back immediately that a workflow has started. That's great for MOST cases. But
what if for some reason you want your application to know when the Workflow has FINISHED and with what STATUS? That
is when you have to implement a polling mechanism. This example is to be used on longer running Workflows that run for a longer period of time
that require a "check in" instead of blocking on the Web Service for completion.

### How it works

1. Javascript uses axios to POST to the backend web services and passes two parameters. These two parameters are a regular Input, and a UUID that is generated by the uuidv4 function
2. PHP initiates cURL and hits workflow. A Text file is created in this format...GUID_STATUS.txt, with initial status being STARTED.
3. Workflow starts and runs in the background, performing it's action. In this example, it's actions are SDKs scripts and delays. The SDK scripts
updates a text file with a UUID every minute and "esclates" the text file in levels. I.E. LevelOne, LevelTwo, Success.
4. axios sends a POST to a PHP web service every 30 seconds if UUIDs in circulation exists. The PHP web service loops through the directory and updates each UUIDs status and relays it back to the front end
5. Javascript makes the determination of where to put the UUID based on status. If NOT Success, keep it in the Circulation Array. If Success
move it to the Success Array.
6. Javascript updates the textareas correspondingly.

### Screenshots

#### All UUIDS at Level One
![All UUIDS at Level One](https://github.com/CabarrusCo/Laserfiche-Workflow-Web-Services-API-Examples/blob/master/Javascript-To-PHP-POLL-WORKFLOW-Javascript/Screenshots/LEVELONE.PNG)

#### All UUIDS at Level Two
![All UUIDS at Level Two](https://github.com/CabarrusCo/Laserfiche-Workflow-Web-Services-API-Examples/blob/master/Javascript-To-PHP-POLL-WORKFLOW-Javascript/Screenshots/LevelTwo.PNG)

#### All UUIDS at SUCCESS

![All UUIDS at Success](https://github.com/CabarrusCo/Laserfiche-Workflow-Web-Services-API-Examples/blob/master/Javascript-To-PHP-POLL-WORKFLOW-Javascript/Screenshots/UUIDsCompletedSuccessfully.PNG)

#### Laserfiche Workflow

![Laserfiche Workflow](https://github.com/CabarrusCo/Laserfiche-Workflow-Web-Services-API-Examples/blob/master/Javascript-To-PHP-POLL-WORKFLOW-Javascript/Screenshots/Workflow.PNG)

### Things to consider

+ This example assumes that you have to authenticate to IIS. It uses NTLM for this setup(Again, test server to test server). If your server is setup for anonymous auth, you may not NEED the NTLM auth part, or may have to tweak it depending on Auth Method
+ The credentials are stored in script. **THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script.** Please store them somewhere server side encrypted.
+ This example is test server to test server, special tweaks may have to put in place for PHP cURL to work with HTTPS, depending on your server setup
+ This example ties up a Web Service until completion and **SHOULD NOT** be used on long running Workflows. For more advanced Workflows, consider using Javascript/AJAX to poll on an interval