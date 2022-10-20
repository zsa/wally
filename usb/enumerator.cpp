#include <iostream>
#include <libusb.h>
#include <chrono>
#include <thread>

#include "enumerator.hpp"

Enumerator::Enumerator()
{
    libusb_init(NULL);
}

// This is blocking but it is ran within a go routine in the main thread, don't need to implement threading here.
void Enumerator::ListenDevices()
{
    listening_devices = true;

    while (listening_devices == true)
    {
        libusb_device **list;
        ssize_t count = libusb_get_device_list(NULL, &list);
        for (int i = 0; i < count; i++)
        {
            libusb_device *dev = list[i];
            struct libusb_device_descriptor desc;
            int res = libusb_get_device_descriptor(dev, &desc);

            // That device doesn't interest us, ignoring the event
            if (res < 0 || !Device::is_interesting(desc.idVendor, desc.idProduct))
                continue;

            auto fingerprint = reinterpret_cast<std::intptr_t>(dev);
            bool registered = false;
            for (int i = 0; i < this->Devices.size(); i++)
            {
                // Check if we already registered that device
                if (fingerprint == this->Devices[i].fingerprint)
                {
                    registered = true;
                    break;
                }
            }

            // Register the device if it wasn't previously registered
            if (registered == false)
            {
                auto device = Device(dev, desc.idVendor, desc.idProduct);
                this->Devices.push_back(device);
                this->EventObject->handleUSBConnectionEvent(true, device);
            }
        }
        libusb_free_device_list(list, 1);

        // Loop on the registered device list to check for disconnections
        for (int i = 0; i < this->Devices.size(); i++)
        {
            auto device = this->Devices[i];
            int connected = device.check_connected();
            if (connected != LIBUSB_SUCCESS)
            {
                this->EventObject->handleUSBConnectionEvent(false, device);
                this->Devices.erase(this->Devices.begin() + i);
            }
        }
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
