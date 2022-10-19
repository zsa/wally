#include <iostream>
#include <libusb.h>

#include "enumerator.hpp"

Enumerator::Enumerator()
{
    libusb_init(NULL);
}

// This is blocking but it is ran within a go routine in the main thread, don't need to implement threading here.
void Enumerator::Listen()
{
    listening = true;

    auto on_device_event = [](struct libusb_context *ctx, struct libusb_device *dev, libusb_hotplug_event event, void *user_data) -> int
    {
        Enumerator *that = static_cast<Enumerator *>(user_data);
        struct libusb_device_descriptor desc;
        int res = libusb_get_device_descriptor(dev, &desc);

        // That device doesn't interest us, ignoring the event
        if (res < 0 || !Device::is_interesting(desc.idVendor, desc.idProduct))
        {
            return LIBUSB_SUCCESS;
        }

        if (event == LIBUSB_HOTPLUG_EVENT_DEVICE_ARRIVED)
        {
            auto device = Device(dev, desc.idVendor, desc.idProduct);
            that->Devices.push_back(device);
            that->EventObject->handleUSBConnectionEvent(true, device);
        }

        if (event == LIBUSB_HOTPLUG_EVENT_DEVICE_LEFT)
        {
            auto fingerprint = reinterpret_cast<std::intptr_t>(dev);
            for (int i = 0; i < that->Devices.size(); i++)
            {
                if (fingerprint == that->Devices[i].fingerprint)
                {
                    auto device = that->Devices[i];
                    //                    device.close();
                    that->EventObject->handleUSBConnectionEvent(false, device);
                    that->Devices.erase(that->Devices.begin() + i);
                }
            }
        }

        return LIBUSB_SUCCESS;
    };

    libusb_hotplug_register_callback(NULL, static_cast<libusb_hotplug_event>(LIBUSB_HOTPLUG_EVENT_DEVICE_ARRIVED | LIBUSB_HOTPLUG_EVENT_DEVICE_LEFT), LIBUSB_HOTPLUG_ENUMERATE, LIBUSB_HOTPLUG_MATCH_ANY, LIBUSB_HOTPLUG_MATCH_ANY, LIBUSB_HOTPLUG_MATCH_ANY, on_device_event, this, &callback_handle);

    while (listening == true)
    {
        libusb_handle_events(NULL);
    }
}

Enumerator::~Enumerator()
{
    listening = false;
    libusb_hotplug_deregister_callback(NULL, callback_handle);
    libusb_exit(NULL);
}
