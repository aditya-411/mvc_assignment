CREATE TABLE `users` (
  `username` varchar(100) NOT NULL,
  `password` varchar(100) DEFAULT NULL,
  `isadmin` tinyint(1) DEFAULT NULL,
  `admin_request` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;