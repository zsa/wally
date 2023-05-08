#include <iostream>
#include <libusb.h>
#include <chrono>
#include <thread>

#include "enumerator.hpp"

Enumerator::Enumerator()
{
    libusb_init(NULL);
}

// This is blocking but it is ran within a go routine by the main go thread, don't need to implement threading here.

struct device_attributes
{
    int port_number;
    int vid;
    int pid;
};

void Enumerator::ListenDevices()
{
    listening_devices = true;

    while (listening_devices == true)
    {
        libusb_device **list;
        ssize_t count = libusb_get_device_list(NULL, &list);
        device_attributes interesting_devices[count];
        int devices_found = 0;
        for (int i = 0; i < count; i++)
        {
            libusb_device *dev = list[i];
            struct libusb_device_descriptor desc;
            int res = libusb_get_device_descriptor(dev, &desc);

            // That device doesn't interest us, ignoring the event
            if (res < 0 || !Device::is_interesting(desc.idVendor, desc.idProduct))
                continue;
            int port_number = libusb_get_port_number(dev);

            // Save device attributes to check disconnections later
            interesting_devices[devices_found] = device_attributes{.port_number = port_number, .vid = desc.idVendor, .pid = desc.idProduct};
            devices_found++;

            bool registered = false;
            for (int i = 0; i < this->Devices.size(); i++)
            {
                // Check if we already registered that device
                Device *registered_dev = this->Devices[i];

                if (port_number == registered_dev->port_number && desc.idVendor == registered_dev->vid && desc.idProduct == registered_dev->pid)
                {
                    registered = true;
                    break;
                }
            }

            // Register the device if it wasn't previously registered
            if (registered == false)
            {
                libusb_ref_device(dev);
                auto device = new Device(dev, desc.idVendor, desc.idProduct);
                this->Devices.push_back(device);
                this->EventObject->handleUSBConnectionEvent(true, device);
            }
        }

        // Loop on the registered device list to check for disconnections
        for (int i = 0; i < this->Devices.size(); i++)
        {
            Device *registered_dev = this->Devices[i];
            bool connected = false;
            for (int j = 0; j < devices_found; j++)
            {
                device_attributes dev = interesting_devices[j];

                if (dev.port_number == registered_dev->port_number && dev.vid == registered_dev->vid && dev.pid == registered_dev->pid)
                {
                    connected = true;
                    break;
                }
            }

            if (connected == false)
            {
                this->EventObject->handleUSBConnectionEvent(false, registered_dev);
                this->Devices.erase(this->Devices.begin() + i);
            }
        }
        libusb_free_device_list(list, 0);
        std::this_thread::sleep_for(std::chrono::milliseconds(500));
    }
}

void Enumerator::StopListenDevices()
{
    listening_devices = false;
}

void Enumerator::HandleEvents()
{

    handle_events = true;
    while (handle_events == true)
    {
        libusb_handle_events(NULL);
    }
}

Enumerator::~Enumerator()
{
    listening_devices = false;
    handle_events = false;
    libusb_exit(NULL);
}
