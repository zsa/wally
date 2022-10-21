/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: usb.i

#define SWIGMODULE usb
#define SWIG_DIRECTORS

#ifdef __cplusplus
/* SwigValueWrapper is described in swig.swg */
template<typename T> class SwigValueWrapper {
  struct SwigMovePointer {
    T *ptr;
    SwigMovePointer(T *p) : ptr(p) { }
    ~SwigMovePointer() { delete ptr; }
    SwigMovePointer& operator=(SwigMovePointer& rhs) { T* oldptr = ptr; ptr = 0; delete oldptr; ptr = rhs.ptr; rhs.ptr = 0; return *this; }
  } pointer;
  SwigValueWrapper& operator=(const SwigValueWrapper<T>& rhs);
  SwigValueWrapper(const SwigValueWrapper<T>& rhs);
public:
  SwigValueWrapper() : pointer(0) { }
  SwigValueWrapper& operator=(const T& t) { SwigMovePointer tmp(new T(t)); pointer = tmp; return *this; }
  operator T&() const { return *pointer.ptr; }
  T *operator&() { return pointer.ptr; }
};

template <typename T> T SwigValueInit() {
  return T();
}
#endif

/* -----------------------------------------------------------------------------
 *  This section contains generic SWIG labels for method/variable
 *  declarations/attributes, and other compiler dependent labels.
 * ----------------------------------------------------------------------------- */

/* template workaround for compilers that cannot correctly implement the C++ standard */
#ifndef SWIGTEMPLATEDISAMBIGUATOR
# if defined(__SUNPRO_CC) && (__SUNPRO_CC <= 0x560)
#  define SWIGTEMPLATEDISAMBIGUATOR template
# elif defined(__HP_aCC)
/* Needed even with `aCC -AA' when `aCC -V' reports HP ANSI C++ B3910B A.03.55 */
/* If we find a maximum version that requires this, the test would be __HP_aCC <= 35500 for A.03.55 */
#  define SWIGTEMPLATEDISAMBIGUATOR template
# else
#  define SWIGTEMPLATEDISAMBIGUATOR
# endif
#endif

/* inline attribute */
#ifndef SWIGINLINE
# if defined(__cplusplus) || (defined(__GNUC__) && !defined(__STRICT_ANSI__))
#   define SWIGINLINE inline
# else
#   define SWIGINLINE
# endif
#endif

/* attribute recognised by some compilers to avoid 'unused' warnings */
#ifndef SWIGUNUSED
# if defined(__GNUC__)
#   if !(defined(__cplusplus)) || (__GNUC__ > 3 || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4))
#     define SWIGUNUSED __attribute__ ((__unused__))
#   else
#     define SWIGUNUSED
#   endif
# elif defined(__ICC)
#   define SWIGUNUSED __attribute__ ((__unused__))
# else
#   define SWIGUNUSED
# endif
#endif

#ifndef SWIG_MSC_UNSUPPRESS_4505
# if defined(_MSC_VER)
#   pragma warning(disable : 4505) /* unreferenced local function has been removed */
# endif
#endif

#ifndef SWIGUNUSEDPARM
# ifdef __cplusplus
#   define SWIGUNUSEDPARM(p)
# else
#   define SWIGUNUSEDPARM(p) p SWIGUNUSED
# endif
#endif

/* internal SWIG method */
#ifndef SWIGINTERN
# define SWIGINTERN static SWIGUNUSED
#endif

/* internal inline SWIG method */
#ifndef SWIGINTERNINLINE
# define SWIGINTERNINLINE SWIGINTERN SWIGINLINE
#endif

/* exporting methods */
#if defined(__GNUC__)
#  if (__GNUC__ >= 4) || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4)
#    ifndef GCC_HASCLASSVISIBILITY
#      define GCC_HASCLASSVISIBILITY
#    endif
#  endif
#endif

#ifndef SWIGEXPORT
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   if defined(STATIC_LINKED)
#     define SWIGEXPORT
#   else
#     define SWIGEXPORT __declspec(dllexport)
#   endif
# else
#   if defined(__GNUC__) && defined(GCC_HASCLASSVISIBILITY)
#     define SWIGEXPORT __attribute__ ((visibility("default")))
#   else
#     define SWIGEXPORT
#   endif
# endif
#endif

/* calling conventions for Windows */
#ifndef SWIGSTDCALL
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   define SWIGSTDCALL __stdcall
# else
#   define SWIGSTDCALL
# endif
#endif

