FROM golang:latest

LABEL maintainer="ZERO"

RUN apt-get update && \
	go get -u google.golang.org/api/calendar/v3 && \
	go get -u golang.org/x/oauth2/google && \
	go get github.com/bwmarrin/discordgo && \
	#go build ./GGCSDB/Authentication/auth.go && \
	#go build ./GGCSDB/GGCSDB.go

ADD ./GGCSDB /

CMD bash -c "go build ./GGCSDB/Authentication/auth.go && go build ./GGCSDB/GGCSDB.go && \
	/GGCSDB/Authentication/auth && ./GGCSDB/GGCSDB"
