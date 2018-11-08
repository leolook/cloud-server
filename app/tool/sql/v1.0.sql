
drop table `db_info`;
CREATE TABLE `db_info` (
  `id` INTEGER   primary key autoincrement ,
  `name` varchar(200)  NOT NULL ,
  `ip` varchar(20)  NOT NULL ,
  `port` int(11)  NOT NULL ,
  `user_name` varchar(200)  NOT NULL ,
  `password` varchar(200)  NOT NULL ,
  `create_at` int(11) NOT NULL DEFAULT '0',
  `update_at` int(11) NOT NULL DEFAULT '0',
  `delete_at` int(11) NOT NULL DEFAULT '0'
) ;
