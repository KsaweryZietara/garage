INSERT INTO makes (id, name)
VALUES (1, 'Seat'),
       (2, 'Renault'),
       (3, 'Peugeot'),
       (4, 'Dacia'),
       (5, 'Citroën'),
       (6, 'Opel'),
       (7, 'Alfa Romeo'),
       (8, 'Škoda'),
       (9, 'Chevrolet'),
       (10, 'Porsche'),
       (11, 'Honda'),
       (12, 'Subaru'),
       (13, 'Mazda'),
       (14, 'Mitsubishi'),
       (15, 'Lexus'),
       (16, 'Toyota'),
       (17, 'BMW'),
       (18, 'Volkswagen'),
       (19, 'Suzuki'),
       (20, 'Mercedes-Benz'),
       (21, 'Saab'),
       (22, 'Audi'),
       (23, 'Kia'),
       (24, 'Land Rover'),
       (25, 'Dodge'),
       (26, 'Chrysler'),
       (27, 'Ford'),
       (28, 'Hummer'),
       (29, 'Hyundai'),
       (30, 'Infiniti'),
       (31, 'Jaguar'),
       (32, 'Jeep'),
       (33, 'Nissan'),
       (34, 'Volvo'),
       (35, 'Daewoo'),
       (36, 'Fiat'),
       (37, 'MINI'),
       (38, 'Rover'),
       (39, 'Smart')
ON CONFLICT (id) DO NOTHING;

INSERT INTO models (id, name, make_id)
VALUES
-- Models for Renault
(1, 'Alhambra', 1),
(2, 'Altea', 1),
(3, 'Altea XL', 1),
(4, 'Arosa', 1),
(5, 'Cordoba', 1),
(6, 'Cordoba Vario', 1),
(7, 'Exeo', 1),
(8, 'Ibiza', 1),
(9, 'Ibiza ST', 1),
(10, 'Exeo ST', 1),
(11, 'Leon', 1),
(12, 'Inca', 1),
(13, 'Mii', 1),
(14, 'Toledo', 1),

-- Models for Renault
(15, 'Captur', 2),
(16, 'Clio', 2),
(17, 'Clio Grandtour', 2),
(18, 'Espace', 2),
(19, 'Express', 2),
(20, 'Fluence', 2),
(21, 'Grand Espace', 2),
(22, 'Grand Modus', 2),
(23, 'Grand Scenic', 2),
(24, 'Kadjar', 2),
(25, 'Kangoo', 2),
(26, 'Kangoo Express', 2),
(27, 'Koleos', 2),
(28, 'Laguna', 2),
(29, 'Laguna Grandtour', 2),
(30, 'Latitude', 2),
(31, 'Mascott', 2),
(32, 'Mégane', 2),
(33, 'Mégane CC', 2),
(34, 'Mégane Combi', 2),
(35, 'Mégane Grandtour', 2),
(36, 'Mégane Coupé', 2),
(37, 'Mégane Scénic', 2),
(38, 'Scénic', 2),
(39, 'Talisman', 2),
(40, 'Talisman Grandtour', 2),
(41, 'Thalia', 2),
(42, 'Twingo', 2),
(43, 'Wind', 2),
(44, 'Zoé', 2),

-- Models for Peugeot
(45, '1007', 3),
(46, '107', 3),
(47, '106', 3),
(48, '108', 3),
(49, '2008', 3),
(50, '205', 3),
(51, '206', 3),
(52, '206 CC', 3),
(53, '206 SW', 3),
(54, '207', 3),
(55, '207 CC', 3),
(56, '207 SW', 3),
(57, '306', 3),
(58, '307', 3),
(59, '307 CC', 3),
(60, '307 SW', 3),
(61, '308', 3),
(62, '308 CC', 3),
(63, '308 SW', 3),
(64, '406', 3),
(65, '407', 3),
(66, '407 SW', 3),
(67, '5008', 3),
(68, '508', 3),
(69, '508 SW', 3),
(70, '607', 3),
(71, '806', 3),
(72, '807', 3),
(73, 'Bipper', 3),
(74, 'Expert', 3),
(75, 'iOn', 3),

-- Models for Dacia
(76, 'Dokker', 4),
(77, 'Duster', 4),
(78, 'Lodgy', 4),
(79, 'Logan', 4),
(80, 'Logan MCV', 4),
(81, 'Sandero', 4),

