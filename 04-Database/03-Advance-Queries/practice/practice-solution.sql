USE empresa_db;

------------------------------------------------------------
-- Seleccionar el nombre, el puesto del empleado y la localidad de los departamentos donde trabajan los vendedores.
SELECT e.nombre, e.puesto, d.localidad FROM empleados e
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
WHERE e.puesto = 'Vendedor';


------------------------------------------------------------
-- Visualizar los departamentos con más de 1 empleado.
SELECT d.depto_nro, d.localidad FROM empleados e
RIGHT JOIN departamentos d ON e.depto_nro = d.depto_nro
GROUP BY d.depto_nro, d.localidad HAVING COUNT(e.depto_nro) > 1;

SELECT d.depto_nro, d.localidad, e2.cantidad FROM (
	SELECT COUNT(*) as cantidad, e.* FROM empleados e
	GROUP BY e.depto_nro
	HAVING cantidad > 1
) as e2
INNER JOIN departamentos d ON d.depto_nro = e2.depto_nro;

select d.*, e.cantidad
from departamentos d
INNER JOIN (
	select e.depto_nro, count(*) cantidad 
	from empleados e
	group by e.depto_nro
	having cantidad > 1
) e ON d.depto_nro = e.depto_nro;


------------------------------------------------------------
-- Mostrar el nombre, salario del empleado y nombre del departamento que tengan el mismo puesto
-- que ‘Mito Barchuk’.
SELECT e.nombre, e.salario, d.nombre FROM empleados e
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
WHERE e.puesto = (
	SELECT e2.puesto FROM empleados e2 WHERE e2.nombre = 'Mito' AND e2.apellido = 'Barchuk'
);


------------------------------------------------------------
-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad,
-- ordenados por nombre.
SELECT e.* FROM empleados e 
INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
WHERE d.nombre = 'Contabilidad' 
ORDER BY e.nombre;

SELECT e.* FROM empleados e
WHERE e.depto_nro = (
	SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Contabilidad"
)
ORDER BY e.nombre;


------------------------------------------------------------
-- Mostrar el nombre del empleado que tiene el salario más bajo.
SELECT e.nombre FROM empleados e
WHERE e.salario = (
	SELECT MIN(e2.salario) FROM empleados e2
);


------------------------------------------------------------
-- Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
SELECT e.* FROM empleados e
WHERE e.salario = (
	SELECT MAX(e2.salario) FROM empleados e2 WHERE e2.depto_nro = (
		SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Ventas"
    )
) AND e.depto_nro = (
	SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Ventas"
);

SELECT e.* FROM empleados e INNER JOIN departamentos d ON e.depto_nro = d.depto_nro
AND d.nombre_depto = "Ventas" AND e.salario = (
	SELECT MAX(e2.salario) FROM empleados e2 WHERE e2.depto_nro = (
		SELECT d.depto_nro FROM departamentos d WHERE d.nombre_depto = "Ventas"
    )
);

SELECT e.* FROM empleados e INNER JOIN departamentos d ON e.depto_nro = d.depto_nro AND d.nombre_depto = "Ventas"
ORDER BY e.salario DESC
LIMIT 1;