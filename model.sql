CREATE TYPE roles AS ENUM ('user', 'driver', 'admin', 'government', 'business_owner');
CREATE TYPE genders AS ENUM ('male', 'female');
CREATE TYPE statuses AS ENUM ('on', 'off');

CREATE TABLE IF NOT EXISTS Users (
     id uuid PRIMARY KEY,
     first_name VARCHAR(255) NOT NULL,
     last_name VARCHAR(255),
     email VARCHAR(255) UNIQUE NOT NULL,
     phone_number VARCHAR(255) UNIQUE NOT NULL,
     password VARCHAR(255) NOT NULL,
     date_of_birth DATE,
     age INTEGER,
     gender genders,
     role roles NOT NULL,
     is_blocked BOOLEAN DEFAULT FALSE,
     oauth VARCHAR(255) DEFAULT NULL,
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE BusinessOwners (
    id uuid PRIMARY KEY,
    NIK VARCHAR(255),
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Drivers (
    id uuid PRIMARY KEY,
    owner_id uuid,
    registration_number VARCHAR(255) UNIQUE,
    status statuses DEFAULT 'off',
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id) REFERENCES Users(id),
    FOREIGN KEY (owner_id) REFERENCES BusinessOwners(id)
);

CREATE TABLE IF NOT EXISTS Trips (
    id uuid PRIMARY KEY,
    user_id uuid,
    driver_id uuid,
    location VARCHAR(255),
    destination VARCHAR(255),
    trip_date TIMESTAMP,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (driver_id) REFERENCES Drivers(id)
);

CREATE TABLE IF NOT EXISTS Reviews (
    id uuid PRIMARY KEY ,
    user_id uuid,
    driver_id uuid,
    review VARCHAR(255),
    star INT,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (driver_id) REFERENCES Drivers(id)
);

CREATE TABLE IF NOT EXISTS Routes (
    id uuid PRIMARY KEY,
    route_name VARCHAR(255),
    initial_route VARCHAR(255),
    destination_route VARCHAR(255),
    created_at TIMESTAMP,
    FOREIGN KEY (id) REFERENCES Drivers(id)
);

CREATE TABLE IF NOT EXISTS ResetPassword (
    id INT PRIMARY KEY,
    user_id uuid,
    reset_code VARCHAR(255),
    created_at TIMESTAMP
);

CREATE FUNCTION expire_reset_password_links() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
DELETE FROM ResetPassword WHERE current_timestamp < NOW() - INTERVAL '1 minute';
RETURN NEW;
END;
$$;

CREATE TRIGGER expire_reset_password_links
    AFTER INSERT ON ResetPassword
    EXECUTE PROCEDURE expire_reset_password_links();

INSERT INTO "users" VALUES
    ('72335cde-7d00-3be6-ba9e-034286d2f51b','Rosetta','Ondricka','bo29@hotmail.com','004.462.6469x6367','c8513e59b6fdab8f8bc2bb0b8626d6dbae2632e9','1974-01-18',27,'male','user', FALSE ,NULL,'1972-02-04 09:32:37','2023-06-18 16:18:18'),
    ('0ef607b2-aaec-3de2-9243-04aef3037fce','Arlene','Pouros','kutch.gregorio@hotmail.com','(547)822-1466','8d02edd3a84b62594d84740463815581eed4afaf','1971-04-24',41,'male','admin', FALSE ,NULL,'2010-05-07 16:08:25','1984-02-28 21:13:57'),
    ('44cf7361-31b3-3db4-acbb-062fa2433c6a','Kristopher','Davis','efeeney@yahoo.com','1-737-846-9256x6059','1e96fc27c708aa2926f4b6687297bc3260381a82','1986-05-19',73,'female','business_owner', FALSE ,NULL,'1991-11-29 22:17:11','1997-02-02 00:30:43'),
    ('70586f67-8d3d-358d-810a-069aaba04d56','Salvador','Schimmel','nskiles@gmail.com','00277100195','2f66c960075bfffdf65b6313c9c700fc03c42e4e','1971-08-02',72,'male','admin', TRUE, NULL,'1978-02-03 02:23:36','2012-09-05 23:10:04'),
    ('4fa740e8-6551-391c-a719-0823eb8d8084','Devan','Buckridge','sasha.collins@hotmail.com','254.241.2132x67324','cf04fc36b4d91d24908d731c8d3f918ab24f2daa','1995-08-24',1,'male','business_owner', FALSE ,NULL,'2023-08-04 14:47:46','1987-11-27 02:51:00'),
    ('2722047c-468d-3166-aa03-10632cbfadc5','Guido','Stehr','xlegros@yahoo.com','1-826-235-4532x787','f2fa0baa76400f79c8c5c4805c3a97c9bc852dc5','2008-10-14',58,'female','business_owner', TRUE, NULL,'2011-07-05 20:14:19','1986-12-23 19:30:36'),
    ('8581d025-9b2e-3b1d-8a8b-13ad883226b3','Jamel','Purdy','mueller.jedediah@hotmail.com','805-085-1721','ee48ab699c5a4d8eb5d652805f882d701d72b742','1991-01-31',61,'female','user', TRUE, NULL,'1990-11-14 18:11:10','1971-08-03 13:34:27'),
    ('c8821230-7eb5-3ed8-897c-13e2dee91d6b','Hertha','Cartwright','floyd.ebert@hotmail.com','818.259.9491x1506','dffda0bbb7c85c181dbcee41f3f885cbb272a7e2','1982-12-22',64,'male','admin', TRUE, NULL,'1972-08-04 06:49:31','1984-05-06 03:08:12'),
    ('d339176a-dd0e-3229-b9f7-17f22a2f02eb','Elmo','Deckow','kellen38@yahoo.com','(321)301-8368x2614','40cc8058e45885818ad7bdb28c172dde642ec0d6','1988-11-18',74,'male','business_owner', TRUE, NULL,'1970-06-06 18:03:31','1973-11-13 04:27:55'),
    ('9be79a28-c3a7-3e97-bcc4-1bad7abc46d2','Maurice','Adams','cmueller@yahoo.com','(008)598-4482x97713','07c7d6dd5b0531dfe65469d7202f72407b29e7d8','1975-01-07',19,'male','user', FALSE ,NULL,'1982-05-19 15:01:38','1973-12-30 16:30:45'),
    ('82590018-ee56-356f-a236-1e1e68072836','Montana','Schultz','leffler.bruce@yahoo.com','798-946-4713','0a97df676f61a262cd7778ce338f27054991c5cf','1996-03-15',1,'male','admin', TRUE, NULL,'2021-09-25 03:58:57','2012-12-01 21:28:41'),
    ('c2cad4d9-8b91-3f7a-bca5-1fbeda837dda','Shirley','Gaylord','dooley.dallas@yahoo.com','(865)591-6731x052','31903746c16611e2bdd4ecd50bd56c369cba4a22','1986-10-03',79,'male','admin', FALSE ,NULL,'1985-07-18 20:16:52','1975-03-30 07:55:36'),
    ('ce98d4da-cb38-34c9-9f3a-2a454e3fc4b3','Rasheed','Upton','alexandrine91@gmail.com','436-701-5293x11450','6e6541b2584ef9388c39e5c413ae6b78748213fc','1983-08-18',66,'male','government', TRUE, NULL,'2019-12-01 16:19:48','1986-05-17 04:41:25'),
    ('59e66f9e-1885-3bdb-a521-2f449e1064f0','Jettie','Cruickshank','fthiel@yahoo.com','(270)246-2959','1ceb92cc7650d40b22b30b16a44b0526491f77db','2012-05-15',37,'female','admin', FALSE ,NULL,'2020-09-28 01:43:15','1973-08-31 14:43:50'),
    ('ffe1e315-b9cd-3a14-af6e-30d461024fb9','Gaylord','Reichel','isom59@hotmail.com','+16(8)4321735987','5127da067b66ab6d7d7e529c0332e11d4e94cd97','2009-06-24',19,'female','user', FALSE ,NULL,'1993-07-28 10:09:04','1996-06-14 17:04:03'),
    ('da98133b-c212-3d3b-91b8-32f88ac7d8ba','Jesse','Spencer','pfannerstill.rosalyn@hotmail.com','191-227-1337','aea2b647f66f9df2cf9b08ff806b97d394b09a8e','2008-06-02',30,'male','user', FALSE ,NULL,'1984-04-27 15:30:34','1997-09-04 20:32:05'),
    ('7d333ff0-ae8b-3dba-9dd9-334f46dbfed9','Jackie','Gutmann','zemlak.kurt@gmail.com','(834)368-0898','986f69405ffa7f64fe61cb46ebfb1ec5f17e57cf','2019-12-14',74,'female','admin', TRUE, NULL,'2005-10-18 09:12:51','1983-07-19 00:16:10'),
    ('3a20bf0d-58c0-30f4-a8e5-36f9d0829954','Margret','Kohler','gorczany.earl@gmail.com','636.587.0721x7536','1bdb8dfb3d8cd57624cc8d12653c04c9664c2564','1983-04-27',20,'male','business_owner', TRUE, NULL,'2015-06-15 17:16:54','1970-11-03 20:39:50'),
    ('8ece6a50-0f4d-30bb-a16c-37e9f23f1ddc','Amparo','OConnell','gerlach.kareem@gmail.com','1-794-658-9771','4cf7f94e16f8884044c61d7a91758b3e7050f293','1973-12-27',14,'female','business_owner', TRUE, NULL,'2009-02-07 21:29:38','1977-12-27 08:31:04'),
    ('ae45642a-2d7c-3acf-8366-382d727afa56','Alf','Casper','uberge@yahoo.com','+16(5)5389071243','00cd3da1bead7dee1dd08282a92cb50af9328669','2007-12-06',43,'male','business_owner', FALSE ,NULL,'2007-04-24 11:06:37','1993-12-05 04:16:58'),
    ('618d9c35-21ae-34c2-8db1-3b1040516ea3','Rahsaan','Sanford','mauricio79@yahoo.com','876.885.2460x18868','e413b38cfbd519df483dadd5a27f7f04eb843ca4','1987-05-11',12,'male','driver', FALSE ,NULL,'2022-07-31 10:38:17','2001-12-01 08:55:26'),
    ('c49da4bc-4fda-32d5-845f-3d94e8e905d4','Jordon','Leffler','anabelle19@hotmail.com','+12(8)7196348567','4cb783058d9c4118cd3a73729090acc99992238b','1991-05-31',47,'male','government', TRUE, NULL,'1982-06-21 19:39:13','1990-04-30 16:30:01'),
    ('8c8c6b0b-b23d-34f7-affc-40416f936b89','Kirstin','Schaden','ghintz@hotmail.com','630.873.3418x61443','47d82bda9b2755018e41c3f4ad6eb5b681e219e4','1980-09-17',34,'female','driver', FALSE ,NULL,'2004-06-09 22:04:47','2014-07-26 06:21:34'),
    ('a7b97c9b-ca52-3061-9d31-412a5310885d','Germaine','Kreiger','keith53@hotmail.com','1-043-519-4455','6acc6b5015592347b91553432da6e591e238e871','2024-06-08',67,'male','business_owner', TRUE, NULL,'2024-03-19 03:35:15','1990-03-17 07:27:41'),
    ('db2c6719-fea1-30af-b073-413df20dd6e6','Elouise','Little','dorothea.gottlieb@hotmail.com','791.328.4577x186','d8da73be6f15e3a3fa69cdba09a5622edac9d1c4','1989-09-02',45,'female','government', FALSE ,NULL,'1974-12-03 05:29:52','2018-11-30 01:55:07'),
    ('0676454b-6159-3d6c-8f38-424714347fd5','Annabelle','Tromp','rollin20@yahoo.com','999-503-5879','9c818a8fbd7e805e4dcb2cd1efed2a959da7894b','2023-07-02',28,'female','business_owner', FALSE ,NULL,'1987-01-04 10:38:45','1989-04-08 10:01:13'),
    ('5255b8b3-492f-3c70-8ccb-498d260df99f','Berta','Greenfelder','schneider.devyn@yahoo.com','050.445.1095x492','312d0ca5ce3068f4b160b9f49e2aeeb61db07dbc','1995-11-11',75,'female','user', TRUE, NULL,'1995-09-13 07:05:00','1994-08-03 22:38:55'),
    ('4b098881-1e09-37ce-b5ad-4acc16ed9085','Edwin','Mayert','fwest@gmail.com','(199)089-8510x059','b87b5aff56db21b3bf68866ce0c1b6ce759f7123','1995-05-27',25,'male','admin', FALSE ,NULL,'1988-01-29 13:44:02','2020-10-20 12:37:44'),
    ('dfcfc580-79c6-39c7-b71f-4f589e7f1800','Llewellyn','Yost','misael38@gmail.com','405-972-5993x60104','572bb45b8d858e41146b1bf5d6b4b3ab52d16a76','1991-06-04',67,'female','driver', TRUE, NULL,'1971-01-22 18:12:52','2003-01-15 03:30:15'),
    ('8f7b57ad-4bbf-393e-8780-5120619d4667','Kelvin','Bechtelar','rhartmann@yahoo.com','559.296.3053x008','f3e2f80122fa274a7837065177d88dc79e17d559','2011-12-18',37,'male','user', FALSE ,NULL,'2022-03-10 23:08:18','2012-06-05 16:26:22'),
    ('6db6c69d-13d4-3dca-aef8-53cd6f1dc447','Asha','Eichmann','tgerhold@yahoo.com','1-723-219-8068','8cfcfcf1f30f342599dac51be65e4635476635c1','1970-04-26',20,'male','government', TRUE, NULL,'2016-03-19 01:34:59','1992-12-12 14:59:59'),
    ('01567e7a-7708-30a0-9111-56bf49e116ba','Jazmyn','Hills','virgie.hamill@yahoo.com','(331)725-5564x060','0a7a83d1b9528992a50cde9048071f8026b4bc87','2013-08-05',65,'male','driver', FALSE ,NULL,'2021-11-06 11:20:41','1976-10-10 15:43:38'),
    ('d5203fbd-8fa8-3f1a-bc75-6da251475abd','Bernadette','Gibson','ali93@gmail.com','604-977-0451x473','ce83a77d90fe3e951f853f7f75a42da6715b1750','2011-08-04',80,'male','admin', TRUE, NULL,'1972-05-19 12:22:17','2024-03-13 12:58:14'),
    ('5b54f944-713b-3e60-9eaf-71db3f601225','Dax','Murray','gerlach.cullen@yahoo.com','+65(3)5834825401','bc5707c712522f443bf661b13067471ebf6b54fd','2021-06-30',53,'male','driver', TRUE, NULL,'2001-08-29 08:27:12','2016-03-13 18:19:36'),
    ('a07fd0e7-4714-3836-9a7a-7532917ad402','Kris','Fadel','cristobal.stehr@hotmail.com','272-099-0093','dff45bab0e7d4e1c5c4267ec2b5650a0cf2b2c54','2017-09-03',34,'male','business_owner', TRUE, NULL,'1986-07-11 20:52:51','1984-12-05 15:07:17'),
    ('aef42f55-e41e-36fb-9a58-79797d4e6c93','Hilario','Pacocha','oran.littel@hotmail.com','1-734-895-3607','8ac949ae6d40c6a92984bd83fe06dedbd90ece28','2010-01-14',47,'male','business_owner', TRUE, NULL,'1995-09-08 17:48:50','1999-04-23 10:53:38'),
    ('d860f148-a057-3cee-9c0d-7c798643d058','Viola','Gulgowski','zstokes@gmail.com','1-502-483-7779','cc9247de0fd4862a36b93e96876105e815144e79','2015-05-05',28,'female','business_owner', FALSE ,NULL,'1991-06-07 18:59:49','2015-04-13 16:45:03'),
    ('d021f141-7771-3471-9271-7e27682fcdef','Coy','Schmitt','walker.arvilla@hotmail.com','07161927459','2649ecf65d91595e16fe5380e0b4400da8b2227e','2000-03-21',12,'male','business_owner', TRUE, NULL,'2006-02-15 09:01:29','2019-06-23 21:32:56'),
    ('3ebe1f95-1a8b-3a0e-9439-7ef695b6d9f7','Vicky','Wiegand','okey.wehner@hotmail.com','341.829.9873x074','fa8a27b9020ad6356809d27ef3206d50f41590bb','2001-05-30',20,'female','driver', FALSE ,NULL,'2005-05-02 09:59:22','1973-01-09 13:26:47'),
    ('691486a9-2028-3f46-b9c4-81831c2bbd5b','Jevon','Larson','ernie05@gmail.com','952.596.9722','851d2388bd0b8db00f7ae04e530ada042329d699','1984-06-22',25,'female','government', FALSE ,NULL,'1970-05-25 13:40:58','1988-10-20 10:37:25'),
    ('95a947ec-1f86-37ee-b4cd-83281bbf8d53','Adrian','Lehner','lmueller@gmail.com','917-760-7187x4311','50cfeb476d007b29f6eb502c21b7e3a3915b1735','1970-01-05',47,'female','admin', TRUE, NULL,'2019-05-30 09:38:04','2015-03-19 13:10:41'),
    ('aafdd3b3-843f-3c70-b4de-84b91a3254cd','Americo','Bogan','jbrekke@hotmail.com','1-247-600-9486','0f3f4aca2b678799f13e068a0d0a43d264b4bb5b','1995-12-18',57,'male','user', FALSE ,NULL,'1975-07-31 03:00:49','2016-03-07 01:06:49'),
    ('833af3ae-fe5c-343f-9dd0-84fb18ed1632','Deon','Terry','nelle42@hotmail.com','615.624.1759x23161','671c7f686695cf1eb26bc2e5c43c49698c48133f','1988-12-01',18,'female','admin', TRUE, NULL,'2010-07-23 01:58:43','1991-06-22 12:51:02'),
    ('c7c34b50-81b3-35e3-8974-8cca3dee86f6','Kaya','Klein','baylee40@yahoo.com','08654887046','615945b7c7dc99a323efbc48e0b59b876053f940','1980-06-11',64,'male','driver', TRUE, NULL,'2000-01-07 09:37:30','2021-11-02 15:57:14'),
    ('20f398ff-6dca-31e2-9dd0-8e13fbdea4ee','Letitia','Mosciski','dahlia87@gmail.com','(860)587-6130x75893','aafecca5e23ea1f100706db6d05a5650557f94f5','2016-10-28',8,'male','business_owner', FALSE ,NULL,'1971-04-22 23:45:40','1976-03-15 04:49:22'),
    ('221d5dda-267e-3358-85f3-8edf95299fd4','Ola','Johnson','trevor.cremin@hotmail.com','(867)536-7146','515d6c3538b138fa805ae2ea4ddfbcf14858e82c','2004-10-07',2,'male','user', TRUE, NULL,'1993-12-04 22:11:15','2009-06-19 03:24:12'),
    ('95dd31fc-f67b-3c38-912c-90a32a3cbb87','Pete','Murazik','wintheiser.lenna@yahoo.com','(099)290-3895x549','8403da3649a46f7cfaacc22ead3e0aa7d856215d','1978-06-22',74,'female','business_owner', FALSE ,NULL,'2015-12-25 16:15:49','2019-03-06 17:18:38'),
    ('47af7837-f382-3421-b48b-90fa72908b53','Isabell','Bernier','beier.opal@yahoo.com','512-356-8512x113','ef6980c18a2fab51b481a0b575467dafc1ee47b8','1972-09-30',26,'female','user', FALSE ,NULL,'2003-08-01 17:58:32','1973-03-08 07:19:31'),
    ('bc287905-5036-36e1-aa41-9356e9bdd6e4','Sigurd','Champlin','tyrese56@gmail.com','1-846-466-5553x79198','8754bcc9aedfd9e51c86f4a94af89af10780a6e1','1978-12-24',78,'female','user', TRUE, NULL,'2017-01-19 16:59:09','2012-04-15 10:52:00'),
    ('4e8cc39c-38c9-3ac1-b991-93b5e13ecc12','Burley','Conroy','vernice68@gmail.com','437-042-2356x45941','eab8581793bbd14b32f3cfdc26627d51aa64ecdf','1988-12-04',39,'female','user', FALSE ,NULL,'1991-06-04 13:49:35','1979-10-10 14:05:53'),
    ('04ab511a-7bca-33e5-9576-95d80dd42701','Nicholaus','Wunsch','tnikolaus@hotmail.com','874-809-7910','ee894e2fc4043d835c976c3c294e17f4dff3927b','1970-10-30',46,'male','business_owner', TRUE, NULL,'2001-10-31 14:47:27','1999-11-11 11:22:44'),
    ('1192cfac-209c-38d0-a78a-9755f2d1962b','Tyra','Wiza','alex19@yahoo.com','161-110-8396','42786a1fc402f1db9ff524e7229b00e8061be706','2021-03-11',70,'female','user', TRUE, NULL,'2010-09-07 00:45:01','1970-01-13 12:27:19'),
    ('9770ee66-db2c-337d-a86f-98cd369b2a91','Nola','Cummings','felipa87@hotmail.com','(931)285-8258x60926','09d663ec621c137facf3f2f2cbee6d34a58a287f','1972-01-01',12,'female','user', TRUE, NULL,'1990-07-13 05:05:30','2001-12-17 18:50:04'),
    ('6547a234-d6a0-3858-b5dd-9ba869f95824','Avis','Kutch','dax65@hotmail.com','996.118.3618x863','73e2ed71ea4806d90b54157ed20217c05353d166','2010-12-14',25,'female','driver', TRUE, NULL,'2011-10-08 13:39:40','1971-04-16 20:20:25'),
    ('4d2dc3af-5eeb-32cd-bfdf-9bac6bf2c81a','Kellen','Corkery','coby92@yahoo.com','+40(3)4898882539','f7ef90518a352a357e191dc2ef8c24e4492b4e11','1991-03-22',31,'male','driver', TRUE, NULL,'1993-08-22 02:26:52','1983-07-20 18:42:33'),
    ('97e32669-6c69-3dd9-8877-9c96764d0a3f','Lewis','Raynor','barton30@hotmail.com','(468)191-2948','6a31a35bef94a0ab766fa2ca8eca03b1a062ae87','2001-02-23',43,'male','government', FALSE ,NULL,'2008-02-15 10:26:21','1971-08-10 14:48:56'),
    ('0118ca3f-ecf1-3f21-901d-9fcdfe96c6bb','Cheyenne','Watsica','brandi.wehner@gmail.com','622-402-0140x38057','166712aebf5643d260cc83871ddf7ef9e5f2c6ae','2016-08-09',22,'female','driver', TRUE, NULL,'1977-05-14 13:43:54','2005-06-12 07:06:18'),
    ('8f26b57b-b197-3591-879c-a30cdc2cd501','Agustina','Padberg','rhea21@yahoo.com','192-889-0709','1379a6a97b64f7567df7f546c608924c6681073f','1995-02-12',75,'female','user', FALSE ,NULL,'1973-05-31 11:21:37','2005-06-11 01:24:24'),
    ('0718ffba-cb7c-3789-832a-a40fa83f2365','Merle','Kemmer','katharina.gutmann@gmail.com','507.420.9987','d2741e68f6396bbb77baac1e8f637af9095b195c','1999-08-27',7,'male','driver', FALSE ,NULL,'1981-02-07 02:26:10','2019-11-08 18:35:33'),
    ('2bffe4da-b20b-3945-a7fe-a553338e88f5','Ignatius','Cremin','uschoen@hotmail.com','04733128134','597c05105c84c76eea0bb162d45eb7ee1e4db2dc','2017-08-31',19,'male','government', FALSE ,NULL,'1977-10-11 20:10:45','1970-02-15 13:25:38'),
    ('6bca79c2-ed76-34bf-8209-a56398e8a261','Antone','Bechtelar','bauch.okey@gmail.com','(294)418-5456','df036079a7ef5d81be70d29fdf1f8f3a78a6502a','1980-03-03',13,'female','user', TRUE, NULL,'1999-02-12 20:37:44','2001-02-15 13:44:56'),
    ('e9836d4e-48c4-3658-9a5a-a5ade1b311d7','Sim','Hoeger','schamplin@gmail.com','06083878264','cb73b5236101b0bd8242b789319dabed564b0e56','1988-04-25',59,'male','driver', TRUE, NULL,'1976-03-19 02:30:09','1997-02-02 10:59:05'),
    ('33336ca2-d7c4-3a22-9ef5-a9bacdac13dd','Adaline','Conroy','mohamed48@yahoo.com','792-559-3442','11d1864fa78bc21d0559311932d7440b84a72434','2008-11-03',34,'male','admin', FALSE ,NULL,'1997-03-09 22:08:29','2006-10-03 01:29:54'),
    ('119c6e95-42aa-3a04-a2d8-aa84273c9411','Humberto','Robel','leanna00@gmail.com','1-339-062-2218x159','14671c94aef30718f3533aea677a214e68d9bfb7','2006-08-16',27,'male','driver', FALSE ,NULL,'2004-03-09 07:02:48','2004-10-11 16:04:56'),
    ('09d31ba9-7b5a-3613-a32e-ae8075906964','Samson','Graham','stehr.raheem@hotmail.com','1-433-818-6910','73ba65512a732ac26f07e01d38c7db7793881625','1985-10-10',70,'male','admin', TRUE, NULL,'2019-04-02 01:55:41','1997-07-07 19:13:28'),
    ('a8cad78f-6164-3941-9794-aefd9fdd8b5e','Jalen','Flatley','nader.brandon@hotmail.com','648.343.2490','e1bb65b398b24df0492ffed57f6f009a762b9cd9','1972-03-17',44,'female','admin', TRUE, NULL,'1977-06-05 18:32:43','1976-09-10 19:52:20'),
    ('0c00630e-773b-34b0-88af-b022ca137ab1','Berta','VonRueden','justina.schimmel@hotmail.com','(537)252-0108x0593','c6f7127c56043d332da8ebb6830446006a48e860','2008-01-15',44,'female','business_owner', FALSE ,NULL,'2007-06-16 23:59:58','2020-12-24 09:41:22'),
    ('2a33c379-4c9a-3767-be72-b107ab4a3273','Allen','Aufderhar','allen.harvey@gmail.com','(259)192-1117x4977','b298ff0d27fbd5ba44afbf6cbb76514d20163661','2010-07-02',17,'female','admin', TRUE, NULL,'2020-04-16 02:33:04','1999-01-26 11:04:24'),
    ('1388716e-70f4-3391-a9fa-b10dfc413ee0','Moises','Prohaska','xbarrows@gmail.com','1-022-524-1015','cb55ebc9fbc01c36e1130cc16b7b6013e53d5de8','2015-07-28',66,'male','user', TRUE, NULL,'1976-03-05 12:59:16','1984-10-22 15:11:16'),
    ('6ef2d648-222a-3e69-975f-b20379af341d','Green','Raynor','araceli10@yahoo.com','862.975.8797x3133','d15d15511b98adc66b0029640a93cfe7d48ed1de','1973-04-17',66,'male','user', TRUE, NULL,'1973-01-22 12:29:55','2008-09-02 12:49:07'),
    ('bfc784e0-812b-3375-8d2a-b7129ad51987','Wade','Fay','taurean85@gmail.com','1-799-291-3060','0cbb3efbd92e8b33dbae565a921692d41b648502','1971-05-31',9,'male','government', TRUE, NULL,'2022-07-17 16:51:29','1990-07-05 14:24:52'),
    ('7b4c04cc-4057-334e-8103-b9f63036e127','Axel','Douglas','gia.hagenes@gmail.com','+12(3)7972097553','206635aca020bfd88bf58b19886f5702f93b4e03','2002-07-19',49,'female','admin', FALSE ,NULL,'1994-03-12 13:33:58','2005-04-01 10:16:55'),
    ('42fa7d22-db7e-31e4-b332-bb21639f278a','Bonita','Ankunding','labadie.porter@hotmail.com','991-046-9048x4078','4164d19ad52cc03cf8c747a41c25c02ae6679bf6','2015-07-24',50,'female','driver', FALSE ,NULL,'1974-04-06 16:56:54','2019-09-18 09:16:51'),
    ('47b2b605-4b23-3cc5-88d5-be2d9a3c1357','Adaline','Ledner','becker.maureen@yahoo.com','02333861868','b1aae04fc21ef60db108d17ac673129683aec680','2000-02-20',25,'female','driver', TRUE, NULL,'1972-02-13 23:10:24','2014-06-17 22:10:07'),
    ('a63b141a-1ca6-3374-a5f2-be3439ddfa4c','Alisa','Walter','magnolia.hoppe@gmail.com','1-635-921-1276x8076','455f1c4b19d017c08e9a98cda2c4015e6eba9889','1997-10-24',35,'female','admin', TRUE, NULL,'1997-11-25 03:57:01','2019-12-11 17:52:37'),
    ('bd5e9517-efe9-307e-b228-c252439a2a6e','Ewell','Ankunding','lowe.keshaun@yahoo.com','634.155.2773x87725','ae9e82da74bd40eed511a85a32a449404be05b36','2023-12-15',14,'male','admin', FALSE ,NULL,'1982-10-27 20:28:56','1990-08-20 14:20:40'),
    ('ec61e5ba-c5bd-3e04-92d6-c68ed5e0abec','Hubert','Boyle','samanta62@hotmail.com','685.127.1599x64742','d67b7cef668bdaf2ccf464731cd852e64a6e1d62','2020-01-21',38,'female','government', TRUE, NULL,'2024-06-22 18:39:57','1985-02-09 19:14:24'),
    ('8fb5806f-b06d-3bbf-8a4b-c87835cbb912','Cleora','Jacobs','okeefe.jillian@yahoo.com','+62(0)6197454327','888319bbf4b53a7ffa2365998dec40febcd58a8c','1977-09-24',55,'female','driver', TRUE, NULL,'2012-12-03 04:01:27','1978-09-28 11:24:25'),
    ('2e067a8e-7e13-34ab-a15c-c8890c1ae822','Earlene','Stokes','lillie76@yahoo.com','297-239-3475x259','9ee672428e5a36a6bcec41a2388b7abf3a5a30f5','2002-07-07',48,'female','business_owner', FALSE ,NULL,'2000-06-11 09:01:49','2007-04-16 14:29:46'),
    ('f99cebbf-72a9-3292-a9ac-cabbb0320924','Alanis','VonRueden','brown.beatty@hotmail.com','(672)168-0069x170','5c37a1c9d9f7dc280bbad4f8dae5b9b5877a58b7','2007-01-18',40,'male','admin', FALSE ,NULL,'2006-07-20 12:29:48','2018-06-16 15:32:20'),
    ('c11b3755-40c6-38fc-9ea6-d0c3d308bdd1','Tracey','DAmore','enola04@gmail.com','940-580-0761x190','a88563a46dfb93e51c4d282f340c62d449d6165c','2017-09-29',25,'female','user', TRUE, NULL,'1995-05-07 12:35:52','2011-01-27 10:44:51'),
    ('5f13ccbd-c96b-3218-9a6a-d4bc5391e5bf','Kirk','Bode','nvon@gmail.com','01257928115','729296d5679def9650703701ed9969ebfc63e3ad','1984-07-13',67,'female','government', TRUE, NULL,'1976-08-01 19:23:43','1993-10-20 04:39:32'),
    ('d2e40c31-da66-3edc-bd82-d5710589b695','Elvie','Stark','kulas.lucie@hotmail.com','1-760-579-4817','54e54006c4196d05e0c6aadadb48d7e673b90b36','1981-03-23',11,'male','business_owner', FALSE ,NULL,'1971-08-08 09:53:30','1994-06-17 02:44:35'),
    ('30e3a04c-0c37-31d1-93e5-d8cfa34d4a02','Judah','Mertz','chaz30@gmail.com','590-047-7424','0bce2ff6f3a86aedbc64294f123f02fd107205bb','1991-10-09',64,'male','business_owner', FALSE ,NULL,'1991-05-18 03:26:21','1998-09-16 05:07:28'),
    ('06c4d777-caf8-3ff5-955c-db2408cf85c6','Shirley','Dach','coleman88@yahoo.com','1-375-856-4948x24783','003c111960688da46ca194755d2d8cd30c7c2fde','1975-07-11',73,'male','admin', FALSE ,NULL,'1986-06-14 16:08:14','1978-10-11 22:37:34'),
    ('edee1fef-6d42-392c-b536-dbabb7e1772c','Jarret','Blick','herminia.rogahn@yahoo.com','757-143-4583x899','050c3dc37b322339ff736f216ffdd4beb98bd4a8','2007-11-09',46,'female','driver', FALSE ,NULL,'2020-11-16 05:08:28','1975-01-26 17:33:52'),
    ('2a06b0c1-3ab0-33b2-ab7c-dc5769ee18b2','Nico','Bogisich','orville.ullrich@hotmail.com','408.700.2460x0286','6cbcf35023fbe5dae5cd8fa1cccb9821f27a63af','1980-06-06',42,'male','government', TRUE, NULL,'2015-10-27 21:40:28','2008-11-12 11:12:08'),
    ('747f00e2-17dd-3413-a816-e357b8e4a8a3','Susie','Schmidt','shanelle34@hotmail.com','(416)036-8860x3248','29d6f0a28cd597d4f17e8d12045981dd5b7fe196','2006-12-01',35,'male','admin', FALSE ,NULL,'2010-11-17 15:08:06','2006-10-16 11:34:50'),
    ('b314db34-1b7a-3dda-a094-e3b9aaf59920','Jonatan','Kshlerin','hudson.luciano@hotmail.com','347.584.9412x8465','c1132dd06a296c228f15d65b5faf204cd2cdd10c','1997-09-27',51,'male','government', TRUE, NULL,'2006-07-01 02:47:26','2013-01-25 21:09:19'),
    ('ca0df884-c765-3f0a-b215-e4b31ed17225','Natalia','Steuber','pframi@hotmail.com','1-827-188-2486x1075','760a89d69d503e644c86d59ba22e063cef93a433','1976-07-21',24,'male','government', TRUE, NULL,'1975-02-26 01:17:32','2019-11-22 11:52:44'),
    ('7a02bd1c-38a4-3b22-8463-e7201fd24de3','Eulah','Heidenreich','emard.abbie@yahoo.com','(323)578-3971','1c1eab2ac148e9099cb20777ceb6fb37f66b56e4','2004-02-03',44,'female','driver', TRUE, NULL,'2018-12-08 16:46:08','1990-10-26 06:33:36'),
    ('e0178eb8-ab2c-3341-901c-e7b6db20ef8c','Mervin','Stamm','jenkins.rolando@hotmail.com','(463)724-6336x719','c1bdf52d9af668e5034d1c56114c1eb056200c30','2023-07-10',77,'male','admin', TRUE, NULL,'1972-09-22 11:38:33','1979-08-08 17:01:24'),
    ('62abdb35-655e-39d5-87f9-ef10daaccc20','Peter','McCullough','sarina.stokes@yahoo.com','(957)676-5843x7346','8868deac6f3c40e96e14c0545358bb0f26ec880d','1975-07-01',72,'female','admin', TRUE, NULL,'1971-10-27 01:49:50','2004-07-05 16:41:04'),
    ('bd2d8841-ed25-3772-b8a6-ef87dd87d311','Ed','Eichmann','bernie.labadie@hotmail.com','757-795-5747x82040','a4f98be448dc45015b54ea5e87b2a09bd64969ef','2023-06-23',75,'female','admin', FALSE ,NULL,'1991-10-11 02:13:56','1991-01-15 07:25:16'),
    ('628b8a2a-5db9-3c6f-9eac-f17c05ad32c9','Loma','Pollich','shanelle13@gmail.com','317.812.6562','a99f33484ba708044ed6db6cd2e761107295930f','2024-02-08',21,'female','admin', TRUE, NULL,'2021-03-12 04:01:07','1990-05-08 15:22:59'),
    ('98310b88-99ee-3833-8cd7-f4c18f051b7b','Leola','Kovacek','owen06@hotmail.com','274-142-3007x03969','89da70e5ad78614219aecc14c6f5b85d0dbdaf63','1991-04-07',71,'female','driver', FALSE ,NULL,'2020-12-17 20:23:12','2012-08-05 19:45:22'),
    ('45c56d03-d58e-3a98-bd02-f993749ff50e','Eleanore','Considine','caleigh54@gmail.com','446.282.0910x99517','2c4c36ce4298951153c3438cd7a6569fd1598b02','1972-04-01',21,'female','user', FALSE ,NULL,'1995-02-22 09:15:01','2015-02-02 03:46:32'),
    ('bee80fdf-e772-30bf-a13c-fc4bd62c7050','Malinda','Bogisich','gleason.emmalee@yahoo.com','032.524.8305x7686','24fceac96071afba7a86e51df5e3c75ea9e01a9e','2002-04-11',3,'male','government', FALSE ,NULL,'1981-01-13 17:27:26','1985-01-30 22:55:10'),
    ('ab47f183-2b45-3ebe-a799-feb0199ca625','Krystal','Stanton','jermain.macejkovic@yahoo.com','418-136-9993x438','0787249a4f1ba64e26e08f2b84fc3c3bc172559a','2017-01-12',77,'male','government', FALSE ,NULL,'1974-06-23 04:22:15','1989-11-09 14:07:36'),
    ('25a157a5-a0c9-39be-b633-ff690e3cdd3c','Eloy','Hirthe','zachariah.strosin@yahoo.com','863.449.5919x829','a9860b5d8c3383fb7422ebc674f702e4b3cd9d92','2002-07-24',15,'male','driver', FALSE ,NULL,'2016-07-29 22:10:40','1995-09-22 16:43:38');
