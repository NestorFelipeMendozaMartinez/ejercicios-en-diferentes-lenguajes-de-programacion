# Primer bloque: Contador con while
contador = 0
while contador < 30
  contador += 1
  puts "Iteración #{contador}"
end

# Segundo bloque: Combinado con if-else
loop do
  puts "Introduzca un valor mayor a 10:"
  a = gets.to_i
  
  if a > 10
    puts "Es correcto"
  else
    puts "Es incorrecto, pruebe de nuevo"
    break  # Sale del ciclo while si el valor ingresado es incorrecto
  end
end