# Praktikum DBMS    

## Praktikum 1

1. Insert
    1. Insert 5 operators pada tabel operators
   ```sql
   insert into tugas_alta.operators (id, name, created_at, updated_at)
   values  (1, 'Razan', '2022-09-18 19:40:15', '2022-09-18 19:40:17'),
   (2, 'Fawwaz', '2022-09-18 19:40:27', '2022-09-18 19:40:30'),
   (3, 'Ahsan', '2022-09-18 19:40:45', '2022-09-18 19:40:47'),
   (4, 'Setiawan', '2022-09-18 19:41:05', '2022-09-18 19:41:07'),
   (5, 'Gideon', '2022-09-18 19:41:18', '2022-09-18 19:41:19');
   ```
   2. Insert 3 product type
    ```sql
   insert into tugas_alta.product_types (id, name, created_at, updated_at)
   values  (1, 'Makanan Ringan', '2022-09-18 19:43:47', '2022-09-18 19:43:48'),
   (2, 'Makanan Basah', '2022-09-18 19:48:36', '2022-09-18 19:53:41'),
   (3, 'Minuman Ringan', '2022-09-18 19:54:10', '2022-09-18 19:54:10');
    ```
   3. Insert 2 product dengan product type id = 1 dan operators id = 3, Insert 3 product dengan product type id = 2 dan operators id = 1, Insert 3 product dengan product type id = 3 dan operators id = 4
    ```sql
   insert into tugas_alta.products (id, product_type_id, operator_id, code, name, status, created_at, updated_at)
   values  (1, 1, 3, '101', 'Oreo', 1, '2022-09-18 19:55:11', '2022-09-18 19:55:12'),
   (2, 2, 1, '201', 'Nugget Champ', 1, '2022-09-18 19:55:45', '2022-09-18 19:55:46'),
   (3, 3, 4, '401', 'Pulpy Orange', 1, '2022-09-18 19:56:18', '2022-09-18 19:56:19'),
   (4, 1, 3, '102', 'CocoKrunch', 1, '2022-09-18 21:40:09', '2022-09-18 21:40:13'),
   (5, 2, 1, '202', 'Sosis Sonice', 1, '2022-09-18 21:40:42', '2022-09-18 21:40:43'),
   (6, 2, 1, '203', 'Kentang Curah', 1, '2022-09-18 21:41:35', '2022-09-18 21:41:36'),
   (7, 3, 4, '402', 'Frestea', 1, '2022-09-18 21:42:06', '2022-09-18 21:42:06'),
   (8, 3, 4, '403', 'Fruittea', 1, '2022-09-18 21:42:28', '2022-09-18 21:42:31');
    ```
   6. Insert product description pada setiap product
    ```sql
   insert into tugas_alta.product_descriptions (id, description, created_at, updated_at)
   values  (1, 'Jangan lupa dipakein susu!', '2022-09-18 19:58:09', '2022-09-18 19:58:14'),
   (2, 'Goreng dan sajikan!', '2022-09-18 19:58:33', '2022-09-18 19:58:35'),
   (3, 'Glek, Glek, Menyegarkan!', '2022-09-18 19:58:59', '2022-09-18 19:59:01'),
   (4, 'Dapatkan Hadiahnya!', '2022-09-18 21:58:59', '2022-09-18 22:45:01'),
   (5, 'Sosis lejat dan bergiji', '2022-09-18 22:03:59', '2022-09-18 22:48:01'),
   (6, 'Kentang murah nan nikmat', '2022-09-18 22:08:59', '2022-09-18 22:59:01'),
   (7, 'Segarnya terasa', '2022-09-18 22:18:59', '2022-09-18 22:59:01'),
   (8, 'Lebih besar dan nikmat', '2022-09-18 22:38:59', '2022-09-18 22:59:01');
    ```
   7. Insert 3 payment methods
    ```sql
   insert into tugas_alta.payment_methods (id, name, created_at, updated_at)
   values  (1, 'QRIS', '2022-09-18 19:59:15', '2022-09-18 19:59:17'),
   (2, 'Debit Card', '2022-09-18 19:59:21', '2022-09-18 19:59:20'),
   (3, 'Tunai', '2022-09-18 19:59:42', '2022-09-18 19:59:43');
    ```
   8. Insert 5 user pada tabel user
    ```sql
   insert into tugas_alta.users (id, status, dob, gender, created_at, updated_at)
   values  (1, 1, '2022-09-18', 'M', '2022-09-18 20:00:08', '2022-09-18 20:00:09'),
   (2, 1, '2022-09-07', 'F', '2022-09-18 20:00:19', '2022-09-18 21:00:11'),
   (3, 1, '2022-09-05', 'M', '2022-09-18 20:00:46', '2022-09-18 20:00:49'),
   (4, 1, '2022-09-01', 'F', '2022-09-18 20:01:20', '2022-09-18 20:01:23'),
   (5, 1, '2022-05-05', 'M', '2022-09-18 20:01:31', '2022-09-18 21:01:26');
    ```
   9. Insert 3 transaksi di masing-masing user
    ```sql
   insert into tugas_alta.transactions (id, user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at)
   values  (1, 1, 1, 'SUCCESS', 1, 8500.00, '2022-09-18 20:10:15', '2022-09-18 20:10:18'),
   (2, 3, 2, 'SUCCESS', 2, 16000.00, '2022-09-18 20:10:38', '2022-09-18 20:10:40'),
   (3, 4, 2, 'SUCCESS', 2, 9000.00, '2022-09-18 20:11:00', '2022-09-18 20:11:03');
    ```
   10. Insert 3 product di masing-masing transaksi
    ```sql
   insert into tugas_alta.transaction_details (transaction_id, product_id, status, qty, price, created_at, updated_at)
   values  (1, 1, 'SUCCESS', 1, 8500.00, '2022-09-18 20:11:38', '2022-09-18 20:11:39'),
   (2, 2, 'SUCCESS', 2, 16000.00, '2022-09-18 20:11:58', '2022-09-18 20:11:59'),
   (3, 3, 'SUCCESS', 2, 9000.00, '2022-09-18 20:12:13', '2022-09-18 20:12:14');
    ```
