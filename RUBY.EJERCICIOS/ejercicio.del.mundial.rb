# Método para obtener el ganador del partido
def ganadores_octavos(marcador, equipo1, equipo2)
    goles = marcador.split(" - ")
    goles1 = goles[0].to_i
    goles2 = goles[1].to_i
    goles1 > goles2 ? equipo1 : equipo2
  end
  
  # Lista de partidos
  equipos = [
    "1. Países Bajos vs. Estados Unidos.",
    "2. Argentina vs. Australia.",
    "3. Francia vs. Polonia.",
    "4. Inglaterra vs. Senegal.",
    "5. Japón vs. Croacia.",
    "6. Brasil vs. Corea del Sur.",
    "7. Marruecos vs. España.",
    "8. Portugal vs. Suiza."
  ]
  
  # Mapa de resultados
  resultados = {}
  
  # Imprimir el encabezado
  puts "Mundial Qatar 2022\n\nOctavos de final\n"
  
  # Pedir marcadores y almacenar resultados
  equipos.each do |partido|
    puts partido
    print "Ingrese el marcador final (Formato: 2-1): "
    marcador = gets.chomp
    resultados[partido] = marcador
  end
  
  # Imprimir los ganadores
  puts "\nGanadores de la fase de octavos:"
  resultados.each do |partido, marcador|
    equipos_partido = partido.split(" vs. ")
    equipo1 = equipos_partido[0].split(" ")[1]  # Obtener el primer equipo
    equipo2 = equipos_partido[1].split(" ")[0]  # Obtener el segundo equipo
    ganador = ganadores_octavos(marcador, equipo1, equipo2)
    puts "Ganador de #{partido} es: #{ganador}"
  end