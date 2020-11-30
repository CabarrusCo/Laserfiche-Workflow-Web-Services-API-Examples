import time
import requests
from requests_ntlm import HttpNtlmAuth
import webhook_listener

#v=EX09nRA-jE4

def process_get_request(request, *args, **kwargs):

    workflow_url = "http://WORKFLOW_URL_HERE"
    input_parameter = kwargs["inputParameter"]
    
    r = requests.post(workflow_url, json={"ParameterCollection":[{"Name":"testParameter", "Value":input_parameter}]}, auth=HttpNtlmAuth('domain\\XXXXXXX','XXXXXXX'))
    return

webhooks = webhook_listener.Listener(handlers={"GET": process_get_request})

webhooks.start()

while True:
    print("Still alive...")
    time.sleep(300)