2. Select
    1. Tampilkan nama user/pelanggan dengan gender M
   ```sql
   select name from users where gender = 'M';
   ```
   2. Tampilkan product dengan id = 3
   ```sql
    select * from products where id = 3;
   ```
   3. Tampilkan data yang created_at dalam range 7 hari kebelakang dan mempunya nama mengandung kata 'a'
   ```sql
    select * from users where created_at between '2022-09-11' and '2022-09-18' and name like '%a%';
   ```
   4. Hitung jumlah user/pelanggan dengan gender F
   ```sql
    select count(*) from users where gender = 'F';
   ```
   5. Tampilkan data pelanggan dengan urutan sesuai nama abjad
   ```sql
    select * from users order by name;
   ```
   6. Tampilkan 5 data pada data product
   ```sql
    select * from products limit 5;
   ```
   
3. Update
   1. Ubah data product id 1 dengan nama product dummy
   ```sql
    update products set name = 'product dummy' where id = 1;
   ```
   2. Update qty = 3 pada transaction detail dengan product id 1
   ```sql
    update transaction_details set qty = 3 where product_id = 1;
   ```
4. Delete
   1. Delete data pada tabel product dengan id 1
   ```sql
    delete from products where id = 1;
   ```
   2. Delete data pada tabel product dengan product type id 1
   ```sql
    delete from products where product_type_id = 1;
   ```

## Praktikum 2 - Join, Union, Sub-query, Function
1. Gabungkan data transaksi dari user id 1 dan user id 2
   ```sql
    select * from transactions where user_id = 1
    union
    select * from transactions where user_id = 2;
   ```
2. Tampilkan jumlah harga transaksi user id 1
   ```sql
    select sum(total_price) from transactions where user_id = 1;
   ```
3. Tampilkan total transaksi dengan product type 2
   ```sql
    select sum(total_price) from transactions
    join transaction_details on transactions.id = transaction_details.transaction_id
    join products on transaction_details.product_id = products.id
    where products.product_type_id = 2;
   ```
4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
   ```sql
    select products.*, product_types.name from products
    join product_types on products.product_type_id = product_types.id;
   ```
5. Tampilkan semua field table transaction, field name table product, dan field name table user.
   ```sql
    select transactions.*, products.name, users.name from transactions
    join transaction_details on transactions.id = transaction_details.transaction_id
    join products on transaction_details.product_id = products.id
    join users on transactions.user_id = users.id;
   ```
6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
   ```sql
    create function delete_transaction(id int)
    returns int
    begin
    delete from transactions where id = id;
    delete from transaction_details where transaction_id = id;
    return 1;
    end;
   ```
7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
   ```sql
    create function update_total_qty(id int)
    returns int
    begin
    update transactions set total_qty = (select sum(qty) from transaction_details where transaction_id = id) where id = id;
    return 1;
    end;
   ```
8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
   ```sql
    select * from products where id not in (select product_id from transaction_details);
   ```

