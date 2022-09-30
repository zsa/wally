FROM centos:centos8

RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-* && \
    sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-* && \
    yum update -y

RUN dnf install -y dnf-plugins-core && \
    dnf config-manager --set-enabled powertools && \
    dnf install -y \
      gcc gtk3 gtk3-devel \
      libusb-devel \
      nodejs npm pkg-config \
      webkit2gtk3-devel wget && \
    mkdir project && \
    wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz -O go.tar.gz

RUN tar -zxf go.tar.gz && \
    cp -r ./go /usr/local/bin

ENV PATH=$PATH:/usr/local/bin/go/bin
ENV GOPATH=/usr/local/bin/go

WORKDIR project
COPY /*.go ./
COPY /go.mod ./go.mod
COPY /go.sum ./go.sum
COPY /frontend ./frontend
COPY /project.json ./project.json
COPY /wally ./wally

RUN npm i -g yarn
RUN go install github.com/wailsapp/wails/cmd/wails@v1.16.7

RUN wails build
ENTRYPOINT ["sleep", "infinity"]