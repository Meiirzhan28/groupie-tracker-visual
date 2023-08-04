# Groupie-tracker-visualizations
## Description
Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.


## Links
 - [Good practices.](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md)
 - [What is a docker](https://selectel.ru/blog/what-is-docker/#whatis)
 - [How to write a Good readme](https://readme.so/editor)
## Authors:
- @gesti9
- @Meiirzhan28

## Usage:
make run 

      Запуск сервера на http://127.0.0.1:3000

## For the dockerfile:

1.Create image for Dockerfile

  
    docker build -t tracker .

2.Enter this command to start the program

    docker run --name raz -dp 4000:3000 tracker
3.Open the web browser and go to

    http://127.0.0.1:4000/
    