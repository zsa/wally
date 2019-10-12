# Wally

Flash your [ZSA Keyboard](https://ergodox-ez.com) the EZ way.

## Getting started

Wally comes in two flavors, a GUI and a CLI app.
Download the application for your favorite platform from the [release page](https://github.com/zsa/wally/releases).

Note for Linux users, follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install) before running the application.

Note for Mac OS users, the CLI requires libusb to be installed: `brew install libusb`

## Contributing

The following instructions apply only to those who wish to actively _develop_ Wally and contribute new features or bugfixes to this open-source project. If you simply want to flash your board with some fresh firmware, see above.

Found a bug? Open an [issue here](https://github.com/zsa/wally/issues).

Wally is built using [Go](https://golang.org/) at its core and [Preact](https://preactjs.com/) for the UI. The binding between core and UI happens using a [fork](https://github.com/fdidron/webview) of the [webview package](https://github.com/zserge/webview). This guide assumes you have a proper Go and NodeJS development environment running.

### Installing dev dependencies

Wally is compatible with Windows, Linux, and macOS. Developing using each platform requires some extra setup:

#### Windows

1. Install [TDM GCC](http://tdm-gcc.tdragon.net/download)
2. Setup pkg-config - see [http://www.mingw.org/wiki/FAQ](http://www.mingw.org/wiki/FAQ) "How do I get pkg-config installed?"
3. Grab and install the latest version of libusb [from here](http://sourceforge.net/projects/libusb/files/libusb-1.0/)

#### Linux

Follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install)

#### macOS

Install libusb using `brew`:

```
brew install libusb
```

### Serve the UI locally

From the `ui` folder run `npm run serve` or `yarn dev` to run a webpack dev server locally on port `8080`.

### Compile a dev build

Run the following:

```
go build -tags dev -o wally
```

This build will bind its UI with the local webpack server.

### Compile the CLI

Run the following:

```
go build cli/main.go -o wally-cli
```

### Compile a release build

#### Windows

Run `build.win.bat`

#### Linux

Run `build.linux.sh`

#### Mac OS

Run `build.osx.sh`
