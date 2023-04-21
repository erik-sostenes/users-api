# users-api

REST API that receives via http request a .csv file with user data, which will be stored in a Mongo database

### what start ?
``
run the following scripts from the script folder
``

*Step 1*
```
./script.sh
```
*Step 2*
```
curl -X POST localhost:8080/v1/users
   -d @scripts/script.sh 
   -H 'Content-Type: text/csv'
```

<table>
    <tr>
        <th>Name</th>
        <th>Purpose</th>
        <th>Optional</th>
        <th>Default value</th>
    </tr>
    <tr>
        <td>PORT</td>
        <td>Server Port Setting</td>
        <td>YES</td>
        <td>8080</td>
    </tr>
</table>