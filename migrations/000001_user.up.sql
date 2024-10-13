CREATE TABLE `user` (
  `user_id` varchar(40) ,
  `user_name` varchar(40) NOT NULL,
  `email` varchar(40) NOT NULL UNIQUE,
  `password` varchar(40) NOT NULL,
  PRIMARY KEY (user_id)
)