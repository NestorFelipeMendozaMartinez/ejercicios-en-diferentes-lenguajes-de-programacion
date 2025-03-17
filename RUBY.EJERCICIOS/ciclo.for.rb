# Solicitar el primer valor (base)
puts "Ingrese el primer valor (base): "
a = gets.to_i

# Inicializar el valor de B (aunque no se use, se puede mantener como variable si es necesario)
b = 0

# Solicitar el segundo valor (exponente)
puts "Ingrese el segundo valor (exponente): "
c = gets.to_i

# Calcular la potencia
valor = a ** c  # Usamos el operador '**' para calcular la potencia en Ruby

# Mostrar el resultado
puts "La potencia de #{a} sobre #{c} es: #{valor}"