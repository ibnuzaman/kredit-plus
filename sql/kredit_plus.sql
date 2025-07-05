-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: mysql-80
-- Generation Time: Jul 05, 2025 at 02:51 PM
-- Server version: 8.0.41
-- PHP Version: 8.2.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kredit_plus`
--

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint UNSIGNED NOT NULL,
  `nik` varchar(16) NOT NULL,
  `full_name` varchar(100) NOT NULL,
  `email` varchar(128) NOT NULL,
  `password` varchar(72) NOT NULL,
  `legal_name` varchar(100) NOT NULL,
  `place_birth` varchar(32) NOT NULL,
  `date_birth` date NOT NULL,
  `salary` decimal(12,2) NOT NULL,
  `identity_file` varchar(64) NOT NULL,
  `selfie_file` varchar(64) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `nik`, `full_name`, `email`, `password`, `legal_name`, `place_birth`, `date_birth`, `salary`, `identity_file`, `selfie_file`, `created_at`, `updated_at`) VALUES
(1, '3510010101900001', 'Budi Santoso', 'budi@gmail.com', '$2a$12$8rHI4GvSgCr8LDn9c6Y/kO/71biGspofK8Zeyh/jDZsevMJRnB/Jy', 'Budi Santoso', 'Jakarta', '1990-01-01', 5000000.00, '9c35760e-4d3e-45af-9fdc-f4fdd9cf3185.jpg', 'ca7022c1-685d-4903-94f6-1f469c135f72.jpg', '2025-07-05 21:51:01.327', '2025-07-05 21:51:01.327'),
(2, '3510010202920002', 'Annisa Fitriani', 'annisa@gmail.com', '$2a$12$8rHI4GvSgCr8LDn9c6Y/kO/71biGspofK8Zeyh/jDZsevMJRnB/Jy', 'Annisa Fitriani', 'Surabaya', '1992-02-02', 15000000.00, 'f19efd94-484d-4507-b961-19c8d354d7ef.jpg', '220ab59f-b476-4064-9227-234e36842284.jpg', '2025-07-05 21:51:01.340', '2025-07-05 21:51:01.340');

-- --------------------------------------------------------

--
-- Table structure for table `loans`
--

CREATE TABLE `loans` (
  `id` bigint UNSIGNED NOT NULL,
  `customer_id` bigint UNSIGNED NOT NULL,
  `otr` decimal(12,2) NOT NULL,
  `admin_fee` decimal(12,2) NOT NULL,
  `installment_amount` decimal(12,2) NOT NULL,
  `assets_name` varchar(32) NOT NULL,
  `tenor_months` tinyint UNSIGNED NOT NULL,
  `total_paid` bigint NOT NULL DEFAULT '0',
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tenors`
--

CREATE TABLE `tenors` (
  `id` bigint UNSIGNED NOT NULL,
  `customer_id` bigint UNSIGNED NOT NULL,
  `month` tinyint UNSIGNED NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `tenors`
--

INSERT INTO `tenors` (`id`, `customer_id`, `month`, `amount`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 100000.00, '2025-07-05 21:51:01.332', '2025-07-05 21:51:01.332'),
(2, 1, 2, 200000.00, '2025-07-05 21:51:01.334', '2025-07-05 21:51:01.334'),
(3, 1, 3, 500000.00, '2025-07-05 21:51:01.336', '2025-07-05 21:51:01.336'),
(4, 1, 6, 700000.00, '2025-07-05 21:51:01.338', '2025-07-05 21:51:01.338'),
(5, 2, 1, 1000000.00, '2025-07-05 21:51:01.342', '2025-07-05 21:51:01.342'),
(6, 2, 2, 1200000.00, '2025-07-05 21:51:01.345', '2025-07-05 21:51:01.345'),
(7, 2, 3, 1500000.00, '2025-07-05 21:51:01.347', '2025-07-05 21:51:01.347'),
(8, 2, 6, 2000000.00, '2025-07-05 21:51:01.348', '2025-07-05 21:51:01.348');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint UNSIGNED NOT NULL,
  `loan_id` bigint UNSIGNED NOT NULL,
  `customer_id` bigint UNSIGNED NOT NULL,
  `amount` decimal(12,2) NOT NULL,
  `interest_amount` decimal(12,2) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_customers_nik` (`nik`),
  ADD UNIQUE KEY `idx_customers_email` (`email`);

--
-- Indexes for table `loans`
--
ALTER TABLE `loans`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_customers_loans` (`customer_id`);

--
-- Indexes for table `tenors`
--
ALTER TABLE `tenors`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_tenors_customer_id` (`customer_id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_loans_transactions` (`loan_id`),
  ADD KEY `fk_customers_transactions` (`customer_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `loans`
--
ALTER TABLE `loans`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `tenors`
--
ALTER TABLE `tenors`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `loans`
--
ALTER TABLE `loans`
  ADD CONSTRAINT `fk_customers_loans` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`);

--
-- Constraints for table `tenors`
--
ALTER TABLE `tenors`
  ADD CONSTRAINT `fk_customers_tenors` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_customers_transactions` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
  ADD CONSTRAINT `fk_loans_transactions` FOREIGN KEY (`loan_id`) REFERENCES `loans` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
