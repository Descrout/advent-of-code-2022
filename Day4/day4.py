class Pair:
    def __init__(self, txt):
        comma_split = txt.split(",")
        first = comma_split[0].split("-")
        second = comma_split[1].split("-")
        self.min1 = int(first[0])
        self.max1 = int(first[1])
        self.min2 = int(second[0])
        self.max2 = int(second[1])

    def fully_contains(self):
        return (self.min1 >= self.min2 and self.max1 <= self.max2) or (self.min2 >= self.min1 and self.max2 <= self.max1)

    def overlaps(self):
        return (self.min1 >= self.min2 and self.min1 <= self.max2) or (self.min2 >= self.min1 and self.min2 <= self.max1)

file = open('input.txt', 'r')
input_lines = file.readlines()

pairs = []
for line in input_lines:
    line = line.strip()
    pairs.append(Pair(line))

fully_contains = 0
overlaps = 0
for pair in pairs:
    if pair.fully_contains():
        fully_contains += 1
    
    if pair.overlaps():
        overlaps += 1
    
print("[Part1] Fully contained pairs:", fully_contains)
print("[Part2] Overlapped pairs:", overlaps)