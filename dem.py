from bitstream import BitStream
from enum import Enum
import struct

class frame(Enum):
    dem_signon = 1
    dem_packet = 2
    dem_synctick = 3
    dem_consolecmd = 4
    dem_usercmd = 5
    dem_datatables = 6
    dem_stop = 7

index = 0

def read(bitstream, numbytes):
    global index
    index += numbytes
    return bitstream.read(bytes, numbytes)

def nextFrame(bitstream): 
    print(f"index before: {index}")
    command = int.from_bytes(read(bitstream, 1), "little")
    if(command == 0):
        print("skipping bit")
        return nextFrame
    length = int.from_bytes(read(bitstream, 4), "little")
    retVal = read(bitstream, length)
    print(f"index after: {index}")
    print(f"length {length}")
    print(f"command raw {command}")
    print(f"command {frame(command).name}")
    print("="*50)
    return retVal

demo = r"match730_003403354556619293004_0695644917_191.dem"

db = ""
with open(demo, "rb") as f:
    db = f.read()

bs = BitStream(db)
header = read(bs, 8).decode("utf-8")
demo_protocol = int.from_bytes(read(bs, 4), "little")
network_protocol = int.from_bytes(read(bs, 4), "little")
server_name = read(bs, 260).decode("utf-8")
client_name = read(bs, 260).decode("utf-8")
map_name = read(bs, 260).decode("utf-8")
game_directory = read(bs, 260).decode("utf-8")
playback_time = struct.unpack("f", read(bs, 4))
ticks = int.from_bytes(read(bs, 4), "little")
frames = int.from_bytes(read(bs, 4), "little")
sign_on_length = int.from_bytes(read(bs, 4), "little")

print(f" header: {header} \n demo_protocol {demo_protocol} \n network_protocol {network_protocol}\n server_name {server_name}\n client_name {client_name}\n map_name {map_name}\n game_directory {game_directory}\n playback_time {playback_time} \n ticks {ticks}\n frames {frames}\n sign_on_length {sign_on_length}\n index {index}")
print("="*50)

sign_on_bytes = read(bs, sign_on_length)

b1 = nextFrame(bs)
b2 = nextFrame(bs)
b3 = nextFrame(bs)
b4 = nextFrame(bs)








# ========================
import os
from bitstream import BitStream
from enum import Enum
import struct

class frame(Enum):
    dem_signon = 1
    dem_packet = 2
    dem_synctick = 3
    dem_consolecmd = 4
    dem_usercmd = 5
    dem_datatables = 6
    dem_stop = 7

index = 0

def read(bitstream, numbytes):
    global index
    index += numbytes
    return bitstream.read(bytes, numbytes)

os.chdir(r"C:\Users\locker-leger\Desktop\repos\CSGOImpact")
demo = r"match730_003403354556619293004_0695644917_191.dem"

db = ""
with open(demo, "rb") as f:
    db = f.read()

bs = BitStream(db)
header = read(bs, 8).decode("utf-8")
demo_protocol = int.from_bytes(read(bs, 4), "little")
network_protocol = int.from_bytes(read(bs, 4), "little")
server_name = read(bs, 260).decode("utf-8")
client_name = read(bs, 260).decode("utf-8")
map_name = read(bs, 260).decode("utf-8")
game_directory = read(bs, 260).decode("utf-8")
playback_time = struct.unpack("f", read(bs, 4))
ticks = int.from_bytes(read(bs, 4), "little")
frames = int.from_bytes(read(bs, 4), "little")
sign_on_length = int.from_bytes(read(bs, 4), "little")