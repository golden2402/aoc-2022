import sys

lines = [line.rstrip() for line in sys.stdin]

ASCII_A = ord('a')


def to_priority(char):
    lower = char.lower()
    return ord(lower) - ASCII_A + 1 + (26 if char != lower else 0)


# part 1:
print(sum([to_priority(list(set(rucksack[:(len(rucksack) // 2)]) &
      set(rucksack[(len(rucksack) // 2):]))[0]) for rucksack in lines]))

# part 2:
print(sum([to_priority(list(set(chunk[0]) & set(chunk[1]) & set(chunk[2]))[0])
      for chunk in [lines[n:n+3] for n in range(0, len(lines), 3)]]))
