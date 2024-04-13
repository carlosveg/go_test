# API GO

### Descripción

API de registro de usuarios de manera estática

### Endpoints

- GET /: Regresa los usuarios registrados
- POST /: Crea un usuario, recibe los datos en el body
- POST /login: recibe credenciales enel body y retorna un JWT en caso de éxito

### Tareas

- [x] Crear un servicio de registro de usuario que reciba como parámetros usuario, correo, teléfono y contraseña.
- [x] El servicio deberá validar que el correo y telefono no se encuentren registrados, de lo contrario deberá retornar un mensaje “el correo/telefono ya se encuentra registrado”.
- [x] Deberá validar que la contraseña sea de 6 caracteres mínimo y 12 máximo y contener al menos una mayúscula, una minúscula, un carácter especial (@ $ o &) y un número.
- [x] Validar que el teléfono sea a 10 dígitos y el correo tenga un formato válido.
- [x] Crear un servicio login que reciba como parámetros usuario o correo y contraseña.
- [x] El servicio debe devolver un token jwt.
- [x] Deberá validar que el usuario o correo y contraseña sean válidos, de lo contrario retorna un mensaje “usuario / contraseña incorrectos”.
- [x] En ambos servicios se deberá validar que todos los parámetros solicitados vayan en el cuerpo de la petición, de lo contrario retorna un mensaje con el campo faltante.
