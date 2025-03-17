# Crear la lista
lista = []

# Agregar elementos a la lista
lista.push("Lunes")
lista.push("Martes")
lista.push("Miercoles")
lista.push("Jueves")
lista.push("Viernes")
lista.push("Sabado")

# Imprimir el elemento en la posición 4 (comienza en 0)
puts lista[4]  # Trae el dato según posición en la lista comenzando desde cero

# Imprimir la lista completa
puts lista.inspect  # Trae la lista completa de datos

# Sublista: Elementos desde la posición 0 hasta la 3 (sin incluir la 3)
puts lista[0...3].inspect  # Trae la sublista desde la posición 0 hasta la 3 (sin incluir la 3)

# Lista con diferentes tipos de datos
lista_mixta = []

lista_mixta.push("Lunes")
lista_mixta.push("Martes")
lista_mixta.push("Miercoles")
lista_mixta.push("Jueves")
lista_mixta.push("Viernes")
lista_mixta.push("Sabado")
lista_mixta.push(1)
lista_mixta.push(2)
lista_mixta.push(3)

# Crear una lista interna dentro de la lista principal
datos_esteban = []
datos_esteban.push("Esteban")
datos_esteban.push(0.2)
datos_esteban.push(2.25)
datos_esteban.push(true)

# Agregar la lista interna a la lista principal
lista_mixta.push(datos_esteban)

# Imprimir la sublista (primeros 4 elementos) y los primeros 3 elementos de la lista interna
puts lista_mixta[0...4].inspect  # Imprime los primeros 4 elementos de la lista principal
puts lista_mixta[9][0...3].inspect  # Imprime los primeros 3 elementos de la lista interna