from enum import Enum


class Shape(Enum):
    ROCK = 1
    PAPER = 2
    SCISSORS = 3


class Game(Enum):
    WON = 1
    LOST = 2
    DRAW = 3


def main():
    score1 = 0
    score2 = 0
    with open('input2.txt', 'r') as reader:
        line = reader.readline()
        while line != '':
            data = line.strip().split(" ", 2)
            opponent = get_shape(data[0])
            guide1 = guide_to_shape1(data[1])
            guide2 = guide_to_shape2(opponent, data[1])
            score1 = score1 + get_score(opponent, guide1)
            score2 = score2 + get_score(opponent, guide2)
            line = reader.readline()
    print("Task1 ", score1)
    print("Task2 ", score2)


def get_score(opponent, guide):
    shape_score = 0
    if guide == Shape.ROCK:
        shape_score = 1
    if guide == Shape.PAPER:
        shape_score = 2
    if guide == Shape.SCISSORS:
        shape_score = 3

    game_score = 0
    game = get_game(opponent, guide)
    if game == Game.WON:
        game_score = 6
    if game == Game.LOST:
        game_score = 0
    if game == Game.DRAW:
        game_score = 3

    return shape_score + game_score


def get_game(opponent, guide) -> Game:
    if guide == opponent:
        return Game.DRAW
    if (guide == Shape.ROCK and opponent == Shape.SCISSORS) or \
            (guide == Shape.SCISSORS and opponent == Shape.PAPER) or \
            (guide == Shape.PAPER and opponent == Shape.ROCK):
        return Game.WON
    return Game.LOST


def get_shape(code) -> Shape:
    if code == 'A':
        return Shape.ROCK
    if code == 'B':
        return Shape.PAPER
    if code == 'C':
        return Shape.SCISSORS
    raise RuntimeError("Unknown shape " + code)


def guide_to_shape1(code) -> Shape:
    if code == 'X':
        return Shape.ROCK
    if code == 'Y':
        return Shape.PAPER
    if code == 'Z':
        return Shape.SCISSORS
    raise RuntimeError("Unknown shape " + code)


def guide_to_shape2(opponent_shape, code) -> Shape:
    if code == 'X':  # lose
        if opponent_shape == Shape.ROCK:
            return Shape.SCISSORS
        if opponent_shape == Shape.PAPER:
            return Shape.ROCK
        return Shape.PAPER
    if code == 'Y':  # draw
        return opponent_shape
    if code == 'Z':  # win
        if opponent_shape == Shape.ROCK:
            return Shape.PAPER
        if opponent_shape == Shape.PAPER:
            return Shape.SCISSORS
        return Shape.ROCK
    raise RuntimeError("Unknown shape " + code)


# Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
# Rock     - A <- X -> need to lose
# Paper    - B <- Y -> need to end in a draw
# Scissors - C <- Z -> need to need to win


if __name__ == '__main__':
    main()