-- Models for Citroën
(82, 'Berlingo', 5),
(83, 'C-Crosser', 5),
(84, 'C-Elissée', 5),
(85, 'C-Zero', 5),
(86, 'C1', 5),
(87, 'C2', 5),
(88, 'C3', 5),
(89, 'C3 Picasso', 5),
(90, 'C4', 5),
(91, 'C4 Aircross', 5),
(92, 'C4 Cactus', 5),
(93, 'C4 Coupé', 5),
(94, 'C4 Grand Picasso', 5),
(95, 'C4 Picasso', 5),
(96, 'C5', 5),
(97, 'C5 Break', 5),
(98, 'C5 Tourer', 5),
(99, 'C6', 5),
(100, 'C8', 5),
(101, 'DS3', 5),
(102, 'DS4', 5),
(103, 'DS5', 5),
(104, 'Evasion', 5),
(105, 'Jumper', 5),
(106, 'Jumpy', 5),
(107, 'Saxo', 5),
(108, 'Nemo', 5),

-- Models for Opel
(109, 'Agila', 6),
(110, 'Ampera', 6),
(111, 'Antara', 6),
(112, 'Astra', 6),
(113, 'Astra Cabrio', 6),
(114, 'Astra Caravan', 6),
(115, 'Astra Coupé', 6),
(116, 'Astra GTC', 6),
(117, 'Astra Sports Tourer', 6),
(118, 'Astra TwinTop', 6),
(119, 'Combo', 6),
(120, 'Corsa', 6),
(121, 'Frontera', 6),
(122, 'Insignia', 6),
(123, 'Insignia Sports Tourer', 6),
(124, 'Meriva', 6),
(125, 'Mokka', 6),
(126, 'Movano', 6),
(127, 'Omega', 6),
(128, 'Signum', 6),
(129, 'Vectra', 6),
(130, 'Vectra Caravan', 6),
(131, 'Vivaro', 6),
(132, 'Zafira', 6),
(133, 'Zafira Tourer', 6),

-- Models for Alfa Romeo
(134, '145', 7),
(135, '146', 7),
(136, '147', 7),
(137, '155', 7),
(138, '156', 7),
(139, '156 Sportwagon', 7),
(140, '159', 7),
(141, '159 Sportwagon', 7),
(142, '164', 7),
(143, '166', 7),
(144, '4C', 7),
(145, 'Brera', 7),
(146, 'GTV', 7),
(147, 'MiTo', 7),
(148, 'Crosswagon', 7),
(149, 'Spider', 7),
(150, 'GT', 7),
(151, 'Giulietta', 7),
(152, 'Giulia', 7),

-- Models for Škoda
(153, 'Favorit', 8),
(154, 'Felicia', 8),
(155, 'Citigo', 8),
(156, 'Fabia', 8),
(157, 'Fabia Combi', 8),
(158, 'Fabia Sedan', 8),
(159, 'Felicia Combi', 8),
(160, 'Octavia', 8),
(161, 'Octavia Combi', 8),
(162, 'Roomster', 8),
(163, 'Yeti', 8),
(164, 'Rapid', 8),
(165, 'Rapid Spaceback', 8),
(166, 'Superb', 8),
(167, 'Superb Combi', 8),


-- Models for Chevrolet
(168, 'Alero', 9),
(169, 'Aveo', 9),
(170, 'Camaro', 9),
(171, 'Captiva', 9),
(172, 'Corvette', 9),
(173, 'Cruze', 9),
(174, 'Cruze SW', 9),
(175, 'Epica', 9),
(176, 'Equinox', 9),
(177, 'Evanda', 9),
(178, 'HHR', 9),
(179, 'Kalos', 9),
(180, 'Lacetti', 9),
(181, 'Lacetti SW', 9),
(182, 'Lumina', 9),
(183, 'Malibu', 9),
(184, 'Matiz', 9),
(185, 'Monte Carlo', 9),
(186, 'Nubira', 9),
(187, 'Orlando', 9),
(188, 'Spark', 9),
(189, 'Suburban', 9),
(190, 'Tacuma', 9),
(191, 'Tahoe', 9),
(192, 'Trax', 9),

