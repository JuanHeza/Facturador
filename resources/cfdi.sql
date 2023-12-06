/*
SQLyog Ultimate v8.82 
MySQL - 5.7.11-log : Database - cfdi
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`cfdi` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `cfdi`;

/*Table structure for table `catalogo` */

DROP TABLE IF EXISTS `catalogo`;

CREATE TABLE `catalogo` (
  `CATALOGO_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Numero identificador de un catalogo',
  `NOMBRE` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Nombre del catalogo',
  `DESCRIPCION` varchar(5000) COLLATE utf8_spanish_ci DEFAULT NULL COMMENT 'Nombre del catalogo',
  `ESTATUS_ID` bigint(20) NOT NULL COMMENT 'Estatus',
  `ELIMINADO_ID` bigint(20) NOT NULL COMMENT 'Eliminado',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'Fecha en que fue agregada la información por primera vez',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que genera o agrega por primera vez informacion en la tabla',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'Fecha en la que fueron realizados cambios a la Información',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que realizo modificaciones a la tabla',
  PRIMARY KEY (`CATALOGO_ID`),
  KEY `CATALOGO_USUARIO_USUARIOCREACION_FK` (`USUARIO_CREACION`),
  KEY `CATALOGO_USUARIO_USUARIOMODIFICO_FK` (`USUARIO_MODIFICO`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci COMMENT='Almacena la información sobre los diferentes catálogos que e';

/*Data for the table `catalogo` */

insert  into `catalogo`(`CATALOGO_ID`,`NOMBRE`,`DESCRIPCION`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`) values (1,'Estatus','Catálogo de con los valores estándar de Estatus',1,3,'2017-10-18 00:00:00',1,'2019-01-31 18:19:09',1),(2,'Eliminado','Catálogo de con los valores estándar de Eliminado',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1);

/*Table structure for table `catalogo_parametro` */

DROP TABLE IF EXISTS `catalogo_parametro`;

