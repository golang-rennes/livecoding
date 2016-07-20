# THIS IS SAMPLE APP

# Périmètre
 - DB Postgres
 - API REST
 - Gestion des logs
 - Client CLI
 - Gestion de la conf
 - Packaging avec Docker

# Appli Web
## API Standard

15 lignes pour démarrer un server web avec une route
==> ./withoutGin/main.go

## Gin-Gonic

Framework Web :
 - Zero allocation router : httprouter (https://github.com/julienschmidt/httprouter) Mux dans le monde Go(lang)
 - Battle tested
 - Productif
 - "Crystal Clear" API
 - Middleware support (martini/contrib)
 - Crash free (automatic panic recovery)
 - JSON Validation
 - Routes Grouping
 - Built-in rendering (json, xml, etc.)
 - ...

## Design de l'appli
 ````
 +--------------+    +-------------------------+             +---------------------+
 |              |    |                         |             |                     |
 | main()       +---->  config                 |             |  models             |
 |              |    |                         |             |                     |
 +---------+----+    +-------------------------+             +----------+----------+
           |                                                            ^
           |         +-------------------------+             +----------+----------+
           |         |                         |             |                     |
           +--------->  database connection    <-------------+  controllers        |
           |         |                         |             |                     |
           |         +-------------------------+             +----------+----------+
           |                                                            ^
           |         +-------------------------+             +----------+----------+
           |         |                         |             |                     |
           +--------->  web server             +------------->  routes             |
                     |                         |             |                     |
                     +--------------+----------+             +---------------------+
                                    |
                                    |
                            +-------+----------+
                            |                  |
                            | .start()         |
                            |                  |
                            +------------------+
````

### Identification des packages
- main
  - models : les objets métiers, la connexion à la DB
  - server : le server lui même, les routes, ....
  - controllers : la logique de gestion des données métiers
  - config : la config accessible partout (injection de dépendances, tests, blabla)

### Let's code
 1. la config
   - ./config/config.go
 2. le server qui démarre avec une route ANY /
   - ./server/server.go
   - ./server/routes.go
   - main.go
 3. le modele User
   -  ./users/user.go
 4. le controller User avec un get /users qui renvoie une liste vide (rendering c.JSON())

## Accès aux données avec GORM
"The fantastic ORM library for Golang, aims to be developer friendly."

### Let's code
5. La connexion à la BDD, gestion du pool de connexion, etc...
  - ./models/db.go
  - démarrage d'un conteneur postgres (merci docker pour mac)
6. auto-migration, initialisation du modèle
7. implémentons le /users et les autres controlleurs
8. BindJSON, formdata, gestion des erreurs, etc...
9. Authentification basique
10. Groupes

# Gestion des logs

# Configuration de l'appli

# Un client en ligne de commande (CLI)

# Packaging et lancement dans un conteneur docker
