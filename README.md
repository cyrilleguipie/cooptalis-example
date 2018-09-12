# Golang REST API 

Cette API en [Golang](http://www.golang.org/) permet à un utilisateur de :
- récupérer de la liste des menus, des collaborateurs et des clients.
- retrouver un collabateur à partir de son identifiant (ID)
- modifier certains champs (Pays, Client, Statut)

Le SGBD utilisé dans ce projet SQLite, le code source est formaté selon le standard <code>gofmt</code> 
   

## Installation (locale)
   
   - Installer [Golang](https://golang.org/dl/)
   - Installer [Revel](https://revel.github.io/)
   
    go get github.com/revel/revel
    go get github.com/revel/cmd/revel
    
   - Installer [Gorm](http://gorm.io/),une bilbliothèque ORM pour Golang
   
    go get github.com/jinzhu/gorm
   - Installer [Casbin](https://casbin.org/), 
une bibliothèque d'autorisations qui prend en charge des modèles de contrôle d'accès tels que ACL, RBAC, ABAC pour Golang, Java, PHP et Node.js.

    go get github.com/casbin/casbin
   - Installer BCrypt
    
    go get golang.org/x/crypto/bcrypt
   
   - Se placer dans le repertoire $GOPATH/src, et lancer
    
    git clone https://github.com/cyrilleguipie/cooptalis-example.git
    cd cooptalis-example
    
   
   
   
  

#### Execution    
   Une fois l'application installée et parametrée, depuis la repertoire de l'application <code>$GOPATH/src/cooptalis-example</code> 
   vous démarrez le serveur, en executant : <code>revel run</code>
    
   
   - Dans votre navigateur Web, allez sur <code>http://localhost:9000/</code> et vous verrez :

    "Cooptalis API"

GET     /menu                                   Menus.List
GET     /collaborateurs                         Collaboraters.List
GET     /collaborateur/:id                      Collaboraters.FindById
POST    /collaborateur/:id                      Collaboraters.Update
## API
  - [GET /menu](#get-menu) retourne la liste des menus (<code>parentId = 0</code>) et sous-menus visibles (<code>parentId != 0</code> qui contient la valeur de id du menu parent) 
  selon le token (quelques tokens pour le POC <code>ADMIN123, MEMBER123, ANONYMOUS123</code>)
  - [GET /collaborateurs](#get-collaborateurs) retourne la liste des collaborateurs avec paramètres de filtrage <code>k:mot cle, sF:champ de filtrage, sD:ordre de filtrage </code>
  - [GET /collaborateur/[id]](#get-collaborateurid) retourne les détails d'un collaborateur selon son <code>id</code>
  - [POST /collaborateur/[id]](#post-collaborateurid) modifier des champs(pour le POC <code>country, client, status</code>) d'un collaborateur selon son <code>id</code>

### GET /menu
Exemple

    curl -H "Authorization:ADMIN123" http://localhost:9000/menu
    

Reponse:

    {
      "success": true,
      "data": [
        {
          "ID": 4,
          "CreatedAt": "2018-09-11T16:23:08.034892Z",
          "UpdatedAt": "2018-09-11T16:23:08.034892Z",
          "DeletedAt": null,
          "title": "Clients",
          "type": "list",
          "apiURL": "/clients",
          "iconURL": "http://i.imgur.com/BuaklRB.png",
          "parentID": 0
        },
        {
          "ID": 1,
          "CreatedAt": "2018-09-11T16:23:08.017837Z",
          "UpdatedAt": "2018-09-11T16:23:08.017837Z",
          "DeletedAt": null,
          "title": "Collaborateurs",
          "type": "list",
          "apiURL": "/collaborateurs",
          "iconURL": "http://i.imgur.com/DRuOeD0.png",
          "parentID": 0
        },
        {
          "ID": 3,
          "CreatedAt": "2018-09-11T16:23:08.032374Z",
          "UpdatedAt": "2018-09-11T16:23:08.032374Z",
          "DeletedAt": null,
          "title": "Immigration",
          "type": "list",
          "apiURL": "/collaborateurs?k=immigration",
          "iconURL": "http://i.imgur.com/EDCW4OD.png",
          "parentID": 1
        },
        {
          "ID": 2,
          "CreatedAt": "2018-09-11T16:23:08.030131Z",
          "UpdatedAt": "2018-09-11T16:23:08.030131Z",
          "DeletedAt": null,
          "title": "Relocation",
          "type": "list",
          "apiURL": "/collaborateurs?k=relocation",
          "iconURL": "http://i.imgur.com/sRK4tw3.png",
          "parentID": 1
        }
      ],
      "message": ""
    }
    
    
### GET /collaborateurs
Exemple

    curl http://localhost:9000/collaborateurs
    

Reponse:

    {
      "success": true,
      "data": [
        {
          "ID": 3,
          "CreatedAt": "2018-09-11T16:23:08.291464Z",
          "UpdatedAt": "2018-09-11T16:23:08.291464Z",
          "DeletedAt": null,
          "firstname": "Eric",
          "lastname": "Cartman",
          "dateOfEntry": "2018-09-11T16:23:08.291204Z",
          "country": "MAROC",
          "job": "Android Developer",
          "client": "SOPRA",
          "state": "relocation"
        },
        {
          "ID": 4,
          "CreatedAt": "2018-09-11T16:23:08.293225Z",
          "UpdatedAt": "2018-09-11T16:23:08.293225Z",
          "DeletedAt": null,
          "firstname": "Kenny",
          "lastname": "Broflovski",
          "dateOfEntry": "2018-09-11T16:23:08.293055Z",
          "country": "MADAGASCAR",
          "job": "Python Developer",
          "client": "SOPRA",
          "state": "immigration"
        },
        {
          "ID": 2,
          "CreatedAt": "2018-09-11T16:23:08.288853Z",
          "UpdatedAt": "2018-09-11T16:23:08.288853Z",
          "DeletedAt": null,
          "firstname": "Kyle",
          "lastname": "Broflovski",
          "dateOfEntry": "2018-09-11T16:23:08.28863Z",
          "country": "COTE D'IVOIRE",
          "job": "Go Developer",
          "client": "ATOS",
          "state": "immigration"
        },
        {
          "ID": 1,
          "CreatedAt": "2018-09-11T16:23:08.28646Z",
          "UpdatedAt": "2018-09-11T16:23:08.28646Z",
          "DeletedAt": null,
          "firstname": "Stan",
          "lastname": "Marsh",
          "dateOfEntry": "2018-09-11T16:23:08.286282Z",
          "country": "COTE D'IVOIRE",
          "job": "Java Developer",
          "client": "ATOS",
          "state": "relocation"
        }
      ],
      "message": ""
    }
    
### GET /collaborateur/[id]
Exemple

    curl http://localhost:9000/collaborateur/1
    

Reponse:

    {
      "success": true,
      "data": {
        "ID": 1,
        "CreatedAt": "2018-09-11T16:23:08.28646Z",
        "UpdatedAt": "2018-09-11T16:23:08.28646Z",
        "DeletedAt": null,
        "firstname": "Stan",
        "lastname": "Marsh",
        "dateOfEntry": "2018-09-11T16:23:08.286282Z",
        "country": "COTE D'IVOIRE",
        "job": "Java Developer",
        "client": "ATOS",
        "state": "relocation"
      },
      "message": ""
    }
    
### POST /collaborateur/[id]
Exemple

    curl -H "Content-Type: application/json" -d '{"country":"FRANCE","client":"SOPRA"}' http://localhost:9000/collaborateur/1
    

Reponse:

    {
      "success": true,
      "data": {
        "ID": 1,
        "CreatedAt": "2018-09-11T16:23:08.28646Z",
        "UpdatedAt": "2018-09-12T13:12:13.984495Z",
        "DeletedAt": null,
        "firstname": "Stan",
        "lastname": "Marsh",
        "dateOfEntry": "2018-09-11T16:23:08.286282Z",
        "country": "FRANCE",
        "job": "Java Developer",
        "client": "SOPRA",
        "state": "relocation"
      },
      "message": ""
    }
    
    
##Tests


Revel fournit un module de test. Dans ce projet, les tests sont dans le fichier <code>./tests/apptest.go</code>

- Lancer les tests depuis le navigateur sur <code>/@tests</code>
- Lancer les tests en mode commande:
    
    - Pour executer tous les tests de l'application
        
            revel test @GOPATH/src/cooptalis-example dev
            
    - Pour executer un jeu de test ou une methode de test
            
           revel test @GOPATH/src/cooptalis-example dev AppTest
           revel test @GOPATH/src/cooptalis-example dev AppTest.TestThatIndexPageWorks
           
      
     
##Integration Continue, Deploiement Continu 
(***Encore en test***)

####Architecture

   AppEngine, Docker, Github, Travis CI