wmic sysdriver where(name="STTub30") delete
qmk_driver_installer.exe --force --all drivers_list

wdi-simple.exe --type 0 --manufacturer "ZSA" --name "Planck EZ DFU" --vid 0x0483 --pid 0x0DF11 --inf "planck_dfu.inf" --log 0