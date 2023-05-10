#!/bin/bash

touch users.csv
echo "id,name,last_name" > users.csv

for i in {1..100}; do
  id=$(uuidgen)
  name="user_${id}"
  last_name="last_name${id}"
  echo "$id,$name,$last_name" >> users.csv
done