-- Models for Porsche
(193, '911 Carrera', 10),
(194, '911 Carrera Cabrio', 10),
(195, '911 Targa', 10),
(196, '911 Turbo', 10),
(197, '924', 10),
(198, '944', 10),
(199, '997', 10),
(200, 'Boxster', 10),
(201, 'Cayenne', 10),
(202, 'Cayman', 10),
(203, 'Macan', 10),
(204, 'Panamera', 10),


-- Models for Honda
(205, 'Accord', 11),
(206, 'Accord Coupé', 11),
(207, 'Accord Tourer', 11),
(208, 'City', 11),
(209, 'Civic', 11),
(210, 'Civic Aerodeck', 11),
(211, 'Civic Coupé', 11),
(212, 'Civic Tourer', 11),
(213, 'Civic Type R', 11),
(214, 'CR-V', 11),
(215, 'CR-X', 11),
(216, 'CR-Z', 11),
(217, 'FR-V', 11),
(218, 'HR-V', 11),
(219, 'Insight', 11),
(220, 'Integra', 11),
(221, 'Jazz', 11),
(222, 'Legend', 11),
(223, 'Prelude', 11),

-- Models for Subaru
(224, 'BRZ', 12),
(225, 'Forester', 12),
(226, 'Impreza', 12),
(227, 'Impreza Wagon', 12),
(228, 'Justy', 12),
(229, 'Legacy', 12),
(230, 'Legacy Wagon', 12),
(231, 'Legacy Outback', 12),
(232, 'Levorg', 12),
(233, 'Outback', 12),
(234, 'SVX', 12),
(235, 'Tribeca', 12),
(236, 'Tribeca B9', 12),
(237, 'XV', 12),

-- Models for Mazda
(238, '121', 13),
(239, '2', 13),
(240, '3', 13),
(241, '323', 13),
(242, '323 Combi', 13),
(243, '323 Coupé', 13),
(244, '323 F', 13),
(245, '5', 13),
(246, '6', 13),
(247, '6 Combi', 13),
(248, '626', 13),
(249, '626 Combi', 13),
(250, 'B-Fighter', 13),
(251, 'B2500', 13),
(252, 'BT', 13),
(253, 'CX-3', 13),
(254, 'CX-5', 13),
(255, 'CX-7', 13),
(256, 'CX-9', 13),
(257, 'Demio', 13),
(258, 'MPV', 13),
(259, 'MX-3', 13),
(260, 'MX-5', 13),
(261, 'MX-6', 13),
(262, 'Premacy', 13),
(263, 'RX-7', 13),
(264, 'RX-8', 13),
(265, 'Xedox 6', 13),

-- Models for Mitsubishi
(266, '3000 GT', 14),
(267, 'ASX', 14),
(268, 'Carisma', 14),
(269, 'Colt', 14),
(270, 'Colt CC', 14),
(271, 'Eclipse', 14),
(272, 'Fuso canter', 14),
(273, 'Galant', 14),
(274, 'Galant Combi', 14),
(275, 'Grandis', 14),
(276, 'L200', 14),
(277, 'L200 Pick up', 14),
(278, 'L200 Pick up Allrad', 14),
(279, 'L300', 14),
(280, 'Lancer', 14),
(281, 'Lancer Combi', 14),
(282, 'Lancer Evo', 14),
(283, 'Lancer Sportback', 14),
(284, 'Outlander', 14),
(285, 'Pajero', 14),
(286, 'Pajeto Pinin', 14),
(287, 'Pajero Pinin Wagon', 14),
(288, 'Pajero Sport', 14),
(289, 'Pajero Wagon', 14),
(290, 'Space Star', 14),

-- Models for Lexus
(291, 'CT', 15),
(292, 'GS', 15),
(293, 'GS 300', 15),
(294, 'GX', 15),
(295, 'IS', 15),
(296, 'IS 200', 15),
(297, 'IS 250 C', 15),
(298, 'IS-F', 15),
(299, 'LS', 15),
(300, 'LX', 15),
(301, 'NX', 15),
(302, 'RC F', 15),
(303, 'RX', 15),
(304, 'RX 300', 15),
(305, 'RX 400h', 15),
(306, 'RX 450h', 15),
(307, 'SC 430', 15),

