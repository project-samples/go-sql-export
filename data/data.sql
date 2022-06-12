create table if not exists users (
    id varchar(45) not null,
    username varchar(45),
    email varchar(45),
    phone varchar(45),
    status varchar(45),
    createdDate datetime,
    primary key (id)
);
insert into `users` values ('1','abraham59','rory30@example.com','975-283-2267','1','2019-02-20 14:50:08'),('2','jerde.tito','qpacocha@example.com','(738)952-6078x1634','1','1995-06-13 12:49:06'),('3','shanon.gaylord','candelario.grant@example.net','1-691-238-8463','','1983-02-05 06:24:16'),('4','christiana60','louie85@example.org','361-461-5922','','2004-08-21 08:01:17'),('5','shyann52','alford91@example.com','05794709754','','1983-07-12 17:53:53'),('6','stone48','lance40@example.com','+37(4)0884560459','','2005-03-02 09:16:48'),('7','laurence64','colleen65@example.com','+81(3)8218226349','','2020-08-16 10:59:55'),('8','zstanton','stroman.jade@example.com','(423)483-5096','','2005-09-16 10:02:53'),('9','tschultz','kory.kunze@example.net','040-373-7213x048','','2018-06-28 12:15:40');
