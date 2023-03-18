# School Manager API

A simple School management system with fluent CRUD endpoints for the entities (Students and Courses) and admin endpoints that allow you to do the following: 
- Add courses to a student
- Remove courses from a student
- List courses taken by a student
- List students taking a particular course
- Update list of courses taken by a student(bulk updates)

## Prerequisites

- Docker must be installed on your Environment

## SetUp
1. git clone the project from the master branch
2. create a new branch from the master branch
3. Create the Docker Network

```bash
docker network create "sch-manager-network"
```
4. Run "docker-compose" up to start the containers.
```bash
docker-compose up
```

## Documentation

To view the documentation and test all the endpoints, open the following page in the browser with the containers running: 

[School Manager Open API Documentation](http://localhost:8001/docs/index.html)



## Notes:
Please note that the following ports must be free in your network:
* 8080
* 8001  
* 3306
* 3307

Those ports are bound by docker and are required for the containers to run. 

