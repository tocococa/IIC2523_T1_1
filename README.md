# Comentario
Para el algoritmo de busqueda de islas, decidimos implementar dfs, ya que si bien no es la funcion mas eficiente en terminos de tiempo (lineal en el tamaño del input), es relativamente facil de implementar. Si lo comparamos con BFS, usar DFS nos evita el problema de implementar una cola y nos permite trabajar sobre la misma matriz.

Luego, un metodo de implementar esta solución es utilizar una matrix auxiliar que nos ayude a llevar un registro de las casillas visitadas. En nuestro caso, elegimos trabajar sobre la misma matriz y no gastar espacio en memoria (el cual equivaldría a una complejidad de O(filas\*columnas) en el peor caso) puesto que no era necesario conservar la matriz original.

De haber requerido conservar la matriz original en memoria, un mejor approach habría sido usar BFS, dado que la complejidad computacional se mantendría lineal en el tamaño del input, pero la complejidad en el espacio sería menor que el caso de DFS con una matriz auxiliar. En el caso de DFS, el peor caso donde solo tengo 1s en la matriz la complejidad sería O(filas\*columnas), mientras que para BFS sería O(min(filas, columnas)).
