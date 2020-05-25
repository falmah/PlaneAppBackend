/* add users */
INSERT INTO app_db_user (name, surname, phone, email, password, role)
VALUES
    ('John', 'customer', '+3809187129', 'customer@gmail.com', 'qazwsxedc', 'customer'),
    ('Mike', 'operator', '+3809112129', 'operator@gmail.com', 'qazwsxedc', 'operator'),
    ('Nick', 'pilot', '+3801187129', 'pilot@gmail.com', 'qazwsxedc', 'pilot');


INSERT INTO app_db_customer (user_id)
VALUES ((SELECT id FROM app_db_user WHERE email = 'customer@gmail.com'));

INSERT INTO app_db_operator (user_id, company_name, city_id)
VALUES
(
    (SELECT id FROM app_db_user WHERE email = 'operator@gmail.com'), 
    'operator company', 
    (SELECT id FROM app_db_city WHERE name = 'Aomori')
);

INSERT INTO app_db_pilot (user_id, current_location)
VALUES 
(
    (SELECT id FROM app_db_user WHERE email = 'pilot@gmail.com'),
    (SELECT id FROM app_db_city WHERE name = 'Kano')
);

INSERT INTO app_db_ticket (
    customer_id, 
    status,
    cargo_type,   
    title,     
    date_from,      
    date_to,      
    dest_from,       
    dest_to,      
    price,        
    ticket_comment
) VALUES
(
    (SELECT id from app_db_customer WHERE user_id = (select id from app_db_user WHERE email = 'customer@gmail.com')),
    'open',
    'passenger',
    'passenger transportation',
    '2020-03-14',
    '2020-05-05',
    (SELECT id from app_db_airport WHERE name = 'Lake Ell Field'),
    (SELECT id from app_db_airport WHERE name = 'Sixmile Lake Airport'),
    15000,
    'Need to transport some peoples from Lake Ell field to the Sixmile Lake Airport.'
),
(
    (SELECT id from app_db_customer WHERE user_id = (select id from app_db_user WHERE email = 'customer@gmail.com')),
    'open',
    'passenger',
    'Aalborg Airport transport',
    '2020-03-14',
    '2020-05-05',
    (SELECT id from app_db_airport WHERE name = 'Aalborg Airport'),
    (SELECT id from app_db_airport WHERE name = 'Soissons - Courmelles Airport'),
    12000,
    'Need to transport some peoples from Aalborg Airport field to the Soissons - Courmelles Airport.'
),
(
    (SELECT id from app_db_customer WHERE user_id = (select id from app_db_user WHERE email = 'customer@gmail.com')),
    'open',
    'commodity',
    'Cocaine transport',
    '2020-03-14',
    '2020-05-05',
    (SELECT id from app_db_airport WHERE name = 'Ankeny Regional Airport'),
    (SELECT id from app_db_airport WHERE name = 'Fairview Airport' and latitude = 33.095100402832),
    25000,
    'Need to transport cocaine from to the Fairview Airport.'
);

INSERT INTO app_db_plane (name, registration_prefix, registration_id,
                            plane_type, current_location) 
VALUES
(
    'Learjet 23',
    'YV1000',
    'ghtyv-1571846-542ab',
    'two-engine',
    (SELECT id from app_db_airport WHERE name = 'Lake Ell Field')
),
(
    'Cirrus SR22',
    '9A-AAA',
    'iawow-1237123-141aw',
    'single engine',
    (SELECT id from app_db_airport WHERE name = 'Ankeny Regional Airport')
),
(
    'Gulfstream G500',
    'CX-ADA',
    'rftgy-9877654-6v1cz',
    'reactive two-engine',
    (SELECT id from app_db_airport WHERE name = 'Soissons - Courmelles Airport')
),
(
    'AH-225',
    'UR10000',
    'tghtd-3456782-v5h61',
    'transport',
    (SELECT id from app_db_airport WHERE name = 'Tradewind Airport')	
);

INSERT INTO app_db_operator_plane_bridge(operator_id, plane_id) VALUES
(
    (select id from app_db_operator where user_id = (select id from app_db_user where email = 'operator@gmail.com')),
    (select id from app_db_plane where name = 'Learjet 23')
),
(
    (select id from app_db_operator where user_id = (select id from app_db_user where email = 'operator@gmail.com')),
    (select id from app_db_plane where name = 'Cirrus SR22')
),
(
    (select id from app_db_operator where user_id = (select id from app_db_user where email = 'operator@gmail.com')),
    (select id from app_db_plane where name = 'Gulfstream G500')
),
(
    (select id from app_db_operator where user_id = (select id from app_db_user where email = 'operator@gmail.com')),
    (select id from app_db_plane where name = 'AH-225')
);

INSERT INTO app_db_pilot_request (status, operator_id, pilot_id, 
                                    price, required_license, required_visa, deadline, 
                                    request_comment, ticket_id, plane_id)
VALUES
(
    'open',
    (select id from app_db_operator where user_id = (select id from app_db_user where email = 'operator@gmail.com')),
    (select id from app_db_pilot where user_id = (select id from app_db_user where email = 'pilot@gmail.com')),
    5000,
    'license_1',
    'visa_1',
    '2020-03-04',
    'I have job for you. Please let me know if you able to do this.',
    (select id from app_db_ticket where price = 15000),
    (select id from app_db_plane where name = 'Learjet 23')
),
(
    'open',
    (select id from app_db_operator where user_id = (select id from app_db_user where email = 'operator@gmail.com')),
    (select id from app_db_pilot where user_id = (select id from app_db_user where email = 'pilot@gmail.com')),
    10000,
    'license_1',
    'visa_1',
    '2020-03-04',
    'Man.... This is your shine time)))',
    (select id from app_db_ticket where price = 25000),
    (select id from app_db_plane where name = 'Cirrus SR22')
);
