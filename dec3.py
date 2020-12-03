import math
from time import sleep
from typing import List


class Geo:
    def __init__(self, map_):
        self.map = map_
        self.cols = len(map_[0])
        self.rows = len(map_)
        self.current_row = 0
        self.current_col = 0
        self.accumulated_trees = 0

    def get_pos(self, row, col):
        """Get position on map."""
        while col > self.cols:
            remainder = col % self.cols
            if remainder == 0:
                col = self.cols
            else:
                col = remainder
        item = self.map[row - 1][col - 1]

        # self.visualize(row - 1, col - 1)

        if item == "#":
            self.accumulated_trees += 1
        return item

    def traverse_map_and_count_trees(self, slopes):

        tree_list = []

        for row_iter, col_iter in slopes:
            row = 1
            col = 1
            for line in self.map:
                try:
                    self.get_pos(row, col)
                except IndexError as err:
                    pass
                row += row_iter
                col += col_iter
            tree_list.append(self.accumulated_trees)
            self.accumulated_trees = 0  # reset counter

        trees_multiplied = self.multiply_list(tree_list.copy())
        return (tree_list, trees_multiplied)

    def multiply_list(self, lst: List[int], accumulator=1):
        head = lst[0]
        lst.pop(0)
        accumulator = accumulator * head
        if lst == []:
            return accumulator
        return self.multiply_list(lst, accumulator=accumulator)

    def visualize(self, row, col):
        new_line = list([char for char in self.map[row]])
        item = new_line[col]
        if item == "#":
            new_line[col] = "X"
        elif item == ".":
            new_line[col] = "O"
        print("".join(new_line))


with open("dec3/input.txt") as infile:
    map_ = [line for line in infile.read().split("\n")]


geo = Geo(map_)
slopes = [
    (1, 3),
]
tree_list, multiplied = geo.traverse_map_and_count_trees(slopes)
print(f"Part 1: Number of trees found: {sum(tree_list)}")

geo = Geo(map_)
slopes = [(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)]
tree_list, multiplied = geo.traverse_map_and_count_trees(slopes)
print(f"Part 2a: Number of trees found: {tree_list}")
print(f"Part 2b: Number of trees found, multiplied together: {multiplied}")