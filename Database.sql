DROP DATABASE IF EXISTS DEADLINETRACKER;
CREATE DATABASE DEADLINETRACKER DEFAULT CHARACTER SET utf8 COLLATE utf8_hungarian_ci;

USE DEADLINETRACKER;

CREATE TABLE SUBJECTS(
	`SubjectKey` INT AUTO_INCREMENT,
	`SubjectName` VARCHAR(255),
	PRIMARY KEY(`SubjectKey`)
);

CREATE TABLE DEADLINETYPES(
	`DeadlineTypeId` INT AUTO_INCREMENT,
	`DeadlineTypeName` VARCHAR(255),
	PRIMARY KEY(`DeadlineTypeId`)
);

INSERT INTO `deadlinetypes` VALUES (1, 'Röp ZH');
INSERT INTO `deadlinetypes` VALUES (2, 'Nagy Zárthelyi');
INSERT INTO `deadlinetypes` VALUES (3, 'Beadandó');
INSERT INTO `deadlinetypes` VALUES (4, 'Vizsga');
INSERT INTO `deadlinetypes` VALUES (5, 'Házi feladat');
INSERT INTO `deadlinetypes` VALUES (6, 'Javító Zárthelyi');
INSERT INTO `deadlinetypes` VALUES (7, 'Gyakorlati utóvizsga');


INSERT INTO `subjects` VALUES (1, 'Matematika alapok (Gyakorlat)');
INSERT INTO `subjects` VALUES (2, 'Számítógépes rendszerek (Előadás)');
INSERT INTO `subjects` VALUES (3, 'Számítógépes rendszerek (Gyakorlat)');
INSERT INTO `subjects` VALUES (4, 'Imperatív programozás (Előadás)');
INSERT INTO `subjects` VALUES (5, 'Imperatív programozás (Gyakorlat)');
INSERT INTO `subjects` VALUES (6, 'Programozás alapok (Előadás)');
INSERT INTO `subjects` VALUES (7, 'Programozás alapok (Gyakorlat)');
INSERT INTO `subjects` VALUES (8, 'Funkcionális programozás (Előadás)');
INSERT INTO `subjects` VALUES (9, 'Funkcionális programozás (Gyakorlat)');
INSERT INTO `subjects` VALUES (10, 'Jogi ismeretek (Előadás)');


CREATE TABLE DEADLINES(
	`Id` INT AUTO_INCREMENT,
	`SubjectId` INT NOT NULL,
	`Deadline` DATETIME NOT NULL,
    `TypeId` INT NOT NULL,
    `Topic` VARCHAR(255) DEFAULT "",
    `Comments` VARCHAR(1000) DEFAULT "",
	PRIMARY KEY(`Id`)
);

INSERT INTO `deadlines` VALUES (1, 1, '2022-9-19 08:00:00', 1, 'Algebrai kifejezések, nevezetes azonosságok', 'komment');
INSERT INTO `deadlines` VALUES (2, 1, '2022-10-14 16:00:00', 2, '1 - 10. fejezet', 'Jelenléti, papíros');
INSERT INTO `deadlines` VALUES (3, 1, '2022-11-18 16:00:00', 2, '11 - 18. fejezet', 'Jelenléti, papíros');
INSERT INTO `deadlines` VALUES (4, 1, '2022-12-19 16:00:00', 2, '19 - 26. fejezet', 'Formája később');
INSERT INTO `deadlines` VALUES (5, 1, '2023-1-3 14:00:00', 6, 'Követelményrendszer szerint', 'Javító zh, formája később');
INSERT INTO `deadlines` VALUES (6, 1, '2023-1-13 14:00:00', 7, 'Követelményrendszer szerint', 'Gyakorlati utóvizsga, formája később');

CREATE TABLE SUGGESTIONS(
	`SuggestionId` INT NOT NULL,
	`SuggestionSubjectId` INT NOT NULL,
	`SuggestionDeadline` DATETIME NOT NULL,
    `SuggestionTypeId` INT NOT NULL,
    `SuggestionTopic` VARCHAR(255) DEFAULT "",
    `SuggestionComments` VARCHAR(1000) DEFAULT "",
	PRIMARY KEY(`SuggestionId`)
);

CREATE TABLE USERS(
	`Token` VARCHAR(10) NOT NULL,
	`Username` VARCHAR(255) NOT NULL,
	`Password` VARCHAR(255) NOT NULL,
	`Created` DATETIME DEFAULT NOW()
);