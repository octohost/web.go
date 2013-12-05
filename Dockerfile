FROM octohost/go-1.2

RUN go get github.com/hoisie/web
RUN mkdir /srv/www/
RUN cd /srv/www/; curl https://raw.github.com/octohost/web.go/master/hello.go -o hello.go

EXPOSE 9999

CMD cd /srv/www; go run hello.go

