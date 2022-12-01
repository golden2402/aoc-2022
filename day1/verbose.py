def main():
    # initial value:
    elves = [0]
    elf_i = 0

    with open("./sample.txt", "r") as file_data:
        for line in file_data:
            if len(line) > 1:
                elves[elf_i] += int(line)
            else:
                elves.append(0)
                elf_i += 1
    
        elves = sorted(elves)
        elves_count = len(elves)

        # part 1:
        print("Most calories:", elves[elves_count - 1])

        calorie_sum = 0
        for i in range(-3, 0):
            calorie_sum += elves[i]
        
        print("Top 3 total calories:", calorie_sum)


if __name__ == "__main__":
    main()
