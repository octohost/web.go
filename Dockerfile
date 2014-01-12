FROM octohost/go-1.2

RUN go get github.com/hoisie/web
RUN mkdir /srv/www/
ADD . /srv/www

EXPOSE 9999

CMD cd /srv/www; go run hello.go

