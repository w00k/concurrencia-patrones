# Concurrencia y patrones de diseño


## Condición de carrera

Cuando un recurso es usado por 2 procesos concurrentes y no hay certeza de cual termina primero. 

Imaginemos en una variable **balance** igual a 500 y que tenemos una función que espera un número y suma ese valor al **balance**. Ahora supongamos que corremos esta función concurrentemente, ingresando 200 en la primera llamada y 100 en la segunda. Pueden pasar 3 cosas:

1. Que termine el primer proceso y después se ejecute el segundo proceso.
```bash 
500 + 200 = 700 
//y posteriormente 
700 + 100 = 800
``` 

2. Que termine el segundo proceso primero y que déspues termine el primer proceso.
```bash 
500 + 100 = 600 
//y posteriormente 
500 + 200 = 700
```

3. Que termine el primer proceso y el segundo proceso se ejecute antes que termine el primer proceso. 
```bash
500 + 200 + 700 
//y posteriormente 
500 + 100 = 600
```

Go/Golang provee herramientas para superar ls condición de carrera. 
- sync.Mutex.Lock() nos ayudará a bloquear el acceso a valores compartidos en diferentes GoRoutines.
- sync.Mutex.Unlock() desbloqueará nuevamente el valor al que necesitamos acceder.
- sync.RWMutex nos permite bloquear para escritura (Lock() y UnLock()) y para lectura (RLock() y RUnlock()).

