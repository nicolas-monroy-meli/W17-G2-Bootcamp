USE `fescos_db`;

DROP TABLE IF EXISTS `buyers`;

CREATE TABLE `buyers` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `id_card_number` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    `first_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    `last_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `buyer_card_number_unique` (`id_card_number`)
) ENGINE=InnoDB AUTO_INCREMENT=149 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `purchases_orders`;

CREATE TABLE `purchase_orders` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `order_number` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    `order_date` timestamp NULL DEFAULT NULL,
    `tracking_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    `buyer_id` int(10) unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `purcharse_order_number_unique` (`order_number`),
    KEY `purchase_orders_buyer_id_foreign` (`buyer_id`),
    CONSTRAINT `purchase_orders_buyer_id_foreign` FOREIGN KEY (`buyer_id`) REFERENCES `buyers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


DROP TABLE IF EXISTS `order_details`;

CREATE TABLE `order_details` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `clean_liness_status` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    `quantity` int unsigned NULL DEFAULT NULL,
    `temperature` decimal unsigned NULL DEFAULT NULL,
    `product_record_id` int(10) unsigned DEFAULT NULL,
    `purchase_order_id` int(10) unsigned DEFAULT NULL,
    PRIMARY KEY (`id`),
    #KEY `order_details_product_record_id_foreign` (`product_record_id`),
    #CONSTRAINT `order_details_product_record_id_foreign` FOREIGN KEY (`product_record_id`) REFERENCES `product_records` (`id`)
    KEY `order_details_purchase_order_id_foreign` (`purchase_order_id`),
    CONSTRAINT `order_details_purchase_order_id_foreign` FOREIGN KEY (`purchase_order_id`) REFERENCES `purchase_orders` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;