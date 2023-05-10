# users-api

REST API that receives via http request a .csv file with userHandler data, which will be stored in a Mongo database

### how to perform acceptance testing?
``
Run the following internal/apps/backend/handlers/users/scripts from the script folder
``

*Step 1*
```
./script.sh
```
*Step 2: run the image of docker*
```
docker mongodb
```

*Example of how  to perform a request*
```
curl -X POST localhost:8080/v1/users
   -d @scripts/script.sh 
   -H 'Content-Type: text/csv'
```