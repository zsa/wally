FROM centos:centos8

RUN dnf install -y dnf-plugins-core && \
  dnf config-manager --set-enabled powertools && \
  dnf install -y \
    gcc gtk3 gtk3-devel \
    libusb-devel \
    nodejs npm pkg-config \
    webkit2gtk3-devel wget && \
  mkdir project && \
  wget https://golang.org/dl/go1.16.6.linux-amd64.tar.gz -O go.tar.gz

RUN tar -zxf go.tar.gz && \
  cp -r ./go /usr/local/bin

ENV PATH=$PATH:/usr/local/bin/go/bin
ENV GOPATH=/usr/local/bin/go

RUN npm i -g yarn
RUN go get -u github.com/wailsapp/wails/cmd/wails

WORKDIR project
COPY /*.go ./
COPY /go.mod ./go.mod
COPY /go.sum ./go.sum
COPY /frontend ./frontend
COPY /project.json ./project.json
COPY /wally ./wally

RUN wails build
ENTRYPOINT ["sleep", "infinity"]
