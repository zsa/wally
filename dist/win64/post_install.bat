wdi-simple.exe --name "Planck DFU" --vid 0x0483 --pid 0xdf11 --type 0 --manufacturer ZSA.io
%windir%\sysnative\pnputil.exe /add-driver ./usb_driver/usb_device.inf /install