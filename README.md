# hotel web application

### Getting start

if you want to use in development mode
first of all set the database name and pass with defualt
value (see in app files and codes)
and if you want to use in production mode
set -dbname -dbuser -dbpass and -dbport flags when
you run the app

```
    git clone https://github.com/ErfiDev/hotel-web-app.git
    cd hotel-web-app
    ./run.bat
```

### Authentication method

in this application I use the session-based authentication
with scs package
and CSRF token for authenticate POST method requests

### Enviroment variables

if see the .env.example file you can see all neccesary variables and then create a .env file and set all of those
then see the .yml.example file and create database.yml file and set all neccesary variables

### Created by

developer: erfanhanifezade@gmail.com

go programming language and std libraries
scs, nosurf and godotenv pkg
postgresql database
soda cli tool
pgx connection for postgresql connection

### Contribute Section

if you're interested for adding some feature or fixing bug
can create a pull request and i am merge that, thanks
