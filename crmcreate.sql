CREATE TABLE IF NOT EXISTS `users`(
  `email` varchar(255) NOT NULL,
  `agentid` varchar(25) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `agent`(
  `firstname` varchar(255) NOT NULL,
  `secondname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phoneno` varchar(255) NOT NULL,
  `agentid` varchar(25) NOT NULL,
  `active` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `appointments`(
  `agentid` varchar(25) NOT NULL,
  `appointmentid` varchar(25) NOT NULL,
  `title` varchar(25) NOT NULL,
  `description` varchar(255) NOT NULL,
  `done` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `deals`(
  `contactid` varchar(25) NOT NULL,
  `agentid` varchar(25) NOT NULL,
  `dealid` varchar(25) NOT NULL,
  `description` varchar(255) NOT NULL,
  `complete` BOOLEAN,
  `approoved` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `todo`(
  `agentid` varchar(25) NOT NULL,
  `todoid` varchar(25) NOT NULL,
  `description` varchar(255) NOT NULL,
  `complete` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `contact`(
  `creatorid` varchar(25) NOT NULL,
  `contactid` varchar(25) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phoneno` varchar(255) NOT NULL,
  `dateoflastcontact` varchar(255) NOT NULL,
  `typeofcontact` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `company` varchar(255) NOT NULL,
  `notes` varchar(255) NOT NULL,
  `professionaltitle` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `zipcode` INT(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `users` (
  `agentid` varchar(25) NOT NULL,
  `email` varchar(255) NOT NULL,
  `passwword` varchar(255) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `quotation`(
  `quoteid` VARCHAR(255) NOT NULL,
  `ownerid` VARCHAR(255) NOT NULL,
  `productid` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) not NULL,
  `email` VARCHAR(255) NOT NULL,
  `phonenumber` VARCHAR(255) NOT NULL,
  `quotationrequest` VARCHAR(255) NOT NULL,
  `viewed` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `enquiry`(
  `name` VARCHAR(255) not NULL,
  `email` VARCHAR(255) NOT NULL,
  `subject` VARCHAR(255) not NULL,
  `message` VARCHAR(255) NOT NULL,
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `contactme`(
  `ownerid` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) not NULL,
  `email` VARCHAR(255) NOT NULL,
  `phonenumber` VARCHAR(255) NOT NULL,
  `viewed` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `feedback`(
  `ownerid` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) not NULL,
  `email` VARCHAR(255) NOT NULL,
  `phonenumber` VARCHAR(255) NOT NULL,
  `feedback` VARCHAR(255) NOT NULL,
  `viewed` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `news` (
  `title` varchar(25) NOT NULL,
  `to` varchar(255) NOT NULL,
  `from` varchar(25) NOT NULL,
  `data` varchar(255) NOT NULL,
  `viewed` BOOLEAN,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `category`(
  `name` VARCHAR(255) NOT NULL,
  `productid` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `products`(
  `name` VARCHAR(255) NOT NULL,
  `productid` VARCHAR(255) NOT NULL,
  `categoryid` VARCHAR(255) NOT NULL,
  `number` DECIMAL(19,4) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `amount` INT(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
