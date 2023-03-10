# client
client for a postgresql database

Author: Kris Bruenn

'client' is a wrapper for postgresql, so that you can create databases, etc., without using SQL.
It is written in Golang. As you create your database 'name', client creates a 'name'.sql file 
containing the appropriate SQL commands, so you can instantiate your database on another computer.

This is my first repository in Github, so I am still figuring out things like license, etc.

Installation:
-------------

1) You need to install postgresql on your computer

2) You need to install Golang on your computer

3) You need to install this repository on your computer

Note: you need to have a main package in order for Go to run your client.
