# Primer bloque: Verificar si 'a' es igual a 2
a = 2
if a == 2
  puts "a vale 2"
else
  puts "a es diferente de 2"
end

# Segundo bloque: Concatenar condiciones con 'and' y 'not'
nombre = "Esteban"
edad = 28
pais = "Colombia"

if nombre == "Esteban" && edad == 28 && pais == "Colombia"
  puts "Su nombre es #{nombre}, tiene #{edad} y es de #{pais}"
elsif nombre == "Esteban" && edad != 28
  puts "Su nombre es Esteban y no tiene 28 años"
elsif nombre != "Esteban" && edad == 28
  puts "Su nombre no es Esteban y tiene 28 años"
else
  puts "No se llama Esteban ni tiene 28 años"
end