# FamilyBudget example app

A simple sample application that allows user to manage budgets. It allows user to add and remove any number of budgets, add two types of transactions (income and expenses) with categories and share your own budgets with other users.

### Important notes

The application is in constant development. Golang app backend (app2 and app3) is not set as default yet.

## How to install in local environment

### Environment

You must have Docker installed on the machine where you want to run the sample environment.

If you want to run the development environment, you will need at least:
* Python 3.9 with requirements.txt
* Postgres 13
* Node.js 17.1.0
* npm 8.1.2

### Running sample env

#### Build frontend

Before running sample environment, you will need to build frontend for production.
Go to the frontend directory:

`cd family_budget/frontend/budget_app`

Run the following cmd:

`npm run build`

#### Create *run* directory

Create `run` directory inside root project dir, with `etc/db_password`, `etc/envfile.env` files and `data/postgres` directory:

<img width="247" alt="image" src="https://user-images.githubusercontent.com/54007065/204387176-9fcf92bd-bd3c-40b3-9632-45e7eaaf3f97.png">

<img width="310" alt="image" src="https://user-images.githubusercontent.com/54007065/204386203-25d2d3f7-9765-45e6-a010-e5c855c88b8c.png">

`envfile.env` file should contain backend environment vars, eg:

DB_NAME=budget_app

DB_USER=postgres

DB_PASSWORD=postgres

DB_HOST=postgres

DB_PORT=5432

`db_password` file simply contains only database password, eg: `postgres`.

#### Start the environment

`cd family_budget/frontend/budget_app`

`docker compose up`

The application should available on `http://localhost:52000`

#### Login to the app

The application contains initial data with which you can log in.
There are three example users:

`username1:password1`

`username2:password2`

`username3:password3`

Of course, you can also create your own account and use it to log in.
