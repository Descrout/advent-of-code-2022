file = open('input.txt', 'r')
input_lines = file.readlines()

latest_elf_calories = []
top_three_elf_calories = [0, 0, 0]
top_elf_calories_len = 3

def update_top_three(newCalories): 
    for i in range(top_elf_calories_len):
        if newCalories > top_three_elf_calories[i]:
            oldCalori = top_three_elf_calories[i]
            top_three_elf_calories[i] = newCalories
            update_top_three(oldCalori)
            break

for line in input_lines:
    line = line.strip()
    if len(line) == 0:
        total_calori = sum(latest_elf_calories)
        update_top_three(total_calori)
        latest_elf_calories.clear()
        continue

    latest_elf_calories.append(int(line))

print("[Part1] Most carried calories:", top_three_elf_calories[0])
top_three_sum = sum(top_three_elf_calories)
print("[Part2] Top three calories sum:", top_three_sum)

