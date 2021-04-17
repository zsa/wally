# Wally cli

Flash your [ZSA Keyboard](https://ergodox-ez.com) the EZ way.

## Getting started
Download the application for your favorite platform from the [release page](https://github.com/zsa/wally-cli/releases).

Note for Linux users, follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install) before running the application.

Note for Mac OS users, the CLI requires libusb to be installed: `brew install libusb`

You can also compile and install Wally using go's package manager, make sure you follow the `Installing dev dependencies` section for your platform below:

```
go get -u github.com/zsa/wally-cli
```

## Installing dev dependencies
Wally is compatible with Windows, Linux, and macOS. Developing using each platform requires some extra setup:

### Windows
1. Install [TDM GCC](http://tdm-gcc.tdragon.net/download)
2. Setup pkg-config - see [http://www.mingw.org/wiki/FAQ](http://www.mingw.org/wiki/FAQ) "How do I get pkg-config installed?"
3. Grab and install the latest version of libusb [from here](http://sourceforge.net/projects/libusb/files/libusb-1.0/)

### Linux
Follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install)

### macOS
Install libusb using `brew`:

```
brew install libusb
```

### build

```
go build
```
