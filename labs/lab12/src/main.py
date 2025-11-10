def order_by_key(data, key): 
    return sorted(data, key = lambda d: d[key])

def ejercicio1():
    D = [
        {'make': 'Nokia', 'model': 216, 'color': 'Black'},
        {'make': 'Apple', 'model': 2, 'color': 'Silver'},
        {'make': 'Huawei', 'model': 50, 'color': 'Gold'},
        {'make': 'Samsung', 'model': 7, 'color': 'Blue'},
        {'make': 'Xiaomi', 'model': 931, 'color': 'Green'},
    ]
    # print(order_by_key(D, 'color'))
    # print(order_by_key(D, 'make'))
    print(order_by_key(D, 'model'))

def calculate_n_powers(data, n):
    return list(map(lambda x: x**n, data))

def ejercicio2():
    E = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    # print(calculate_n_powers(E, 2))
    # print(calculate_n_powers(E, 3))
    print(calculate_n_powers(E, 4))

def transpose_matrix(matrix):
    return list(map(lambda *row: list(row), zip(*matrix)))

def ejercicio3():
    X = [
        [1, 2, 3, 1],
        [4, 5, 6, 0],
        [7, 8, 9, -1]
    ]
    Y = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
    # print(transpose_matrix(X))
    print(transpose_matrix(Y))

def delete_elements(list1, list2):
    return list(filter(lambda x: x not in list2, list1))

def ejercicio4():
    # F = ['manzana', 'banana', 'amarillo', 'gris', 'rojo', 'naranja', 'morado']
    # G = ['manzana', 'banana', 'naranja', 'rojo', 'mango']
    # print(delete_elements(F, G))
    H = ['perro', 'gato', 'pez', 'loro', 'hamster']
    I = ['gato', 'loro', 'conejo']
    print(delete_elements(H, I))

if __name__ == "__main__":  
    # ejercicio1()
    # ejercicio2()
    # ejercicio3()
    ejercicio4()