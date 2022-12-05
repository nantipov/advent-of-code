def main():
    max_calories = 0
    current_calories = 0
    with open('input1.txt', 'r') as reader:
        line = reader.readline()
        while line != '':
            if len(line.strip()) > 0:
                calories = int(line.strip())
                current_calories = current_calories + calories
            else:
                if current_calories > max_calories:
                    max_calories = current_calories
                current_calories = 0
            line = reader.readline()
    print(max_calories)


if __name__ == '__main__':
    main()
