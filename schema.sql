CREATE TABLE IF NOT EXISTS movie (
	id integer unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
	title varchar(255),
	released_year varchar(255),
    rating integer,
	created_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted boolean NOT NULL DEFAULT false	
) ENGINE=InnoDB;



//for testing the rating serach we need this data.
insert into movie (title,released_year,rating) values ("movie_gretaer_than_3_rating","2018",4);
insert into movie (title,released_year,rating) values ("movie_gretaer_than_4_rating","2018",5);
insert into movie (title,released_year,rating) values ("movie_gretaer_than_5_rating","2017",8);
insert into movie (title,released_year,rating) values ("movie_gretaer_than_5_rating","2017",7);