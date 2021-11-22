SELECT user.*, product.nama_parent as ParentUserName 
FROM user LEFT JOIN product ON user.id_parent = product.id_parent