# Wally

Flash your [ZSA Keyboard](https://ergodox-ez.com) the EZ way.

ℹ️ Windows users: [There's a new version of Wally](https://github.com/zsa/wally-win/) rewritten from the ground up using native technologies.

ErgoDox EZ users: [Teensy Loader](https://www.pjrc.com/teensy/loader.html) will still work for flashing your ErgoDox EZ (as will Wally — you've got options!).

## Getting started

⚠️ Warning starting from firmware v19, Wally needs to be updated to v2.1.0 in order to detect your keyboard. Linux users should update their udev rules using the latest version from [the wiki](https://github.com/zsa/wally/wiki/Live-training-on-Linux).

Wally comes in two flavors, a GUI and a CLI app.
Download the application for your favorite platform from the relevant release page: [GUI](https://github.com/zsa/wally/releases) / [CLI](https://github.com/zsa/wally-cli/releases).

Note for Linux users, follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install) before running the application.

Note for Mac OS users, the CLI requires libusb to be installed: `brew install libusb`

## Contributing

The following instructions apply only to those who wish to actively _develop_ Wally and contribute new features or bugfixes to this open-source project. If you simply want to flash your board with some fresh firmware, see above.

Wally is built using [Wails](https://wails.app/) at its core and [Preact](https://preactjs.com/) for the UI. This guide assumes you have a sane [Wails environment](https://wails.app/gettingstarted/) setup.

### Installing dev dependencies

Wally is compatible with Windows, Linux, and macOS. Developing using each platform requires some extra setup:

#### Windows

1. Install [Wails](https://wails.app/gettingstarted/windows/)
2. Setup pkg-config - see [http://www.mingw.org/wiki/FAQ](http://www.mingw.org/wiki/FAQ) "How do I get pkg-config installed?"
3. Grab and install the latest version of libusb [from here](http://sourceforge.net/projects/libusb/files/libusb-1.0/)
4. At the root of the project run `wails build`, the resulting binary will be available in the `build` folder.

#### Linux

The easiest way to compile locally is to use Docker:

Run `./build.linux.sh`, the resulting binary will be available in the `dist/linux64` directory.

An alternative method:

1. Follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install).
2. Install [Wails](https://wails.app/gettingstarted/linux/)
3. At the root of the project run `wails build`, the resulting binary will be available in the `build` folder.

#### macOS

1. Install libusb using `brew`:

```
brew install libusb
```
2. Install [Wails](https://wails.app/gettingstarted/mac/)
3. At the root of the project run `wails build`, the resulting binary will be available in the `build` folder.

Note: the GUI app won't include libusb so it needs to be installed on the computer running it. To embed libusb into the binary, install [dylibbundler](https://github.com/auriamg/macdylibbundler/) and run:

`dylibbundler -of -b -x ./dist/osx/Wally.app/Contents/MacOS/Wally -d ./dist/osx/Wally.app/Contents/libs/`
