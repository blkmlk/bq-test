CREATE DATABASE IF NOT EXISTS bq;
use bq;
CREATE TABLE IF NOT EXISTS records (
    id INT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    f_symb VARCHAR(20),
    t_symb VARCHAR(20),
    change_24_hour DOUBLE,
    change_pct_24_hour DOUBLE,
    open_24_hour DOUBLE,
    volume_24_hour DOUBLE,
    low_24_hour DOUBLE,
    high_24_hour DOUBLE,
    price DOUBLE,
    supply DOUBLE,
    mktcap DOUBLE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
