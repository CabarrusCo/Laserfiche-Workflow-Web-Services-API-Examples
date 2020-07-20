### Go Post to Laserfiche Workflow

This example shows how to POST to Laserfiche Workflow and Start a workflow using Go. This example uses structs to generate the JSON used during the POST.

### Things to consider

+ This example assumes that you have to authenticate to IIS. It uses the go-ntlmssp package for handling auth for this setup(test server to test server). If your server is setup for anonymous auth, you may not NEED this package.
+ The credentials are stored in script. **THIS IS FOR EXAMPLE ONLY. DO NOT STORE credentials in the real life script.** Please store them somewhere server side encrypted.

### Running and build

```
Run - go run main.go
Build - go build main.go
```