/* Deal with Microsoft's attempt at deprecating C standard runtime functions */
#if !defined(SWIG_NO_CRT_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_CRT_SECURE_NO_DEPRECATE)
# define _CRT_SECURE_NO_DEPRECATE
#endif

/* Deal with Microsoft's attempt at deprecating methods in the standard C++ library */
#if !defined(SWIG_NO_SCL_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_SCL_SECURE_NO_DEPRECATE)
# define _SCL_SECURE_NO_DEPRECATE
#endif

/* Deal with Apple's deprecated 'AssertMacros.h' from Carbon-framework */
#if defined(__APPLE__) && !defined(__ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES)
# define __ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES 0
#endif

/* Intel's compiler complains if a variable which was never initialised is
 * cast to void, which is a common idiom which we use to indicate that we
 * are aware a variable isn't used.  So we just silence that warning.
 * See: https://github.com/swig/swig/issues/192 for more discussion.
 */
#ifdef __INTEL_COMPILER
# pragma warning disable 592
#endif


#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>



typedef long long intgo;
typedef unsigned long long uintgo;


# if !defined(__clang__) && (defined(__i386__) || defined(__x86_64__))
#   define SWIGSTRUCTPACKED __attribute__((__packed__, __gcc_struct__))
# else
#   define SWIGSTRUCTPACKED __attribute__((__packed__))
# endif



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;




#define swiggo_size_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];
#define swiggo_size_assert(t, n) swiggo_size_assert_eq(sizeof(t), n, swiggo_sizeof_##t##_is_not_##n)

swiggo_size_assert(char, 1)
swiggo_size_assert(short, 2)
swiggo_size_assert(int, 4)
typedef long long swiggo_long_long;
swiggo_size_assert(swiggo_long_long, 8)
swiggo_size_assert(float, 4)
swiggo_size_assert(double, 8)

#ifdef __cplusplus
extern "C" {
#endif
extern void crosscall2(void (*fn)(void *, int), void *, int);
extern char* _cgo_topofstack(void) __attribute__ ((weak));
extern void _cgo_allocate(void *, int);
extern void _cgo_panic(void *, int);
#ifdef __cplusplus
}
#endif

static char *_swig_topofstack() {
  if (_cgo_topofstack) {
    return _cgo_topofstack();
  } else {
    return 0;
  }
}

static void _swig_gopanic(const char *p) {
  struct {
    const char *p;
  } SWIGSTRUCTPACKED a;
  a.p = p;
  crosscall2(_cgo_panic, &a, (int) sizeof a);
}




#define SWIG_contract_assert(expr, msg) \
  if (!(expr)) { _swig_gopanic(msg); } else


static _gostring_ Swig_AllocateString(const char *p, size_t l) {
  _gostring_ ret;
  ret.p = (char*)malloc(l);
  memcpy(ret.p, p, l);
  ret.n = l;
  return ret;
}

/* -----------------------------------------------------------------------------
 * director_common.swg
 *
 * This file contains support for director classes which is common between
 * languages.
 * ----------------------------------------------------------------------------- */

/*
  Use -DSWIG_DIRECTOR_STATIC if you prefer to avoid the use of the
  'Swig' namespace. This could be useful for multi-modules projects.
*/
#ifdef SWIG_DIRECTOR_STATIC
/* Force anonymous (static) namespace */
#define Swig
#endif
/* -----------------------------------------------------------------------------
 * director.swg
 *
 * This file contains support for director classes so that Go proxy
 * methods can be called from C++.
 * ----------------------------------------------------------------------------- */

#include <exception>
#include <map>

namespace Swig {

  class DirectorException : public std::exception {
  };
}

/* Handle memory management for directors.  */

namespace {
  struct GCItem {
    virtual ~GCItem() {}
  };

  struct GCItem_var {
    GCItem_var(GCItem *item = 0) : _item(item) {
    }

    GCItem_var& operator=(GCItem *item) {
      GCItem *tmp = _item;
      _item = item;
      delete tmp;
      return *this;
    }

    ~GCItem_var() {
      delete _item;
    }

    GCItem* operator->() {
      return _item;
    }

    private:
      GCItem *_item;
  };

  template <typename Type>
  struct GCItem_T : GCItem {
    GCItem_T(Type *ptr) : _ptr(ptr) {
    }

