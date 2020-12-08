import copy
import re
import sys
from functools import lru_cache

from loguru import logger

logger.remove()
logger.add(sys.stderr, level="INFO")


@lru_cache
def get_input():
    with open("dec8/input.txt") as infile:
        return [line for line in infile.read().split("\n")]


@lru_cache
def get_input_as_string():
    with open("dec8/input.txt") as infile:
        return infile.read()


def get_instructions(input_=get_input()):
    instructions = []
    for item in input_:
        instruction, argument = item.split()
        instructions.append(
            {"instruction": instruction, "argument": argument, "executed": False}
        )
    return instructions


def iterator(pos, accumulator, instructions):
    i = instructions[pos]
    instruction = i["instruction"]
    arg = int(i["argument"])
    if i["executed"]:
        return (pos, accumulator)
    elif pos == len(instructions) - 1:
        return (pos, accumulator)
    i["executed"] = True
    if instruction == "nop":
        logger.debug(f"{pos}: nop {arg}")
        return iterator(pos + 1, accumulator, instructions)
    elif instruction == "jmp":
        new_pos = pos + arg
        logger.debug(f"{pos}: jmp {arg}")
        return iterator(new_pos, accumulator, instructions)
    elif instruction == "acc":
        new_accumulator = accumulator + arg
        logger.debug(f"{pos}: acc {arg}")
        return iterator(pos + 1, new_accumulator, instructions)


def part1():
    x = iterator(pos=0, accumulator=0, instructions=get_instructions())
    pos = x[0]
    accumulator = x[1]
    logger.info(f"Returning part 1 pos={pos}, accumulator={accumulator}")


def part2():
    nops = len([line for line in get_input() if "nop" in str(line)])
    jmps = len([line for line in get_input() if "jmp" in str(line)])
    logger.debug(f"Found {nops} nops")
    logger.debug(f"Found {jmps} jmps")
    input_ = get_input()
    for search, replace in [("nop", "jmp"), ("jmp", "nop")]:
        stop = False
        for pos, line in enumerate(input_):
            if not stop:
                if search in line:
                    modified_input = input_.copy()
                    modified_input[pos] = line.replace(search, replace)
                    modified_instructions = get_instructions(modified_input)
                    x = iterator(
                        pos=0, accumulator=0, instructions=modified_instructions
                    )
                    pos = x[0]
                    accumulator = x[1]
                    if pos == len(get_input()) - 1:
                        logger.info(
                            f"Returning part 2 pos={pos}, accumulator={accumulator}"
                        )
                        stop = True


def main():
    part1()
    part2()


if __name__ == "__main__":
    main()
