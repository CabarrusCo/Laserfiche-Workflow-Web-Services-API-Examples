import requests
from requests_ntlm import HttpNtlmAuth

workflow_url = "http://WORKFLOW_URL_HERE"

r = requests.post(workflow_url, json={"ParameterCollection":[{"Name":"message", "Value":"HELLO WORLD!"}]}, auth=HttpNtlmAuth('domain_user\\xxxxxx','xxxxxxxx'))
print(r.status_code)