// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {state} from '../models';
import {usb} from '../models';

export function DownloadUpdate(arg1:state.Update):Promise<string>;

export function GetAppVersion():Promise<string>;

export function HandleUSBConnectionEvent(arg1:boolean,arg2:usb.Device):Promise<void>;

export function InitUSB():Promise<void>;

export function InstallUpdate():Promise<void>;

export function Log(arg1:string,arg2:string):Promise<void>;

export function Quit():Promise<void>;

export function Reset():Promise<void>;

export function SelectDevice(arg1:number):Promise<void>;

export function SelectFirmware():Promise<void>;

export function SetStep(arg1:state.Step):Promise<void>;

export function SetUpdateCheck(arg1:boolean):Promise<void>;

export function StartFlashing():Promise<void>;

export function Teardown():Promise<void>;
