### Python Example With Watchdog
This example shows how to setup Python Watchdog to monitor a folder, then execute a workflow based on changes in the folder

## Things to consider
+ This example assumes that you have to authenticate to IIS. It uses requests_nltm for this setup(test server to test server). If your server is setup for anonymous auth, you may not NEED the requests_ntlm auth part or may have to tweak it if you're using Kerberos.
+ The credentials are stored in script. THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script. Please store them somewhere server side encrypted.
+ Cool thing about this script is you can set it up to run as Windows Service to constantly monitor a folder in the background, much like ImportAgent does, and pass data along to work.

## Extra documentation

+ Python Watchdog https://pypi.org/project/watchdog/
+ Setting up a Python script as a Windows Service http://thepythoncorner.com/dev/how-to-create-a-windows-service-in-python/
