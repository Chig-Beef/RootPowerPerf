# Copyright 2024 The original author
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

'''
DEPRECATED: TOO SLOW, DON'T USE, USE
`generate.exe` INSTEAD YOU CAN CREATE
`generate.exe` BY COMPILING `main.go`
IF YOU DON'T CARE, JUST GO TOUCH GRASS
WHILE YOU WAIT

ALSO, THIS WILL USE A LOT OF MEMORY,
THE FACT IS THAT ONE MILLION ROWS TIMES
ONE THOUSAND COLUMNS CREATES A MINIMUM
OF A GIGABYTE, BUT THAT IS ABSOLUTE
BEST CASE, WORST CASE IS 5GB, NOT
INCLUDING ALL THE EXTRA MEMORY PYTHON
USES TO STORE STRINGS
'''

import random
import sys

# Use a different character if they want
args = sys.argv
if len(args) > 1:
    if args[1] == "nl":
        line_character = '\n'
    else:
        line_character = args[1][0]
else :
    line_character = '\n'


count_lines = 1_000_000
count_nums = 1_000

print("generating")

data: list[list[str]] = []

for i in range(count_lines):
    data.append([])
    for j in range(count_nums):
        n = random.getrandbits(16) # 16-bit unsigned integer
        data[i].append(str(n))

print("joining lines")

output = []
for i in range(len(data)):
    output.append(','.join(data[i]))

print("saving")

with open("../runner/data.txt", "wb") as file:
    file.write(bytes(line_character.join(output), "ascii"))

print("done")
