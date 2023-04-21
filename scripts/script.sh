#!/bin/bash

touch users.csv
echo "id,name,last_name" > users.csv

for i in {1..10}; do
  id=$(uuidgen)
  name="user_${i}"
  last_name="last_name${i}"
  echo "$id,$name,$last_name" >> users.csv
done