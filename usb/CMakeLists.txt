cmake_minimum_required(VERSION 3.23)

project("usblib")

find_package(SWIG 4.0.2)
find_package(PkgConfig)
pkg_check_modules(libusb-1.0 REQUIRED libusb-1.0)

get_cmake_property(_variableNames VARIABLES)
list (SORT _variableNames)
foreach (_variableName ${_variableNames})
    message(STATUS "${_variableName}=${${_variableName}}")
endforeach()

if(SWIG_FOUND)
  message("Swig found, creating Go <-> C++ interface.")

  include (UseSWIG)
  set(SWIG_MODULE_NAME "usb")
  set(SWIG_FILE, "usb.i")
  set(SWIG_USE_TARGET_INCLUDE_DIRECTORIES true)
  set(CMAKE_SWIG_FLAGS -c++ -go -cgo -intgosize 64 ${libusb-1.0_CFLAGS})

  swig_add_library(
    ${SWIG_MODULE_NAME}
    TYPE STATIC
    LANGUAGE go
    SOURCES usb.i
    )

  swig_link_libraries(${SWIG_MODULE_NAME} ${libusb-1.0_STATIC_LIBRARIES})

  #target_include_directories(${SWIG_MODULE_NAME} ${libusb-1.0_INCLUDE_DIRS})

else()
  message("Swig not found, make sure swig v 4.0.2+ is installed.")
endif()
