FROM ubuntu:18.04

RUN apt-get update
RUN apt-get install wget -y

RUN wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz 


RUN apt-get install -y build-essential libssl-dev libseccomp-dev libpcap-dev libelf-dev
# if which go; then echo "Skipping go (already installed)"; else ${SUDO} apt install -y golang-go; fi; \
# 			if which gcc; then echo "Skipping gcc (already installed)"; else ${SUDO} apt install -y build-essential; fi; \
# 			if which cmake; then echo "Skipping cmake (already installed)"; else ${SUDO} apt install -y cmake; fi; \
# 			if which flex; then echo "Skipping flex (already installed)"; else ${SUDO} apt install -y flex; fi; \
# 			if which bison; then echo "Skipping bison (already installed)"; else ${SUDO} apt install -y bison; fi; \
# 			if which pkg-config; then echo "Skipping pkg-config (already installed)"; else ${SUDO} apt install -y pkg-config; fi; \
# 			if [ -d /usr/share/doc/libssl-dev ]; then echo "Skipping OpenSSL headers (already installed)"; else ${SUDO} apt install -y libssl-dev; fi; \
# 			if [ -d /usr/share/doc/libseccomp-dev ]; then echo "Skipping libseccomp headers (already installed)"; else ${SUDO} apt-get install -y libseccomp-dev; fi; \
# 			if [ -d /usr/share/doc/libpcap-dev ]; then echo "Skipping libpcap-dev headers (already installed)"; else ${SUDO} apt-get install -y libpcap-dev; fi; \
# 			if [ ! -f /usr/include/libelf.h ]; then \
# 				${SUDO} apt-get install -y libelf-dev; \
# 			fi; \

WORKDIR /tcpdump
COPY pkg /tcpdump/pkg
COPY cmd /tcpdump/cmd
COPY go.mod /tcpdump/go.mod 
COPY go.sum /tcpdump/go.sum
COPY main.go /tcpdump/main.go
COPY /run.sh /run.sh
RUN chmod 755 /run.sh

#RUN CGO_LDFLAGS="-static -w -s -Wl,--dynamic-linker=/vorteil/ld-linux-x86-64.so.2 -Wl,-rpath,/vorteil" /usr/local/go/bin/go build -v -tags netgo

CMD ["/run.sh"]