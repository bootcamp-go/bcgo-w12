# Lista de Códigos de Error Usuales de MySQL

## Code 1064 - You have an error in your SQL syntax
Este error suele ocurrir cuando hay un error de sintaxis en tu declaración SQL, como la falta de una palabra clave, paréntesis o comillas incorrectas.

## Code 1054 - Unknown column in 'field list'
Este error indica que estás intentando hacer referencia a una columna que no existe en la tabla que estás consultando o actualizando.

---

## Code 1048 - Column cannot be null
Este error ocurre cuando intentas insertar un registro con una restricción NOT NULL en una columna, pero proporcionas un valor nulo para esa columna.

## Code 1062 - Duplicate entry for key
Este error indica que estás intentando insertar un registro con un valor que ya existe en una columna con una restricción única, como una clave primaria o un índice único.

## Code 1451 - Cannot delete or update a parent row: foreign key constraint fails
Este error ocurre cuando intentas eliminar o actualizar un registro que tiene vinculado a él otros registros a través de alguna clave externa (foreign key) con restricciones.

## Code 1452 - Cannot add or update a child row: a foreign key constraint fails
Este error ocurre cuando intentas insertar o actualizar alguna clave externa (foreign key) con un registro que no existe.