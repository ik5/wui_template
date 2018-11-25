# WUI Template

The following source repo is a bootstrap for creating a server side Web system
using the Go (Golang) language.

It is set of directory and file structure that does all the repeated bootstrap
to get started using a web server.


## Why?

Go has many HTTP servers, some of the servers known more, others known less.
All of them require some boilerplate code, that does some bootstrapping in order
to get started.

While tools such as [Buffalo](https://gobuffalo.io/) that provides a set of tools, they are very complete.
This tool is just the start, that allows the develop to choose most things, except for:

  - Database
  - Ability to work with API
  - Configuration system

The following template provides the above, however, they are not 'set in stone', and
can be changed if/when needed.


## Used packages

The template is using the following packages:

  - [Chi](https://github.com/go-chi/chi) - Web server
  -  [Pq](https://github.com/lib/pq) - Postgresql driver
  - [Sqlx](https://github.com/jmoiron/sqlx) - Extended [database/sql](https://golang.org/pkg/database/sql/)
  - [viper](https://github.com/spf13/viper) - configuration system
  - [Logrus](https://github.com/sirupsen/logrus) - Logging system
  - [dep](https://github.com/golang/dep) - Go dependency management tool


