#!/bin/bash
set -e
if [[ "$USER" == 'root' ]]; then
    echo -e "\n\tPlease run script as non-root user!\n"
    exit 1
else
    cd $HOME
fi

# GLOBALS
goUrl='https://dl.google.com/go/go1.13.7.linux-armv6l.tar.gz'
wallyBin='https://configure.ergodox-ez.com/wally/linux'
wallyCli='github.com/zsa/wally-cli'

# INSTALL
declare -A packageAA=(
    ['apt-get']='
        gtk+3.0
        libwebkit2gtk-4.0
        libusb-dev
        libusb-1.0
    '
    ['yum']='
        gtk3
        webkit2gtk3
        libusb
    '
    ['pacman']='
        gtk3
        webkit2gtk
        libusb
    '
)
for key in ${!packageAA[@]}; do
    which $key && sudo $key install -y ${packageAA[$key]}
done

# UDEV RULES
install -m644 -t /etc/udev/rules.d/ \
    dist/linux64/50-oryx.rules dist/linux64/50-wally.rules

# HARDWARE PLATFORM DEPENDENT WALLY
if [[ "$(uname -i)" =~ 'x86' ]]; then
    curl -OSL $wallyBin
    sudo chmod +x wally
    echo -e "You can launch the GUI via\n\t./wally "
else
    curl -SL $goUrl -O go.tar.gz
    sudo tar -C /usr/local -xzf go.tar.gz
    cat <<\EOF | sed -r 's/^ *//' >>$HOME/.bashrc
        alias wally-cli='$HOME/go/bin/linux_arm64/wally-cli'
        export GOPATH=$HOME/go
        export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
        export CC=aarch64-linux-gnu-gcc
        export GOARCH=arm64
        export GOOS=linux
        export CGO_ENABLED=1
EOF
    . $HOME/.bashrc
    go get -u $wallyCli
    echo -e "Now you can run wally like this:\n\twally-cli <firmware-file>"
fi

echo 'Done.'
