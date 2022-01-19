#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:company .
#docker run -p 8080:8080 -d mercurius:company

FROM scratch

ADD company /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale


#to be added to mercurius
EXPOSE 8080
#end to be added to mercurius
CMD [ "/company" ]
