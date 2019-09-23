# Microservices
---

Les grands principes de programmation ne s'envolent pas avec les microservices

---

Le but d'une __bonne architecture__ n'est pas de trouver la solution ultime (elle n'existe pas) mais de __permettre de toujours pouvoir évoluer progressivement__

---

## Unix philosophy

---

> Write programs that do one thing and do it well. 
>
> Write programs to work together. 
>
> Write programs to handle text streams, because that is a universal interface.

__Douglas McIlroy (1978)__

---

@ul
- OS = Cluster
- programmes = containers / unité de déploiement
- systeme = script
- Text stream interface = HTTP
@ulend

---
## Monolith vs Microservices vs monolith distribué
---

Monolithe : clean architecture, Architecture hexagonale, ports & adapters

@ul
- Separation stricte des niveaux d'abstraction et de responsabilité
- Inversion of Control (dépendance et abstraction pointent dans la même direction)
@ulend

---

Ces principes deviennent vitaux au niveau d'une architecture microservices

Note:
- vital au niveau du système, moins au niveau du service lui-même

---

sinon, on construit un monolithe distribué

---


services de base, services métier (réalisent une User Story en utilisant les services de base)
penser résilience, pas DRY (couplage)

- chaque microservice peut etre déployé indépendamment des autres
- un service de base n'a aucune connaissance des autres services
- un service métier n'a aucune conscience de la localisation et de l'implémentation des autres
- chaque microservice peut être codé dans un langage différent (faire un choix en fonction du probleme, du moment, de la team)
- il est assez petit pour que l'on puisse le recoder from scratch sans crainte
- il ne fait qu'une seule chose et la fait bien

Note:
- service de base : n'a pas a connaitre son usage, reste simple 
- service métier : injection de dépendance via paramétrisation des hosts qui implémentent une interface


---

## Systemes distribués
- Ne pas croire dans les "network fallacies"
- Ca va fail, c'est certain : messages malformés, drop d'une partie des messages, messages reçus plusieurs fois
- Tout doit étre automatisé : pas de script à la main : trouver les outils les plus simples permettant de faire le travail sans customization
- la CI test l'applicatif en black box (test d'intégration), le chaos engineering test le systeme
- le code est éventuellement on teste le code avec
- update distribué : problème des généraux byzantins (transacation distribuée, 2PC, 3PC, saga ?)

## Methodes de travail
- simplifier la CI ! (évitez le mono-repo). Le point de contact entre votre code et votre stack, c'est l'artefact repository
- faire du trunk-based developpement : si vous pensez que votre projet est trop gros pour cela : c'est que vous n'etes plus en train de faire un microservice
- doit étre idempotent
- CQRS peut aider a rester simple 

## Vous allez devoir debugger un code réparti sur plusieurs machines, il faut de nouveaux outils :
### Detecter
- alerting based on machine & services monitoring
### Debug & improve performance
- Distributed tracing (observability)
- Logging centralisé, 


## Abstraction du routing
Service discovery : vous ne pointez pas vers une instance mais vers n instances healthy

## Sécurité :
Zero-trust networks
Docker image scanning, distroless base images
Host security

Vous devez gérer les retries & le circuit-breaking, mais cela ne doit pas être au niveau du code, ce n'est pas son niveau d'abstraction

## CI/CD ?
Integration : création d'une unité de déploiement : artefact auto-suffisants
Déploiement : run d'un artefact dans la stack
Release : routing du traffic vers un service
