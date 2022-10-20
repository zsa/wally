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
                Device registered_dev = this->Devices[i];

                int port_number = libusb_get_port_number(dev);
                if (port_number = registered_dev.port_number && desc.idVendor == registered_dev.vid && desc.idProduct == registered_dev.pid)
                {
                    registered = true;
                    break;
                }
            }

            // Register the device if it wasn't previously registered
            if (registered == false)
            {
                std::cout << "register" << std::endl;
                libusb_ref_device(dev);
                auto device = Device(dev, desc.idVendor, desc.idProduct);
                this->Devices.push_back(device);
                this->EventObject->handleUSBConnectionEvent(true, device);
            }
        }

        // Loop on the registered device list to check for disconnections
        for (int i = 0; i < this->Devices.size(); i++)
        {
            bool connected = false;
            Device registered_dev = this->Devices[i];
            for (int j = 0; j < count; j++)
            {
                libusb_device *dev = list[i];
                struct libusb_device_descriptor desc;
                int res = libusb_get_device_descriptor(dev, &desc);

                // That device doesn't interest us, ignoring
                if (res < 0 || !Device::is_interesting(desc.idVendor, desc.idProduct))
                    continue;

                int port_number = libusb_get_port_number(dev);
                if (port_number = registered_dev.port_number && desc.idVendor == registered_dev.vid && desc.idProduct == registered_dev.pid)
                {
                    connected = true;
                    break;
                }
            }

            if(connected == false) {
                std::cout << "unregister" << std::endl;
                this->EventObject->handleUSBConnectionEvent(false, registered_dev);
                this->Devices.erase(this->Devices.begin() + i);
            }
        }
        libusb_free_device_list(list, 1);
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
