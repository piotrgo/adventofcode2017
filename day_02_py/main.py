spreadsheet = []
checksum = []
i = 0

# open file, read content as lines
with open('input') as f:
    read_data = f.readlines()

# parse lines into a list
while i < len(read_data):
    spreadsheet.append(read_data[i].split())
    i=i+1
# map str to int on the list
spreadsheet = [list(map(int, x)) for x in spreadsheet]

# calculate row checksums
for row in spreadsheet:
    checksum.append(max(row)-min(row))

#calculate ss checksum
checksum = sum(checksum)
print(checksum)