    virtual ~GCItem_T() {
      delete _ptr;
    }

  private:
    Type *_ptr;
  };
}

class Swig_memory {
public:
  template <typename Type>
  void swig_acquire_pointer(Type* vptr) {
    if (vptr) {
      swig_owner[vptr] = new GCItem_T<Type>(vptr);
    }
  }
private:
  typedef std::map<void *, GCItem_var> swig_ownership_map;
  swig_ownership_map swig_owner;
};

template <typename Type>
static void swig_acquire_pointer(Swig_memory** pmem, Type* ptr) {
  if (!pmem) {
    *pmem = new Swig_memory;
  }
  (*pmem)->swig_acquire_pointer(ptr);
}

static void Swig_free(void* p) {
  free(p);
}

static void* Swig_malloc(int c) {
  return malloc(c);
}


#include <string>


#include <stdint.h>		// Use the C99 official header


#include "device.hpp"
#include "enumerator.hpp"


// C++ director class methods.
#include "usb_wrap.h"

SwigDirector_HIDPacketHandler::SwigDirector_HIDPacketHandler(int swig_p)
    : HIDPacketHandler(),
      go_val(swig_p), swig_mem(0)
{ }

extern "C" void Swiggo_DeleteDirector_HIDPacketHandler_usb_4f19f1d7d83a7073(intgo);
SwigDirector_HIDPacketHandler::~SwigDirector_HIDPacketHandler()
{
  Swiggo_DeleteDirector_HIDPacketHandler_usb_4f19f1d7d83a7073(go_val);
  delete swig_mem;
}

extern "C" void Swig_DirectorHIDPacketHandler_callback_handleIncomingPacket_usb_4f19f1d7d83a7073(int, char *packet);
void SwigDirector_HIDPacketHandler::handleIncomingPacket(signed char *packet) {
  char *swig_packet;
  
  *(signed char **)&swig_packet = (signed char *)packet; 
  Swig_DirectorHIDPacketHandler_callback_handleIncomingPacket_usb_4f19f1d7d83a7073(go_val, swig_packet);
}

SwigDirector_EventHandler::SwigDirector_EventHandler(int swig_p)
    : EventHandler(),
      go_val(swig_p), swig_mem(0)
{ }

extern "C" void Swiggo_DeleteDirector_EventHandler_usb_4f19f1d7d83a7073(intgo);
SwigDirector_EventHandler::~SwigDirector_EventHandler()
{
  Swiggo_DeleteDirector_EventHandler_usb_4f19f1d7d83a7073(go_val);
  delete swig_mem;
}

extern "C" void Swig_DirectorEventHandler_callback_handleUSBConnectionEvent_usb_4f19f1d7d83a7073(int, bool connected, Device *dev);
void SwigDirector_EventHandler::handleUSBConnectionEvent(bool connected, Device *dev) {
  bool swig_connected;
  Device *swig_dev;
  
  swig_connected = (bool)connected; 
  *(Device **)&swig_dev = (Device *)dev; 
  Swig_DirectorEventHandler_callback_handleUSBConnectionEvent_usb_4f19f1d7d83a7073(go_val, swig_connected, swig_dev);
}

#ifdef __cplusplus
extern "C" {
#endif

void _wrap_Swig_free_usb_4f19f1d7d83a7073(void *_swig_go_0) {
  void *arg1 = (void *) 0 ;
  
  arg1 = *(void **)&_swig_go_0; 
  
  Swig_free(arg1);
  
}


void *_wrap_Swig_malloc_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  int arg1 ;
  void *result = 0 ;
  void *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = (void *)Swig_malloc(arg1);
  *(void **)&_swig_go_result = (void *)result; 
  return _swig_go_result;
}


