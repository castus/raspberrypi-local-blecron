#!/usr/bin/python

import subprocess

bluetooth_address = "26:1A:DC:1C:6D:75"

output = subprocess.run('echo "info ' + bluetooth_address + '" | bluetoothctl | grep "Connected: "', shell=True, capture_output=True, text=True)
isConnected = output.stdout.split(" ")[1].find("yes") == 0
if isConnected == False:
    print("#no#")
    subprocess.run('echo "connect ' + bluetooth_address + '" | bluetoothctl', shell=True)
else:
    print("#yes#")
