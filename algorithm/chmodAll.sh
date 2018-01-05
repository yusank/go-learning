#!/bin/bash

for dir in `ls .`
do 
	if [ -f $dir ]
	then
		chmod 644 $dir
	fi
done
