# Solicitar al usuario sus datos personales

# Pedir al usuario que ingrese sus datos
print "Escriba sus nombres completos: "
a = gets.chomp  # Lee el nombre completo

print "Escriba sus Apellidos completos: "
b = gets.chomp  # Lee los apellidos completos

print "Escriba su profesion: "
c = gets.chomp  # Lee la profesión

print "Escriba su año de nacimiento: "
d = gets.to_i  # Lee el año de nacimiento

# Calcular la edad en base al año actual (2025)
e = 2025 - d

# Mostrar el resultado
puts "El (La) #{c} #{a} #{b} tiene #{e} años"