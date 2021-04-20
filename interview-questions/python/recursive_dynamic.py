
def e8_1_stairway(steps):
    return find(steps)

cache = {}

def find(left):
    if left in cache:
        return cache[left]

    if left==0:
        return 1

    else:
        c = 0
        for i in range(1,4):
            if i > left:
                break
            c += find(left-i)
        cache[left] = c
        return c
