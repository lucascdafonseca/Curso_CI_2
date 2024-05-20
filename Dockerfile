FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV HOST=localhost PORT=5432

ENV USER=root PASSWORD=root DBNAME=root 

COPY ./main.exe .

#ENTRYPOINT ou CMD têm a mesma função à princípio
CMD [ "./main.exe" ]