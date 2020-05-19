# Laserfiche Workflow API-Open Source Examples

### About Cabarrus County
---
Cabarrus is an ever-growing county in the southcentral area of North Carolina. Cabarrus is part of the Charlotte/Concord/Gastonia NC-SC Metropolitan Statistical Area and has a population of about 210,000. Cabarrus is known for its rich stock car racing history and is home to Reed Gold Mine, the site of the first documented commercial gold find in the United States.

### About our team
---
The Business & Location Innovative Services (BLIS) team for Cabarrus County consists of five members:

+ Joseph Battinelli - Team Supervisor
+ Mark McIntyre - Software Developer
+ Landon Patterson - Software Developer
+ Brittany Yoder - Software Developer
+ Marci Jones - Software Developer

Our team is responsible for software development and support for the [County](https://www.cabarruscounty.us/departments/information-technology). We work under the direction of the Chief Information Officer.

### About Laserfiche Workflow Web Services
---

The Laserfiche Workflow Web Services API allows you to run and call Workflow from any application, Web or otherwise. The problem is there is hardly any documentation out there for the Workflow Web Services, and the documentation that is out there is mainly written around .NET technologies.

In this repository, we provide examples in

+ PHP
+ Python
+ Javascript to PHP
+ Python watchdog for monitoring folders.

All these example are Windows Test Server to Windows Test Server only. Other special considerations may have to be made depending on.

+ What type of Auth system you use for the Workflow Web Services(If Any). We use NTLM auth in these examples because they are just Test Server to Test Server, so these examples have to be tweaked if using Kerberos for Auth.
+ Restructuring certain aspects of cURL in certain examples.
+ For incorporating SSPI into the mix instead of just plain ol NTLM auth, please check out this package located here https://pypi.org/project/requests-negotiate-sspi/

Each example will be a separate directory and each directory will have it's own README.md file.

### Examples we would love to see added
---

+ Rust
+ Golang
