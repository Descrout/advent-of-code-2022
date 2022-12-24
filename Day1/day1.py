file = open('input.txt', 'r')
input_lines = file.readlines()

latest_elf_calories = []
max_elf_calori = 0

for line in input_lines:
    line = line.strip()
    if len(line) == 0:
        total_calori = sum(latest_elf_calories)
        if total_calori > max_elf_calori:
            max_elf_calori = total_calori
        latest_elf_calories.clear()
        continue

    latest_elf_calories.append(int(line))

print('Most carried calories:', max_elf_calori)
