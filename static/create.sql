CREATE TABLE IF NOT EXISTS `user` (
  `account` varchar(40) NOT NULL ,
  `name` varchar(40),
  `city` varchar(40),
  `address` varchar(40),
  `phone_call` varchar(40),
  `mobile` varchar(40),
  `email` varchar(40),
  `job` varchar(40),
  `contact` varchar(40),
  `contact_phone` varchar(40),
  `income` int(32),
  `credit_rating` varchar(40),
  PRIMARY KEY (`account`)  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `balance` (
  `_id` int(32) NOT NULL AUTO_INCREMENT,
  `account` varchar(40),
  `amount` int(32),
  `interest` int(32),
  CONSTRAINT FOREIGN KEY (`account`) REFERENCES `user`(`account`)
  ON DELETE  RESTRICT  ON UPDATE CASCADE,
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `savingRecord` (
  `_id` int(32) NOT NULL AUTO_INCREMENT,    
  `account` varchar(40),
  `name` varchar(40),
  `bank` varchar(40),
  `amount` int(32),
  `save_time` varchar(40),
  `save_type` varchar(40),
  CONSTRAINT FOREIGN KEY (`account`) REFERENCES `user`(`account`)
  ON DELETE  RESTRICT  ON UPDATE CASCADE,
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `withdrawRecord` (
  `_id` int(32) NOT NULL AUTO_INCREMENT,    
  `account` varchar(40),
  `name` varchar(40),
  `bank` varchar(40),
  `amount` int(32),
  `save_time` varchar(40),
  `save_type` varchar(40),
  CONSTRAINT FOREIGN KEY (`account`) REFERENCES `user`(`account`)
  ON DELETE  RESTRICT  ON UPDATE CASCADE,
  PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



