# Solicitar al usuario los valores de voltaje y resistencia

# Pedir al usuario que ingrese los valores de voltaje y resistencia
print "Ingrese el valor del voltaje del circuito: "
voltaje = gets.to_i  # Lee el valor del voltaje

print "Ingrese el valor de la resistencia del circuito: "
resistencia = gets.to_i  # Lee el valor de la resistencia

# Calcular la intensidad (amperaje)
intensidad = voltaje.to_f / resistencia

# Mostrar el resultado
puts "Al conectar un resistor de R #{resistencia} ohm a una fuente de V #{voltaje} voltaje circulará una corriente de #{intensidad} amperios"