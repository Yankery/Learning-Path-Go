DROP PROCEDURE IF EXISTS addActor;
DELIMITER $$
CREATE PROCEDURE addActor(firstname varChar(100), lastname varChar(100))
BEGIN
SET @actor_id = 0;
SELECT actor_id
INTO @actor_id
FROM actor
WHERE first_name = firstname AND last_name = lastname;
IF @actor_id = 0 THEN 
	BEGIN
		INSERT INTO actor (first_name, last_name)
        VALUES (firstname, lastname);
        SET @actor_id = LAST_INSERT_ID();
	END;
    END IF;
    
    SELECT actor_id, first_name, last_name
    FROM actor
    WHERE actor_id = @actor_id;
END $$
DELIMITER ;