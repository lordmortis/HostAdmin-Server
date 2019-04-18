#!/bin/sh

TIMESTAMP=`date +%s`

UP=${TIMESTAMP}_${1}.up.sql
DOWN=${TIMESTAMP}_${1}.down.sql

echo "creating '$UP'"
echo "creating '$DOWN'"

read -n1 -r -p "Press any key to continue..." key

touch datasource/migrations/$UP
touch datasource/migrations/$DOWN