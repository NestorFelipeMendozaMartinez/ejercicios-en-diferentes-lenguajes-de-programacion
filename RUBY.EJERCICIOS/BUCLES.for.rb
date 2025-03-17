# Crear la lista de nombres
nombres = ["Esteban", "Hans", "Jhon", "Juan Pablo"]

# Imprimir los nombres
nombres.each do |nombre|
  puts nombre
end

# Crear la lista de personas (como lista de Hashes)
personas = []

# Crear los Hashes para cada persona
a = { "nombre" => "Esteban", "Edad" => 28 }
b = { "nombre" => "Hans", "Edad" => 27 }
c = { "nombre" => "Jhon", "Edad" => 41 }
d = { "nombre" => "Juan Pablo", "Edad" => 23 }

# Agregar los Hashes a la lista
personas.push(a, b, c, d)

# Imprimir los nombres y edades
personas.each do |persona|
  puts "#{persona["nombre"]} #{persona["Edad"]}"
end