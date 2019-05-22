FROM golang:latest

RUN apt-get update && \
	go get -u google.golang.org/api/calendar/v3 && \
	go get -u golang.org/x/oauth2/google && \
	go get github.com/joho/godotenv && \
	go get github.com/bwmarrin/discordgo

ADD ./GGCSDB /go

CMD bash -c "go run Authentication/auth.go"
#"go run GGCSDB.go"
