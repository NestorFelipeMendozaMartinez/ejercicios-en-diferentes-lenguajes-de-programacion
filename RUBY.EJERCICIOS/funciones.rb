# Función 1: Definición y llamada
def mostrar_texto
    puts "Hola"
  end
  
  # Función 2: Multiplicar 5 y 8 e imprimir el resultado
  def multiplicar
    a = 5
    b = 8
    puts a * b
  end
  
  # Función 3: Multiplicar con variables globales
  def multiplicar_con_variables_globales(a, b)
    puts a * b
  end
  
  # Función 4: Multiplicar con return y usar el resultado fuera de la función
  def multiplicar_con_return
    a = 5
    b = 8
    return a * b
  end
  
  # Función 5: Validar si el valor de a es 5
  def validar_dato(a)
    if a == 5
      return true
    else
      return false
    end
  end
  
  # Código principal
  a = [1, 2, 3, 4, 5]
  b = [6, 7, 8, 9, 10]
  c = []
  
  # Multiplicar los elementos correspondientes de a y b y almacenarlos en c
  a.each_with_index do |elem_a, i|
    c << elem_a * b[i]
  end
  
  puts c.inspect
  
  # Llamada a las funciones
  mostrar_texto
  multiplicar
  
  a_global = 5
  b_global = 8
  multiplicar_con_variables_globales(a_global, b_global)
  
  resultado = multiplicar_con_return
  puts resultado + 5
  
  dato = validar_dato(a_global)
  puts dato