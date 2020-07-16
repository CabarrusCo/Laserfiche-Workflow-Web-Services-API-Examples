import time
from watchdog.observers import Observer
from watchdog.events import PatternMatchingEventHandler
import os
import requests
from requests_ntlm import HttpNtlmAuth

if __name__ == "__main__":
    patterns = "*.txt"
    ignore_patterns = ""
    ignore_directories = False
    case_sensitive = True
    my_event_handler = PatternMatchingEventHandler(patterns, ignore_patterns, ignore_directories, case_sensitive)

    def on_created(event):
        newly_created_file = event.src_path
        newly_created_file_data = open(newly_created_file)
        input_parameter = newly_created_file_data.readline()
        newly_created_file_data.close()

        workflow_url = "http://WORKFLOW_URL_HERE"

        r = requests.post(workflow_url, json={"ParameterCollection":[{"Name":"testParameter", "Value":input_parameter}]}, auth=HttpNtlmAuth('domain\\XXXXXX','XXXXXX'))
        print(r.status_code)

        if r.status_code == 200:
            print("WORKFLOW FIRED")


    my_event_handler.on_created = on_created

    path = r"C:\TESTWATCHDIRECTORY"
    go_recursively = False
    my_observer = Observer()
    my_observer.schedule(my_event_handler, path, recursive=go_recursively)

    my_observer.start()
    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        my_observer.stop()
        my_observer.join()