void _wrap_TransferStatus_status_code_set_usb_4f19f1d7d83a7073(TransferStatus *_swig_go_0, intgo _swig_go_1) {
  TransferStatus *arg1 = (TransferStatus *) 0 ;
  int arg2 ;
  
  arg1 = *(TransferStatus **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  if (arg1) (arg1)->status_code = arg2;
  
}


intgo _wrap_TransferStatus_status_code_get_usb_4f19f1d7d83a7073(TransferStatus *_swig_go_0) {
  TransferStatus *arg1 = (TransferStatus *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(TransferStatus **)&_swig_go_0; 
  
  result = (int) ((arg1)->status_code);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_TransferStatus_status_error_set_usb_4f19f1d7d83a7073(TransferStatus *_swig_go_0, _gostring_ _swig_go_1) {
  TransferStatus *arg1 = (TransferStatus *) 0 ;
  std::string *arg2 = 0 ;
  
  arg1 = *(TransferStatus **)&_swig_go_0; 
  
  std::string arg2_str(_swig_go_1.p, _swig_go_1.n);
  arg2 = &arg2_str;
  
  
  if (arg1) (arg1)->status_error = *arg2;
  
}


_gostring_ _wrap_TransferStatus_status_error_get_usb_4f19f1d7d83a7073(TransferStatus *_swig_go_0) {
  TransferStatus *arg1 = (TransferStatus *) 0 ;
  std::string *result = 0 ;
  _gostring_ _swig_go_result;
  
  arg1 = *(TransferStatus **)&_swig_go_0; 
  
  result = (std::string *) & ((arg1)->status_error);
  _swig_go_result = Swig_AllocateString((*result).data(), (*result).length()); 
  return _swig_go_result;
}


TransferStatus *_wrap_new_TransferStatus_usb_4f19f1d7d83a7073() {
  TransferStatus *result = 0 ;
  TransferStatus *_swig_go_result;
  
  
  result = (TransferStatus *)new TransferStatus();
  *(TransferStatus **)&_swig_go_result = (TransferStatus *)result; 
  return _swig_go_result;
}


void _wrap_delete_TransferStatus_usb_4f19f1d7d83a7073(TransferStatus *_swig_go_0) {
  TransferStatus *arg1 = (TransferStatus *) 0 ;
  
  arg1 = *(TransferStatus **)&_swig_go_0; 
  
  delete arg1;
  
}


HIDPacketHandler *_wrap__swig_NewDirectorHIDPacketHandlerHIDPacketHandler_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  int arg1 ;
  HIDPacketHandler *result = 0 ;
  HIDPacketHandler *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = new SwigDirector_HIDPacketHandler(arg1);
  *(HIDPacketHandler **)&_swig_go_result = (HIDPacketHandler *)result; 
  return _swig_go_result;
}


void _wrap_DeleteDirectorHIDPacketHandler_usb_4f19f1d7d83a7073(HIDPacketHandler *_swig_go_0) {
  HIDPacketHandler *arg1 = (HIDPacketHandler *) 0 ;
  
  arg1 = *(HIDPacketHandler **)&_swig_go_0; 
  
  delete arg1;
  
}


void _wrap_delete_HIDPacketHandler_usb_4f19f1d7d83a7073(HIDPacketHandler *_swig_go_0) {
  HIDPacketHandler *arg1 = (HIDPacketHandler *) 0 ;
  
  arg1 = *(HIDPacketHandler **)&_swig_go_0; 
  
  delete arg1;
  
}


void _wrap_HIDPacketHandler_handleIncomingPacket_usb_4f19f1d7d83a7073(HIDPacketHandler *_swig_go_0, char *_swig_go_1) {
  HIDPacketHandler *arg1 = (HIDPacketHandler *) 0 ;
  signed char *arg2 = (signed char *) 0 ;
  
  arg1 = *(HIDPacketHandler **)&_swig_go_0; 
  arg2 = *(signed char **)&_swig_go_1; 
  
  (arg1)->handleIncomingPacket(arg2);
  
}


intgo _wrap_PROTOCOL_UNKNOWN_Device_usb_4f19f1d7d83a7073() {
  Device::flash_protocol result;
  intgo _swig_go_result;
  
  
  result = Device::PROTOCOL_UNKNOWN;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_HALFKAY_Device_usb_4f19f1d7d83a7073() {
  Device::flash_protocol result;
  intgo _swig_go_result;
  
  
  result = Device::HALFKAY;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_DFU_Device_usb_4f19f1d7d83a7073() {
  Device::flash_protocol result;
  intgo _swig_go_result;
  
  
  result = Device::DFU;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FORMAT_UNKNOWN_Device_usb_4f19f1d7d83a7073() {
  Device::firmware_format result;
  intgo _swig_go_result;
  
  
  result = Device::FORMAT_UNKNOWN;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_HEX_Device_usb_4f19f1d7d83a7073() {
  Device::firmware_format result;
  intgo _swig_go_result;
  
  
  result = Device::HEX;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_BIN_Device_usb_4f19f1d7d83a7073() {
  Device::firmware_format result;
  intgo _swig_go_result;
  
  
  result = Device::BIN;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


void _wrap_Device_file_format_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  Device::firmware_format arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (Device::firmware_format)_swig_go_1; 
  
  if (arg1) (arg1)->file_format = arg2;
  
}


intgo _wrap_Device_file_format_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  Device::firmware_format result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (Device::firmware_format) ((arg1)->file_format);
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


void _wrap_Device_protocol_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  Device::flash_protocol arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (Device::flash_protocol)_swig_go_1; 
  
  if (arg1) (arg1)->protocol = arg2;
  
}


intgo _wrap_Device_protocol_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  Device::flash_protocol result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (Device::flash_protocol) ((arg1)->protocol);
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


void _wrap_Device_packet_handler_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, HIDPacketHandler *_swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  HIDPacketHandler *arg2 = (HIDPacketHandler *) 0 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = *(HIDPacketHandler **)&_swig_go_1; 
  
  if (arg1) (arg1)->packet_handler = arg2;
  
}


HIDPacketHandler *_wrap_Device_packet_handler_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  HIDPacketHandler *result = 0 ;
  HIDPacketHandler *_swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (HIDPacketHandler *) ((arg1)->packet_handler);
  *(HIDPacketHandler **)&_swig_go_result = (HIDPacketHandler *)result; 
  return _swig_go_result;
}


void _wrap_Device_bootloader_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, bool _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  bool arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (bool)_swig_go_1; 
  
  if (arg1) (arg1)->bootloader = arg2;
  
}


bool _wrap_Device_bootloader_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (bool) ((arg1)->bootloader);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_Device_pid_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  int arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  if (arg1) (arg1)->pid = arg2;
  
}


intgo _wrap_Device_pid_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (int) ((arg1)->pid);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_Device_port_number_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  int arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  if (arg1) (arg1)->port_number = arg2;
  
}


intgo _wrap_Device_port_number_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (int) ((arg1)->port_number);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_Device_address_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, char _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  uint8_t arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (uint8_t)_swig_go_1; 
  
  if (arg1) (arg1)->address = arg2;
  
}


