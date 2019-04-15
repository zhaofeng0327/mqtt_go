function connectSql() {
var mysql = cep_node.require('mysql');

var con = mysql.createConnection({
  host: "localhost",
  user: "phpmyadmin",
  password: "PHP@password123"
});

con.connect(function(err) {
  if (err) throw err;
  console.log("Connected!");
});

}
