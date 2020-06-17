import pycurl
import json


if __name__ == "__main__":
    
    url = "http://WORKFLOW_URL_HERE"
    
    data = json.dumps({"ParameterCollection":[{"Name":"InputParameter", "Value":"Hello World!"}]})

    c = pycurl.Curl()
    c.setopt(pycurl.URL, url)

    c.setopt(pycurl.HTTPHEADER, ['Content-Type: application/json', 
                                 'Accept: application/json']) #v=Urj7Uw4TodA

    c.setopt(pycurl.POST, 1)
    c.setopt(pycurl.POSTFIELDS, data)
    c.setopt(pycurl.VERBOSE, 1)
    c.setopt(pycurl.HTTPAUTH, pycurl.HTTPAUTH_NTLM)
    c.setopt(pycurl.USERPWD, "DOMAIN\\XXXXXXXXX:XXXXXXXX")
    c.perform()
    c.close()