char _wrap_Device_address_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  uint8_t result;
  char _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (uint8_t) ((arg1)->address);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_Device_vid_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  int arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  if (arg1) (arg1)->vid = arg2;
  
}


intgo _wrap_Device_vid_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (int) ((arg1)->vid);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_Device_fingerprint_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  std::intptr_t arg2 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (std::intptr_t)_swig_go_1; 
  
  if (arg1) (arg1)->fingerprint = arg2;
  
}


intgo _wrap_Device_fingerprint_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  std::intptr_t result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (std::intptr_t) ((arg1)->fingerprint);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_Device_friendly_name_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, _gostring_ _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  std::string *arg2 = 0 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  std::string arg2_str(_swig_go_1.p, _swig_go_1.n);
  arg2 = &arg2_str;
  
  
  if (arg1) (arg1)->friendly_name = *arg2;
  
}


_gostring_ _wrap_Device_friendly_name_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  std::string *result = 0 ;
  _gostring_ _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (std::string *) & ((arg1)->friendly_name);
  _swig_go_result = Swig_AllocateString((*result).data(), (*result).length()); 
  return _swig_go_result;
}


void _wrap_Device_model_set_usb_4f19f1d7d83a7073(Device *_swig_go_0, _gostring_ _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  std::string *arg2 = 0 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  std::string arg2_str(_swig_go_1.p, _swig_go_1.n);
  arg2 = &arg2_str;
  
  
  if (arg1) (arg1)->model = *arg2;
  
}


_gostring_ _wrap_Device_model_get_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  std::string *result = 0 ;
  _gostring_ _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (std::string *) & ((arg1)->model);
  _swig_go_result = Swig_AllocateString((*result).data(), (*result).length()); 
  return _swig_go_result;
}


