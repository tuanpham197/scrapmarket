use sendo_db;

ALTER TABLE shop_category_products
DROP FOREIGN KEY fk_product_category_shop_2;

ALTER TABLE shop_category_products
ADD CONSTRAINT fk_product_category_shop_2
FOREIGN KEY (shop_category_id)
REFERENCES shop_categories (id)