import re


class Passport:
    byr = None
    iyr = None
    eyr = None
    hgt = None
    hcl = None
    ecl = None
    pid = None
    cid = None

    def is_valid_v1(self):
        if (
            self.byr
            and self.iyr
            and self.eyr
            and self.hgt
            and self.hcl
            and self.ecl
            and self.pid
        ):
            return True
        return False

    def __str__(self):
        return (
            "Passport("
            f"byr={self.byr}, iyr={self.iyr}, eyr={self.eyr}, "
            f"hgt={self.hgt}, hcl={self.hcl}, ecl={self.ecl}, pid={self.pid}"
            ")"
        )

    def byr_is_valid(self):
        if (
            not self.byr
            or (not len(self.byr) == 4)
            or (int(self.byr) not in range(1920, 2002 + 1))
        ):
            return False
        return True

    def iyr_is_valid(self):
        if (
            not self.iyr
            or (not len(self.iyr) == 4)
            or (int(self.iyr) not in range(2010, 2020 + 1))
        ):
            return False
        return True

    def eyr_is_valid(self):
        if (
            not self.eyr
            or (not len(self.eyr) == 4)
            or (int(self.eyr) not in range(2020, 2030 + 1))
        ):
            return False
        return True

    def hgt_is_valid(self):
        if not self.hgt:
            return False
        if not (self.hgt.endswith("cm") or self.hgt.endswith("in")):
            return False
        if self.hgt.endswith("cm") and int(self.hgt.replace("cm", "")) not in range(
            150, 193 + 1
        ):
            return False
        if self.hgt.endswith("in") and int(self.hgt.replace("in", "")) not in range(
            59, 76 + 1
        ):
            return False
        return True

    def hcl_is_valid(self):
        if not self.hcl:
            return False
        matches = re.findall(r"^#[a-f0-9]{6}$", self.hcl)
        if not matches:
            return False
        return True

    def ecl_is_valid(self):
        if self.ecl not in ("amb", "blu", "brn", "gry", "grn", "hzl", "oth"):
            return False
        return True

    def pid_is_valid(self):
        if not self.pid:
            return False
        matches = re.findall(r"^[\d]{9}$", self.pid)
        if not matches:
            return False
        return True

    def is_valid_v2(self):
        if not self.byr_is_valid():
            return False
        if not self.iyr_is_valid():
            return False
        if not self.eyr_is_valid():
            return False
        if not self.hgt_is_valid():
            return False
        if not self.hcl_is_valid():
            return False
        if not self.ecl_is_valid():
            return False
        if not self.pid_is_valid():
            return False
        return True


def format_passports(passports_input):
    passports = []
    for passport_input in passports_input:
        p = Passport()
        entries = passport_input.replace("\n", " ").split(" ")
        for entry in entries:
            key_value = entry.split(":")
            key = key_value[0]
            value = key_value[1]
            # print(f"Setting {key} = {value}")
            setattr(p, f"{key}", value)
        passports.append(p)
    return passports


with open("dec4/input.txt") as infile:
    passports_input = [line for line in infile.read().split("\n\n")]
passports = format_passports(passports_input)

valid_passports = sum([passport.is_valid_v1() for passport in passports])
print(f"Number of valid passports v1: {valid_passports}")

valid_passports = sum([passport.is_valid_v2() for passport in passports])
print(f"Number of valid passports v2: {valid_passports}")
