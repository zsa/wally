# Wally

Flash your [ZSA Keyboard](https://ergodox-ez.com) the EZ way.

ErgoDox EZ users: [Teensy Loader](https://www.pjrc.com/teensy/loader.html) will still work for flashing your ErgoDox EZ (as will Wally — you've got options!).

## Getting started

⚠️ Warning starting from firmware v22, Wally needs to be updated to v3.x.x in order to detect your keyboard.

Wally comes in two flavors, a GUI and a CLI app.
Download the application for your favorite platform from the relevant release page: [GUI](https://github.com/zsa/wally/releases) / [CLI](https://github.com/zsa/wally-cli/releases).

Note for Linux users, make sure your udev rules match the latest version [from the wiki](https://github.com/zsa/wally/wiki/Linux-install).

Note for Mac OS users, the CLI requires libusb to be installed: `brew install libusb`

## Contributing

The following instructions apply only to those who wish to actively _develop_ Wally and contribute new features or bugfixes to this open-source project. If you simply want to flash your board with some fresh firmware, see above.

Wally is built using [Wails](https://wails.io/) at its core and [Svelte](https://svelte.dev/) for the UI. This guide assumes you have a sane [Wails environment](https://wails.io/docs/gettingstarted/installation) setup.

### Installing dev dependencies

Wally is compatible with Windows, Linux, and macOS. Developing using each platform requires some extra setup:

#### Windows

Wally for Windows is built using [MSYS2](https://www.msys2.org)

1. Install MSYS2 and install Golang, pkg-config, libusb, hidapi, golang, nodejs. After installing nodejs, install pnpm globally: `npm i -g pnpm`.
2. Install [Wails](https://wails.app/gettingstarted/windows/)
3. At the root of the project run `wails build`, the resulting binary will be available in the `build` folder.

#### Linux

The easiest way to compile locally is to use Docker:

Run `./build.linux.sh`, the resulting binary will be available in the `dist/linux64` directory.

An alternative method:

1. Follow the instructions from our [wiki page](https://github.com/zsa/wally/wiki/Linux-install).
2. Install [Wails](https://wails.app/gettingstarted/linux/)
3. At the root of the project run `wails build`, the resulting binary will be available in the `build` folder.

#### macOS

1. Install libusb using `brew`:

In order to build a universal version of Wally, you first need to have universal builds of libusb and hidapi. Make sure you DO MOT install both libraries with homebrew.

### libusb universal

Clone libusbr`s directory

2. Install [Wails](https://wails.app/gettingstarted/mac/)
3. At the root of the project run `wails build`, the resulting binary will be available in the `build` folder.

Note: the GUI app won't include libusb so it needs to be installed on the computer running it. To embed libusb into the binary, install [dylibbundler](https://github.com/auriamg/macdylibbundler/) and run:

`dylibbundler -of -b -x ./dist/osx/Wally.app/Contents/MacOS/Wally -d ./dist/osx/Wally.app/Contents/libs/`

### Sending feedback

As you may have noticed, we do not have GitHub Issues enabled for this project. Instead, please submit all feedback via email to contact@zsa.io — you will find us very responsive. Thank you for your help with Wally!
