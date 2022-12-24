file = open('input.txt', 'r')
input_lines = file.readlines()

def get_priority(letter):
    return letter - 96 if letter > 96 else letter - 38

def find_error(compartment):
    occurence = set()
    half = len(compartment) / 2

    for i in range(len(compartment)):
        code_unit = ord(compartment[i])
        if i < half:
            occurence.add(code_unit)
        elif code_unit in occurence:
            return code_unit

    raise Exception("could not find any recurring letter");

def get_unique_ascii(text):
    unique_letters = set()
    for letter in text:
        unique_letters.add(letter)

    return unique_letters

def find_recurring_in_group(group):
    occurences = dict()

    for compartment in group:
        compartment_unqiue = get_unique_ascii(compartment.strip())
        for letter in compartment_unqiue:
            occurence = occurences.get(letter, 0)
            occurences[letter] = occurence + 1

    return get_three_occurence_letter(occurences)

def get_three_occurence_letter(map):
    for key in map:
        if map[key] == 3:
            return key

    raise Exception("could not find 3 times recurring letter in a group")

sum_errors = 0
for line in input_lines:
    line = line.strip()
    sum_errors += get_priority(find_error(line))

print("[Part1] Sum of priorities:", sum_errors)

sum_recurring = 0
for i in range(0, len(input_lines), 3):
    recurring = find_recurring_in_group([
        input_lines[i],
        input_lines[i + 1],
        input_lines[i + 2],
    ])
    sum_recurring += get_priority(ord(recurring))

print("[Part2] Sum of all groups:", sum_recurring)