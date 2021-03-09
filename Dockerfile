FROM centos 
WORKDIR /build/
COPY uploadvoice /build/
COPY conf	/build/conf/
EXPOSE  8081
ENTRYPOINT ["/build/uploadvoice"]