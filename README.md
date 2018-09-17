GenConRX: Recieves commands from a Sega Genesis controler over TCP and emulates windows keypresses")

use `nc -lk 9000 < /dev/ttyUSB*` on the remote server connected to the arduino.
run `genconrx` to recieve the keypresses.