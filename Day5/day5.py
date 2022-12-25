class Procedure:
    def __init__(self, txt):
        splitted = txt.split(" ")
        self.move = int(splitted[1])
        self.frm = int(splitted[3]) - 1
        self.to = int(splitted[5]) - 1

class Crates:
    def __init__(self, size):
        self.crates = []
        for _ in range(size):
            self.crates.append([])

    def copy(self):
        length = len(self.crates)
        crates = Crates(length)
        for i in range(length):
            crates.crates[i] = self.crates[i].copy()
        return crates

    def fill_crates(self, line):
        j = 0
        for i in range(1, len(line), 4):
            if line[i] != " ":
                self.crates[j].insert(0, line[i])
            j += 1

    def top_of_stacks(self):
        return "".join(map(lambda x: x[-1], self.crates))

    def apply_proc(self, proc, is_9001 = False):
        sublist = self.crates[proc.frm][-proc.move:]
        self.crates[proc.to].extend(sublist if is_9001 else reversed(sublist))
        self.crates[proc.frm] = self.crates[proc.frm][:-proc.move]
        
file = open('input.txt', 'r')
input_lines = file.readlines()

crates = Crates(9)
procedures = []

for i in range(0, len(input_lines)):
    line = input_lines[i].strip("\n")
    if i < 8:
        crates.fill_crates(line)
        continue
    if i > 9:
        procedures.append(Procedure(line))

crates9001 = crates.copy()

for proc in procedures:
    crates.apply_proc(proc)
    crates9001.apply_proc(proc, True)

print("[Part1] Top of each stack CrateMover9000:", crates.top_of_stacks())
print("[Part2] Top of each stack CrateMover9001:", crates9001.top_of_stacks())