-- Models for Toyota
(308, '4-Runner', 16),
(309, 'Auris', 16),
(310, 'Avensis', 16),
(311, 'Avensis Combi', 16),
(312, 'Avensis Van Verso', 16),
(313, 'Aygo', 16),
(314, 'Camry', 16),
(315, 'Carina', 16),
(316, 'Celica', 16),
(317, 'Corolla', 16),
(318, 'Corolla Combi', 16),
(319, 'Corolla sedan', 16),
(320, 'Corolla Verso', 16),
(321, 'FJ Cruiser', 16),
(322, 'GT86', 16),
(323, 'Hiace', 16),
(324, 'Hiace Van', 16),
(325, 'Highlander', 16),
(326, 'Hilux', 16),
(327, 'Land Cruiser', 16),
(328, 'MR2', 16),
(329, 'Paseo', 16),
(330, 'Picnic', 16),
(331, 'Prius', 16),
(332, 'RAV4', 16),
(333, 'Sequoia', 16),
(334, 'Starlet', 16),
(335, 'Supra', 16),
(336, 'Tundra', 16),
(337, 'Urban Cruiser', 16),
(338, 'Verso', 16),
(339, 'Yaris', 16),
(340, 'Yaris Verso', 16),

-- Models for BMW
(341, 'i3', 17),
(342, 'i8', 17),
(343, 'M3', 17),
(344, 'M4', 17),
(345, 'M5', 17),
(346, 'M6', 17),
(347, 'Rad 1', 17),
(348, 'Rad 1 Cabrio', 17),
(349, 'Rad 1 Coupé', 17),
(350, 'Rad 2', 17),
(351, 'Rad 2 Active Tourer', 17),
(352, 'Rad 2 Coupé', 17),
(353, 'Rad 2 Gran Tourer', 17),
(354, 'Rad 3', 17),
(355, 'Rad 3 Cabrio', 17),
(356, 'Rad 3 Compact', 17),
(357, 'Rad 3 Coupé', 17),
(358, 'Rad 3 GT', 17),
(359, 'Rad 3 Touring', 17),
(360, 'Rad 4', 17),
(361, 'Rad 4 Cabrio', 17),
(362, 'Rad 4 Gran Coupé', 17),
(363, 'Rad 5', 17),
(364, 'Rad 5 GT', 17),
(365, 'Rad 5 Touring', 17),
(366, 'Rad 6', 17),
(367, 'Rad 6 Cabrio', 17),
(368, 'Rad 6 Coupé', 17),
(369, 'Rad 6 Gran Coupé', 17),
(370, 'Rad 7', 17),
(371, 'Rad 8 Coupé', 17),
(372, 'X1', 17),
(373, 'X3', 17),
(374, 'X4', 17),
(375, 'X5', 17),
(376, 'X6', 17),
(377, 'Z3', 17),
(378, 'Z3 Coupé', 17),
(379, 'Z3 Roadster', 17),
(380, 'Z4', 17),
(381, 'Z4 Roadster', 17),

-- Models for Volkswagen
(382, 'Amarok', 18),
(383, 'Beetle', 18),
(384, 'Bora', 18),
(385, 'Bora Variant', 18),
(386, 'Caddy', 18),
(387, 'Caddy Van', 18),
(388, 'Life', 18),
(389, 'California', 18),
(390, 'Caravelle', 18),
(391, 'CC', 18),
(392, 'Crafter', 18),
(393, 'Crafter Van', 18),
(394, 'Crafter Kombi', 18),
(395, 'CrossTouran', 18),
(396, 'Eos', 18),
(397, 'Fox', 18),
(398, 'Golf', 18),
(399, 'Golf Cabrio', 18),
(400, 'Golf Plus', 18),
(401, 'Golf Sportvan', 18),
(402, 'Golf Variant', 18),
(403, 'Jetta', 18),
(404, 'LT', 18),
(405, 'Lupo', 18),
(406, 'Multivan', 18),
(407, 'New Beetle', 18),
(408, 'New Beetle Cabrio', 18),
(409, 'Passat', 18),
(410, 'Passat Alltrack', 18),
(411, 'Passat CC', 18),
(412, 'Passat Variant', 18),
(413, 'Passat Variant Van', 18),
(414, 'Phaeton', 18),
(415, 'Polo', 18),
(416, 'Polo Van', 18),
(417, 'Polo Variant', 18),
(418, 'Scirocco', 18),
(419, 'Sharan', 18),
(420, 'T4', 18),
(421, 'T4 Caravelle', 18),
(422, 'T4 Multivan', 18),
(423, 'T5', 18),
(424, 'T5 Caravelle', 18),
(425, 'T5 Multivan', 18),
(426, 'T5 Transporter Shuttle', 18),
(427, 'Tiguan', 18),
(428, 'Touareg', 18),
(429, 'Touran', 18),

