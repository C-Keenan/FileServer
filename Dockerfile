FROM ubuntu:latest
RUN apt update && apt upgrade -y && apt autoremove -y
RUN apt install -y wget ca-certificates
RUN mkdir -p /app/fileserver/views
COPY ./fileserver /app/fileserver/fileserver
COPY ./views /app/fileserver/views
COPY ./static /app/fileserver/static
COPY ./entrypoint.sh /app/fileserver/entrypoint.sh
RUN chmod +x /app/fileserver/entrypoint.sh
RUN chmod +x /app/fileserver/fileserver
WORKDIR /app/fileserver
#ENTRYPOINT ["./entrypoint.sh"]
ENTRYPOINT ["./fileserver"]
EXPOSE 3000