def recsum(ints):
    head = ints[0]
    tail = ints[1:]
    for i in tail:
        if head + i == 2020:
            return head * i
    return recsum(tail)


def dirtysum(ints):
    for i1 in ints:
        for i2 in ints:
            for i3 in ints:
                if i1 + i2 + i3 == 2020:
                    return i1 * i2 * i3


with open("dec1/input.txt") as infile:
    ints = [int(n) for n in infile.read().split("\n")]
print(f"Part 1: {recsum(ints)}")
print(f"Part 2: {dirtysum(ints)}")