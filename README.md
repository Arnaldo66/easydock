Développement easydock

Permet de créer des environnements de développement sur docker facilement
Inspiré de stakkr d'Emmanuel
Itulisation pour le moment des images d'Emmanuel
A voir ensuite si j'en crée des spécifique

Commencer par faire pour developpement php
Puis javascript -> node

Pour docker pb reste avec les droits entre le conteneur et dossier local.
Si fichier crée dans container il est créer en root
à Part avec l'utilisateur 1000.
Voir comment améliorer ce point


Fonctionnalités prévues:


- Création d'un nouveau projet (ex: create projectname)
- Choix environnement de dev (langage, db, elastic ... versions)
- Création des containers (start project)
- Arrêt des containers (stop project)
- Redemarrage (restart project)
- listing des container(list project)
- Listing des service possible (list service)
- Ajout de service(add service)
- Accéder au conteneur (exec containername)



Langages et services prévus:
- php à partir de la version 7 avec composer et git dans container
- apache ou nginx
- mysql ou postgresql
- elasticsearch, logstash et kibana
- traefik pour les url
- Redis, memcached
- portainer
- adminer

à Prévoir ensuite:
- Ruby
- javascript
- go
- python
- varnish
- solidity
- faire une bonne documentation
- Créer un installateur pour linux

Package go:
- un package command pour gérer les interactions avec l'utilisateur
- un package docker qui sera un wrapper des fonctionnalité docker
- un package generate qui servira a generer les fichiers pour docker
- un package nomenclature qui servira gérer le nom du projet, des images et des containers







