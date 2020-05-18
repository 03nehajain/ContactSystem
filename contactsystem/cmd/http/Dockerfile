FROM centos
LABEL Author="neha.jain@exotel.in"
ADD contactsystem.tar.gz /usr/local
WORKDIR /usr/local/build
RUN set -x ; yum update -y ; yum install make -y; make install
EXPOSE 80
CMD ["/usr/local/target/contactsystem", "-port", "80", "-host","0.0.0.0"]