CREATE TABLE `books` (
  `title` varchar(100) NOT NULL,
  `author` varchar(50) DEFAULT NULL,
  `publisher` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `books` VALUES ('123','1234','12345'),('title1','author1','publisher1'),('title2','author2','publisher2'),('title3','author3','publisher3'),('title4','author4','publisher4'),('title5','author5','publisher5');
