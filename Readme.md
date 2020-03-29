
# Developing

## Prerequisites

  * [Golang development environment](https://golang.org/dl/)
  * Postgres 10 server

For data model changes you'll need:
  * [golang-migrate cli tool](https://github.com/golang-migrate/migrate/releases) installed and in your path
  * [sqlboiler](https://github.com/volatiletech/sqlboiler#download) installed
  
Also recommended is [fresh](https://github.com/gravityblast/fresh) for automatic recompiling.

## Downloading 

## Data Model

### Domains
  - all domains have a base TLD which is the name (i.e. `sektorseven.net`)
  - domains can have one or more user accounts associated, with permissions
    - by default, if a user account is associated with a domain, it can _see_ all the information
    - admin permission allows users to be added to the domain
    - email permission allows the modification of email users and forwards    

## Config file setup

  1. Copy `sqlboiler.sample.toml.sample` and `config.sample.json` to new files - removing the `.sample`
  2. You'll need to at least change the Database Name (`<DATABASE>`), DB User (`<DBUSER>`) and DB Password (`<DBPASS>`) in those files to a suitable config for your system.
     * It's a good idea to use a unique database and user and a random password to avoid issues with your dev environment
  3. You may need to change the host and port entries to reflect your local setup. The example fields in `config.sample.json` are the defaults. 
  4. `config.sample.json` is used by the live application, while `sqlboiler.sample.toml` is only used to generate new datamodels from database data.
  5. the root "development" flag is only used to perform migrations from files.  
  
## Helper tool 

In the subdirectory `helperTool` of this repo is package that compiles a helper tool.
This should be built via `go build` and the executable moved into your path.

  - By default the helper tool will look for your config file in the working directory
  - By default the helper tool will use the postgres user `postgres` (this needs to be a root user)
  - By default the helper tool will use the postgres password `rootpassword`
  - If you've changed the docker config these may need to be specified - run the tool with the `/h` option to get help

The following sections will make use of the helper tool to create databases and users.
  
## Database Setup

  1. Setup the config file with the correct:
     1. database host
     2. database port
     3. the desired database name to use (this will be created and if this exists, it will be overwritten)
     4. the desired user to user (this will be created and if this exists, it will be overwritten)
     5. the desired user's password
  2. run the helper tool with the `createDB` command
  3. if no errors are encountered, the base setup is complete!

## Creating a user

  1. You'll need to have completed the database setup step above.
  2. You'll need to have run the server at least once to have crated the required tables
  3. run the helper tool with the `createUser` command. This command takes the following parameters:
      1. `username` - the desired username for the new user
      2. `password` - the password for the new user
      3. `email` - the email for the new user
      4. (optional) `superadmin` - if the user is a super/root administrator. defaults to false
  4. if no errors are encountered, the user was created!

## Changing data models

  1. Ensure you have the development flag turned on to test your migration scripts
  2. run the helper tool with the `createMigration` command. This command takes the following parameters:
      1. `directory` - the root directory of this repository (if the helpertool is being run in the helperTool directory, 
      then this will be either `..\` or `../` depending on your os)
      2. `name` - the name of the migration to create. Should be related to what the migration does, should probably be 
      enclosed in quotes. The tool will replace whitespace with `_`, and non-alpha numeric characters with `-` 
  3. Specify the migration syntax in the new files created in the `datasource/migrations/` directory
  4. Test your migration via either:
      1. compiling and running the server with the development flag turned on
      2. manually running the migration via `migrate -source file://datasource/migrations -database postgres://<DBUSER>:<DBPASS>@localhost:15432/<DBNAME>?sslmode=disable up`
  5. Once your migration has the correct SQL, regenerate the bindata using the helperTool's `updateBindata` command
      1. `directory` - the root directory of this repository (if the helpertool is being run in the helperTool directory, 
      then this will be either `..\` or `../` depending on your os)
  6. Then generate the models:
     1. Ensure you have copied `sqlboiler.sample.toml` to `sqlboler.toml` and changed the various directives to your local ones.
     2. Run `sqlboiler psql -o datamodels_raw -p datamodels_raw -c sqlboiler.toml` from the root of the repo.