CREATE TABLE `catalogo_parametro` (
  `CATALOGO_PARAMETRO_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Numero identificador para los parametros del sistema',
  `CATALOGO_ID` bigint(20) NOT NULL COMMENT 'Numero identificador de un catalogo',
  `NOMBRE` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Nombre del Catalogo de Parametros',
  `ORDEN` bigint(20) DEFAULT NULL,
  `ESTATUS_ID` bigint(20) NOT NULL COMMENT 'Numero identificador de la tabla de Estatus',
  `ELIMINADO_ID` bigint(20) DEFAULT NULL COMMENT 'Numero identificador del eliminado logico de una Tabla',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'Fecha en que fue agregada la información por primera vez',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que genera o agrega por primera vez informacion en la tabla',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'Fecha en la que fueron realizados cambios a la Información',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que realizo modificaciones a la tabla',
  PRIMARY KEY (`CATALOGO_PARAMETRO_ID`),
  KEY `CAPA_CATA_FK_I` (`CATALOGO_ID`),
  KEY `CATALOGOPARAMETRO_USUARIO_USUARIOCREACION_FK` (`USUARIO_CREACION`),
  KEY `CATALOGOPARAMETRO_USUARIO_USUARIOMODIFICO_FK` (`USUARIO_MODIFICO`),
  CONSTRAINT `CATALOGOPARAMETRO_CATALOGO_CATALOGOID_FK` FOREIGN KEY (`CATALOGO_ID`) REFERENCES `catalogo` (`CATALOGO_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci COMMENT='Almacena la información de los parametros de los catálogos ';

/*Data for the table `catalogo_parametro` */

insert  into `catalogo_parametro`(`CATALOGO_PARAMETRO_ID`,`CATALOGO_ID`,`NOMBRE`,`ORDEN`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`) values (1,1,'Activo',1,1,3,'2017-10-18 00:00:00',1,'2019-01-31 18:19:09',1),(2,1,'Inactivo',2,1,3,'2017-10-18 00:00:00',1,'2019-01-31 18:19:09',1),(3,2,'No Eliminado',1,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(4,2,'Eliminado',2,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1);

/*Table structure for table `configuracion` */

DROP TABLE IF EXISTS `configuracion`;

CREATE TABLE `configuracion` (
  `CONFIGURACION_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Número que identifica al registro de  configuración',
  `NOMBRE` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Nombre del parámetro de configuración',
  `DESCRIPCION` varchar(5000) COLLATE utf8_spanish_ci DEFAULT NULL COMMENT 'Nombre del catalogo',
  `ESTATUS_ID` bigint(20) NOT NULL COMMENT 'Numero identificador de la tabla de Estatus',
  `ELIMINADO_ID` bigint(20) DEFAULT NULL COMMENT 'Numero identificador del Estatus de una Tabla',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'Fecha en que fue agregada la información por primera vez',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que genera o agrega por primera vez informacion en la tabla',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'Fecha en la que fueron realizados cambios a la Información',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que realizo modificaciones a la tabla',
  PRIMARY KEY (`CONFIGURACION_ID`),
  KEY `CONFIGURACION_CATALOGOPARAMETRO_ELIMINADOID_FK` (`ELIMINADO_ID`),
  KEY `CONFIGURACION_CATALOGOPARAMETRO_ESTATUSID_FK` (`ESTATUS_ID`),
  KEY `CONFIGURACION_USUARIO_USUARIOCREACION_FK` (`USUARIO_CREACION`),
  KEY `CONFIGURACION_USUARIO_USUARIOMODIFICO_FK` (`USUARIO_MODIFICO`),
  CONSTRAINT `CONFIGURACION_CATALOGOPARAMETRO_ELIMINADOID_FK` FOREIGN KEY (`ELIMINADO_ID`) REFERENCES `catalogo_parametro` (`CATALOGO_PARAMETRO_ID`),
  CONSTRAINT `CONFIGURACION_CATALOGOPARAMETRO_ESTATUSID_FK` FOREIGN KEY (`ESTATUS_ID`) REFERENCES `catalogo_parametro` (`CATALOGO_PARAMETRO_ID`),
  CONSTRAINT `CONFIGURACION_USUARIO_USUARIOCREACION_FK` FOREIGN KEY (`USUARIO_CREACION`) REFERENCES `usuario` (`USUARIO_ID`),
  CONSTRAINT `CONFIGURACION_USUARIO_USUARIOMODIFICO_FK` FOREIGN KEY (`USUARIO_MODIFICO`) REFERENCES `usuario` (`USUARIO_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci COMMENT='Tabla donde se registran los valores que pueden alterarse pa';

/*Data for the table `configuracion` */

insert  into `configuracion`(`CONFIGURACION_ID`,`NOMBRE`,`DESCRIPCION`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`) values (1,'Correo','Configuración para envío de correo.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(2,'Mensajes recuperar contrasena','Mensajes para la recuperación de contraseña.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(3,'Aplicacion','Datos generales de la aplicación.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(4,'upload','Configuración para subida de archivos',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(5,'Desarrollador','Datos generales del desarrollador.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(6,'Cliente','Datos generales del cliente.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1);

/*Table structure for table `configuracion_parametro` */

DROP TABLE IF EXISTS `configuracion_parametro`;

CREATE TABLE `configuracion_parametro` (
  `CONFIGURACION_PARAMETRO_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Número que identifica al registro de  configuración',
  `CONFIGURACION_ID` bigint(20) DEFAULT NULL,
  `NOMBRE` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Nombre del parámetro de configuración',
  `VALOR` longtext COLLATE utf8_spanish_ci COMMENT 'Valor del parámetro de configuración',
  `DESCRIPCION` varchar(5000) COLLATE utf8_spanish_ci DEFAULT NULL COMMENT 'Nombre del catalogo',
  `ESTATUS_ID` bigint(20) NOT NULL COMMENT 'Numero identificador de la tabla de Estatus',
  `ELIMINADO_ID` bigint(20) DEFAULT NULL COMMENT 'Numero identificador del Estatus de una Tabla',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'Fecha en que fue agregada la información por primera vez',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que genera o agrega por primera vez informacion en la tabla',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'Fecha en la que fueron realizados cambios a la Información',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'Nombre de Usuario que realizo modificaciones a la tabla',
  PRIMARY KEY (`CONFIGURACION_PARAMETRO_ID`),
  KEY `CONFIGURACIONPARAMETRO_CATALOGOPARAMETRO_ELIMINADOID_FK` (`ELIMINADO_ID`),
  KEY `CONFIGURACIONPARAMETRO_CATALOGOPARAMETRO_ESTATUSID_FK` (`ESTATUS_ID`),
  KEY `CONFIGURACIONPARAMETRO_CONFIGURACION_CONFIGURACIONID_FK` (`CONFIGURACION_ID`),
  KEY `CONFIGURACIONPARAMETRO_USUARIO_USUARIOCREACION_FK` (`USUARIO_CREACION`),
  KEY `CONFIGURACIONPARAMETRO_USUARIO_USUARIOMODIFICO_FK` (`USUARIO_MODIFICO`),
  CONSTRAINT `CONFIGURACIONPARAMETRO_CATALOGOPARAMETRO_ELIMINADOID_FK` FOREIGN KEY (`ELIMINADO_ID`) REFERENCES `catalogo_parametro` (`CATALOGO_PARAMETRO_ID`),
  CONSTRAINT `CONFIGURACIONPARAMETRO_CATALOGOPARAMETRO_ESTATUSID_FK` FOREIGN KEY (`ESTATUS_ID`) REFERENCES `catalogo_parametro` (`CATALOGO_PARAMETRO_ID`),
  CONSTRAINT `CONFIGURACIONPARAMETRO_CONFIGURACION_CONFIGURACIONID_FK` FOREIGN KEY (`CONFIGURACION_ID`) REFERENCES `configuracion` (`CONFIGURACION_ID`),
  CONSTRAINT `CONFIGURACIONPARAMETRO_USUARIO_USUARIOCREACION_FK` FOREIGN KEY (`USUARIO_CREACION`) REFERENCES `usuario` (`USUARIO_ID`),
  CONSTRAINT `CONFIGURACIONPARAMETRO_USUARIO_USUARIOMODIFICO_FK` FOREIGN KEY (`USUARIO_MODIFICO`) REFERENCES `usuario` (`USUARIO_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

/*Data for the table `configuracion_parametro` */

insert  into `configuracion_parametro`(`CONFIGURACION_PARAMETRO_ID`,`CONFIGURACION_ID`,`NOMBRE`,`VALOR`,`DESCRIPCION`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`) values (1,1,'Remitente','noreply.spacectpermits@lloko.mx','Dirección desde donde se envía el correo',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(2,1,'Contrasena','Novasys2023','Contraseña de la dirección de correo',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(3,1,'SMTP','mail.lloko.mx','Cliente SMTP',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(4,1,'Puerto','465','Puerto del cliente SMTP',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(5,1,'Autenticacion','true','Autenticacion',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(6,2,'Error usuario','El nombre de usuario <b>[user]</b> no es válido.','Mensaje del sistema al usuario cuando no se puede encontrar el usuario en la base de datos. El texto es leído como html.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(7,2,'Error envio','Error al procesar la recuperación de contraseña, reintente nuevamente.Si el problema persiste consulte con el administrador del sistema.','Mensaje del sistema al usuario cuando hubo un error al procesar la recuperación de contraseña. El texto es leído como html.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(8,2,'Envio ok','La contraseña ha sido enviada a:<br><br><b>[email]</b><br><br>Revise su bandeja de entrada para recuperar su contraseña.','Mensaje del sistema al usuario cuando el proceso de recuperación de contraseña se efectuó correctamente. El texto es leído como html.',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(9,2,'Envio mensaje','Ha solicitado la recuperación de contraseña. Su contraseña es: [password]','Es el mensaje del correo que se envía cuando el proceso de recuperación de contraseña se efectuó correctamente',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(10,2,'Envio titulo','Recuperación de contraseña','Es el titulo del correo que se envía cuando el proceso de recuperación de contraseña se efectuó correctamente',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(11,3,'Aplicacion nombre abreviado','SOT','Es el nombre abreviado de la aplicación',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(12,3,'Aplicacion nombre completo','M.S','Es el nombre completo de la aplicación',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(13,3,'Aplicacion version','Ver.: 1.0.1','Es la version de la aplicacion',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(14,6,'Cliente nombre','Nombre','Es el nombre del cliente',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(15,6,'Cliente direccion','...','Es la dirección completa del cliente',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(16,6,'Cliente telefono','...','Es el teléfono del cliente',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(17,5,'Desarrollador nombre','Desarrollado por: Novasys','Nombre del desarrollador',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(18,5,'Desarrollador web','www.novasys.com.mx','Sitio web del desarrollador',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(19,5,'Desarrollador correo','novasys@novasys.com.mx','Es el correo electrónico del desarrollador',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(20,4,'ruta_imagenes','/home/cfdi','Es la ruta donde se guardan los uploads para imagenes. Utilizar como separador el caracter \"/\".',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(21,2,'Envio mensaje contrasena','Ha cambiado la contraseña, Su nueva contraseña es: [password] para el sistema de Cotizador','Es el mensaje del correo que se envía cuando el proceso de recuperación de contraseña se efectuó correctamente',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(22,2,'Envio titulo contrasena','Cambio de contraseña','Es el titulo del correo que se envía cuando el proceso de recuperación de contraseña se efectuó correctamente',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(23,4,'Url_servicios','http://desarrollo.novasys.com.mx/cfdiWS/Facturacion/','URL de proyecto de servicios',1,4,'2023-11-03 18:44:00',1,'2023-11-03 18:44:00',1);

/*Table structure for table `factura` */

DROP TABLE IF EXISTS `factura`;

CREATE TABLE `factura` (
  `FACTURA_ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `NEGOCIO_ID` bigint(20) NOT NULL,
  `RECEPTOR_ID` bigint(20) NOT NULL,
  `FOLIO_TICKET` varchar(20) NOT NULL,
  `RUTA_PDF` varchar(500) NOT NULL,
  `ESTATUS_ID` tinyint(2) DEFAULT NULL,
  `ELIMINADO_ID` tinyint(2) DEFAULT NULL,
  `FECHA_CREACION` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `USUARIO_CREACION` bigint(20) DEFAULT NULL,
  `FECHA_MODIFICO` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL,
  `RAZON_SOCIAL_RECEPTOR` varchar(200) DEFAULT NULL COMMENT 'RAZON SOCIAL DEL RECEPTOR',
  `RFC_RECEPTOR` varchar(20) DEFAULT NULL COMMENT 'RFC DEL RECEPTOR',
  `CORREO_RECEPTOR` varchar(100) DEFAULT NULL COMMENT 'CORREO DEL RECEPTOR',
  `REGIMEN_FISCAL_RECEPTOR` bigint(5) DEFAULT NULL COMMENT 'REGIMEN FISCAL DEL RECEPTOR',
  `USO_CFDI_RECEPTOR` varchar(5) DEFAULT NULL COMMENT 'USO DEL CFDI DEL RECEPTOR',
  `TIPO_VENTA_RECEPTOR` bigint(5) DEFAULT NULL COMMENT 'TIPO DE VENTA',
  `CALLE_RECEPTOR` varchar(100) DEFAULT NULL COMMENT 'CALLE DEL RECEPTOR',
  `NUMERO_EXTERIOR_RECEPTOR` mediumint(20) DEFAULT NULL COMMENT 'NUMERO ESXTERIOR DEL RECEPTOR',
  `NUMERO_INTERIOR_RECEPTOR` mediumint(20) DEFAULT NULL COMMENT 'NUMERO INTERIOR DEL RECEPTOR',
  `COLONIA_RECEPTOR` varchar(100) DEFAULT NULL COMMENT 'COLONIA DEL RECEPTOR',
  `CODIGO_POSTAL_RECEPTOR` varchar(5) DEFAULT NULL COMMENT 'CODIGO POSTAL DEL RECEPTOR',
  `MUNICIPIO_RECEPTOR` varchar(100) DEFAULT NULL COMMENT 'MUNICIPIO DEL RECEPTOR',
  `ESTADO_RECEPTOR` varchar(100) DEFAULT NULL COMMENT 'ESTADO DEL RECEPTOR',
  `PAIS_RECEPTOR` varchar(100) DEFAULT NULL COMMENT 'PAIS DEL RECEPTOR',
  `SUCURSAL_ID` bigint(20) DEFAULT NULL,
  `RFC_EMISOR` varchar(20) DEFAULT NULL,
  `NOMBRE_EMISOR` varchar(200) DEFAULT NULL,
  `REGIMEN_FISCAL_EMISOR` varchar(5) DEFAULT NULL,
  `FACTURA_UUID` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`FACTURA_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8;

/*Data for the table `factura` */

insert  into `factura`(`FACTURA_ID`,`NEGOCIO_ID`,`RECEPTOR_ID`,`FOLIO_TICKET`,`RUTA_PDF`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`,`RAZON_SOCIAL_RECEPTOR`,`RFC_RECEPTOR`,`CORREO_RECEPTOR`,`REGIMEN_FISCAL_RECEPTOR`,`USO_CFDI_RECEPTOR`,`TIPO_VENTA_RECEPTOR`,`CALLE_RECEPTOR`,`NUMERO_EXTERIOR_RECEPTOR`,`NUMERO_INTERIOR_RECEPTOR`,`COLONIA_RECEPTOR`,`CODIGO_POSTAL_RECEPTOR`,`MUNICIPIO_RECEPTOR`,`ESTADO_RECEPTOR`,`PAIS_RECEPTOR`,`SUCURSAL_ID`,`RFC_EMISOR`,`NOMBRE_EMISOR`,`REGIMEN_FISCAL_EMISOR`,`FACTURA_UUID`) values (23,3,17,'1521','/galeria/facturador/XOJI740919U48-CACX7605101P8-dIqnAvkdA6FRcmRye-xDRg2',1,3,'2023-11-09 16:26:56',0,'2023-11-09 16:26:56',0,'XOCHILT CASAS CHAVEZ','CACX7605101P8','JUAN.HERNANDEZ@NOVASYS.COM.MX',605,'CP01',29,'ATOTONILCO',806,NULL,'CHEPEVERA','36257','015','GUA','MEX',2,'XOJI740919U48','INGRID XODAR JIMENEZ','626','dIqnAvkdA6FRcmRye-xDRg2');

/*Table structure for table `negocio` */

DROP TABLE IF EXISTS `negocio`;

CREATE TABLE `negocio` (
  `NEGOCIO_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'LLAVE PRIMARIA',
  `NOMBRE` varchar(200) DEFAULT NULL COMMENT 'NOMBRE DE EMPRESA',
  `URL_LOGO_NEGOCIO` varchar(500) DEFAULT NULL COMMENT 'URL DE LOGO DE EMPRESA',
  `URL_IMAGEN_NEGOCIO` varchar(500) DEFAULT NULL COMMENT 'URL DE IMAGEN DE EMPRESA',
  `MENU_TXT_TITULO_1` varchar(500) DEFAULT NULL COMMENT 'TEXTO DEL TITULO 1 EN EL ENCABEZADO',
  `MENU_URL_TITULO_1` varchar(500) DEFAULT NULL COMMENT 'URL DEL TITULO 1 EN EL ENCABEZADO ',
  `MENU_TXT_TITULO_2` varchar(500) DEFAULT NULL COMMENT 'TEXTO DEL TITULO 2 EN EL ENCABEZADO',
  `MENU_URL_TITULO_2` varchar(500) DEFAULT NULL COMMENT 'URL DEL TITULO 2 EN EL ENCABEZADO',
  `MENU_TXT_TITULO_3` varchar(500) DEFAULT NULL COMMENT 'TEXTO DEL TITULO 3 EN EL ENCABEZADO ',
  `MENU_URL_TITULO_3` varchar(500) DEFAULT NULL COMMENT 'URL DEL TITULO 3 EN EL ENCABEZADO ',
  `CLAVE` varchar(100) DEFAULT NULL COMMENT 'CLAVE DE EMPRESA UTILIZADA EN EL INDEXACTION PARA SACAR LA INFO DEL NEGOCIO',
  `COLOR_FONDO_ENFASIS` varchar(7) DEFAULT NULL COMMENT 'COLOR EN FORMATO HEXADECIMAL, P.E. #4B0082',
  `COLOR_TEXTO_ENFASIS` varchar(7) DEFAULT NULL COMMENT 'COLOR EN FORMATO HEXADECIMAL, P.E. #75CEDE',
  `ACERCA_DE_TITULO` varchar(100) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `ACERCA_DE_TEXTO` varchar(1000) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `PIE_LINK_TXT_1` varchar(100) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `PIE_LINK_URL_1` varchar(1000) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `PIE_LINK_TXT_2` varchar(100) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `PIE_LINK_URL_2` varchar(1000) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `PIE_LINK_TXT_3` varchar(100) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `PIE_LINK_URL_3` varchar(1000) DEFAULT NULL COMMENT 'TEXTO EN PANTALLA',
  `ESTATUS_ID` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `ELIMINADO_ID` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `CONCEPTO_GENERICO_id` bigint(10) DEFAULT NULL,
  `REGIMEN_FISCAL` varchar(5) DEFAULT NULL,
  `RFC` varchar(13) DEFAULT NULL,
  `IMPUESTO` bigint(3) DEFAULT NULL,
  `CONCEPTO_GENERICO_TEXTO` varchar(100) DEFAULT NULL,
  `CODIGO_POSTAL` varchar(5) DEFAULT NULL,
  `CODIGO_UNITARIO` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`NEGOCIO_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

/*Data for the table `negocio` */

insert  into `negocio`(`NEGOCIO_ID`,`NOMBRE`,`URL_LOGO_NEGOCIO`,`URL_IMAGEN_NEGOCIO`,`MENU_TXT_TITULO_1`,`MENU_URL_TITULO_1`,`MENU_TXT_TITULO_2`,`MENU_URL_TITULO_2`,`MENU_TXT_TITULO_3`,`MENU_URL_TITULO_3`,`CLAVE`,`COLOR_FONDO_ENFASIS`,`COLOR_TEXTO_ENFASIS`,`ACERCA_DE_TITULO`,`ACERCA_DE_TEXTO`,`PIE_LINK_TXT_1`,`PIE_LINK_URL_1`,`PIE_LINK_TXT_2`,`PIE_LINK_URL_2`,`PIE_LINK_TXT_3`,`PIE_LINK_URL_3`,`ESTATUS_ID`,`ELIMINADO_ID`,`USUARIO_CREACION`,`FECHA_CREACION`,`USUARIO_MODIFICO`,`FECHA_MODIFICO`,`CONCEPTO_GENERICO_id`,`REGIMEN_FISCAL`,`RFC`,`IMPUESTO`,`CONCEPTO_GENERICO_TEXTO`,`CODIGO_POSTAL`,`CODIGO_UNITARIO`) values (1,'EVIL PANDA','https://raw.githubusercontent.com/enjeck/libre-logos/9370c927ee9c50de5234cdcb16afa35bb3eba6fc/src/images/logos/24-abstract-box-hexagon-arrows.png','https://placekitten.com/501/501','ACERCA DE','#','NOTICIAS','#','CONTACTO','#','evilpanda','#4B0082','#EBEBEB','CONOCENOS','AQUI VA UN TEXTO DESCRIPTIVO PERO NO SE ME OCURREN COSAS QUE PONER','ENLACE 1','#','ENLACE 2','#','ENLACE 3','#',1,3,1,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(2,'empty','https://www.sangabriel.mx/images/logo-sangabriel-2.png','https://placekitten.com/501/501','FECHAS','#','PRODUCTOS','#','INFORMACION','#','empty','','','NOSOTROS','NUESTRA EMPRESA ES UNA MUY GRANDE CON MUCHO TIEMPO DE EXPERIENCIA EN EL RAMO','RED 1','#','RED 2','#','RED 3','#',1,3,1,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(3,'INGRID XODAR JIMENEZ','https://www.sangabriel.mx/images/logo-sangabriel-2.png','https://www.sangabriel.mx/images/admin/categorias/grandes/625.jpg','FECHAS','#','PRODUCTOS','#','INFORMACION','#','sangabriel','','','NOSOTROS','NUESTRA EMPRESA ES UNA MUY GRANDE CON MUCHO TIEMPO DE EXPERIENCIA EN EL RAMO','RED 1','#','RED 2','#','RED 3','#',1,3,1,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',43232406,'626','XOJI740919U48',NULL,'prueba','66450','H87'),(4,'INGRID XODAR JIMENEZ','https://novasys.com.mx/portal/img/custom/logo_ns.png','https://placekitten.com/501/501','ACERCA DE','#','NOTICIAS','#','CONTACTO','#','novasys','#00a1e0','#EBEBEB','CONOCENOS','AQUI VA UN TEXTO DESCRIPTIVO PERO NO SE ME OCURREN COSAS QUE PONER','ENLACE 1','#','ENLACE 2','#','ENLACE 3','#',1,3,1,'2023-10-25 18:32:13',1,'2023-10-25 18:32:13',43232406,'626','XOJI740919U48',NULL,'Software de pruebas de programas','64030','H87');

/*Table structure for table `pagina` */

DROP TABLE IF EXISTS `pagina`;

CREATE TABLE `pagina` (
  `PAGINA_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Id de la página',
  `NOMBRE` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Nombre de la página como aparecerá en el menú',
  `ANIDAR` bigint(20) DEFAULT NULL COMMENT 'Id de otra página donde se anida ésta página',
  `URL` varchar(250) COLLATE utf8_spanish_ci DEFAULT NULL COMMENT 'Url de la página',
  `ICONO` char(250) COLLATE utf8_spanish_ci DEFAULT NULL,
  `ORDEN` bigint(20) NOT NULL COMMENT 'Orden de la página',
  `ESTATUS_ID` bigint(20) NOT NULL COMMENT 'Estatus',
  `ELIMINADO_ID` bigint(20) NOT NULL COMMENT 'Eliminado',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'Fecha en que fue agregada la información por primera vez',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'Usuario que genera o agrega por primera vez informacion en la tabla',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'Fecha en la que fueron realizados modificaciones a la información',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'Usuario que realizó modificaciones',
  PRIMARY KEY (`PAGINA_ID`),
  KEY `PAGINA_CATALOGOPARAMETRO_ELIMINADOID_FK` (`ELIMINADO_ID`),
  KEY `PAGINA_CATALOGOPARAMETRO_ESTATUSID_FK` (`ESTATUS_ID`),
  KEY `PAGINA_USUARIO_USUARIOCREACION_FK` (`USUARIO_CREACION`),
  KEY `PAGINA_USUARIO_USUARIOMODIFICO_FK` (`USUARIO_MODIFICO`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci COMMENT='Tabla de paginas';

/*Data for the table `pagina` */

insert  into `pagina`(`PAGINA_ID`,`NOMBRE`,`ANIDAR`,`URL`,`ICONO`,`ORDEN`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`) values (1,'Dashboard Admin Sistema',0,'/cfdi/admin/dashboard.action','dashboard',1,1,3,'2017-10-18 00:00:00',1,'2018-01-12 13:07:39',1),(2,'Sistema',0,'','done',4,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(3,'Aplicacion Web',0,'','done',2,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(4,'Catalogos Cliente',0,'','done',3,1,3,'2017-10-18 00:00:00',1,'2018-01-12 13:09:06',1),(5,'Usuarios',3,'/cfdi/admin/usuario.action','person',5,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(6,'Pagina',2,'/cfdi/admin/pagina.action','web_asset',6,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(7,'Perfil',2,'/cfdi/admin/perfil.action','rule',7,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(8,'Video Temporada',4,'/cfdi/admin/videoTemporada.action','videocam',8,1,4,'2017-10-18 00:00:00',1,'2023-07-26 13:49:14',1);

/*Table structure for table `perfil` */

DROP TABLE IF EXISTS `perfil`;

CREATE TABLE `perfil` (
  `PERFIL_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Numero identificador de Usuarios',
  `NOMBRE` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Nombre real',
  `ESTATUS_ID` bigint(20) NOT NULL COMMENT 'Estatus',
  `ELIMINADO_ID` bigint(20) NOT NULL COMMENT 'Eliminado',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'Fecha en que fue agregada la información por primera vez',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'Usuario que genera o agrega por primera vez informacion en la tabla',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'Fecha en la que fueron realizados modificaciones a la información',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'Usuario que realizó modificaciones',
  `PANTALLA_INICIO` char(250) COLLATE utf8_spanish_ci DEFAULT NULL,
  PRIMARY KEY (`PERFIL_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

/*Data for the table `perfil` */

insert  into `perfil`(`PERFIL_ID`,`NOMBRE`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`,`PANTALLA_INICIO`) values (1,'Sistemas',1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1,'/cfdi/admin/dashboard.action');

/*Table structure for table `perfil_pagina` */

DROP TABLE IF EXISTS `perfil_pagina`;

CREATE TABLE `perfil_pagina` (
  `PERFIL_PAGINA_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Id de Usuario-Página',
  `PERFIL_ID` bigint(20) NOT NULL COMMENT 'Id de Usuario',
  `PAGINA_ID` bigint(20) NOT NULL,
  PRIMARY KEY (`PERFIL_PAGINA_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci;

/*Data for the table `perfil_pagina` */

insert  into `perfil_pagina`(`PERFIL_PAGINA_ID`,`PERFIL_ID`,`PAGINA_ID`) values (1,1,1),(2,1,2),(3,1,3),(4,1,4),(5,1,5),(6,1,6),(7,1,7),(8,1,8);

/*Table structure for table `receptor` */

DROP TABLE IF EXISTS `receptor`;

CREATE TABLE `receptor` (
  `RECEPTOR_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'LLAVE PRIMARIA',
  `RAZON_SOCIAL` varchar(200) DEFAULT NULL COMMENT 'RAZON SOCIAL DEL RECEPTOR',
  `RFC` varchar(20) DEFAULT NULL COMMENT 'RFC DEL RECEPTOR',
  `CORREO` varchar(100) DEFAULT NULL COMMENT 'CORREO DEL RECEPTOR',
  `REGIMEN_FISCAL` bigint(5) DEFAULT NULL COMMENT 'REGIMEN FISCAL DEL RECEPTOR',
  `USO_CFDI` varchar(5) DEFAULT NULL COMMENT 'USO DEL CFDI DEL RECEPTOR',
  `TIPO_VENTA` varchar(5) DEFAULT NULL COMMENT 'TIPO DE VENTA',
  `CALLE` varchar(100) DEFAULT NULL COMMENT 'CALLE DEL RECEPTOR',
  `NUMERO_EXTERIOR` mediumint(20) DEFAULT NULL COMMENT 'NUMERO ESXTERIOR DEL RECEPTOR',
  `NUMERO_INTERIOR` mediumint(20) DEFAULT NULL COMMENT 'NUMERO INTERIOR DEL RECEPTOR',
  `COLONIA` varchar(100) DEFAULT NULL COMMENT 'COLONIA DEL RECEPTOR',
  `CODIGO_POSTAL` varchar(5) DEFAULT NULL COMMENT 'CODIGO POSTAL DEL RECEPTOR',
  `MUNICIPIO` varchar(100) DEFAULT NULL COMMENT 'MUNICIPIO DEL RECEPTOR',
  `ESTADO` varchar(100) DEFAULT NULL COMMENT 'ESTADO DEL RECEPTOR',
  `PAIS` varchar(100) DEFAULT NULL COMMENT 'PAIS DEL RECEPTOR',
  `ESTATUS_ID` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `ELIMINADO_ID` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'CAMPOS DE CONTROL',
  PRIMARY KEY (`RECEPTOR_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;

/*Data for the table `receptor` */

insert  into `receptor`(`RECEPTOR_ID`,`RAZON_SOCIAL`,`RFC`,`CORREO`,`REGIMEN_FISCAL`,`USO_CFDI`,`TIPO_VENTA`,`CALLE`,`NUMERO_EXTERIOR`,`NUMERO_INTERIOR`,`COLONIA`,`CODIGO_POSTAL`,`MUNICIPIO`,`ESTADO`,`PAIS`,`ESTATUS_ID`,`ELIMINADO_ID`,`USUARIO_CREACION`,`FECHA_CREACION`,`USUARIO_MODIFICO`,`FECHA_MODIFICO`) values (10,'ARCELORMITTAL TUBULAR PRODUCTS SA DE CV','ARTP20018392L','JUAN.HERNANDEZ@NOVASYS.COM.MX',601,'P01','26','ATOTONILCO',806,NULL,'CHEPEVERA','66450','039','NLE','MEX',1,3,0,'2023-11-07 18:50:49',0,'2023-11-08 19:01:05'),(16,'ARCELORMITTAL TUBULAR PRODUCTS SA DE CV','ARTP20018392R','JUAN.HERNANDEZ@NOVASYS.COM.MX',601,'P01','26','ATOTONILCO',806,NULL,'CHEPEVERA','66450','039','NLE','MEX',1,3,0,'2023-11-08 11:52:12',0,'2023-11-08 11:52:12'),(17,'XOCHILT CASAS CHAVEZ','CACX7605101P8','JUAN.HERNANDEZ@NOVASYS.COM.MX',605,'CP01','29','ATOTONILCO',806,NULL,'CHEPEVERA','36257','015','GUA','MEX',1,3,0,'2023-11-08 19:02:57',0,'2023-11-09 16:26:54'),(18,'JORGE ESTRADA LIMON','EALJ760205CN6','JORGE.ESTRADA1976@GMAIL.COM',612,'G03','01','DELTA ',200,NULL,'DOS','67134','026','NLE','MEX',1,3,0,'2023-11-09 15:28:51',0,'2023-11-09 15:28:51');

/*Table structure for table `ticket` */

DROP TABLE IF EXISTS `ticket`;

CREATE TABLE `ticket` (
  `TICKET_ID` bigint(100) AUTO_INCREMENT NOT NULL,
  `NEGOCIO_ID` bigint(10) NOT NULL,
  `FOLIO` varchar(20) NOT NULL,
  `FECHA` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `TOTAL` float NOT NULL,
  `ESTATUS_ID` bigint(2) NOT NULL,
  `ELIMINADO_ID` bigint(2) NOT NULL,
  `FECHA_CREACION` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `USUARIO_CREACION` bigint(20) NOT NULL,
  `FECHA_MODIFICO` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `USUARIO_MODIFICO` bigint(20) NOT NULL,
  `SUCURSAL_ID` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`TICKET_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `ticket` */

insert  into `ticket`(`TICKET_ID`,`NEGOCIO_ID`,`FOLIO`,`FECHA`,`TOTAL`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`,`SUCURSAL_ID`) values (1,4,'5826','2023-11-09 13:57:02',99,1,3,'2023-10-19 15:46:57',1,'2023-10-19 15:46:57',1,1),(2,3,'1521','2023-11-09 15:31:43',99,1,3,'2023-10-19 15:46:57',1,'2023-10-19 15:46:57',1,2);

/*Table structure for table `usuario` */

DROP TABLE IF EXISTS `usuario`;

CREATE TABLE `usuario` (
  `USUARIO_ID` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Numero identificador de Usuarios',
  `USUARIO` char(50) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Nombre de usuario',
  `CONTRASENA` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Contraseña',
  `NOMBRE` char(250) COLLATE utf8_spanish_ci DEFAULT NULL COMMENT 'Nombre real',
  `APELLIDO_PATERNO` char(250) COLLATE utf8_spanish_ci DEFAULT NULL COMMENT 'Apellido paterno',
  `APELLIDO_MATERNO` char(250) COLLATE utf8_spanish_ci DEFAULT NULL COMMENT 'Apellido materno',
  `CORREO` char(250) COLLATE utf8_spanish_ci NOT NULL COMMENT 'Correo electrónico',
  `PERFIL_ID` bigint(20) NOT NULL COMMENT 'Id de Usuario-Página',
  `ESTATUS_ID` bigint(20) NOT NULL COMMENT 'Estatus',
  `ELIMINADO_ID` bigint(20) NOT NULL COMMENT 'Eliminado',
  `FECHA_CREACION` datetime DEFAULT NULL COMMENT 'Fecha en que fue agregada la información por primera vez',
  `USUARIO_CREACION` bigint(20) DEFAULT NULL COMMENT 'Usuario que genera o agrega por primera vez informacion en la tabla',
  `FECHA_MODIFICO` datetime DEFAULT NULL COMMENT 'Fecha en la que fueron realizados modificaciones a la información',
  `USUARIO_MODIFICO` bigint(20) DEFAULT NULL COMMENT 'Usuario que realizó modificaciones',
  PRIMARY KEY (`USUARIO_ID`),
  KEY `FK_REFERENCE_41` (`PERFIL_ID`),
  KEY `USUARIO_CATALOGOPARAMETRO_ELIMINADOID_FK` (`ELIMINADO_ID`),
  KEY `USUARIO_CATALOGOPARAMETRO_ESTATUSID_FK` (`ESTATUS_ID`),
  CONSTRAINT `FK_REFERENCE_41` FOREIGN KEY (`PERFIL_ID`) REFERENCES `perfil` (`PERFIL_ID`),
  CONSTRAINT `USUARIO_CATALOGOPARAMETRO_ELIMINADOID_FK` FOREIGN KEY (`ELIMINADO_ID`) REFERENCES `catalogo_parametro` (`CATALOGO_PARAMETRO_ID`),
  CONSTRAINT `USUARIO_CATALOGOPARAMETRO_ESTATUSID_FK` FOREIGN KEY (`ESTATUS_ID`) REFERENCES `catalogo_parametro` (`CATALOGO_PARAMETRO_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_spanish_ci COMMENT='Tabla de usuarios';

/*Data for the table `usuario` */

insert  into `usuario`(`USUARIO_ID`,`USUARIO`,`CONTRASENA`,`NOMBRE`,`APELLIDO_PATERNO`,`APELLIDO_MATERNO`,`CORREO`,`PERFIL_ID`,`ESTATUS_ID`,`ELIMINADO_ID`,`FECHA_CREACION`,`USUARIO_CREACION`,`FECHA_MODIFICO`,`USUARIO_MODIFICO`) values (1,'admin','184006DEC458997AD56AF53CF67F4B6D','Administrador','Administrador','Administrador','juan.estrada@novasys.com.mx',1,1,3,'2017-10-18 00:00:00',1,'2017-10-18 00:00:00',1),(6,'dummy','139F218227F0312D1C2DB5F26FC9FC23','juan','zavala','hernandez','juan@heza.com',1,1,3,'2023-10-19 15:46:57',0,'2023-10-19 15:46:57',0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
