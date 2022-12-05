def main():
    top_elves = [0, 0, 0]
    min_elf = 0
    current_calories = 0
    with open('input1.txt', 'r') as reader:
        line = reader.readline()
        while line != '':
            if len(line.strip()) > 0:
                calories = int(line.strip())
                current_calories = current_calories + calories
            else:
                if current_calories > min_elf:
                    min_elf = put_elf_into_top(top_elves, current_calories, min_elf)
                current_calories = 0
            line = reader.readline()
        if current_calories > min_elf:
            put_elf_into_top(top_elves, current_calories, min_elf)
    print(sum(top_elves))


def put_elf_into_top(top_elves, current_elf, min_elf) -> int:
    put_done = False
    new_min_elf = -1
    i = 0
    while i < len(top_elves):
        if not put_done and top_elves[i] == min_elf:
            top_elves[i] = current_elf
            put_done = True

        if new_min_elf < 0 or top_elves[i] < new_min_elf:
            new_min_elf = top_elves[i]

        i = i + 1

    return new_min_elf


if __name__ == '__main__':
    main()
