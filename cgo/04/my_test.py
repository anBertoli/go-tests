import time

def concat(name):
    res = "Hello from " + name + "."
    print(res)
    return res


def greet():
    res = "Hello!"
    print(res)
    return res


def get_num():
    return round(time.time() * 1000)

