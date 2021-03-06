import usb_cdc
import board
from digitalio import DigitalInOut, Direction
from time import sleep

sleep(5)

data_port = usb_cdc.data

# We're using inverse digital logic
PRESSED = False
RELEASED = True


def _init_pin(pin):
    res = DigitalInOut(pin)
    res.direction = Direction.OUTPUT
    res.value = RELEASED
    return res


# The ordering needs to line up with the enum ordering in the Go client
output_pins = [
    board.GP14,  # A
    board.GP26,  # B,
    board.GP21,  # X
    board.GP22,  # Y
    board.GP19,  # Z
    board.GP5,  # L
    board.GP27,  # R
    board.GP0,  # START
    board.GP17,  # UP
    board.GP3,  # DOWN
    board.GP4,  # LEFT
    board.GP2,  # RIGHT
    board.GP12,  # CUP
    board.GP15,  # CDOWN
    board.GP13,  # CLEFT
    board.GP16,  # CRIGHT
]

outputs = [_init_pin(pin) for pin in output_pins]


def _get_and_validate_pin(packet):
    try:
        pin_no = int(packet[1:])

        if pin_no >= len(outputs):
            return

        return outputs[pin_no]
    except Exception:
        return


while True:
    packet = data_port.readline().decode().strip()
    # print(packet)

    if packet == "RACK":
        data_port.write(b"ACKGCN")
        continue

    pin = _get_and_validate_pin(packet)

    if pin is None:
        continue

    if packet.startswith("p"):
        pin.value = PRESSED

    if packet.startswith("r"):
        pin.value = RELEASED
