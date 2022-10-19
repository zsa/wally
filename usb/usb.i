%module (directors="1") usb
%feature("director") EventHandler;
%feature("director") HIDPacketHandler;

%include "std_string.i"
%include "stdint.i"

using namespace std;
typedef std::string String;

%insert(cgo_comment) %{
#cgo pkg-config: libusb-1.0
#cgo pkg-config: hidapi
#include <libusb.h>
#include <hidapi.h> %}

%{
#include "device.hpp"
#include "enumerator.hpp"
%}

%include "device.hpp"
%include "enumerator.hpp"