-- Models for Suzuki
(430, 'Alto', 19),
(431, 'Baleno', 19),
(432, 'Baleno kombi', 19),
(433, 'Grand Vitara', 19),
(434, 'Grand Vitara XL-7', 19),
(435, 'Ignis', 19),
(436, 'Jimny', 19),
(437, 'Kizashi', 19),
(438, 'Liana', 19),
(439, 'Samurai', 19),
(440, 'Splash', 19),
(441, 'Swift', 19),
(442, 'SX4', 19),
(443, 'SX4 Sedan', 19),
(444, 'Vitara', 19),
(445, 'Wagon R+', 19),

-- Models for Mercedes-Benz
(446, '100 D', 20),
(447, '115', 20),
(448, '124', 20),
(449, '126', 20),
(450, '190', 20),
(451, '190 D', 20),
(452, '190 E', 20),
(453, '200 - 300', 20),
(454, '200 D', 20),
(455, '200 E', 20),
(456, '210 Van', 20),
(457, '210 kombi', 20),
(458, '310 Van', 20),
(459, '310 kombi', 20),
(460, '230 - 300 CE Coupé', 20),
(461, '260 - 560 SE', 20),
(462, '260 - 560 SEL', 20),
(463, '500 - 600 SEC Coupé', 20),
(464, 'Trieda A', 20),
(465, 'A', 20),
(466, 'A L', 20),
(467, 'AMG GT', 20),
(468, 'Trieda B', 20),
(469, 'Trieda C', 20),
(470, 'C', 20),
(471, 'C Sportcoupé', 20),
(472, 'C T', 20),
(473, 'Citan', 20),
(474, 'CL', 20),
(475, 'CLA', 20),
(476, 'CLC', 20),
(477, 'CLK Cabrio', 20),
(478, 'CLK Coupé', 20),
(479, 'CLS', 20),
(480, 'Trieda E', 20),
(481, 'E', 20),
(482, 'E Cabrio', 20),
(483, 'E Coupé', 20),
(484, 'E T', 20),
(485, 'Trieda G', 20),
(486, 'G Cabrio', 20),
(487, 'GL', 20),
(488, 'GLA', 20),
(489, 'GLC', 20),
(490, 'GLE', 20),
(491, 'GLK', 20),
(492, 'Trieda M', 20),
(493, 'MB 100', 20),
(494, 'Trieda R', 20),
(495, 'Trieda S', 20),
(496, 'S', 20),
(497, 'S Coupé', 20),
(498, 'SL', 20),
(499, 'SLC', 20),
(500, 'SLK', 20),
(501, 'SLR', 20),
(502, 'Sprinter', 20),

-- Models for Saab
(503, '9-3', 21),
(504, '9-3 Cabriolet', 21),
(505, '9-3 Coupé', 21),
(506, '9-3 SportCombi', 21),
(507, '9-5', 21),
(508, '9-5 SportCombi', 21),
(509, '900', 21),
(510, '900 C', 21),
(511, '900 C Turbo', 21),
(512, '9000', 21),

