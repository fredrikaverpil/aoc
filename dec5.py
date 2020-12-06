class Seat:
    def __init__(self, seat):
        self.row_formula = seat[:7]
        self.column_formula = seat[7:]
        self.rows = [i for i in range(0, 127 + 1)]
        self.columns = [i for i in range(0, 7 + 1)]

        self.validate()
        self.row = self.divide_and_conquer(self.rows.copy(), self.row_formula)
        self.column = self.divide_and_conquer(self.columns.copy(), self.column_formula)
        self.seat_id = self.get_seat_id()

    def validate(self):
        row_formula = "FBFBBFF"
        row = self.divide_and_conquer(self.rows.copy(), row_formula)
        assert row == 44

        column_formula = "RLR"
        column = self.divide_and_conquer(self.columns.copy(), column_formula)
        assert column == 5

    @staticmethod
    def divide_and_conquer(rows, formula):
        # print("*/* ->", rows)
        for letter in formula:
            half = len(rows) // 2
            if letter in ("F", "L"):
                rows = rows[:half]
                # print("F/L ->", rows)
            elif letter in ("B", "R"):
                rows = rows[half:]
                # print("B/L ->", rows)
        # print(f"Row: {rows[0]}")
        return rows[0]

    def get_seat_id(self):
        return (self.row * 8) + self.column


def missing_elements(lst):
    start, end = lst[0], lst[-1]
    return sorted(set(range(start, end + 1)).difference(lst))


with open("dec5/input.txt") as infile:
    seats = [line for line in infile.read().split("\n")]

seat_ids = []
for seat in seats:
    seat_id = Seat(seat).get_seat_id()
    seat_ids.append(seat_id)
print(f"Seat ID with highest number: {max(seat_ids)}")

sorted_seat_ids = sorted(seat_ids)
missing_seat_ids = missing_elements(sorted_seat_ids)
print(f"Missing seat ID: {missing_seat_ids}")
