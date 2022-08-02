---
### Starting 

#### Before the start: 
* If you **do not have postgres image** in your docker,
  use the following command:

  ```docker pull postgres```

* After this **run the docker postgres container** using makefile
in this repository or run the following command:

  ```docker run --name=dbname -e POSTGRES_PASSWORD='mypassword' -p 5432:5432 -d --rm postgres```
* In next step you need to up **sql migrations**, you also can use makefile: 

  ```make m-up ``` 
  
  or 

  write following command in terminal: 

  ```migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5434/postgres?sslmode=disable' up```

  If you run container using another port, change in this command port from 5434 on yours.

* Now you have started your container and postgres on it. If you want check data in you database, you can use makefile: 

  ```make d-exec```

  or 

  use this command in terminal:

  ```docker exec -it containerID /bin/bash```

  **for running this command you need to get your container id using command:**

  ```docker ps```

  And put in container id into command.

### Now you are ready to work with database
