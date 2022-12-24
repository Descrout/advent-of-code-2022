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
    myScore = scores[mine]
    otherScore = scores[other]

    diff = myScore - otherScore
    if diff == 0:
        myScore += 3
    elif diff == 1 or diff == -2:
        myScore += 6

    return myScore

def get_my_score_part2(other, mine):
    symbolIdx = scores[other]

    match mine:
        case "X":
            return get_my_score(other, symbols[symbolIdx - 1])
        case "Z":
            return get_my_score(other, symbols[symbolIdx + 1])
        case _:
            return get_my_score(other, other)

scorePart1 = 0
scorePart2 = 0
for line in input_lines:
    line = line.strip()
    splitted = line.split(" ")
    other = splitted[0]
    mine = splitted[1]
    scorePart1 += get_my_score(other, mine)
    scorePart2 += get_my_score_part2(other, mine)

print("[Part1] Your score is:", scorePart1)
print("[Part2] Your score is:", scorePart2)