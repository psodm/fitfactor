$Server = "localhost"
$Port = "5432"
$DB = "fitfactor"
$Uid = "postgres"
$Pass = "password"

$DSN = "Driver={PostgreSQL UNICODE(x64)};Server=$Server;Port=$Port;Database=$DB;Uid=$Uid;Pwd=$Pass;"
$DBConn = New-Object System.Data.Odbc.OdbcConnection;
$DBConn.ConnectionString = $DSN;
$DBConn.Open();
$DBCmd = $DBConn.CreateCommand();
$DBCmd.CommandText = "SELECT * FROM users;";
$DBCmd.ExecuteReader();
$DBConn.Close();