TransferStatus *_wrap_Device_usb_transfer_usb_4f19f1d7d83a7073(Device *_swig_go_0, char _swig_go_1, char _swig_go_2, short _swig_go_3, short _swig_go_4, char *_swig_go_5, short _swig_go_6, intgo _swig_go_7) {
  Device *arg1 = (Device *) 0 ;
  uint8_t arg2 ;
  uint8_t arg3 ;
  uint16_t arg4 ;
  uint16_t arg5 ;
  unsigned char *arg6 = (unsigned char *) 0 ;
  uint16_t arg7 ;
  int arg8 ;
  TransferStatus result;
  TransferStatus *_swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (uint8_t)_swig_go_1; 
  arg3 = (uint8_t)_swig_go_2; 
  arg4 = (uint16_t)_swig_go_3; 
  arg5 = (uint16_t)_swig_go_4; 
  arg6 = *(unsigned char **)&_swig_go_5; 
  arg7 = (uint16_t)_swig_go_6; 
  arg8 = (int)_swig_go_7; 
  
  result = (arg1)->usb_transfer(arg2,arg3,arg4,arg5,arg6,arg7,arg8);
  *(TransferStatus **)&_swig_go_result = new TransferStatus(result); 
  return _swig_go_result;
}


bool _wrap_Device_hid_listen_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (bool)(arg1)->hid_listen();
  _swig_go_result = result; 
  return _swig_go_result;
}


bool _wrap_Device_hid_open_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  int arg2 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  result = (bool)(arg1)->hid_open(arg2);
  _swig_go_result = result; 
  return _swig_go_result;
}


bool _wrap_Device_usb_claim_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (bool)(arg1)->usb_claim();
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_Device_check_connected_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (int)(arg1)->check_connected();
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_Device_send_hid_packet_usb_4f19f1d7d83a7073(Device *_swig_go_0, char *_swig_go_1, intgo _swig_go_2) {
  Device *arg1 = (Device *) 0 ;
  unsigned char *arg2 = (unsigned char *) 0 ;
  int arg3 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = *(unsigned char **)&_swig_go_1; 
  arg3 = (int)_swig_go_2; 
  
  result = (int)(arg1)->send_hid_packet(arg2,arg3);
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_Device_usb_auto_detach_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  result = (int)(arg1)->usb_auto_detach();
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_Device_usb_claim_interface_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  int arg2 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  result = (int)(arg1)->usb_claim_interface(arg2);
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_Device_usb_set_configuration_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  int arg2 ;
  int result;
  intgo _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  result = (int)(arg1)->usb_set_configuration(arg2);
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_Device_get_firmware_format_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  Device::flash_protocol arg1 ;
  Device::firmware_format result;
  intgo _swig_go_result;
  
  arg1 = (Device::flash_protocol)_swig_go_0; 
  
  result = (Device::firmware_format)Device::get_firmware_format(arg1);
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_Device_get_flashing_protocol_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  int arg1 ;
  Device::flash_protocol result;
  intgo _swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = (Device::flash_protocol)Device::get_flashing_protocol(arg1);
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


bool _wrap_Device_is_bootloader_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  int arg1 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = (bool)Device::is_bootloader(arg1);
  _swig_go_result = result; 
  return _swig_go_result;
}


bool _wrap_Device_is_interesting_usb_4f19f1d7d83a7073(intgo _swig_go_0, intgo _swig_go_1) {
  int arg1 ;
  int arg2 ;
  bool result;
  bool _swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  result = (bool)Device::is_interesting(arg1,arg2);
  _swig_go_result = result; 
  return _swig_go_result;
}


_gostring_ _wrap_Device_get_friendly_name_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  int arg1 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = Device::get_friendly_name(arg1);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


_gostring_ _wrap_Device_get_model_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  int arg1 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = Device::get_model(arg1);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


_gostring_ _wrap_Device_get_dfu_string_usb_4f19f1d7d83a7073(Device *_swig_go_0, intgo _swig_go_1) {
  Device *arg1 = (Device *) 0 ;
  int arg2 ;
  std::string result;
  _gostring_ _swig_go_result;
  
  arg1 = *(Device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  
  result = (arg1)->get_dfu_string(arg2);
  _swig_go_result = Swig_AllocateString((&result)->data(), (&result)->length()); 
  return _swig_go_result;
}


void _wrap_Device_close_hid_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  (arg1)->close_hid();
  
}


void _wrap_Device_usb_close_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  (arg1)->usb_close();
  
}


Device *_wrap_new_Device_usb_4f19f1d7d83a7073(libusb_device *_swig_go_0, intgo _swig_go_1, intgo _swig_go_2) {
  libusb_device *arg1 = (libusb_device *) 0 ;
  int arg2 ;
  int arg3 ;
  Device *result = 0 ;
  Device *_swig_go_result;
  
  arg1 = *(libusb_device **)&_swig_go_0; 
  arg2 = (int)_swig_go_1; 
  arg3 = (int)_swig_go_2; 
  
  result = (Device *)new Device(arg1,arg2,arg3);
  *(Device **)&_swig_go_result = (Device *)result; 
  return _swig_go_result;
}


