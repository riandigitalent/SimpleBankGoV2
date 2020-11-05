-- phpMyAdmin SQL Dump
-- version 4.0.4.2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Nov 05, 2020 at 12:29 PM
-- Server version: 5.6.13
-- PHP Version: 5.4.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `simplebankdb`
--
CREATE DATABASE IF NOT EXISTS `simplebankdb` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `simplebankdb`;

-- --------------------------------------------------------

--
-- Table structure for table `accounts`
--

CREATE TABLE IF NOT EXISTS `accounts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `id_account` longtext,
  `name` longtext,
  `password` longtext,
  `account_number` bigint(20) DEFAULT NULL,
  `saldo` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=3 ;

--
-- Dumping data for table `accounts`
--

INSERT INTO `accounts` (`id`, `id_account`, `name`, `password`, `account_number`, `saldo`) VALUES
(1, 'id-447', 'rian', '$2a$10$/RfaVmsdjXUHXHlltoGu/.G2VSR6tP5v0cm5RqmNyi7OSLamT5D0i', 98511093, 1980000),
(2, 'id-2799', 'rani', '$2a$10$ehLPBIBacFHZNBsFVYN0OeA39kHuhEtYHJbVz10NBpCvC49hc21Zu', 27146861, 8000000);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE IF NOT EXISTS `transactions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `transaction_type` bigint(20) DEFAULT NULL,
  `transaction_description` longtext,
  `sender` bigint(20) DEFAULT NULL,
  `amount` bigint(20) DEFAULT NULL,
  `recipient` bigint(20) DEFAULT NULL,
  `timestamp` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=9 ;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `transaction_type`, `transaction_description`, `sender`, `amount`, `recipient`, `timestamp`) VALUES
(1, 2, 'gajian', 98511093, 9000000, 98511093, 1604578799),
(2, 2, 'lemburan', 98511093, 1000000, 0, 1604579022),
(3, 2, 'uang makan', 98511093, 980000, 0, 1604579075),
(4, 0, 'nafkah', 98511093, 8000000, 27146861, 1604579169),
(5, 2, 'tarik tunai', 98511093, 500000, 0, 1604579271),
(6, 1, 'tarik tunai', 98511093, 500000, 0, 1604579306),
(7, 1, 'tarik tunai2', 98511093, 500000, 0, 1604579312),
(8, 1, 'tarik tunai3', 98511093, 500000, 0, 1604579325);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
