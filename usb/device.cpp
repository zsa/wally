#include <libusb.h>
#include <iostream>
#include "device.hpp"

bool Device::is_interesting(int vid, int pid)
{
    // Add new vendor ids here
    switch (vid)
    {
    case 0xFEED: // QMK default
    case 0x3297: // ZSA
    case 0x0483: // STM32 Bootloader
    case 0x16C0:
        break; // Half kay
    default:
        return false;
    }
    return (get_friendly_name(pid) != "Unknown");
}

std::string Device::get_friendly_name(int pid)
{
    switch (pid)
    {
    // Add new keyboard models in here
    case 0x1307:
        return "Ergodox EZ";
    case 0x4974:
        return "Ergodox EZ Original";
    case 0x4975:
        return "Ergodox EZ Shine";
    case 0x4976:
        return "Ergodox EZ Glow";
    case 0x6060:
        return "Planck EZ";
    case 0xC6CE:
        return "Planck EZ Standard";
    case 0xC6CF:
        return "Planck EZ Glow";
    case 0x1969:
        return "Moonlander MK1";
    case 0x0478:
        return "Ergodox in Reset Mode";
    case 0xDF11:
        return "Keyboard in Reset Mode";
    default:
        return "unknown";
    };
}

std::string Device::get_model(int pid)
{
    switch (pid)
    {
    // Add new keyboard models in here
    case 0x1307:
    case 0x4974:
    case 0x4975:
    case 0x4976:
    case 0x0478:
        return "ergodox";
    case 0x6060:
    case 0xC6CE:
    case 0xC6CF:
        return "planck";
    case 0x1969:
        return "moonlander";
    case 0xDF11:
        return "stm32";
    default:
        return "unknown";
    };
}

Device::flash_protocol Device::get_flashing_protocol(int pid)
{
    switch (pid)
    {
    case 0x1307:
    case 0x4974:
    case 0x4975:
    case 0x4976:
    case 0x0478:
        return HALFKAY;
    case 0x6060:
    case 0xC6CE:
    case 0xC6CF:
    case 0x1969:
    case 0xDF11:
        return DFU;
    default:
        return PROTOCOL_UNKNOWN;
    }
}

Device::firmware_format Device::get_firmware_format(Device::flash_protocol protocol)
{
    switch (protocol)
    {
    case HALFKAY:
        return HEX;
    case DFU:
        return BIN;
    default:
        return FORMAT_UNKNOWN;
    }
}

bool Device::is_bootloader(int pid)
{
    switch (pid)
    {
    case 0x0478:
    case 0xDF11:
        return true;
    default:
        return false;
    }
}

Device::Device(libusb_device *dev, int vid, int pid) : usb_device(dev), vid(vid), pid(pid)
{
    friendly_name = Device::get_friendly_name(pid);
    bootloader = Device::is_bootloader(pid);
    protocol = Device::get_flashing_protocol(pid);
    model = get_model(pid);
    file_format = Device::get_firmware_format(protocol);
    port_number = libusb_get_port_number(dev);
    address = libusb_get_device_address(dev);

    // Using the libusb device pointer address as a unique indentifier to the device.
    // This is useful to find the device within the context of a disconnection event to remove it from the enumerator's list of devices.
    fingerprint = reinterpret_cast<std::intptr_t>(dev);
}

Device::~Device()
{
    std::cout << "the great destroyer" << std::endl;
}

bool Device::usb_claim()
{
    if (libusb_open(usb_device, &usb_handle))
        return false;

    if (usb_auto_detach())
        return false;

    claimed = true;
    return true;
}

TransferStatus Device::usb_transfer(uint8_t bmRequestType, uint8_t bRequest, uint16_t wValue, uint16_t wIndex, unsigned char *data, uint16_t wLength, int timeout)
{
    TransferStatus transfer_status;

    int res = libusb_control_transfer(usb_handle, bmRequestType, bRequest, wValue, wIndex, data, wLength, timeout);

    if (res == wLength)
    {
        transfer_status.status_code = 0;
    }
    else
    {
        std::string status(libusb_error_name(res));
        transfer_status.status_error = status;
        transfer_status.status_code = res;
    }
    return transfer_status;
    /*
    int ret = 0;
    unsigned char *buf;
    struct libusb_transfer *transfer;
    transfer = libusb_alloc_transfer(0);
    struct TransferStatus transfer_status;

    transfer_status.transferring = true;

    auto transfer_callback = [](struct libusb_transfer *control_transfer)
    {
        TransferStatus *status = static_cast<TransferStatus *>(control_transfer->user_data);
        status->transferring = false;
        status->status_code = control_transfer->status;
        if (status->status_code != LIBUSB_SUCCESS)
        {
            std::string status_error(libusb_error_name(control_transfer->status));
            status->status_error = status_error;
        }
        unsigned char *data = libusb_control_transfer_get_data(control_transfer);
        for (int i = 0; i < USB_BUFFER_SIZE; i++)
        {
            if (i == control_transfer->length)
            {
                break;
            }
            status->buf[i] = data[i];
        }
    };

    buf = (unsigned char *)malloc(LIBUSB_CONTROL_SETUP_SIZE + wLength);
    libusb_fill_control_setup(buf, bmRequestType, bRequest, wValue, wIndex, wLength);
    memcpy(buf + LIBUSB_CONTROL_SETUP_SIZE, data, wLength);
    libusb_fill_control_transfer(transfer, usb_handle, buf, transfer_callback, &transfer_status, timeout);

    int res = libusb_submit_transfer(transfer);

    // Submission failed, returning the error
    if (res != LIBUSB_SUCCESS)
    {
        std::string status(libusb_error_name(res));
        transfer_status.status_error = status;
        transfer_status.status_code = res;

        return transfer_status;
    }

    while (transfer_status.transferring)
    {
    }
    // transfer_status.status_code = 1;
    // return transfer_status;

    libusb_free_transfer(transfer);

    free(buf);
    return transfer_status;
    */
}

