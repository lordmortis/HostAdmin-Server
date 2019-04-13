
# Developing

## Prerequisites

  * [Golang development environment](https://golang.org/dl/)
  * Postgres 10 server

## Downloading 

## Config file setup

  1. Copy `sqlboiler.sample.toml.sample` and `config.sample.json` to new files - removing the `.sample`
  2. You'll need to at least change the Database Name (`<DATABASE>`), DB User (`<DBUSER>`) and DB Password (`<DBPASS>`) in those files to a suitable config for your system.
     * It's a good idea to use a unique database and user and a random password to avoid issues with your dev environment
  3. You may need to change the host and port entries to reflect your local setup. The example fields in `config.sample.json` are the defaults. 
  4. `config.sample.json` is used by the live application, while `sqlboiler.sample.toml` is only used to generate new datamodels from database data.  
  
## Database Setup

  1. You'll need to have a postgres install to connect to, and the ability to log in as the  `postgres` user.
  2. Create the login role from the settings you did above: 
     * `CREATE ROLE "<DBUSER>" WITH PASSWORD '<DBPASS>' LOGIN` (the difference in quoting is intentional)
  3. Create the database via:
     * `CREATE DATABASE "<DBNAME>" OWNER "<DBUSER>"`
  4. The application should create any apropriate database fields on bootup!
