FROM scratch

MAINTAINER Ryan Brewster <ryanpbrewster@gmail.com>

ADD hello /Production/hello/hello.exe

EXPOSE 80

CMD ["/Production/hello/hello.exe"]
