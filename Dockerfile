FROM golang:latest

RUN apt-get update && \
	go get -u google.golang.org/api/calendar/v3 && \
	go get -u golang.org/x/oauth2/google && \
	go get github.com/joho/godotenv && \
	go get github.com/bwmarrin/discordgo && \
	go get github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Authentication && \
	go get github.com/ZEROssk/GCS-get-Discord-bot/GGCSDB/Get-Schedule

ADD ./GGCSDB /go

CMD ["go", "run", "GGCSDB.go"]