void _wrap_delete_Device_usb_4f19f1d7d83a7073(Device *_swig_go_0) {
  Device *arg1 = (Device *) 0 ;
  
  arg1 = *(Device **)&_swig_go_0; 
  
  delete arg1;
  
}


EventHandler *_wrap__swig_NewDirectorEventHandlerEventHandler_usb_4f19f1d7d83a7073(intgo _swig_go_0) {
  int arg1 ;
  EventHandler *result = 0 ;
  EventHandler *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = new SwigDirector_EventHandler(arg1);
  *(EventHandler **)&_swig_go_result = (EventHandler *)result; 
  return _swig_go_result;
}


void _wrap_DeleteDirectorEventHandler_usb_4f19f1d7d83a7073(EventHandler *_swig_go_0) {
  EventHandler *arg1 = (EventHandler *) 0 ;
  
  arg1 = *(EventHandler **)&_swig_go_0; 
  
  delete arg1;
  
}


void _wrap_delete_EventHandler_usb_4f19f1d7d83a7073(EventHandler *_swig_go_0) {
  EventHandler *arg1 = (EventHandler *) 0 ;
  
  arg1 = *(EventHandler **)&_swig_go_0; 
  
  delete arg1;
  
}


void _wrap_EventHandler_handleUSBConnectionEvent_usb_4f19f1d7d83a7073(EventHandler *_swig_go_0, bool _swig_go_1, Device *_swig_go_2) {
  EventHandler *arg1 = (EventHandler *) 0 ;
  bool arg2 ;
  Device *arg3 = (Device *) 0 ;
  
  arg1 = *(EventHandler **)&_swig_go_0; 
  arg2 = (bool)_swig_go_1; 
  arg3 = *(Device **)&_swig_go_2; 
  
  (arg1)->handleUSBConnectionEvent(arg2,arg3);
  
}


void _wrap_Enumerator_EventObject_set_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0, EventHandler *_swig_go_1) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  EventHandler *arg2 = (EventHandler *) 0 ;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  arg2 = *(EventHandler **)&_swig_go_1; 
  
  if (arg1) (arg1)->EventObject = arg2;
  
}


EventHandler *_wrap_Enumerator_EventObject_get_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  EventHandler *result = 0 ;
  EventHandler *_swig_go_result;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  
  result = (EventHandler *) ((arg1)->EventObject);
  *(EventHandler **)&_swig_go_result = (EventHandler *)result; 
  return _swig_go_result;
}


Enumerator *_wrap_new_Enumerator_usb_4f19f1d7d83a7073() {
  Enumerator *result = 0 ;
  Enumerator *_swig_go_result;
  
  
  result = (Enumerator *)new Enumerator();
  *(Enumerator **)&_swig_go_result = (Enumerator *)result; 
  return _swig_go_result;
}


void _wrap_delete_Enumerator_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  
  delete arg1;
  
}


void _wrap_Enumerator_ListenDevices_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  
  (arg1)->ListenDevices();
  
}


void _wrap_Enumerator_StopListenDevices_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  
  (arg1)->StopListenDevices();
  
}


void _wrap_Enumerator_HandleEvents_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  
  (arg1)->HandleEvents();
  
}


void _wrap_Enumerator_Devices_set_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0, std::vector< Device * > *_swig_go_1) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  std::vector< Device * > *arg2 = (std::vector< Device * > *) 0 ;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  arg2 = *(std::vector< Device * > **)&_swig_go_1; 
  
  if (arg1) (arg1)->Devices = *arg2;
  
}


std::vector< Device * > *_wrap_Enumerator_Devices_get_usb_4f19f1d7d83a7073(Enumerator *_swig_go_0) {
  Enumerator *arg1 = (Enumerator *) 0 ;
  std::vector< Device * > *result = 0 ;
  std::vector< Device * > *_swig_go_result;
  
  arg1 = *(Enumerator **)&_swig_go_0; 
  
  result = (std::vector< Device * > *)& ((arg1)->Devices);
  *(std::vector< Device * > **)&_swig_go_result = (std::vector< Device * > *)result; 
  return _swig_go_result;
}


#ifdef __cplusplus
}
#endif