int Device::check_connected()
{
    try
    {
        int res = libusb_open(usb_device, &usb_handle);
        if (res == LIBUSB_SUCCESS)
        {
            // libusb_close(usb_handle);
        }
        return res;
    }
    catch (const std::exception &)
    {
        return 1;
    }
}

int Device::usb_auto_detach()
{
    int res = libusb_set_auto_detach_kernel_driver(usb_handle, true);
    // Current os auto detach not supported, ignoring.
    if (res == LIBUSB_ERROR_NOT_SUPPORTED)
    {
        return 0;
    }
    return res;
}

int Device::usb_set_configuration(int config)
{
    return libusb_set_configuration(usb_handle, config);
}

int Device::usb_claim_interface(int interface)
{
    return libusb_claim_interface(usb_handle, interface);
}

void Device::usb_close()
{
    // libusb_close(usb_handle);
    claimed = false;
}

void Device::close_hid()
{
    hid_close(hid_handle);
    hid_opened = false;
    hid_exit();
}

bool Device::hid_open(int usage_page)
{
    int res = hid_init();
    if (res < 0)
    {
        return false;
    }

    struct hid_device_info *devs, *cur_dev;
    devs = hid_enumerate(vid, pid);
    cur_dev = devs;

    while (cur_dev)
    {
        if (cur_dev->usage_page == usage_page)
        {
            break;
        }
        else
        {
            cur_dev = cur_dev->next;
        }
    }

    if (!cur_dev)
        return false;

    hid_handle = hid_open_path(cur_dev->path);

    hid_free_enumeration(devs);

    if (hid_handle != nullptr)
    {
        hid_opened = true;
        return true;
    }

    return false;
}

bool Device::hid_listen()
{
    if (hid_opened == false)
    {
        return false;
    }

    unsigned char buf[HID_PACKET_SIZE];
    int status = 0;

    while (hid_opened == true && status != -1)
    {
        status = hid_read(hid_handle, buf, HID_PACKET_SIZE);
        if (status != -1)
        {
            signed char *packet = reinterpret_cast<signed char *>(buf);
            packet_handler->handleIncomingPacket(packet);
        }
        else
        {
            close_hid();
        }
    }

    return true;
}

std::string Device::get_dfu_string(int cfg_idx)
{

    int rc;
    struct libusb_config_descriptor *cfg;
    const struct libusb_interface_descriptor *iface_desc;
    const struct libusb_interface *iface;

    rc = libusb_get_config_descriptor_by_value(usb_device, cfg_idx, &cfg);
    if (rc)
    {
        return "";
    }

    for (int iface_idx = 0; iface_idx < cfg->bNumInterfaces; iface_idx++)
    {
        iface = &cfg->interface[iface_idx];
        if (!iface)
            break;
        for (int alt_idx = 0; alt_idx < iface->num_altsetting; alt_idx++)
        {
            iface_desc = &iface->altsetting[alt_idx];
            if (!iface_desc)
                break;
            if (iface_desc->bInterfaceClass == 0xfe &&
                iface_desc->bInterfaceSubClass == 1)
            {
                if (iface_desc->iInterface)
                {
                    unsigned char *data;
                    libusb_get_string_descriptor_ascii(usb_handle, iface_desc->iInterface, data, 128);
                    std::string dfu_string = std::string(data, data + 128);
                    libusb_free_config_descriptor(cfg);
                    return dfu_string;
                }
            }
        }
    }

    libusb_free_config_descriptor(cfg);
    return "";
}

int Device::send_hid_packet(unsigned char *packet, int len)
{
    return hid_write(hid_handle, packet, len);
}
