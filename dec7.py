import re
import sys
from functools import lru_cache

from loguru import logger


RE_CONTAIN = re.compile(r"contain ([\d\w\s\,]+)\.")
RE_AMOUNT_AND_COLOR = re.compile(r"([\d]+) ([\w\s]+) bag")

logger.remove()
logger.add(sys.stderr, level="DEBUG")


@lru_cache(maxsize=None)
class Node:
    def __init__(self, color, use_bag_count):
        self.color = color
        self.children = []
        self.parent = None
        self.use_bag_count = use_bag_count
        self.create_children()

    def __str__(self):
        return f"{self.color.title().replace(' ', '')}"

    def __repr__(self):
        return f"{self.color.title().replace(' ', '')}"

    @lru_cache(maxsize=None)
    def get_input(self):
        with open("dec7/input.txt") as infile:
            return [line for line in infile.read().split("\n")]

    def create_children(self):
        for entry in self.get_input():
            if entry.startswith(self.color):
                matches = re.findall(RE_CONTAIN, entry)
                for match in matches:
                    if match == "no other bags":
                        return
                    parts = match.split(", ")
                    for part in parts:
                        amount, color = re.findall(RE_AMOUNT_AND_COLOR, part)[0]
                        for _ in range(0, int(amount)):
                            new_node = Node(color, self.use_bag_count)
                            self.add_child(new_node)
                            new_node.set_parent(self)
                            if not self.use_bag_count:
                                break

    def set_parent(self, node):
        self.parent = node

    def add_child(self, node):
        self.children.append(node)

    @lru_cache(maxsize=None)
    def get_children(
        self,
    ):
        children = self.children
        for child in self.children:
            children += child.children
        return children

    def contains_shiny_gold_bag(self):
        for child in self.get_children():
            if child.color == "shiny gold":
                return child.parent
        return False


def get_bag_colors(entries):
    colors = []
    for entry in entries:
        matches = re.findall(r"(\w+\s\w+) bags", entry)
        colors += matches
    colors = list(set(colors))
    return colors


def get_bags(use_bag_count):
    bags = []
    colors = get_bag_colors(Node("dummy color", use_bag_count).get_input())
    counter = 0
    for color in colors:
        if color != "no other":
            counter += 1
            logger.info(
                f"Creating bag with color '{color}' ({counter}/{len(colors)})",
                flush=True,
            )
            node = Node(color, use_bag_count)
            bags.append(node)
        # break  # debug
    return bags


def task_1():
    bags = get_bags(use_bag_count=False)
    bags_containing_shiny_gold_bag = 0
    counter = 0
    for bag in bags:
        logger.info(
            f"Checking '{bag}' for 'shiny gold' ({counter}/{len(bags)})", flush=True
        )
        if bag.contains_shiny_gold_bag():
            bags_containing_shiny_gold_bag += 1
        counter += 1
    logger.info(
        f"Found {bags_containing_shiny_gold_bag} bags to eventually "
        "contain at least one shiny gold bag."
    )
    assert bags_containing_shiny_gold_bag == 208


def task_2():
    node = Node("shiny gold", use_bag_count=True)
    children = node.get_children()
    logger.info("Found {len(children)} folders under the 'shiny gold' bag.")
    assert len(children) == 1664


def main():
    task_1()
    task_2()


if __name__ == "__main__":
    main()
