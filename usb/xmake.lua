set_allowedarchs("x64", "arm64")
set_languages("c99", "c++20")

add_requires("libusb", "hidapi")

target("usbtest")
  set_kind("binary")
  add_files("./*.cpp")
  add_packages("libusb", "hidapi")
