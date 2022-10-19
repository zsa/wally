#pragma once
#include <libusb.h>
#include <hidapi.h>
#include <string>

#define HID_PACKET_SIZE 33
#define USB_BUFFER_SIZE 2048

struct TransferStatus
{
    bool transferring = false;
    int status_code;
    std::string status_error;
    unsigned char buf[USB_BUFFER_SIZE];
};

class HIDPacketHandler
{
public:
    virtual ~HIDPacketHandler(){};
    virtual void handleIncomingPacket(signed char *packet) = 0;
};

class Device
{
public:
    enum flash_protocol
    {
        PROTOCOL_UNKNOWN = -1,
        HALFKAY,
        DFU
    };

    enum firmware_format
    {
        FORMAT_UNKNOWN = -1,
        HEX,
        BIN
    };

    int vid;
    int pid;
    std::intptr_t fingerprint;
    int port_number;
    std::string friendly_name;
    Device::flash_protocol protocol;
    Device::firmware_format file_format;
    bool bootloader;
    HIDPacketHandler *packet_handler;

    bool hid_open(int usage_page);
    bool hid_listen();
    int send_hid_packet(unsigned char *packet, int len);
    void close_hid();

    bool usb_claim();
    int usb_auto_detach();
    TransferStatus usb_transfer(uint8_t bmRequestType, uint8_t bRequest, uint16_t wValue, uint16_t wIndex, unsigned char *data, uint16_t wLength, int timeout);
    int usb_set_configuration(int config);
    int usb_claim_interface(int interface);
    std::string get_dfu_string(int cfg_idx);
    void usb_close();

    static bool is_interesting(int vid, int pid);
    static std::string get_friendly_name(int pid);
    static Device::flash_protocol get_flashing_protocol(int pid);
    static Device::firmware_format get_firmware_format(Device::flash_protocol protocol);
    static bool is_bootloader(int pid);

    Device(libusb_device *dev, int vid, int pid);
    ~Device();

private:
    libusb_device *usb_device;
    libusb_device_handle *usb_handle;
    hid_device *hid_handle;
    bool hid_opened = false;
    bool claimed = false;
    bool transferring = false;
};
