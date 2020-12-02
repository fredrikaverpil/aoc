import re
from collections import namedtuple

RuleSet = namedtuple("RuleSet", ["min", "max", "letter", "password", "is_valid"])


def format_rules(rules_input):
    rule_sets = []
    regexp = re.compile(r"(\d+)-(\d+) (\w): (\w+)")
    for rule_set in rules_input:
        matches = re.match(regexp, rule_set)
        min_ = int(matches.group(1))
        max_ = int(matches.group(2))
        letter = matches.group(3)
        password = matches.group(4)
        rule_sets.append(RuleSet(min_, max_, letter, password, None))
    return rule_sets


def password_is_valid_v1(rule_set):
    letter_count = rule_set.password.count(rule_set.letter)
    if rule_set.min <= letter_count <= rule_set.max:
        return True
    return False


def password_is_valid_v2(rule_set):
    hits = 0
    if rule_set.password[rule_set.min - 1] == rule_set.letter:
        hits += 1
    if rule_set.password[rule_set.max - 1] == rule_set.letter:
        hits += 1
    if hits == 1:
        return True
    return False


with open("dec2/input.txt") as infile:
    rules_input = infile.readlines()

valid_passwords_v1 = sum(
    [password_is_valid_v1(rule_set) for rule_set in format_rules(rules_input)]
)
valid_passwords_v2 = sum(
    [password_is_valid_v2(rule_set) for rule_set in format_rules(rules_input)]
)
print(f"Valid number of passwords v1: {valid_passwords_v1}")
print(f"Valid number of passwords v2: {valid_passwords_v2}")
