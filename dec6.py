class Group:
    def __init__(self, group_string):
        self.answers = sorted(group_string.split("\n"))
        self.yes_count, self.yesses = self.get_yes_count()
        self.collective_yes_count = self.remove_non_collective_answer()

    def get_yes_count(self):
        yesses = []
        for answer in self.answers:
            for letter in answer:
                if letter not in yesses:
                    yesses.append(letter)
        return len(yesses), sorted(yesses)

    def remove_non_collective_answer(self):
        collective_yesses = self.yesses.copy()
        for letter in sorted(collective_yesses):
            for answer in self.answers:
                if letter not in answer:
                    collective_yesses.remove(letter)
                    break
        return len(collective_yesses)


with open("dec6/input.txt") as infile:
    group_strings = [line for line in infile.read().split("\n\n")]

sum_of_yes_counts = 0
sum_of_collective_yes_counts = 0
for group_string in group_strings:
    g = Group(group_string)
    sum_of_yes_counts += g.yes_count
    sum_of_collective_yes_counts += g.collective_yes_count

print(f"Sum of yes count: {sum_of_yes_counts}")
print(f"Sum of collective yes count: {sum_of_collective_yes_counts}")
