# GenConRX

Recieves commands from a Sega Genesis controller over TCP and emulates windows keypresses.

## Usage

1. TODO: arduino to DB9 wiring diagram

2. Flash sega.ino to an Arduino Uno

3. use `nc -lk 9000 < /dev/ttyUSB*` on the remote server connected to the arduino.

4. build and run `genconrx` to recieve the keypresses on another machine. Note: You will need to update `ip` in the genconrx.go source and rebuild each time the IP of the machine with the attached controller changes.