-- Models for Audi
(513, '100', 22),
(514, '100 Avant', 22),
(515, '80', 22),
(516, '80 Avant', 22),
(517, '80 Cabrio', 22),
(518, '90', 22),
(519, 'A1', 22),
(520, 'A2', 22),
(521, 'A3', 22),
(522, 'A3 Cabriolet', 22),
(523, 'A3 Limuzina', 22),
(524, 'A3 Sportback', 22),
(525, 'A4', 22),
(526, 'A4 Allroad', 22),
(527, 'A4 Avant', 22),
(528, 'A4 Cabriolet', 22),
(529, 'A5', 22),
(530, 'A5 Cabriolet', 22),
(531, 'A5 Sportback', 22),
(532, 'A6', 22),
(533, 'A6 Allroad', 22),
(534, 'A6 Avant', 22),
(535, 'A7', 22),
(536, 'A8', 22),
(537, 'A8 Long', 22),
(538, 'Q3', 22),
(539, 'Q5', 22),
(540, 'Q7', 22),
(541, 'R8', 22),
(542, 'RS4 Cabriolet', 22),
(543, 'RS4/RS4 Avant', 22),
(544, 'RS5', 22),
(545, 'RS6 Avant', 22),
(546, 'RS7', 22),
(547, 'S3/S3 Sportback', 22),
(548, 'S4 Cabriolet', 22),
(549, 'S4/S4 Avant', 22),
(550, 'S5/S5 Cabriolet', 22),
(551, 'S6/RS6', 22),
(552, 'S7', 22),
(553, 'S8', 22),
(554, 'SQ5', 22),
(555, 'TT Coupé', 22),
(556, 'TT Roadster', 22),
(557, 'TTS', 22),

-- Models for Kia
(558, 'Avella', 23),
(559, 'Besta', 23),
(560, 'Carens', 23),
(561, 'Carnival', 23),
(562, 'Cee`d', 23),
(563, 'Cee`d SW', 23),
(564, 'Cerato', 23),
(565, 'K 2500', 23),
(566, 'Magentis', 23),
(567, 'Opirus', 23),
(568, 'Optima', 23),
(569, 'Picanto', 23),
(570, 'Pregio', 23),
(571, 'Pride', 23),
(572, 'Pro Cee`d', 23),
(573, 'Rio', 23),
(574, 'Rio Combi', 23),
(575, 'Rio sedan', 23),
(576, 'Sephia', 23),
(577, 'Shuma', 23),
(578, 'Sorento', 23),
(579, 'Soul', 23),
(580, 'Sportage', 23),
(581, 'Venga', 23),

-- Models for Land Rover
(582, '109', 24),
(583, 'Defender', 24),
(584, 'Discovery', 24),
(585, 'Discovery Sport', 24),
(586, 'Freelander', 24),
(587, 'Range Rover', 24),
(588, 'Range Rover Evoque', 24),
(589, 'Range Rover Sport', 24),

-- Models for Dodge
(590, 'Avenger', 25),
(591, 'Caliber', 25),
(592, 'Challenger', 25),
(593, 'Charger', 25),
(594, 'Grand Caravan', 25),
(595, 'Journey', 25),
(596, 'Magnum', 25),
(597, 'Nitro', 25),
(598, 'RAM', 25),
(599, 'Stealth', 25),
(600, 'Viper', 25),

-- Models for Chrysler
(601, '300 C', 26),
(602, '300 C Touring', 26),
(603, '300 M', 26),
(604, 'Crossfire', 26),
(605, 'Grand Voyager', 26),
(606, 'LHS', 26),
(607, 'Neon', 26),
(608, 'Pacifica', 26),
(609, 'Plymouth', 26),
(610, 'PT Cruiser', 26),
(611, 'Sebring', 26),
(612, 'Sebring Convertible', 26),
(613, 'Stratus', 26),
(614, 'Stratus Cabrio', 26),
(615, 'Town & Country', 26),
(616, 'Voyager', 26),

-- Models for Ford
(617, 'Aerostar', 27),
(618, 'B-Max', 27),
(619, 'C-Max', 27),
(620, 'Cortina', 27),
(621, 'Cougar', 27),
(622, 'Edge', 27),
(623, 'Escort', 27),
(624, 'Escort Cabrio', 27),
(625, 'Escort kombi', 27),
(626, 'Explorer', 27),
(627, 'F-150', 27),
(628, 'F-250', 27),
(629, 'Fiesta', 27),
(630, 'Focus', 27),
(631, 'Focus C-Max', 27),
(632, 'Focus CC', 27),
(633, 'Focus kombi', 27),
(634, 'Fusion', 27),
(635, 'Galaxy', 27),
(636, 'Grand C-Max', 27),
(637, 'Ka', 27),
(638, 'Kuga', 27),
(639, 'Maverick', 27),
(640, 'Mondeo', 27),
(641, 'Mondeo Combi', 27),
(642, 'Mustang', 27),
(643, 'Orion', 27),
(644, 'Puma', 27),
(645, 'Ranger', 27),
(646, 'S-Max', 27),
(647, 'Sierra', 27),
(648, 'Street Ka', 27),
(649, 'Tourneo Connect', 27),
(650, 'Tourneo Custom', 27),
(651, 'Transit', 27),
(652, 'Transit', 27),
(653, 'Transit Bus', 27),
(654, 'Transit Connect LWB', 27),
(655, 'Transit Courier', 27),
(656, 'Transit Custom', 27),
(657, 'Transit kombi', 27),
(658, 'Transit Tourneo', 27),
(659, 'Transit Valnik', 27),
(660, 'Transit Van', 27),
(661, 'Transit Van 350', 27),
(662, 'Windstar', 27),

