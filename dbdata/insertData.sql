INSERT INTO
  bareskindb.User (user_id, uid, username, pass, fullName, email, phoneNumber, membership, inviteCode, created_date)
VALUES
  (1, '12345', 'admin', 'abc123', 'bareskin admin', 'admin@gmail.com', '1234567', 'distributor', 'abcde123', CAST('2017-09-19' AS DATETIME)),
  (2, '23456', 'distributor123', 'abc123', 'Juan Dela Cruz', 'jdelacruz@gmail.com', '6961234', 'distributor', 'invite123', CAST('2017-09-19' AS DATETIME));

INSERT INTO
  bareskindb.Order (order_id, user_id, name, address, mobile, payment_channel, payment_date, reference_number, order_date, order_status)
VALUES
  (1, 2, 'Joseph Paulo Los Banos', 'Antipolo City', '09778313782', 'BDO', CAST('2017-01-02' AS DATETIME), 'bb12312v2V', CAST('2017-01-01' AS DATETIME), 'pending');

INSERT INTO
  bareskindb.OrderedProducts (order_id, product_id, product_quantity)
VALUES
  (1, 1, 2),
  (1, 2, 3),
  (1, 4, 5);

INSERT INTO
  bareskindb.Sale (sale_id, user_id, reference_number, buyer_name, buyer_number, sale_date)
VALUES
  (1, 1, 'n1ui23b12iu', 'Jose Los Banos', '6966571', CAST('2017-01-03' AS DATETIME));

INSERT INTO
  bareskindb.SoldProducts (sale_id, product_id, product_quantity)
VALUES
  (1, 1, 3),
  (1, 3, 1),
  (1, 5, 2);

INSERT INTO
  bareskindb.Product (id, product_id, product_name, product_cost_unit, product_srp_unit, product_image_name, product_status)
VALUES
  (1, '1', 'Keratin Pure', 30000, 45000, '1.png', 'active'),
  (2, '2', 'Oatmeal Soap', 7500, 15000, '2.png', 'active'),
  (3, '3', 'Lip Scrub', 5000, 7500, '3.png', 'active'),
  (4, '4', 'Intense Whitening Soap', 7500, 12000, '4.png', 'active'),
  (5, '5', 'Luminous Whitening Soap', 7500, 12000, '5.png', 'active'),
  (6, '6', 'Apricot Scrub', 10000, 16500, '6.png', 'active'),
  (7, '7', 'Luminous Set', 30000, 45000, '7.png', 'active'),
  (8, '8', 'Pure White Glow Soap', 7500, 12000, '8.png', 'active'),
  (9, '9', 'Strawberry Lip Scrub', 10000, 16500, '9.png', 'active'),
  (10, '10', 'DEO Whitening', 10000, 16500, '10.png', 'active');

INSERT INTO
  bareskindb.Tag (tag_id, tag_name)
VALUES
  (1, 'Soaps'),
  (2, 'Scrubs'),
  (3, 'Sets'),
  (4, 'Deodorant');

INSERT INTO
  bareskindb.ProductTags (product_id, tag_id)
VALUES
  (1, 1),
  (2, 1),
  (3, 2),
  (4, 1),
  (5, 1),
  (6, 2),
  (7, 3),
  (8, 1),
  (9, 2),
  (10, 4);

INSERT INTO
  bareskindb.PaymentChannel (payment_channel_name, payment_channel_account)
VALUES
  ('BDO', '011280010574'),
  ('BPI', '4050-0123-74');

