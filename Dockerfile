FROM debian:buster
RUN mkdir /var/phil
COPY phil /var/phil/phil
EXPOSE 8080
CMD ["/var/phil/phil"]
