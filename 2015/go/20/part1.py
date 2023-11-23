def get_presents(house):
    presents = 0
    for elf in range(1, house + 1):
        if house % elf == 0:
            presents += elf * 10
    return presents

def find_house(target):
    house = 1
    while True:
        presents = get_presents(house)
        if presents >= target:
            return house
        house += 1

target = 33100000
house = find_house(target)
print("The lowest house number of the house to get at least as many presents as", target, "is", house)
