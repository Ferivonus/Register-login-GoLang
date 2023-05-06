<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<style>
* {
  box-sizing: border-box;
}

.menu {
  float: left;
  width: 20%;
  text-align: center;
}

.menu a {
  background-color: #e5e5e5;
  padding: 8px;
  margin-top: 7px;
  display: block;
  width: 100%;
  color: black;
}

.main {
  float: left;
  width: 60%;
  padding: 0 20px;
}

.right {
  background-color: #e5e5e5;
  float: left;
  width: 20%;
  padding: 15px;
  margin-top: 7px;
  text-align: center;
}

@media only screen and (max-width: 620px) {
  /* For mobile phones: */
  .menu, .main, .right {
    width: 100%;
  }
}

footer {
            background-color: #F9F9F9;
            text-align: center;
            font-size: 12px;
            margin-top: 50px;
            padding: 15px;
        }
</style>
</head>
<body style="font-family:Verdana;color:#410000;">

<div style="background-color:#e5e5e5;padding:15px;text-align:center;">
  <h2>Welcome to my homework GUI</h2>
  <h2>Yeah,it is wierd to say that But I just said it.</h2>
</div>

<div style="overflow:auto">
  <div class="menu">
    <a href="/Account">Safe main section</a>
    <a href="/Uaccount">Unsafe section</a>
    <a href="/Application">Target website</a>
    <a href="/extra">Extras</a>
  </div>

  <div class="main">
    <h3>While considering the most appropriate programming language, we had some back and forth before finally deciding to use Go-lang."</h3>
    <p>it's connecting with MySql database on local, there is a table which is users and passwords are cripted on database </p>
  </div>

  <div class="right">
    <h2>About</h2>
    <p>Fahrettin Baştürk</p>
    <p>Turgut Özfidaner</p>
  </div>
</div>

    <footer>
        <p> Copyright © 2023, fahrettin basturk & Turgut Özfidaner. All rights reserved. </p>
    </footer>
</body>
</html>

