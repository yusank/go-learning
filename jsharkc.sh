#!/bin/sh
# author by jsharkc
arr=`ls`

for var in ${arr[*]}
do
 if [ -d $var ];then
  echo "$var is dir"
 elif [ -f $var ]; then
  echo "$var is file"
 fi
done
