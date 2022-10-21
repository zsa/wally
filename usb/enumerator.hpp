#pragma once
#include <libusb.h>
#include <vector>
#include "device.hpp"

class EventHandler
{
public:
  virtual ~EventHandler(){};
  virtual void handleUSBConnectionEvent(bool connected, Device *dev) = 0;
};

class Enumerator
{
public:
  EventHandler *EventObject;

  Enumerator();
  ~Enumerator();
  void ListenDevices();
  void StopListenDevices();

  void HandleEvents();

  std::vector<Device *> Devices;

private:
  libusb_context *ctx;
  pthread_t event_thread;
  libusb_hotplug_callback_handle callback_handle;
  bool listening_devices;
  bool handle_events;
};