-- Models for Hummer
(663, 'H2', 28),
(664, 'H3', 28),

-- Models for Hyundai
(665, 'Accent', 29),
(666, 'Atos', 29),
(667, 'Atos Prime', 29),
(668, 'Coupé', 29),
(669, 'Elantra', 29),
(670, 'Galloper', 29),
(671, 'Genesis', 29),
(672, 'Getz', 29),
(673, 'Grandeur', 29),
(674, 'H 350', 29),
(675, 'H1', 29),
(676, 'H1 Bus', 29),
(677, 'H1 Van', 29),
(678, 'H200', 29),
(679, 'i10', 29),
(680, 'i20', 29),
(681, 'i30', 29),
(682, 'i30 CW', 29),
(683, 'i40', 29),
(684, 'i40 CW', 29),
(685, 'ix20', 29),
(686, 'ix35', 29),
(687, 'ix55', 29),
(688, 'Lantra', 29),
(689, 'Matrix', 29),
(690, 'Santa Fe', 29),
(691, 'Sonata', 29),
(692, 'Terracan', 29),
(693, 'Trajet', 29),
(694, 'Tucson', 29),
(695, 'Veloster', 29),

-- Models for Infiniti
(696, 'EX', 30),
(697, 'FX', 30),
(698, 'G', 30),
(699, 'G Coupé', 30),
(700, 'M', 30),
(701, 'Q', 30),
(702, 'QX', 30),

-- Models for Jaguar
(703, 'Daimler', 31),
(704, 'F-Pace', 31),
(705, 'F-Type', 31),
(706, 'S-Type', 31),
(707, 'Sovereign', 31),
(708, 'X-Type', 31),
(709, 'X-type Estate', 31),
(710, 'XE', 31),
(711, 'XF', 31),
(712, 'XJ', 31),
(713, 'XJ12', 31),
(714, 'XJ6', 31),
(715, 'XJ8', 31),
(716, 'XJR', 31),
(717, 'XK', 31),
(718, 'XK8 Convertible', 31),
(719, 'XKR', 31),
(720, 'XKR Convertible', 31),

-- Models for Jeep
(721, 'Cherokee', 32),
(722, 'Commander', 32),
(723, 'Compass', 32),
(724, 'Grand Cherokee', 32),
(725, 'Patriot', 32),
(726, 'Renegade', 32),
(727, 'Wrangler', 32),

-- Models for Nissan
(728, '100 NX', 33),
(729, '200 SX', 33),
(730, '350 Z', 33),
(731, '350 Z Roadster', 33),
(732, '370 Z', 33),
(733, 'Almera', 33),
(734, 'Almera Tino', 33),
(735, 'Cabstar E - T', 33),
(736, 'Cabstar TL2 Valnik', 33),
(737, 'e-NV200', 33),
(738, 'GT-R', 33),
(739, 'Insterstar', 33),
(740, 'Juke', 33),
(741, 'King Cab', 33),
(742, 'Leaf', 33),
(743, 'Maxima', 33),
(744, 'Maxima QX', 33),
(745, 'Micra', 33),
(746, 'Murano', 33),
(747, 'Navara', 33),
(748, 'Note', 33),
(749, 'NP300 Pickup', 33),
(750, 'NV200', 33),
(751, 'NV400', 33),
(752, 'Pathfinder', 33),
(753, 'Patrol', 33),
(754, 'Patrol GR', 33),
(755, 'Pickup', 33),
(756, 'Pixo', 33),
(757, 'Primastar', 33),
(758, 'Primastar Combi', 33),
(759, 'Primera', 33),
(760, 'Primera Combi', 33),
(761, 'Pulsar', 33),
(762, 'Qashqai', 33),
(763, 'Serena', 33),
(764, 'Sunny', 33),
(765, 'Terrano', 33),
(766, 'Tiida', 33),
(767, 'Trade', 33),
(768, 'Vanette Cargo', 33),
(769, 'X-Trail', 33),

