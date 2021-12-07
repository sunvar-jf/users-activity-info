# users-activity-info

**User Login Information & User Last Login Time**

This tool will help to get User's Last Login Information and also will get the login activities of users over 2 year time. This tool will help an organisation to check the user activity for Artifactory. Also the tool will return the last login date time for an user. This will help in audit of users too.

**How does this program work?**

User provides multiple access-request.log(rename files but with .log extension only) of artifactory as an input for the program. The program parses the log and gets unique users with their login information. Along with the access-request.log user has to provide the path of the graphoutput.xlsx file as an input for the program too. The output graphs and relevant data shows up in the excel sheet.

**How to run the tool**

1.  Download the code from git.
2.  Run the following to build the code: 
    go build src/main/main.go
3.  ./main (folderpath for access-request.log) (file path for the excel sheet i.e graphoutput.xlsx)


**What can be improved further?**

Parsing of access-request.log inside a zip file. Search for users logged on a specific day or between a time range. The current logic in code can be easily extended to achieve the day or time ranged based logins easily.

Enjoy Reading and Coding!

