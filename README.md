# simplebank

docker command:
````
docker run -itd -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432 --mount source=simplebank,target=/data --name postgresql postgres