-- Models for Volvo
(770, '240', 34),
(771, '340', 34),
(772, '360', 34),
(773, '460', 34),
(774, '850', 34),
(775, '850 kombi', 34),
(776, 'C30', 34),
(777, 'C70', 34),
(778, 'C70 Cabrio', 34),
(779, 'C70 Coupé', 34),
(780, 'S40', 34),
(781, 'S60', 34),
(782, 'S70', 34),
(783, 'S80', 34),
(784, 'S90', 34),
(785, 'V40', 34),
(786, 'V50', 34),
(787, 'V60', 34),
(788, 'V70', 34),
(789, 'V90', 34),
(790, 'XC60', 34),
(791, 'XC70', 34),
(792, 'XC90', 34),

-- Models for Daewoo
(793, 'Espero', 35),
(794, 'Kalos', 35),
(795, 'Lacetti', 35),
(796, 'Lanos', 35),
(797, 'Leganza', 35),
(798, 'Lublin', 35),
(799, 'Matiz', 35),
(800, 'Nexia', 35),
(801, 'Nubira', 35),
(802, 'Nubira kombi', 35),
(803, 'Racer', 35),
(804, 'Tacuma', 35),
(805, 'Tico', 35),

-- Models for Fiat
(806, '1100', 36),
(807, '126', 36),
(808, '500', 36),
(809, '500L', 36),
(810, '500X', 36),
(811, '850', 36),
(812, 'Barchetta', 36),
(813, 'Brava', 36),
(814, 'Cinquecento', 36),
(815, 'Coupé', 36),
(816, 'Croma', 36),
(817, 'Doblo', 36),
(818, 'Doblo Cargo', 36),
(819, 'Doblo Cargo Combi', 36),
(820, 'Ducato', 36),
(821, 'Ducato Van', 36),
(822, 'Ducato Kombi', 36),
(823, 'Ducato Podvozok', 36),
(824, 'Florino', 36),
(825, 'Florino Combi', 36),
(826, 'Freemont', 36),
(827, 'Grande Punto', 36),
(828, 'Idea', 36),
(829, 'Linea', 36),
(830, 'Marea', 36),
(831, 'Marea Weekend', 36),
(832, 'Multipla', 36),
(833, 'Palio Weekend', 36),
(834, 'Panda', 36),
(835, 'Panda Van', 36),
(836, 'Punto', 36),
(837, 'Punto Cabriolet', 36),
(838, 'Punto Evo', 36),
(839, 'Punto Van', 36),
(840, 'Qubo', 36),
(841, 'Scudo', 36),
(842, 'Scudo Van', 36),
(843, 'Scudo Kombi', 36),
(844, 'Sedici', 36),
(845, 'Seicento', 36),
(846, 'Stilo', 36),
(847, 'Stilo Multiwagon', 36),
(848, 'Strada', 36),
(849, 'Talento', 36),
(850, 'Tipo', 36),
(851, 'Ulysse', 36),
(852, 'Uno', 36),
(853, 'X1/9', 36),

-- Models for MINI
(854, 'Cooper', 37),
(855, 'Cooper Cabrio', 37),
(856, 'Cooper Clubman', 37),
(857, 'Cooper D', 37),
(858, 'Cooper D Clubman', 37),
(859, 'Cooper S', 37),
(860, 'Cooper S Cabrio', 37),
(861, 'Cooper S Clubman', 37),
(862, 'Countryman', 37),
(863, 'Mini One', 37),
(864, 'One D', 37),

-- Models for Rover
(865, '200', 38),
(866, '214', 38),
(867, '218', 38),
(868, '25', 38),
(869, '400', 38),
(870, '414', 38),
(871, '416', 38),
(872, '620', 38),
(873, '75', 38),

-- Models for Smart
(874, 'Cabrio', 39),
(875, 'City-Coupé', 39),
(876, 'Compact Pulse', 39),
(877, 'Forfour', 39),
(878, 'Fortwo cabrio', 39),
(879, 'Fortwo coupé', 39),
(880, 'Roadster', 39)
ON CONFLICT (id) DO NOTHING;
