
# Solicitar y leer los valores de A y B
print "Ingrese valor para A: "
A = gets.to_i
print "Ingrese valor para B: "
B = gets.to_i

# Realizar las operaciones
suma = A + B
puts "La suma de los números es: #{suma}"

res = A - B
puts "La resta de los números es: #{res}"

multi = A * B
puts "La multiplicación de los números es: #{multi}"

# Se asume que la división debe manejarse para evitar división por 0
if B != 0
  div = A.to_f / B
  puts "La división de los números es: #{div}"
else
  puts "Error: División por cero no permitida"
end

# Comparaciones
igual = (A == B)
puts "¿El número es igual? #{igual}"

mayor = (A > B)
puts "¿El número mayor es A? #{mayor}"

# Mostrar el número menor y mayor
if A > B
  puts "El número mayor es: #{A}"
  puts "El número menor es: #{B}"
else
  puts "El número mayor es: #{B}"
  puts "El número menor es: #{A}"
end
}