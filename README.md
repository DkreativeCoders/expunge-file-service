# Expunge File Service
A service that automates deleting of files in a specified manner 

# Getting Started
1. [Download the exe file here](https://drive.google.com/file/d/16Vhs0YDAsZcdehs96upr8Ca7vOoavTt7/view?usp=sharing
   )

2. Create a config folder in the root directory of the project
   
3. Set up your env file in the root directory of the project
   
4. Replicate the FileCleaner.Json file and configure

5. start the service by running the service name on cmd

6. To do a one time clean up hit the end <code>http://localhost:8600/api/v1/flush</code>


# Setting up env file

```go

    PORT =8600
    CONFIG.PATH=config\FileCleaner.json
    LOG.NAME=expungeservice.log
    ENABLE.CRONJOB=true
    CRON.INTERVAL=0/2 * * * *

```

set up the port the service should run on

configure the FileCleaner.json file path

configure the name of the log file 

set ENABLE.CRONJOB to true if you desire the service to run at a scheduled timing

set CRON.INTERVAL to the cron format


# Setting up FileCLeaner.Json
1. Set up general configs
    ```json
   "generalConfig": {
   "fileAgeLastModifiedInDays": 10,
   "enableFileMovementToBackupFolder": true,
   "pathToBackupFile": "path\to\backup\folder"
   }
   ```
   
    **fileAgeLastModifiedInDays** defines the age of the files that should be considered to be deleted
   
    **enableFileMovementToBackupFolder** set this to true if you want to back up the files before deleting
   
    **pathToBackupFile** set the path to the back up folder
   
2. set up service Configs

     ```json
           "serviceConfigs": [
            {
            "serviceName": "ebills",
            "rootPath": "path\to\backup\folder",
            "excludeFolders": ["ebooks","scr","best"],
            "excludeExtensions": [".jar",".bat",".DS_Store",".pdf"],
            "excludeSpecificFileNames" : ["do-not-delete","delete-not","intermediate code generation part 1"],
            "excludeFileNamesContaining" : ["tomcat","tomcat"],
            "enableRecursiveDepth" : false,         
            "useGeneralConfig": false,
            "fileAgeLastModifiedInDays": 5,
            "enableFileMovementToBackupFolder": false,
            "pathToBackUpFolder":"C:\\Users\\dell\\Documents\\daniel\\backupForEbillService"
            },
            {
            "serviceName": "another service",
            "rootPath": "path\to\backup\folder",
            "excludeFolders": ["ebooks","scr","best"],
            "excludeExtensions": [".jar",".bat",".DS_Store",".pdf"],
            "excludeSpecificFileNames" : ["do-not-delete","delete-not","intermediate code generation part 1"],
            "excludeFileNamesContaining" : ["tomcat","tomcat"],
            "enableRecursiveDepth" : false,         
            "useGeneralConfig": false,
            "fileAgeLastModifiedInDays": 5,
            "enableFileMovementToBackupFolder": false,
            "pathToBackUpFolder":"C:\\Users\\dell\\Documents\\daniel\\backupForEbillService"
            }
            
        ] 
     ```

This contains an array of setup that would be processed in the order of insertion


**serviceName** refers to the id 
   
**rootPath** this specifies the directory the deleting process should begin from
    
**excludeFolders** this specifies the directory names that should be ignored in the deleting process

**excludeExtensions** this specifies the extensions of files that should be ignored in the deleting process
   
**excludeSpecificFileNames** this defines specific files names that should be ignored in the deleting process, the file full name must be added. 

**excludeFileNamesContaining** this specifies files names that should be ignored in the deleting process, base on the meeting certain conditions in the file name. 
   e.g a file named "web dev series.mp4" would be ignored if "series" is added to excludeFileNamesContaining

**enableRecursiveDepth** this allows the system search through folders recursively found in its path in the deleting process

**useGeneralConfig** if this is set to true it would ignore the service configuration for 
"fileAgeLastModifiedInDays,enableFileMovementToBackupFolder,pathToBackUpFolder" and use the general configuration

**fileAgeLastModifiedInDays** defines the age of the files that should be considered to be deleted

**enableFileMovementToBackupFolder** set this to true if you want to back up the files before deleting

**pathToBackupFile** set the path to the back up folder
