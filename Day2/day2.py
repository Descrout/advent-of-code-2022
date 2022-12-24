file = open('input.txt', 'r')
input_lines = file.readlines()

scores = {
    "A": 1,
    "B": 2,
    "C": 3,
    "X": 1,
    "Y": 2,
    "Z": 3,
}

# Some languages do negative number modulus different
# So instead of % wrap around
# I put "C" as first and "A" as last
symbols = ["C", "A", "B", "C", "A"]

def get_my_score(other, mine):
    my_score = scores[mine]
    other_score = scores[other]

    diff = my_score - other_score
    if diff == 0:
        my_score += 3
    elif diff == 1 or diff == -2:
        my_score += 6

    return my_score

def get_my_score_part2(other, mine):
    symbol_idx = scores[other]

    match mine:
        case "X":
            return get_my_score(other, symbols[symbol_idx - 1])
        case "Z":
            return get_my_score(other, symbols[symbol_idx + 1])
        case _:
            return get_my_score(other, other)

score_p1 = 0
score_p2 = 0
for line in input_lines:
    line = line.strip()
    splitted = line.split(" ")
    other = splitted[0]
    mine = splitted[1]
    score_p1 += get_my_score(other, mine)
    score_p2 += get_my_score_part2(other, mine)

print("[Part1] Your score is:", score_p1)
print("[Part2] Your score is:", score_p2)